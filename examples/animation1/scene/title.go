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
func (t *Title) Initialize(sim simra.Simraer) {
	t.simra = sim
	t.simra.SetDesiredScreenSize(ScreenWidth, ScreenHeight)
	t.initialize()
}

func (t *Title) initialize() {
	sprite := t.simra.NewSprite()
	sprite.SetScale(ScreenWidth, 80)
	sprite.SetPosition(ScreenWidth/2, ScreenHeight/2)

	animationSet := simra.NewAnimationSet()
	animationSet.AddTexture(t.simra.NewTextTexture("a", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.GetScale().W), int(sprite.GetScale().H))))
	animationSet.AddTexture(t.simra.NewTextTexture("n", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.GetScale().W), int(sprite.GetScale().H))))
	animationSet.AddTexture(t.simra.NewTextTexture("i", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.GetScale().W), int(sprite.GetScale().H))))
	animationSet.AddTexture(t.simra.NewTextTexture("m", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.GetScale().W), int(sprite.GetScale().H))))
	animationSet.AddTexture(t.simra.NewTextTexture("a", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.GetScale().W), int(sprite.GetScale().H))))
	animationSet.AddTexture(t.simra.NewTextTexture("t", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.GetScale().W), int(sprite.GetScale().H))))
	animationSet.AddTexture(t.simra.NewTextTexture("i", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.GetScale().W), int(sprite.GetScale().H))))
	animationSet.AddTexture(t.simra.NewTextTexture("o", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.GetScale().W), int(sprite.GetScale().H))))
	animationSet.AddTexture(t.simra.NewTextTexture("n", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.GetScale().W), int(sprite.GetScale().H))))
	animationSet.AddTexture(t.simra.NewTextTexture("t", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.GetScale().W), int(sprite.GetScale().H))))
	animationSet.AddTexture(t.simra.NewTextTexture("e", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.GetScale().W), int(sprite.GetScale().H))))
	animationSet.AddTexture(t.simra.NewTextTexture("s", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.GetScale().W), int(sprite.GetScale().H))))
	animationSet.AddTexture(t.simra.NewTextTexture("t", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.GetScale().W), int(sprite.GetScale().H))))

	animationSet.SetInterval(12) // frames

	sprite.AddAnimationSet("animation test", animationSet)
	t.simra.AddSprite(sprite)
	tex := t.simra.NewTextTexture("animation test",
		60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.GetScale().W), int(sprite.GetScale().H)))
	sprite.ReplaceTexture(tex)

	t.simra.AddTouchListener(t)
	t.text = sprite
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (t *Title) Drive() {
}

// OnTouchBegin is called when Title scene is Touched.
func (t *Title) OnTouchBegin(x, y float32) {
}

// OnTouchMove is called when Title scene is Touched and moved.
func (t *Title) OnTouchMove(x, y float32) {
}

// OnTouchEnd is called when Title scene is Touched and it is released.
func (t *Title) OnTouchEnd(x, y float32) {
	if t.isAnimating {
		t.text.StopAnimation()
		t.isAnimating = false
	} else {
		t.text.StartAnimation("animation test", true, func() {})
		t.isAnimating = true
	}
}
