package scene

import (
	"image"
	"math"

	"github.com/pankona/gomo-simra/examples/sample2/scene/config"
	"github.com/pankona/gomo-simra/simra"
)

// CtrlTrial represents a scene object for CtrlTrial
type CtrlTrial struct {
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

// Initialize initializes CtrlTrial scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (ctrltrial *CtrlTrial) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// add global touch listener to catch touch end event
	simra.GetInstance().AddTouchListener(ctrltrial)

	// TODO: when goes to next scene, remove global touch listener
	// simra.GetInstance().RemoveTouchListener(ctrltrial)

	// initialize sprites
	ctrltrial.initSprites()
	ctrltrial.buttonReplaced = false

	simra.LogDebug("[OUT]")
}

// OnTouchBegin is called when CtrlTrial scene is Touched.
func (ctrltrial *CtrlTrial) OnTouchBegin(x, y float32) {
	// nop
}

// OnTouchMove is called when CtrlTrial scene is Touched and moved.
func (ctrltrial *CtrlTrial) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when CtrlTrial scene is Touched and it is released.
func (ctrltrial *CtrlTrial) OnTouchEnd(x, y float32) {
	ctrltrial.buttonState = ctrlNop
}

func (ctrltrial *CtrlTrial) initSprites() {
	ctrltrial.initBall()
	ctrltrial.initctrlDown()
	ctrltrial.initctrlUp()
	ctrltrial.initButtonBlue()
	ctrltrial.initButtonRed()
}

func (ctrltrial *CtrlTrial) initBall() {
	// set size of ball
	ctrltrial.ball.SetScale(48, 48)

	// put center of screen at start
	ctrltrial.ball.SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)

	simra.GetInstance().AddSprite(ctrltrial.ball)
	tex := simra.NewImageTexture("ball.png",
		image.Rect(0, 0, int(ctrltrial.ball.GetScale().W), int(ctrltrial.ball.GetScale().H)))
	ctrltrial.ball.ReplaceTexture(tex)

}

const (
	ctrlMarginLeft      = 10
	ctrlMarginBottom    = 10
	ctrlMarginBetween   = 10
	buttonMarginRight   = 20
	buttonMarginBottom  = 20
	buttonMarginBetween = 10
)

// ctrlUp
type ctrlUpTouchListener struct {
	parent *CtrlTrial
}

func (ctrltrial *ctrlUpTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("[IN] ctrlUp Begin!")

	ctrl := ctrltrial.parent
	ctrl.buttonState = ctrlUp

	simra.LogDebug("[OUT]")
}

func (ctrltrial *ctrlUpTouchListener) OnTouchMove(x, y float32) {
	simra.LogDebug("[IN] ctrlUp Move!")

	ctrl := ctrltrial.parent
	ctrl.buttonState = ctrlUp

	simra.LogDebug("[OUT]")
}

func (ctrltrial *ctrlUpTouchListener) OnTouchEnd(x, y float32) {
	simra.LogDebug("[IN] ctrlUp End")

	ctrl := ctrltrial.parent
	ctrl.buttonState = ctrlNop

	simra.LogDebug("[OUT]")
}

func (ctrltrial *CtrlTrial) initctrlUp() {
	// set size of ctrlUp
	ctrltrial.ctrlup.SetScale(120, 120)

	// put ctrlUp on left bottom
	ctrltrial.ctrlup.SetPosition(
		ctrltrial.ctrlup.GetScale().W/2+ctrlMarginLeft,
		ctrlMarginBottom+ctrltrial.ctrldown.GetScale().H+ctrlMarginBetween+ctrltrial.ctrlup.GetScale().H/2)

	// add sprite to glpeer
	simra.GetInstance().AddSprite(ctrltrial.ctrlup)
	tex := simra.NewImageTexture("arrow.png",
		image.Rect(0, 0, int(ctrltrial.ctrlup.GetScale().W), int(ctrltrial.ctrlup.GetScale().H)))
	ctrltrial.ctrlup.ReplaceTexture(tex)

	// add touch listener for sprite
	ctrlup := &ctrlUpTouchListener{}
	ctrltrial.ctrlup.AddTouchListener(ctrlup)
	ctrlup.parent = ctrltrial
}

