package scene

import (
	"image"

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
	simra         simra.Simraer
	effect        simra.Spriter
	initialSprite *simra.Texture
	isAnimating   bool
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
	sprite.SetPosition(ScreenWidth/2, ScreenHeight/2)
	sprite.SetScale(240, 240)

	animationSet := simra.NewAnimationSet()
	title.initialSprite = title.simra.NewImageTexture("effect.png", image.Rect(0, 0, 239, sprite.GetScale().H))
	for i := 0; i < 13; i++ {
		animationSet.AddTexture(title.simra.NewImageTexture("effect.png",
			image.Rect(sprite.GetScale().W*i, 0, (sprite.GetScale().W*(i+1))-1, sprite.GetScale().H)))
	}
	animationSet.SetInterval(6)
	sprite.AddAnimationSet("animation test", animationSet)

	title.simra.AddSprite(sprite)
	sprite.ReplaceTexture(title.initialSprite)
	title.simra.AddTouchListener(title)
	title.effect = sprite
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
		title.effect.StopAnimation()
		title.isAnimating = false
	} else {
		simra.LogDebug("start animation")
		shouldLoop := true
		title.effect.StartAnimation("animation test", shouldLoop, func() {
			simra.LogDebug("animation end")
			title.effect.ReplaceTexture(title.initialSprite)
		})
		title.isAnimating = true
	}
}
