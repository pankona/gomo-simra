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
func (t *Title) Initialize(sim simra.Simraer) {
	t.simra = sim
	t.simra.SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)
	t.initialize()
}

func (t *Title) initialize() {
	t.background = t.simra.NewSprite()
	t.background.SetScale(config.ScreenWidth, config.ScreenHeight)
	// put center of screen
	t.background.SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)
	// add background sprite
	t.simra.AddSprite(t.background)

	t.background.AddTouchListener(t)
	tex := t.simra.NewImageTexture("title.png",
		image.Rect(0, 0, int(t.background.GetScale().W), int(t.background.GetScale().H)))
	t.background.ReplaceTexture(tex)

}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (t *Title) Drive() {
}

// OnTouchBegin is called when Title scene is Touched.
// It is caused by calling AddtouchListener for title.background sprite.
func (t *Title) OnTouchBegin(x, y float32) {
}

// OnTouchMove is called when Title scene is Touched and moved.
// It is caused by calling AddtouchListener for title.background sprite.
func (t *Title) OnTouchMove(x, y float32) {
}

// OnTouchEnd is called when Title scene is Touched and it is released.
// It is caused by calling AddtouchListener for title.background sprite.
func (t *Title) OnTouchEnd(x, y float32) {
	// scene end. go to next scene
	t.simra.SetScene(&sample{})
}
