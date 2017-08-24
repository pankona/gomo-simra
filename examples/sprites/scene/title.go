package scene

import (
	"image"
	"image/color"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/pankona/gomo-simra/simra"
)

// Title represents a scene object for Title
type Title struct {
	screenWidth  int
	screenHeight int
	sprites      []*simra.Sprite
	numOfSprite  *simra.Sprite
	fps          int
	fpsText      *simra.Sprite
	mu           sync.Mutex
	kokeshiTex   *simra.Texture
}

// Initialize initializes title scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (title *Title) Initialize() {
	simra.LogDebug("[IN]")

	title.screenHeight = 1080 / 2
	title.screenWidth = 1920 / 2
	simra.GetInstance().SetDesiredScreenSize((float32)(title.screenHeight), (float32)(title.screenWidth))

	// initialize sprites
	title.initialize()

	title.numOfSprite = simra.NewSprite()
	title.numOfSprite.X = (float32)(title.screenWidth / 2)
	title.numOfSprite.Y = 100
	title.numOfSprite.W = (float32)(title.screenWidth)
	title.numOfSprite.H = 80
	simra.GetInstance().AddSprite(title.numOfSprite)

	tex := simra.NewTextTexture("0",
		60, color.RGBA{255, 255, 255, 255}, image.Rect(0, 0, title.screenWidth, 80))
	title.numOfSprite.ReplaceTexture(tex)

	title.fpsText = simra.NewSprite()
	title.fpsText.X = (float32)(title.screenWidth / 4)
	title.fpsText.Y = 100
	title.fpsText.W = (float32)(title.screenWidth)
	title.fpsText.H = 80
	simra.GetInstance().AddSprite(title.fpsText)

	tex = simra.NewTextTexture("0",
		60, color.RGBA{255, 255, 255, 255}, image.Rect(0, 0, title.screenWidth, 80))
	title.fpsText.ReplaceTexture(tex)
	go func() {
		for {
			<-time.After(1 * time.Second)
			tex = simra.NewTextTexture(strconv.Itoa(title.fps),
				60, color.RGBA{255, 255, 255, 255}, image.Rect(0, 0, title.screenWidth, 80))
			title.fpsText.ReplaceTexture(tex)
			title.mu.Lock()
			title.fps = 0
			title.mu.Unlock()
		}
	}()

	title.kokeshiTex = simra.NewImageTexture("sample.png", image.Rect(0, 0, 384, 384))

	simra.LogDebug("[OUT]")
}

func (title *Title) initialize() {
	simra.GetInstance().AddTouchListener(title)
}

var degree int

// Drive is called from simra.
// This is used to update sprites position.
// Thsi will be called 60 times per sec.
func (title *Title) Drive() {
	degree = (degree - 1) % 360
	for i := range title.sprites {
		r := float32(degree) * math.Pi / 180
		title.sprites[i].R = (float32)(r)
	}
	title.mu.Lock()
	title.fps++
	title.mu.Unlock()
	//runtime.GC()
}

// OnTouchBegin is called when Title scene is Touched.
// It is caused by calling AddtouchListener for title.background sprite.
func (title *Title) OnTouchBegin(x, y float32) {
}

// OnTouchMove is called when Title scene is Touched and moved.
// It is caused by calling AddtouchListener for title.background sprite.
func (title *Title) OnTouchMove(x, y float32) {
}

// OnTouchEnd is called when Title scene is Touched and it is released.
// It is caused by calling AddtouchListener for title.background sprite.
func (title *Title) OnTouchEnd(x, y float32) {
	// scene end. go to next scene
	sprite := simra.NewSprite()
	sprite.W = 128
	sprite.H = 128
	sprite.X = x
	sprite.Y = y
	simra.GetInstance().AddSprite(sprite)
	title.sprites = append(title.sprites, sprite)
	sprite.ReplaceTexture(title.kokeshiTex)

	tex := simra.NewTextTexture(strconv.Itoa(len(title.sprites)),
		60, color.RGBA{255, 255, 255, 255}, image.Rect(0, 0, title.screenWidth, 80))
	title.numOfSprite.ReplaceTexture(tex)
}
