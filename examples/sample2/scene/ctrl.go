package scene

import (
	"math"

	"github.com/pankona/gomo-simra/examples/sample2/scene/config"
	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/image"
)

// sample represents a scene object for this sample
type sample struct {
	simra    simra.Simraer
	ball     simra.Spriter
	ctrlup   simra.Spriter
	ctrldown simra.Spriter
	// buttonState represents which ctrl is pressed (or no ctrl pressed)
	buttonState int

	buttonRed      simra.Spriter
	buttonBlue     simra.Spriter
	buttonReplaced bool
}

const (
	ctrlNop = iota
	ctrlUp
	ctrlDown
)

// Initialize initializes sample scene
// This is called from simra.
// simra.SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (s *sample) Initialize(sim simra.Simraer) {
	s.simra = sim

	s.simra.SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// add global touch listener to catch touch end event
	s.simra.AddTouchListener(s)

	// initialize sprites
	s.initSprites()
	s.buttonReplaced = false
}

// OnTouchBegin is called when CtrlTrial scene is Touched.
func (s *sample) OnTouchBegin(x, y float32) {
	// nop
}

// OnTouchMove is called when CtrlTrial scene is Touched and moved.
func (s *sample) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when CtrlTrial scene is Touched and it is released.
func (s *sample) OnTouchEnd(x, y float32) {
	s.buttonState = ctrlNop
}

func (s *sample) initSprites() {
	s.initBall()
	s.initctrlDown()
	s.initctrlUp()
	s.initButtonBlue()
	s.initButtonRed()
}

func (s *sample) initBall() {
	s.ball = s.simra.NewSprite()
	// set size of ball
	s.ball.SetScale(48, 48)
	// put center of screen at start
	s.ball.SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)

	s.simra.AddSprite(s.ball)
	tex := s.simra.NewImageTexture("ball.png",
		image.Rect(0, 0, s.ball.GetScale().W, s.ball.GetScale().H))
	s.ball.ReplaceTexture(tex)

}

const (
	ctrlMarginLeft      = 10
	ctrlMarginBottom    = 10
	ctrlMarginBetween   = 10
	buttonMarginRight   = 20
	buttonMarginBottom  = 20
	buttonMarginBetween = 10
)

func (s *sample) initctrlUp() {
	s.ctrlup = s.simra.NewSprite()
	// set size of ctrlUp
	s.ctrlup.SetScale(120, 120)
	// put ctrlUp on left bottom
	s.ctrlup.SetPosition(
		s.ctrlup.GetScale().W/2+ctrlMarginLeft,
		ctrlMarginBottom+s.ctrldown.GetScale().H+ctrlMarginBetween+s.ctrlup.GetScale().H/2)

	// add sprite to glpeer
	s.simra.AddSprite(s.ctrlup)
	tex := s.simra.NewImageTexture("arrow.png",
		image.Rect(0, 0, s.ctrlup.GetScale().W, s.ctrlup.GetScale().H))
	s.ctrlup.ReplaceTexture(tex)

	// add touch listener for sprite
	ctrlup := &ctrlUpTouchListener{}
	s.ctrlup.AddTouchListener(ctrlup)
	ctrlup.sample = s
}

func (s *sample) initctrlDown() {
	s.ctrldown = s.simra.NewSprite()
	// set size of ctrlDown
	s.ctrldown.SetScale(120, 120)
	// put ctrlDown on left bottom
	s.ctrldown.SetPosition(
		s.ctrldown.GetScale().W/2+ctrlMarginLeft,
		ctrlMarginBottom+s.ctrldown.GetScale().H/2)

	// rotate arrow to indicate down control
	s.ctrldown.SetRotate(math.Pi)

	// add sprite to glpeer
	s.simra.AddSprite(s.ctrldown)
	tex := s.simra.NewImageTexture("arrow.png",
		image.Rect(0, 0, s.ctrldown.GetScale().W, s.ctrldown.GetScale().H))
	s.ctrldown.ReplaceTexture(tex)

	// add touch listener for sprite
	ctrldown := &ctrlDownTouchListener{}
	s.ctrldown.AddTouchListener(ctrldown)
	ctrldown.sample = s
}

func (s *sample) replaceButtonColor() {
	var tex *simra.Texture

	// red changes to blue
	tex = s.simra.NewImageTexture("blue_circle.png",
		image.Rect(0, 0, s.buttonBlue.GetScale().W, s.buttonBlue.GetScale().H))
	s.buttonRed.ReplaceTexture(tex)

	// blue changes to red
	tex = s.simra.NewImageTexture("red_circle.png",
		image.Rect(0, 0, s.buttonRed.GetScale().W, s.buttonRed.GetScale().H))
	s.buttonBlue.ReplaceTexture(tex)

	s.buttonReplaced = true
}

func (s *sample) originalButtonColor() {
	var tex *simra.Texture

	// set red button to buttonRed
	tex = s.simra.NewImageTexture("red_circle.png",
		image.Rect(0, 0, s.buttonBlue.GetScale().W, s.buttonBlue.GetScale().H))
	s.buttonRed.ReplaceTexture(tex)

	// set blue button to buttonBlue
	tex = s.simra.NewImageTexture("blue_circle.png",
		image.Rect(0, 0, s.buttonRed.GetScale().W, s.buttonRed.GetScale().H))
	s.buttonBlue.ReplaceTexture(tex)

	s.buttonReplaced = false
}

func (s *sample) initButtonBlue() {
	s.buttonBlue = s.simra.NewSprite()
	// set size of button blue
	s.buttonBlue.SetScale(80, 80)

	// put button red on right bottom
	s.buttonBlue.SetPosition(
		config.ScreenWidth-buttonMarginRight-s.buttonBlue.GetScale().W/2,
		buttonMarginBottom+(80)+buttonMarginBetween+s.buttonBlue.GetScale().W/2)

	// add sprite to glpeer
	s.simra.AddSprite(s.buttonBlue)
	tex := s.simra.NewImageTexture("blue_circle.png",
		image.Rect(0, 0, s.buttonBlue.GetScale().W, s.buttonBlue.GetScale().H))
	s.buttonBlue.ReplaceTexture(tex)

	// add touch listener for sprite
	listener := &ButtonBlueTouchListener{}
	s.buttonBlue.AddTouchListener(listener)
	listener.sample = s
}

func (s *sample) initButtonRed() {
	s.buttonRed = s.simra.NewSprite()
	// set size of button red
	s.buttonRed.SetScale(80, 80)
	// put button red on right bottom
	s.buttonRed.SetPosition(
		config.ScreenWidth-buttonMarginRight-s.buttonBlue.GetScale().W-buttonMarginBetween-s.buttonRed.GetScale().W/2,
		buttonMarginBottom+(s.buttonRed.GetScale().H/2))

	// add sprite to glpeer
	s.simra.AddSprite(s.buttonRed)
	tex := s.simra.NewImageTexture("red_circle.png",
		image.Rect(0, 0, s.buttonRed.GetScale().W, s.buttonRed.GetScale().H))
	s.buttonRed.ReplaceTexture(tex)

	// add touch listener for sprite
	listener := &ButtonRedTouchListener{}
	s.buttonRed.AddTouchListener(listener)
	listener.sample = s
}

var degree float32

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (s *sample) Drive() {
	degree++
	if degree >= 360 {
		degree = 0
	}

	p := s.ball.GetPosition()
	switch s.buttonState {
	case ctrlUp:
		s.ball.SetPositionY(p.Y + 1)
	case ctrlDown:
		s.ball.SetPositionY(p.Y - 1)
	}

	s.ball.SetRotate(float32(degree) * math.Pi / 180)
}
