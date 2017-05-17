package scene

import (
	"image"
	"image/color"
	"time"

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
	text        *simra.Sprite
	isAnimating bool
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
	sprite := simra.NewSprite()
	sprite.W = ScreenWidth
	sprite.H = 80
	sprite.X = ScreenWidth / 2
	sprite.Y = ScreenHeight / 2

	animationSet := simra.NewAnimationSet()
	animationSet.AddTexture(simra.NewTextTexture("a", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.W), int(sprite.H))))
	animationSet.AddTexture(simra.NewTextTexture("n", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.W), int(sprite.H))))
	animationSet.AddTexture(simra.NewTextTexture("i", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.W), int(sprite.H))))
	animationSet.AddTexture(simra.NewTextTexture("m", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.W), int(sprite.H))))
	animationSet.AddTexture(simra.NewTextTexture("a", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.W), int(sprite.H))))
	animationSet.AddTexture(simra.NewTextTexture("t", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.W), int(sprite.H))))
	animationSet.AddTexture(simra.NewTextTexture("i", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.W), int(sprite.H))))
	animationSet.AddTexture(simra.NewTextTexture("o", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.W), int(sprite.H))))
	animationSet.AddTexture(simra.NewTextTexture("n", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.W), int(sprite.H))))
	animationSet.AddTexture(simra.NewTextTexture("t", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.W), int(sprite.H))))
	animationSet.AddTexture(simra.NewTextTexture("e", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.W), int(sprite.H))))
	animationSet.AddTexture(simra.NewTextTexture("s", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.W), int(sprite.H))))
	animationSet.AddTexture(simra.NewTextTexture("t", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(sprite.W), int(sprite.H))))

	animationSet.SetInterval(200 * time.Millisecond)

	sprite.AddAnimationSet("animation test", animationSet)
	simra.GetInstance().AddTextSprite("animation test",
		60, // fontsize
		color.RGBA{255, 0, 0, 255},
		image.Rect(0, 0, int(sprite.W), int(sprite.H)),
		sprite)
	simra.GetInstance().AddTouchListener(title)
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
		title.text.StartAnimation("animation test")
		title.isAnimating = true
	}
}
