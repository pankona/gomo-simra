package peer

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"sync"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/pankona/gomo-simra/simra/config"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/asset"
	"golang.org/x/mobile/exp/app/debug"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/exp/sprite/glsprite"
	"golang.org/x/mobile/gl"
)

// GLer interface represents interface of GL
type GLer interface {
	// Initialize initializes GLPeer.
	// This function must be called inadvance of using GLPeer
	Initialize(glctx gl.Context)
	// LoadTexture return texture that is loaded by the information of arguments.
	// Loaded texture can assign using AddSprite function.
	LoadTexture(assetName string, rect image.Rectangle) sprite.SubTex
	// MakeTextureByText createst and return texture by speicied text
	// Loaded texture can assign using AddSprite function.
	// TODO: font parameterize
	MakeTextureByText(text string, fontsize float64, fontcolor color.RGBA, rect image.Rectangle) sprite.SubTex
	// Finalize finalizes GLPeer.
	// This is called at termination of application.
	Finalize()
	// Update updates screen.
	// This is called 60 times per 1 sec.
	Update(publishFunc func() app.PublishResult)
	// Reset resets current gl context.
	// All sprites are also cleaned.
	// This is called at changing of scene, and
	// this function is for clean previous scene.
	Reset()
	// NewTexture returns a new Texture instance
	NewTexture(s sprite.SubTex) *Texture
	// ReleaseTexture releases specified texture
	ReleaseTexture(t *Texture)
}

var (
	glPeer    = &GLPeer{}
	startTime = time.Now()
)

// GLPeer represents gl context.
// Singleton.
type GLPeer struct {
	glctx  gl.Context
	images *glutil.Images
	fps    *debug.FPS
	eng    sprite.Engine
	scene  *sprite.Node
	mu     sync.Mutex
}

// GetGLPeer returns a instance of GLPeer.
// Since GLPeer is singleton, it is necessary to
// call this function to get GLPeer instance.
func GetGLPeer() GLer {
	return glPeer
}

// Initialize initializes GLPeer.
// This function must be called inadvance of using GLPeer
func (glpeer *GLPeer) Initialize(glctx gl.Context) {
	LogDebug("IN")
	glpeer.mu.Lock()
	defer glpeer.mu.Unlock()
	glpeer.glctx = glctx

	// transparency of png
	glpeer.glctx.Enable(gl.BLEND)
	glpeer.glctx.BlendEquation(gl.FUNC_ADD)
	glpeer.glctx.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	glpeer.images = glutil.NewImages(glctx)
	glpeer.fps = debug.NewFPS(glpeer.images)
	glpeer.initEng()

	LogDebug("OUT")
}

func (glpeer *GLPeer) initEng() {
	if glpeer.eng != nil {
		glpeer.eng.Release()
	}
	glpeer.eng = glsprite.Engine(glpeer.images)
	glpeer.scene = &sprite.Node{}
	glpeer.eng.Register(glpeer.scene)
	glpeer.eng.SetTransform(glpeer.scene, f32.Affine{
		{1, 0, 0},
		{0, 1, 0},
	})
}

type arrangerFunc func(e sprite.Engine, n *sprite.Node, t clock.Time)

func (a arrangerFunc) Arrange(e sprite.Engine, n *sprite.Node, t clock.Time) { a(e, n, t) }

func (glpeer *GLPeer) newNode(fn arrangerFunc) *sprite.Node {
	glpeer.mu.Lock()
	defer glpeer.mu.Unlock()
	n := &sprite.Node{Arranger: arrangerFunc(fn)}
	glpeer.eng.Register(n)
	glpeer.scene.AppendChild(n)
	return n
}

func (glpeer *GLPeer) appendChild(n *sprite.Node) {
	glpeer.mu.Lock()
	defer glpeer.mu.Unlock()
	glpeer.scene.AppendChild(n)
}

func (glpeer *GLPeer) removeChild(n *sprite.Node) {
	glpeer.mu.Lock()
	defer glpeer.mu.Unlock()
	glpeer.scene.RemoveChild(n)
}

// LoadTexture return texture that is loaded by the information of arguments.
// Loaded texture can assign using AddSprite function.
func (glpeer *GLPeer) LoadTexture(assetName string, rect image.Rectangle) sprite.SubTex {
	LogDebug("IN")
	glpeer.mu.Lock()
	defer glpeer.mu.Unlock()

	a, err := asset.Open(assetName)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := a.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	img, _, err := image.Decode(a)
	if err != nil {
		log.Fatal(err)
	}
	t, err := glpeer.eng.LoadTexture(img)
	if err != nil {
		log.Fatal(err)
	}

	LogDebug("OUT")
	return sprite.SubTex{T: t, R: rect}
}

