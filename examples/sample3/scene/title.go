package scene

import (
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/examples/sample2/scene/config"
	"github.com/pankona/gomo-simra/simra"
)

// Title represents a scene object for Title
type Title struct {
	background simra.Sprite
	text       simra.Sprite
}

// Initialize initializes title scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (title *Title) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// initialize sprites
	title.initialize()

	simra.LogDebug("[OUT]")
}

func (title *Title) initialize() {
	// add background sprite
	title.background.W = float32(config.ScreenWidth)
	title.background.H = float32(config.ScreenHeight)

	// put center of screen
	title.background.X = config.ScreenWidth / 2
	title.background.Y = config.ScreenHeight / 2

	simra.GetInstance().AddSprite("title.png",
		image.Rect(0, 0, int(title.background.W), int(title.background.H)),
		&title.background)

	title.text.W = 320
	title.text.H = 80
	title.text.X = title.text.W / 2
	title.text.Y = title.text.H / 2
	simra.GetInstance().AddTextSprite("text sample",
		60, // fontsize
		color.RGBA{255, 0, 0, 255},
		image.Rect(0, 0, int(title.text.W), int(title.text.H)),
		&title.text)

	title.background.AddTouchListener(title)
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (title *Title) Drive() {
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
	simra.GetInstance().SetScene(&Stage1{})
}
