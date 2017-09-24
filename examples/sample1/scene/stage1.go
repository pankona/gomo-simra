package scene

import (
	"image"
	"math"

	"github.com/pankona/gomo-simra/simra"
)

// Stage1 represents a scene of Stage1
type Stage1 struct {
	simra  simra.Simraer
	gopher simra.Spriter
}

// Initialize initializes Stage1 scene.
// This is called from simra.
// simra.SetDesiredScreenSize should be called to determine
// screen size of this scene.
// If SetDesiredScreenSize is already called in previous scene, this scene may not call the function.
func (stage1 *Stage1) Initialize(sim simra.Simraer) {
	stage1.simra = sim
	stage1.simra.SetDesiredScreenSize(1080/2, 1920/2)
	stage1.initSprite()
}

func (stage1 *Stage1) initSprite() {
	stage1.initGopher()
}

func (stage1 *Stage1) initGopher() {
	stage1.gopher = stage1.simra.NewSprite()
	// add gopher sprite
	stage1.gopher.SetScale(140, 90)

	// put center of screen at start
	stage1.gopher.SetPosition(1080/2/2, 1920/2/2)

	stage1.simra.AddSprite(stage1.gopher)
	tex := stage1.simra.NewImageTexture("waza-gophers.jpeg",
		image.Rect(152, 10, 152+int(stage1.gopher.GetScale().W), 10+int(stage1.gopher.GetScale().H)))
	stage1.gopher.ReplaceTexture(tex)

	stage1.gopher.AddTouchListener(stage1)
}

var degree float32

// Drive is called from simra.
// This is used to update sprites position.
// Thsi will be called 60 times per sec.
func (stage1 *Stage1) Drive() {
	degree++
	if degree >= 360 {
		degree = 0
	}
	stage1.gopher.SetRotate(float32(degree) * math.Pi / 180)
}

// OnTouchBegin is called when Title scene is Touched.
// It is caused by calling AddtouchListener for stage1.gopher sprite.
func (stage1 *Stage1) OnTouchBegin(x, y float32) {
	stage1.gopher.SetPosition((int)(x), (int)(y))
}

// OnTouchMove is called when Title scene is Touched and moved.
// It is caused by calling AddtouchListener for stage1.gopher sprite.
func (stage1 *Stage1) OnTouchMove(x, y float32) {
	stage1.gopher.SetPosition((int)(x), (int)(y))
}

// OnTouchEnd is called when Title scene is Touched and it is released.
// It is caused by calling AddtouchListener for stage1.gopher sprite.
func (stage1 *Stage1) OnTouchEnd(x, y float32) {
	stage1.gopher.SetPosition((int)(x), (int)(y))
}
