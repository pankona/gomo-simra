package scene

import (
	"image"
	"image/color"

	"golang.org/x/mobile/asset"

	"github.com/pankona/gomo-simra/simra"
)

const (
	// ScreenWidth is screen width
	ScreenWidth = 1080 / 2
	// ScreenHeight is screen height
	ScreenHeight = 1920 / 2
)

// Title represents a scene object for Title
type Title struct {
	simra     simra.Simraer
	audio     simra.Audioer
	isPlaying bool
}

// Initialize initializes title scene
// This is called from simra.
// simra.SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (title *Title) Initialize(sim simra.Simraer) {
	simra.LogDebug("[IN]")
	title.simra = sim
	title.simra.SetDesiredScreenSize(ScreenWidth, ScreenHeight)
	// initialize sprites
	title.initialize()
	simra.LogDebug("[OUT]")
}

func (title *Title) initialize() {

	title.simra.AddTouchListener(title)

	sprite := title.simra.NewSprite()
	sprite.SetScale(ScreenWidth, 80)
	sprite.SetPosition(ScreenWidth/2, ScreenHeight/2)
	title.simra.AddSprite(sprite)
	s := sprite.GetScale()
	tex := title.simra.NewTextTexture("tap to play sound",
		60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, s.W, s.H))
	sprite.ReplaceTexture(tex)

}

// Drive is called from simra.
// This is used to update sprites position.
// Thsi will be called 60 times per sec.
func (title *Title) Drive() {
}

// OnTouchBegin is called when Title scene is Touched.
func (title *Title) OnTouchBegin(x, y float32) {
}

// OnTouchMove is called when Title scene is Touched and moved.
func (title *Title) OnTouchMove(x, y float32) {
}

// OnTouchEnd is called when Title scene is Touched and it is released.
func (title *Title) OnTouchEnd(x, y float32) {
	if title.isPlaying {
		err := title.audio.Stop()
		if err != nil {
			panic(err.Error())
		}
		title.isPlaying = false

	} else {
		title.audio = simra.NewAudio()
		resource, err := asset.Open("test_se.mp3")
		if err != nil {
			panic(err.Error())
		}

		err = title.audio.Play(resource, true, func(err error) {
			simra.LogDebug("playback complete callback. %s\n", err)
			title.isPlaying = false
		})
		if err != nil {
			panic(err)
		}
		title.isPlaying = true
	}
}
