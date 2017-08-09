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
	effect        *simra.Sprite
	initialSprite *simra.Texture
	isAnimating   bool
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
	sprite.W = 240
	sprite.H = 240
	sprite.X = ScreenWidth / 2
	sprite.Y = ScreenHeight / 2

	animationSet := simra.NewAnimationSet()
	title.initialSprite = simra.NewImageTexture("effect.png", image.Rect(0, 0, 239, int(sprite.H)))
	for i := 0; i < 13; i++ {
		animationSet.AddTexture(simra.NewImageTexture("effect.png",
			image.Rect((int)(sprite.W)*i, 0, ((int)(sprite.W)*(i+1))-1, int(sprite.H))))
	}
	animationSet.SetInterval(6)
	sprite.AddAnimationSet("animation test", animationSet)

	simra.GetInstance().AddSprite2(sprite)
	sprite.ReplaceTexture2(title.initialSprite)
	simra.GetInstance().AddTouchListener(title)
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
			title.effect.ReplaceTexture2(title.initialSprite)
		})
		title.isAnimating = true
	}
}
