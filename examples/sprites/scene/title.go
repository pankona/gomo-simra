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
func (title *Title) Initialize(sim simra.Simraer) {
	title.simra = sim

	title.screenHeight = 1080 / 2
	title.screenWidth = 1920 / 2
	title.simra.SetDesiredScreenSize((float32)(title.screenHeight), (float32)(title.screenWidth))

	// initialize sprites
	title.initialize()

	title.numOfSprite = title.simra.NewSprite()
	title.numOfSprite.SetPosition(title.screenWidth/2, 100)
	title.numOfSprite.SetScale(title.screenWidth, 80)
	title.simra.AddSprite(title.numOfSprite)

	tex := title.simra.NewTextTexture("0",
		60, color.RGBA{255, 255, 255, 255}, image.Rect(0, 0, title.screenWidth, 80))
	title.numOfSprite.ReplaceTexture(tex)

	title.fpsText = title.simra.NewSprite()
	title.fpsText.SetPosition(title.screenWidth/4, 100)
	title.fpsText.SetScale(title.screenWidth, 80)
	title.simra.AddSprite(title.fpsText)

	tex = title.simra.NewTextTexture("0",
		60, color.RGBA{255, 255, 255, 255}, image.Rect(0, 0, title.screenWidth, 80))
	title.fpsText.ReplaceTexture(tex)
	go func() {
		for {
			<-time.After(1 * time.Second)
			tex = title.simra.NewTextTexture(strconv.Itoa(title.fps),
				60, color.RGBA{255, 255, 255, 255}, image.Rect(0, 0, title.screenWidth, 80))
			title.fpsText.ReplaceTexture(tex)
			title.mu.Lock()
			title.fps = 0
			title.mu.Unlock()
		}
	}()

	title.kokeshiTex = title.simra.NewImageTexture("sample2.png", image.Rect(0, 0, 64, 64))
}

func (title *Title) initialize() {
	title.simra.AddTouchListener(title)
}

var degree int

// Drive is called from simra.
// This is used to update sprites position.
// Thsi will be called 60 times per sec.
func (title *Title) Drive() {
	degree = (degree - 1) % 360
	for i := range title.sprites {
		r := float32(degree) * math.Pi / 180
		title.sprites[i].SetRotate(r)
	}
	title.mu.Lock()
	title.fps++
	title.mu.Unlock()
	//runtime.GC()
}

// OnTouchBegin is called when Title scene is Touched.
// It is caused by calling AddtouchListener for title.background sprite.
func (title *Title) OnTouchBegin(x, y float32) {
	title.spawnKokeshi(x, y)
}

// OnTouchMove is called when Title scene is Touched and moved.
// It is caused by calling AddtouchListener for title.background sprite.
func (title *Title) OnTouchMove(x, y float32) {
	title.spawnKokeshi(x, y)
}

// OnTouchEnd is called when Title scene is Touched and it is released.
// It is caused by calling AddtouchListener for title.background sprite.
func (title *Title) OnTouchEnd(x, y float32) {
	title.spawnKokeshi(x, y)
}

func (title *Title) spawnKokeshi(x, y float32) {
	// scene end. go to next scene
	sprite := title.simra.NewSprite()
	sprite.SetPosition((int)(x), (int)(y))
	sprite.SetScale(128, 128)
	title.simra.AddSprite(sprite)
	title.sprites = append(title.sprites, sprite)
	sprite.ReplaceTexture(title.kokeshiTex)

	tex := title.simra.NewTextTexture(strconv.Itoa(len(title.sprites)),
		60, color.RGBA{255, 255, 255, 255}, image.Rect(0, 0, title.screenWidth, 80))
	title.numOfSprite.ReplaceTexture(tex)

}
