package scene

import (
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/examples/sample2/scene/config"
	"github.com/pankona/gomo-simra/simra"
)

// Title represents a scene object for Title
type Title struct {
	background simra.Spriter
	text       simra.Spriter
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
	title.background = simra.NewSprite()
	title.background.SetScale(config.ScreenWidth, config.ScreenHeight)
	title.background.SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)
	simra.GetInstance().AddSprite(title.background)

	title.text = simra.NewSprite()
	title.text.SetScale(320, 80)
	title.text.SetPosition(title.text.GetScale().W/2, title.text.GetScale().H/2)
	simra.GetInstance().AddSprite(title.text)

	p := title.background.GetScale()
	var tex *simra.Texture
	tex = simra.NewImageTexture("title.png", image.Rect(0, 0, p.W, p.H))
	title.background.ReplaceTexture(tex)

	p = title.text.GetScale()
	tex = simra.NewTextTexture("text sample",
		60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, p.W, p.H))
	title.text.ReplaceTexture(tex)

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
