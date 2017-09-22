package scene

import (
	"image"

	"github.com/pankona/gomo-simra/examples/sample2/scene/config"
	"github.com/pankona/gomo-simra/simra"
)

// Title represents a scene object for Title
type Title struct {
	simra      simra.Simraer
	background simra.Spriter
}

// Initialize initializes title scene
// This is called from simra.
// simra.SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (title *Title) Initialize(sim simra.Simraer) {
	simra.LogDebug("[IN]")
	title.simra = sim

	title.simra.SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// initialize sprites
	title.initialize()

	simra.LogDebug("[OUT]")
}

func (title *Title) initialize() {
	title.background = title.simra.NewSprite()
	// add background sprite
	title.background.SetScale(config.ScreenWidth, config.ScreenHeight)
	// put center of screen
	title.background.SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)

	title.simra.AddSprite(title.background)

	title.background.AddTouchListener(title)
	tex := title.simra.NewImageTexture("title.png",
		image.Rect(0, 0, int(title.background.GetScale().W), int(title.background.GetScale().H)))
	title.background.ReplaceTexture(tex)

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
	title.simra.SetScene(&CtrlTrial{})
}
