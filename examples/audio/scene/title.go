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
	audio     simra.Audioer
	isPlaying bool
}

// Initialize initializes title scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (title *Title) Initialize() {
	simra.LogDebug("[IN]")
	simra.GetInstance().SetDesiredScreenSize(ScreenWidth, ScreenHeight)
	// initialize sprites
	title.initialize()
	simra.LogDebug("[OUT]")
}

func (title *Title) initialize() {

	simra.GetInstance().AddTouchListener(title)

	sprite := simra.NewSprite()
	sprite.W = ScreenWidth
	sprite.H = 80
	sprite.X = ScreenWidth / 2
	sprite.Y = ScreenHeight / 2
	simra.GetInstance().AddTextSprite("tap to play sound",
		60,
		color.RGBA{255, 0, 0, 255},
		image.Rect(0, 0, int(sprite.W), int(sprite.H)),
		sprite)
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
		resource, err := asset.Open("bgm_maoudamashii_8bit28.mp3")
		if err != nil {
			panic(err.Error())
		}

		err = title.audio.Play(resource, true, func() {
			title.isPlaying = false
		})
		if err != nil {
			panic(err)
		}
		title.isPlaying = true
	}
}