// MakeTextureByText createst and return texture by speicied text
// Loaded texture can assign using AddSprite function.
// TODO: font parameterize
func (glpeer *GLPeer) MakeTextureByText(text string, fontsize float64, fontcolor color.RGBA, rect image.Rectangle) sprite.SubTex {
	LogDebug("IN")
	glpeer.mu.Lock()
	defer glpeer.mu.Unlock()

	dpi := float64(72)
	width := rect.Dx()
	height := rect.Dy()
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	fg, bg := image.NewUniform(fontcolor), image.Transparent
	draw.Draw(img, img.Bounds(), bg, image.Point{}, draw.Src)

	// Draw the text.
	h := font.HintingNone

	gofont, _ := truetype.Parse(goregular.TTF)

	d := &font.Drawer{
		Dst: img,
		Src: fg,
		Face: truetype.NewFace(gofont, &truetype.Options{
			Size:    fontsize,
			DPI:     dpi,
			Hinting: h,
		}),
	}

	textWidth := d.MeasureString(text)

	d.Dot = fixed.Point26_6{
		X: fixed.I(width/2) - textWidth/2,
		Y: fixed.I(int(fontsize * dpi / 72)),
	}
	d.DrawString(text)

	t, err := glpeer.eng.LoadTexture(img)
	if err != nil {
		log.Fatal(err)
	}

	LogDebug("OUT")
	return sprite.SubTex{T: t, R: rect}
}

// Finalize finalizes GLPeer.
// This is called at termination of application.
func (glpeer *GLPeer) Finalize() {
	LogDebug("IN")
	glpeer.mu.Lock()
	defer glpeer.mu.Unlock()

	GetSpriteContainer().RemoveSprites()
	glpeer.eng.Release()
	glpeer.fps.Release()
	glpeer.images.Release()
	glpeer.glctx = nil
	LogDebug("OUT")
}

// Update updates screen.
// This is called 60 times per 1 sec.
func (glpeer *GLPeer) Update(publishFunc func() app.PublishResult) {
	glpeer.mu.Lock()
	defer glpeer.mu.Unlock()

	if glpeer.glctx == nil {
		return
	}
	glpeer.glctx.ClearColor(0, 0, 0, 1) // black background
	glpeer.glctx.Clear(gl.COLOR_BUFFER_BIT)
	now := clock.Time(time.Since(startTime) * 60 / time.Second)

	glpeer.apply()

	glpeer.eng.Render(glpeer.scene, now, screensize.sz)
	if config.DEBUG {
		glpeer.fps.Draw(screensize.sz)
	}

	// app.Publish() calls glctx.Flush, it should be called within this mutex locking.
	publishFunc()
}

// Reset resets current gl context.
// All sprites are also cleaned.
// This is called at changing of scene, and
// this function is for clean previous scene.
func (glpeer *GLPeer) Reset() {
	LogDebug("IN")
	glpeer.mu.Lock()
	defer glpeer.mu.Unlock()
	GetSpriteContainer().RemoveSprites()
	glpeer.initEng()
	LogDebug("OUT")
}

func (glpeer *GLPeer) apply() {

	snpairs := spriteContainer.spriteNodePairs

	for i := range snpairs {
		sc := snpairs[i]
		if sc.sprite == nil || !sc.inuse {
			continue
		}

		affine := &f32.Affine{
			{1, 0, 0},
			{0, 1, 0},
		}
		affine.Translate(affine,
			sc.sprite.X*screensize.scale-sc.sprite.W/2*screensize.scale+screensize.marginWidth/2,
			(screensize.height-sc.sprite.Y)*screensize.scale-sc.sprite.H/2*screensize.scale+screensize.marginHeight/2)
		if sc.sprite.R != 0 {
			affine.Translate(affine,
				0.5*sc.sprite.W*screensize.scale,
				0.5*sc.sprite.H*screensize.scale)
			affine.Rotate(affine, sc.sprite.R)
			affine.Translate(affine,
				-0.5*sc.sprite.W*screensize.scale,
				-0.5*sc.sprite.H*screensize.scale)
		}
		affine.Scale(affine,
			sc.sprite.W*screensize.scale,
			sc.sprite.H*screensize.scale)
		glpeer.eng.SetTransform(sc.node, *affine)
	}
}

// Texture represents a texture object that contains subTex
type Texture struct {
	glPeer *GLPeer
	subTex sprite.SubTex
}

// NewTexture returns a new Texture instance
func (glpeer *GLPeer) NewTexture(s sprite.SubTex) *Texture {
	return &Texture{
		glPeer: glpeer,
		subTex: s,
	}
}

// ReleaseTexture releases specified texture
func (glpeer *GLPeer) ReleaseTexture(t *Texture) {
	glpeer.mu.Lock()
	glpeer.mu.Unlock()
	t.subTex.T.Release()
}
