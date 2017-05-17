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
	temporary := simra.NewSprite()
	temporary.W = ScreenWidth
	temporary.H = 80
	temporary.X = ScreenWidth / 2
	temporary.Y = ScreenHeight / 2

	animationSet := simra.NewAnimationSet()
	animationSet.AddTexture(simra.NewTextTexture("a", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(temporary.W), int(temporary.H))))
	animationSet.AddTexture(simra.NewTextTexture("n", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(temporary.W), int(temporary.H))))
	animationSet.AddTexture(simra.NewTextTexture("i", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(temporary.W), int(temporary.H))))
	animationSet.AddTexture(simra.NewTextTexture("m", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(temporary.W), int(temporary.H))))
	animationSet.AddTexture(simra.NewTextTexture("a", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(temporary.W), int(temporary.H))))
	animationSet.AddTexture(simra.NewTextTexture("t", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(temporary.W), int(temporary.H))))
	animationSet.AddTexture(simra.NewTextTexture("i", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(temporary.W), int(temporary.H))))
	animationSet.AddTexture(simra.NewTextTexture("o", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(temporary.W), int(temporary.H))))
	animationSet.AddTexture(simra.NewTextTexture("n", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(temporary.W), int(temporary.H))))
	animationSet.AddTexture(simra.NewTextTexture("t", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(temporary.W), int(temporary.H))))
	animationSet.AddTexture(simra.NewTextTexture("e", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(temporary.W), int(temporary.H))))
	animationSet.AddTexture(simra.NewTextTexture("s", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(temporary.W), int(temporary.H))))
	animationSet.AddTexture(simra.NewTextTexture("t", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(temporary.W), int(temporary.H))))
	animationSet.SetInterval(1 * time.Second)

	temporary.AddAnimationSet("animation test", animationSet)
	simra.GetInstance().AddTextSprite("animation test",
		60, // fontsize
		color.RGBA{255, 0, 0, 255},
		image.Rect(0, 0, int(temporary.W), int(temporary.H)),
		temporary)
	err := temporary.StartAnimation("animation test")
	if err != nil {
		panic("failed to start animation")
	}
}

// Drive is called from simra.
// This is used to update sprites position.
// Thsi will be called 60 times per sec.
func (title *Title) Drive() {
}
