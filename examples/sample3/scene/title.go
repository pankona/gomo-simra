package scene

import (
	"image/color"

	"github.com/pankona/gomo-simra/examples/sample3/scene/config"
	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/image"
)

// Title represents a scene object for Title
type Title struct {
	simra      simra.Simraer
	background simra.Spriter
	text       simra.Spriter
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
	t.background.SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)
	t.simra.AddSprite(t.background)

	t.text = t.simra.NewSprite()
	t.text.SetScale(320, 80)
	t.text.SetPosition(t.text.GetScale().W/2, t.text.GetScale().H/2)
	t.simra.AddSprite(t.text)

	p := t.background.GetScale()
	var tex *simra.Texture
	tex = t.simra.NewImageTexture("t.png", image.Rect(0, 0, p.W, p.H))
	t.background.ReplaceTexture(tex)

	p = t.text.GetScale()
	tex = t.simra.NewTextTexture("text sample",
		60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, p.W, p.H))
	t.text.ReplaceTexture(tex)

	t.background.AddTouchListener(t)
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
