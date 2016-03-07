package scene

import (
	"image"

	"github.com/pankona/gomo-simra/simra"
)

// Title represents a scene object for Title
type Title struct {
	background simra.Sprite
}

// Initialize initializes title scene
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this app
func (title *Title) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(1080/2, 1920/2)

	// initialize sprites
	title.initialize()

	simra.LogDebug("[OUT]")
}

func (title *Title) initialize() {
	// add background sprite
	title.background.W = float32(1080 / 2)
	title.background.H = float32(1920 / 2)

	// put center of screen
	title.background.X = 1080 / 2 / 2
	title.background.Y = 1920 / 2 / 2

	simra.GetInstance().AddSprite("title.png",
		image.Rect(0, 0, int(title.background.W), int(title.background.H)),
		&title.background)

	title.background.AddTouchListener(title)
}

// Drive is called from simra.
// This is used to update sprites position.
// Thsi will be called 60 times per sec.
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
