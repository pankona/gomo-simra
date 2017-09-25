package scene

import (
	"image"
	"image/color"

	"golang.org/x/mobile/asset"

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
	simra     simra.Simraer
	audio     simra.Audioer
	isPlaying bool
}

// Initialize initializes title scene
// This is called from simra.
// simra.SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (t *Title) Initialize(sim simra.Simraer) {
	t.simra = sim
	t.simra.SetDesiredScreenSize(ScreenWidth, ScreenHeight)
	// initialize sprites
	t.initialize()
}

func (t *Title) initialize() {

	t.simra.AddTouchListener(t)

	sprite := t.simra.NewSprite()
	sprite.SetScale(ScreenWidth, 80)
	sprite.SetPosition(ScreenWidth/2, ScreenHeight/2)
	t.simra.AddSprite(sprite)
	s := sprite.GetScale()
	tex := t.simra.NewTextTexture("tap to play sound",
		60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, s.W, s.H))
	sprite.ReplaceTexture(tex)

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
	if t.isPlaying {
		err := t.audio.Stop()
		if err != nil {
			panic(err.Error())
		}
		t.isPlaying = false

	} else {
		t.audio = simra.NewAudio()
		resource, err := asset.Open("test_se.mp3")
		if err != nil {
			panic(err.Error())
		}

		err = t.audio.Play(resource, true, func(err error) {
			t.isPlaying = false
		})
		if err != nil {
			panic(err)
		}
		t.isPlaying = true
	}
}
