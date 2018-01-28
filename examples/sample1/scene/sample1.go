package scene

import (
	"math"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/image"
)

// sample represents a scene of sample
type sample struct {
	simra  simra.Simraer
	gopher simra.Spriter
}

// Initialize initializes sample scene.
// This is called from simra.
// simra.SetDesiredScreenSize should be called to determine
// screen size of this scene.
// If SetDesiredScreenSize is already called in previous scene, this scene may not call the function.
func (s *sample) Initialize(sim simra.Simraer) {
	s.simra = sim
	s.simra.SetDesiredScreenSize(1080/2, 1920/2)
	s.initSprite()
}

func (s *sample) initSprite() {
	s.initGopher()
}

func (s *sample) initGopher() {
	s.gopher = s.simra.NewSprite()
	// add gopher sprite
	s.gopher.SetScale(140, 90)

	// put center of screen at start
	s.gopher.SetPosition(1080/2/2, 1920/2/2)

	s.simra.AddSprite(s.gopher)
	tex := s.simra.NewImageTexture("waza-gophers.jpeg",
		image.Rect(152, 10, 152+s.gopher.GetScale().W, 10+s.gopher.GetScale().H))
	s.gopher.ReplaceTexture(tex)

	s.gopher.AddTouchListener(s)
}

var degree float32

// Drive is called from simra.
// This is used to update sprites position.
// This function will be called 60 times per sec.
func (s *sample) Drive() {
	degree++
	if degree >= 360 {
		degree = 0
	}
	s.gopher.SetRotate(degree * math.Pi / 180)
}

// OnTouchBegin is called when Title scene is Touched.
// It is caused by calling AddtouchListener for e.gopher sprite.
func (s *sample) OnTouchBegin(x, y float32) {
	s.gopher.SetPosition(x, y)
}

// OnTouchMove is called when Title scene is Touched and moved.
// It is caused by calling AddtouchListener for e.gopher sprite.
func (s *sample) OnTouchMove(x, y float32) {
	s.gopher.SetPosition(x, y)
}

// OnTouchEnd is called when Title scene is Touched and it is released.
// It is caused by calling AddtouchListener for e.gopher sprite.
func (s *sample) OnTouchEnd(x, y float32) {
	s.gopher.SetPosition(x, y)
}
