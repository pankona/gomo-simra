package peer

import (
	"image"
	"log"
	"time"

	"golang.org/x/mobile/asset"
	"golang.org/x/mobile/exp/app/debug"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/exp/sprite/glsprite"
	"golang.org/x/mobile/gl"
)

var glPeer *GLPeer

var startTime = time.Now()

type PeerSprite struct {
	W float32
	H float32
	X float32
	Y float32
	R float32
}

type peerSpriteContainer struct {
	peerSprite *PeerSprite
	node       *sprite.Node
}

type GLPeer struct {
	glctx                gl.Context
	images               *glutil.Images
	fps                  *debug.FPS
	eng                  sprite.Engine
	scene                *sprite.Node
	peerSpriteContainers []*peerSpriteContainer
}

func GetGLPeer() *GLPeer {
	LogDebug("IN")
	if glPeer == nil {
		glPeer = &GLPeer{}
	}
	LogDebug("OUT")
	return glPeer
}

func (self *GLPeer) Initialize(in_glctx gl.Context) {
	LogDebug("IN")
	self.glctx = in_glctx

	// transparency of png
	self.glctx.Enable(gl.BLEND)
	self.glctx.BlendEquation(gl.FUNC_ADD)
	self.glctx.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	self.images = glutil.NewImages(in_glctx)
	self.fps = debug.NewFPS(self.images)
	self.initEng()
	LogDebug("OUT")
}

func (self *GLPeer) initEng() {
	if self.eng != nil {
		self.eng.Release()
	}
	self.eng = glsprite.Engine(self.images)
	self.scene = &sprite.Node{}
	self.eng.Register(self.scene)
	self.eng.SetTransform(self.scene, f32.Affine{
		{1, 0, 0},
		{0, 1, 0},
	})
}

func (self *GLPeer) newNode() *sprite.Node {
	n := &sprite.Node{}
	self.eng.Register(n)
	self.scene.AppendChild(n)
	return n
}

func (self *GLPeer) LoadTexture(assetName string, rect image.Rectangle) sprite.SubTex {
	LogDebug("IN")
	a, err := asset.Open(assetName)
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	img, _, err := image.Decode(a)
	if err != nil {
		log.Fatal(err)
	}
	t, err := self.eng.LoadTexture(img)
	if err != nil {
		log.Fatal(err)
	}

	LogDebug("OUT")
	return sprite.SubTex{t, rect}
}

func (self *GLPeer) Finalize() {
	LogDebug("IN")
	self.eng.Release()
	self.fps.Release()
	self.images.Release()
	self.glctx = nil
	LogDebug("OUT")
}

func (self *GLPeer) Update() {
	if self.glctx == nil {
		return
	}
	self.glctx.ClearColor(1, 1, 1, 1) // white background
	self.glctx.Clear(gl.COLOR_BUFFER_BIT)
	now := clock.Time(time.Since(startTime) * 60 / time.Second)

	self.apply()

	self.eng.Render(self.scene, now, sz)
	self.fps.Draw(sz)
}

func (self *GLPeer) AddSprite(ps *PeerSprite, subTex sprite.SubTex) {
	LogDebug("IN")
	var psc peerSpriteContainer
	psc.peerSprite = ps
	psc.node = self.newNode()
	self.peerSpriteContainers = append(self.peerSpriteContainers, &psc)
	self.eng.SetSubTex(psc.node, subTex)
	LogDebug("OUT")
}

func (self *GLPeer) Reset() {
	LogDebug("IN")
	self.peerSpriteContainers = nil
	self.initEng()
	LogDebug("OUT")
}

func (self *GLPeer) apply() {

	peerSpriteContainers := self.peerSpriteContainers

	for i := range peerSpriteContainers {
		psc := peerSpriteContainers[i]
		if psc.peerSprite == nil {
			continue
		}

		affine := &f32.Affine{
			{1, 0, 0},
			{0, 1, 0},
		}
		affine.Translate(affine,
			psc.peerSprite.X*desiredScreenSize.scale-psc.peerSprite.W/2*desiredScreenSize.scale,
			psc.peerSprite.Y*desiredScreenSize.scale-psc.peerSprite.H/2*desiredScreenSize.scale)
		if psc.peerSprite.R != 0 {
			affine.Translate(affine,
				0.5*psc.peerSprite.W*desiredScreenSize.scale,
				0.5*psc.peerSprite.H*desiredScreenSize.scale)
			affine.Rotate(affine, psc.peerSprite.R)
			affine.Translate(affine,
				-0.5*psc.peerSprite.W*desiredScreenSize.scale,
				-0.5*psc.peerSprite.H*desiredScreenSize.scale)
		}
		affine.Scale(affine,
			psc.peerSprite.W*desiredScreenSize.scale,
			psc.peerSprite.H*desiredScreenSize.scale)
		self.eng.SetTransform(psc.node, *affine)
	}
}

type arrangerFunc func(e sprite.Engine, n *sprite.Node, t clock.Time)

func (a arrangerFunc) Arrange(e sprite.Engine, n *sprite.Node, t clock.Time) { a(e, n, t) }
