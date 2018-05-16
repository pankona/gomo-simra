package scene

import (
	"image/color"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/image"
)

// Title represents a scene object for Title
type Title struct {
	simra        simra.Simraer
	screenWidth  int
	screenHeight int
	sprites      []simra.Spriter
	numOfSprite  simra.Spriter
	fps          int
	fpsText      simra.Spriter
	mu           sync.Mutex
	kokeshiTex   *simra.Texture
}

// Initialize initializes title scene
// This is called from simra.
// simra.SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (t *Title) Initialize(sim simra.Simraer) {
	t.simra = sim

	t.screenHeight = 1080 / 2
	t.screenWidth = 1920 / 2
	t.simra.SetDesiredScreenSize((float32)(t.screenHeight), (float32)(t.screenWidth))

	// initialize sprites
	t.initialize()

	t.numOfSprite = t.simra.NewSprite()
	t.numOfSprite.SetPosition(float32(t.screenWidth/2), 100)
	t.numOfSprite.SetScale(float32(t.screenWidth), 80)
	t.simra.AddSprite(t.numOfSprite)

	tex := t.simra.NewTextTexture("0",
		60, color.RGBA{255, 255, 255, 255}, image.Rect(0, 0, float32(t.screenWidth), 80))
	t.numOfSprite.ReplaceTexture(tex)

	t.fpsText = t.simra.NewSprite()
	t.fpsText.SetPosition(float32(t.screenWidth/4), 100)
	t.fpsText.SetScale(float32(t.screenWidth), 80)
	t.simra.AddSprite(t.fpsText)

	tex = t.simra.NewTextTexture("0",
		60, color.RGBA{255, 255, 255, 255}, image.Rect(0, 0, float32(t.screenWidth), 80))
	t.fpsText.ReplaceTexture(tex)
	go func() {
		for {
			<-time.After(1 * time.Second)
			tex = t.simra.NewTextTexture(strconv.Itoa(t.fps),
				60, color.RGBA{255, 255, 255, 255}, image.Rect(0, 0, float32(t.screenWidth), 80))
			t.fpsText.ReplaceTexture(tex)
			t.mu.Lock()
			t.fps = 0
			t.mu.Unlock()
		}
	}()

	t.kokeshiTex = t.simra.NewImageTexture("sample2.png", image.Rect(0, 0, 64, 64))
}

func (t *Title) initialize() {
	t.simra.AddTouchListener(t)
}

var degree int

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (t *Title) Drive() {
	degree = (degree - 1) % 360
	for i := range t.sprites {
		r := float32(degree) * math.Pi / 180
		t.sprites[i].SetRotate(r)
	}
	t.mu.Lock()
	t.fps++
	t.mu.Unlock()
	//runtime.GC()
}

// OnTouchBegin is called when Title scene is Touched.
// It is caused by calling AddtouchListener for title.background sprite.
func (t *Title) OnTouchBegin(x, y float32) {
	t.spawnKokeshi(x, y)
}

// OnTouchMove is called when Title scene is Touched and moved.
// It is caused by calling AddtouchListener for title.background sprite.
func (t *Title) OnTouchMove(x, y float32) {
	t.spawnKokeshi(x, y)
}

// OnTouchEnd is called when Title scene is Touched and it is released.
// It is caused by calling AddtouchListener for title.background sprite.
func (t *Title) OnTouchEnd(x, y float32) {
	t.spawnKokeshi(x, y)
}

func (t *Title) spawnKokeshi(x, y float32) {
	sprite := t.simra.NewSprite()
	sprite.SetPosition(x, y)
	sprite.SetScale(128, 128)
	t.simra.AddSprite(sprite)
	t.sprites = append(t.sprites, sprite)
	sprite.ReplaceTexture(t.kokeshiTex)

	tex := t.simra.NewTextTexture(strconv.Itoa(len(t.sprites)),
		60, color.RGBA{255, 255, 255, 255}, image.Rect(0, 0, float32(t.screenWidth), 80))
	t.numOfSprite.ReplaceTexture(tex)

	// later sprite goes far side
	t.simra.SetZIndex(sprite, len(t.sprites))
}