// ctrlDown
type ctrlDownTouchListener struct {
	parent *CtrlTrial
}

func (ctrltrial *ctrlDownTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("[IN] ctrlDown Begin!")

	ctrl := ctrltrial.parent
	ctrl.buttonState = ctrlDown

	simra.LogDebug("[OUT]")
}

func (ctrltrial *ctrlDownTouchListener) OnTouchMove(x, y float32) {
	simra.LogDebug("[IN] ctrlDown Move!")

	ctrl := ctrltrial.parent
	ctrl.buttonState = ctrlDown

	simra.LogDebug("[OUT]")
}

func (ctrltrial *ctrlDownTouchListener) OnTouchEnd(x, y float32) {
	simra.LogDebug("[IN] ctrlDown End")

	ctrl := ctrltrial.parent
	ctrl.buttonState = ctrlNop

	simra.LogDebug("[OUT]")
}

func (ctrltrial *CtrlTrial) initctrlDown() {
	// set size of ctrlDown
	ctrltrial.ctrldown.SetScale(120, 120)

	// put ctrlDown on left bottom
	ctrltrial.ctrldown.SetPosition(
		ctrltrial.ctrldown.GetScale().W/2+ctrlMarginLeft,
		ctrlMarginBottom+ctrltrial.ctrldown.GetScale().H/2)

	// rotate arrow to indicate down control
	ctrltrial.ctrldown.SetRotate(math.Pi)

	// add sprite to glpeer
	simra.GetInstance().AddSprite(ctrltrial.ctrldown)
	tex := simra.NewImageTexture("arrow.png",
		image.Rect(0, 0, int(ctrltrial.ctrldown.GetScale().W), int(ctrltrial.ctrldown.GetScale().H)))
	ctrltrial.ctrldown.ReplaceTexture(tex)

	// add touch listener for sprite
	ctrldown := &ctrlDownTouchListener{}
	ctrltrial.ctrldown.AddTouchListener(ctrldown)
	ctrldown.parent = ctrltrial
}

func (ctrltrial *CtrlTrial) replaceButtonColor() {
	simra.LogDebug("IN")
	var tex *simra.Texture

	// red changes to blue
	tex = simra.NewImageTexture("blue_circle.png",
		image.Rect(0, 0, int(ctrltrial.buttonBlue.GetScale().W), int(ctrltrial.buttonBlue.GetScale().H)))
	ctrltrial.buttonRed.ReplaceTexture(tex)

	// blue changes to red
	tex = simra.NewImageTexture("red_circle.png",
		image.Rect(0, 0, int(ctrltrial.buttonRed.GetScale().W), int(ctrltrial.buttonRed.GetScale().H)))
	ctrltrial.buttonBlue.ReplaceTexture(tex)

	ctrltrial.buttonReplaced = true
	simra.LogDebug("OUT")
}

func (ctrltrial *CtrlTrial) originalButtonColor() {
	simra.LogDebug("IN")
	var tex *simra.Texture

	// set red button to buttonRed
	tex = simra.NewImageTexture("red_circle.png",
		image.Rect(0, 0, int(ctrltrial.buttonBlue.GetScale().W), int(ctrltrial.buttonBlue.GetScale().H)))
	ctrltrial.buttonRed.ReplaceTexture(tex)

	// set blue button to buttonBlue
	tex = simra.NewImageTexture("blue_circle.png",
		image.Rect(0, 0, int(ctrltrial.buttonRed.GetScale().W), int(ctrltrial.buttonRed.GetScale().H)))
	ctrltrial.buttonBlue.ReplaceTexture(tex)

	ctrltrial.buttonReplaced = false
	simra.LogDebug("OUT")
}

// ButtonBlueTouchListener represents a listener object
// to notify touch event of Blue Button
type ButtonBlueTouchListener struct {
	parent *CtrlTrial
}

// OnTouchBegin is called when Blue Button is Touched.
func (ctrltrial *ButtonBlueTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("IN")
	if ctrltrial.parent.buttonReplaced {
		ctrltrial.parent.originalButtonColor()
	} else {
		ctrltrial.parent.replaceButtonColor()
	}

	simra.GetInstance().RemoveSprite(ctrltrial.parent.ball)
	simra.LogDebug("OUT")
}

