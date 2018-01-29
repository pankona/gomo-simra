package scene

import (
	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/image"
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
	t.simra.SetDesiredScreenSize(1080/2, 1920/2)
	t.initialize()
	sim.SetOnStopCallback(func() {
		// nop
	})
}

func (t *Title) initialize() {
	t.background = t.simra.NewSprite()

	// add background sprite
	t.background.SetScale(1080/2, 1920/2)
	// put center of screen
	t.background.SetPosition(1080/2/2, 1920/2/2)

	t.simra.AddSprite(t.background)
	tex := t.simra.NewImageTexture("title.png",
		image.Rect(0, 0, t.background.GetScale().W, t.background.GetScale().H))
	t.background.ReplaceTexture(tex)

	t.background.AddTouchListener(t)
}

// Drive is called from simra.
// This is used to update sprites position.
// This function will be called 60 times per sec.
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
	t.simra.SetScene(&filestore{})
}
