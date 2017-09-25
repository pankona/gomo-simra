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
func (t *Title) Initialize(sim simra.Simraer) {
	t.simra = sim
	t.simra.SetDesiredScreenSize(ScreenWidth, ScreenHeight)
	t.initialize()
}

func (t *Title) initialize() {
	sprite := t.simra.NewSprite()
	sprite.SetPosition(ScreenWidth/2, ScreenHeight/2)
	sprite.SetScale(240, 240)

	animationSet := simra.NewAnimationSet()
	t.initialSprite = t.simra.NewImageTexture("effect.png", image.Rect(0, 0, 239, sprite.GetScale().H))
	for i := 0; i < 13; i++ {
		animationSet.AddTexture(t.simra.NewImageTexture("effect.png",
			image.Rect(sprite.GetScale().W*i, 0, (sprite.GetScale().W*(i+1))-1, sprite.GetScale().H)))
	}
	animationSet.SetInterval(6)
	sprite.AddAnimationSet("animation test", animationSet)

	t.simra.AddSprite(sprite)
	sprite.ReplaceTexture(t.initialSprite)
	t.simra.AddTouchListener(t)
	t.effect = sprite
}

// Drive is called from simra.
// This is used to update sprites position.
// Thsi will be called 60 times per sec.
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
		t.effect.StopAnimation()
		t.isAnimating = false
	} else {
		shouldLoop := true
		t.effect.StartAnimation("animation test", shouldLoop, func() {
			t.effect.ReplaceTexture(t.initialSprite)
		})
		t.isAnimating = true
	}
}
