package scene

import (
	"image"
	"image/color"

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
	simra       simra.Simraer
	text        simra.Spriter
	isAnimating bool
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
	sprite := title.simra.NewSprite()
	sprite.SetScale(ScreenWidth, 80)
	sprite.SetPosition(ScreenWidth/2, ScreenHeight/2)

	animationSet := simra.NewAnimationSet()
	animationSet.AddTexture(title.simra.NewTextTexture("a", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, sprite.GetScale().W, sprite.GetScale().H)))
	animationSet.AddTexture(title.simra.NewTextTexture("n", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, sprite.GetScale().W, sprite.GetScale().H)))
	animationSet.AddTexture(title.simra.NewTextTexture("i", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, sprite.GetScale().W, sprite.GetScale().H)))
	animationSet.AddTexture(title.simra.NewTextTexture("m", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, sprite.GetScale().W, sprite.GetScale().H)))
	animationSet.AddTexture(title.simra.NewTextTexture("a", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, sprite.GetScale().W, sprite.GetScale().H)))
	animationSet.AddTexture(title.simra.NewTextTexture("t", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, sprite.GetScale().W, sprite.GetScale().H)))
	animationSet.AddTexture(title.simra.NewTextTexture("i", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, sprite.GetScale().W, sprite.GetScale().H)))
	animationSet.AddTexture(title.simra.NewTextTexture("o", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, sprite.GetScale().W, sprite.GetScale().H)))
	animationSet.AddTexture(title.simra.NewTextTexture("n", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, sprite.GetScale().W, sprite.GetScale().H)))
	animationSet.AddTexture(title.simra.NewTextTexture("t", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, sprite.GetScale().W, sprite.GetScale().H)))
	animationSet.AddTexture(title.simra.NewTextTexture("e", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, sprite.GetScale().W, sprite.GetScale().H)))
	animationSet.AddTexture(title.simra.NewTextTexture("s", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, sprite.GetScale().W, sprite.GetScale().H)))
	animationSet.AddTexture(title.simra.NewTextTexture("t", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, sprite.GetScale().W, sprite.GetScale().H)))

	animationSet.SetInterval(12) // frames

	sprite.AddAnimationSet("animation test", animationSet)
	title.simra.AddSprite(sprite)
	tex := title.simra.NewTextTexture("animation test",
		60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, sprite.GetScale().W, sprite.GetScale().H))
	sprite.ReplaceTexture(tex)

	title.simra.AddTouchListener(title)
	title.text = sprite
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
	if title.isAnimating {
		simra.LogDebug("stop animation")
		title.text.StopAnimation()
		title.isAnimating = false
	} else {
		simra.LogDebug("start animation")
		title.text.StartAnimation("animation test", true, func() {})
		title.isAnimating = true
	}
}