// OnTouchMove is called when Blue Button is Touched and moved.
func (ctrltrial *ButtonBlueTouchListener) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when Blue Button is Touched and it is released.
func (ctrltrial *ButtonBlueTouchListener) OnTouchEnd(x, y float32) {
	// nop
}

func (ctrltrial *CtrlTrial) initButtonBlue() {
	simra.LogDebug("IN")
	// set size of button blue
	ctrltrial.buttonBlue.SetScale(80, 80)

	// put button red on right bottom
	ctrltrial.buttonBlue.SetPosition(
		config.ScreenWidth-buttonMarginRight-ctrltrial.buttonBlue.GetScale().W/2,
		buttonMarginBottom+(80)+buttonMarginBetween+ctrltrial.buttonBlue.GetScale().W/2)

	// add sprite to glpeer
	simra.GetInstance().AddSprite(ctrltrial.buttonBlue)
	tex := simra.NewImageTexture("blue_circle.png",
		image.Rect(0, 0, int(ctrltrial.buttonBlue.GetScale().W), int(ctrltrial.buttonBlue.GetScale().H)))
	ctrltrial.buttonBlue.ReplaceTexture(tex)

	// add touch listener for sprite
	listener := &ButtonBlueTouchListener{}
	ctrltrial.buttonBlue.AddTouchListener(listener)
	listener.parent = ctrltrial
	simra.LogDebug("OUT")
}

// ButtonRedTouchListener represents a listener object
// to notify touch event of Red Button
type ButtonRedTouchListener struct {
	parent *CtrlTrial
}

// OnTouchBegin is called when Red Button is Touched.
func (ctrltrial *ButtonRedTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("IN")
	if ctrltrial.parent.buttonReplaced {
		ctrltrial.parent.originalButtonColor()
	} else {
		ctrltrial.parent.replaceButtonColor()
	}
	simra.GetInstance().AddSprite(ctrltrial.parent.ball)
	tex := simra.NewImageTexture("ball.png",
		image.Rect(0, 0, int(ctrltrial.parent.ball.GetScale().W), int(ctrltrial.parent.ball.GetScale().H)))
	ctrltrial.parent.ball.ReplaceTexture(tex)
	simra.LogDebug("OUT")
}

// OnTouchMove is called when Red Button is Touched and moved.
func (ctrltrial *ButtonRedTouchListener) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when Red Button is Touched and it is released.
func (ctrltrial *ButtonRedTouchListener) OnTouchEnd(x, y float32) {
	// nop
}

func (ctrltrial *CtrlTrial) initButtonRed() {
	// set size of button red
	ctrltrial.buttonRed.SetScale(80, 80)

	// put button red on right bottom
	ctrltrial.buttonRed.SetPosition(
		config.ScreenWidth-buttonMarginRight-ctrltrial.buttonBlue.GetScale().W-buttonMarginBetween-ctrltrial.buttonRed.GetScale().W/2,
		buttonMarginBottom+(ctrltrial.buttonRed.GetScale().H/2))

	// add sprite to glpeer
	simra.GetInstance().AddSprite(ctrltrial.buttonRed)
	tex := simra.NewImageTexture("red_circle.png",
		image.Rect(0, 0, int(ctrltrial.buttonRed.GetScale().W), int(ctrltrial.buttonRed.GetScale().H)))
	ctrltrial.buttonRed.ReplaceTexture(tex)

	// add touch listener for sprite
	listener := &ButtonRedTouchListener{}
	ctrltrial.buttonRed.AddTouchListener(listener)
	listener.parent = ctrltrial
}

var degree float32

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (ctrltrial *CtrlTrial) Drive() {
	degree++
	if degree >= 360 {
		degree = 0
	}

	p := ctrltrial.ball.GetPosition()
	switch ctrltrial.buttonState {
	case ctrlUp:
		ctrltrial.ball.SetPositionY(p.Y + 1)
	case ctrlDown:
		ctrltrial.ball.SetPositionY(p.Y - 1)
	}

	ctrltrial.ball.SetRotate(float32(degree) * math.Pi / 180)
}
