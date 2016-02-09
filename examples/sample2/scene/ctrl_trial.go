package scene

import (
	"image"
	"math"

	"github.com/pankona/gomo-simra/examples/sample2/scene/config"
	"github.com/pankona/gomo-simra/simra"
)

type CtrlTrial struct {
	ball     simra.Sprite
	ctrlup   simra.Sprite
	ctrldown simra.Sprite
	// buttonState represents which ctrl is pressed (or no ctrl pressed)
	buttonState int

	buttonRed      simra.Sprite
	buttonBlue     simra.Sprite
	buttonReplaced bool
}

const (
	CTRL_NOP = iota
	CTRL_UP
	CTRL_DOWN
)

func (self *CtrlTrial) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.SCREEN_WIDTH, config.SCREEN_HEIGHT)

	// add global touch listener to catch touch end event
	simra.GetInstance().AddTouchListener(self)

	// TODO: when goes to next scene, remove global touch listener
	// simra.GetInstance().RemoveTouchListener(self)

	// initialize sprites
	self.initSprites()
	self.buttonReplaced = false

	simra.LogDebug("[OUT]")
}

func (self *CtrlTrial) OnTouchBegin(x, y float32) {
	// nop
}

func (self *CtrlTrial) OnTouchMove(x, y float32) {
	// nop
}

func (self *CtrlTrial) OnTouchEnd(x, y float32) {
	// nop
	self.buttonState = CTRL_NOP
}

func (self *CtrlTrial) initSprites() {
	self.initBall()
	self.initCtrlDown()
	self.initCtrlUp()
	self.initButtonBlue()
	self.initButtonRed()
}

func (self *CtrlTrial) initBall() {
	// set size of ball
	self.ball.W = float32(48)
	self.ball.H = float32(48)

	// put center of screen at start
	self.ball.X = config.SCREEN_WIDTH / 2
	self.ball.Y = config.SCREEN_HEIGHT / 2

	simra.GetInstance().AddSprite("ball.png",
		image.Rect(0, 0, int(self.ball.W), int(self.ball.H)),
		&self.ball)
}

const (
	CTRL_MARGIN_LEFT      = 10
	CTRL_MARGIN_BOTTOM    = 10
	CTRL_MARGIN_BETWEEN   = 10
	BUTTON_MARGIN_RIGHT   = 20
	BUTTON_MARGIN_BOTTOM  = 20
	BUTTON_MARGIN_BETWEEN = 10
)

// CtrlUp
type CtrlUpTouchListener struct {
	parent *CtrlTrial
}

func (self *CtrlUpTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("[IN] CtrlUp Begin!")

	ctrl := self.parent
	ctrl.buttonState = CTRL_UP

	simra.LogDebug("[OUT]")
}

func (self *CtrlUpTouchListener) OnTouchMove(x, y float32) {
	simra.LogDebug("[IN] CtrlUp Move!")

	ctrl := self.parent
	ctrl.buttonState = CTRL_UP

	simra.LogDebug("[OUT]")
}

func (self *CtrlUpTouchListener) OnTouchEnd(x, y float32) {
	simra.LogDebug("[IN] CtrlUp End")

	ctrl := self.parent
	ctrl.buttonState = CTRL_NOP

	simra.LogDebug("[OUT]")
}

func (self *CtrlTrial) initCtrlUp() {
	// set size of CtrlUp
	self.ctrlup.W = float32(120)
	self.ctrlup.H = float32(120)

	// put CtrlUp on left bottom
	self.ctrlup.X = (self.ctrlup.W / 2) + CTRL_MARGIN_LEFT
	self.ctrlup.Y = CTRL_MARGIN_BOTTOM + self.ctrldown.H + CTRL_MARGIN_BETWEEN + (self.ctrlup.H / 2)

	// add sprite to glpeer
	simra.GetInstance().AddSprite("arrow.png",
		image.Rect(0, 0, int(self.ctrlup.W), int(self.ctrlup.H)),
		&self.ctrlup)

	// add touch listener for sprite
	ctrlup := &CtrlUpTouchListener{}
	self.ctrlup.AddTouchListener(ctrlup)
	ctrlup.parent = self
}

// CtrlDown
type CtrlDownTouchListener struct {
	parent *CtrlTrial
}

func (self *CtrlDownTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("[IN] CtrlDown Begin!")

	ctrl := self.parent
	ctrl.buttonState = CTRL_DOWN

	simra.LogDebug("[OUT]")
}

func (self *CtrlDownTouchListener) OnTouchMove(x, y float32) {
	simra.LogDebug("[IN] CtrlDown Move!")

	ctrl := self.parent
	ctrl.buttonState = CTRL_DOWN

	simra.LogDebug("[OUT]")
}

func (self *CtrlDownTouchListener) OnTouchEnd(x, y float32) {
	simra.LogDebug("[IN] CtrlDown End")

	ctrl := self.parent
	ctrl.buttonState = CTRL_NOP

	simra.LogDebug("[OUT]")
}

func (self *CtrlTrial) initCtrlDown() {
	// set size of CtrlDown
	self.ctrldown.W = float32(120)
	self.ctrldown.H = float32(120)

	// put CtrlDown on left bottom
	self.ctrldown.X = (self.ctrldown.W / 2) + CTRL_MARGIN_LEFT
	self.ctrldown.Y = CTRL_MARGIN_BOTTOM + (self.ctrldown.H / 2)

	// rotate arrow to indicate down control
	self.ctrldown.R = math.Pi

	// add sprite to glpeer
	simra.GetInstance().AddSprite("arrow.png",
		image.Rect(0, 0, int(self.ctrldown.W), int(self.ctrldown.H)),
		&self.ctrldown)

	// add touch listener for sprite
	ctrldown := &CtrlDownTouchListener{}
	self.ctrldown.AddTouchListener(ctrldown)
	ctrldown.parent = self
}

func (self *CtrlTrial) replaceButtonColor() {
	simra.LogDebug("IN")
	// red changes to blue
	self.buttonRed.ReplaceTexture("blue_circle.png",
		image.Rect(0, 0, int(self.buttonBlue.W), int(self.buttonBlue.H)))
	// blue changes to red
	self.buttonBlue.ReplaceTexture("red_circle.png",
		image.Rect(0, 0, int(self.buttonRed.W), int(self.buttonRed.H)))

	self.buttonReplaced = true
	simra.LogDebug("OUT")
}

func (self *CtrlTrial) originalButtonColor() {
	simra.LogDebug("IN")
	// set red button to buttonRed
	self.buttonRed.ReplaceTexture("red_circle.png",
		image.Rect(0, 0, int(self.buttonBlue.W), int(self.buttonBlue.H)))
	// set blue button to buttonBlue
	self.buttonBlue.ReplaceTexture("blue_circle.png",
		image.Rect(0, 0, int(self.buttonRed.W), int(self.buttonRed.H)))

	self.buttonReplaced = false
	simra.LogDebug("OUT")
}

// button blue
type ButtonBlueTouchListener struct {
	parent *CtrlTrial
}

func (self *ButtonBlueTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("IN")
	if self.parent.buttonReplaced {
		self.parent.originalButtonColor()
	} else {
		self.parent.replaceButtonColor()
	}

	simra.GetInstance().RemoveSprite(&self.parent.ball)
	simra.LogDebug("OUT")
}

func (self *ButtonBlueTouchListener) OnTouchMove(x, y float32) {
	// nop
}

func (self *ButtonBlueTouchListener) OnTouchEnd(x, y float32) {
	// nop
}

func (self *CtrlTrial) initButtonBlue() {
	simra.LogDebug("IN")
	// set size of button blue
	self.buttonBlue.W = float32(80)
	self.buttonBlue.H = float32(80)

	// put button red on right bottom
	self.buttonBlue.X = config.SCREEN_WIDTH - BUTTON_MARGIN_RIGHT - self.buttonBlue.W/2
	self.buttonBlue.Y = BUTTON_MARGIN_BOTTOM + (80) + BUTTON_MARGIN_BETWEEN + self.buttonBlue.W/2

	// add sprite to glpeer
	simra.GetInstance().AddSprite("blue_circle.png",
		image.Rect(0, 0, int(self.buttonBlue.W), int(self.buttonBlue.H)),
		&self.buttonBlue)

	// add touch listener for sprite
	listener := &ButtonBlueTouchListener{}
	self.buttonBlue.AddTouchListener(listener)
	listener.parent = self
	simra.LogDebug("OUT")
}

// button red
type ButtonRedTouchListener struct {
	parent *CtrlTrial
}

func (self *ButtonRedTouchListener) OnTouchBegin(x, y float32) {
	simra.LogDebug("IN")
	if self.parent.buttonReplaced {
		self.parent.originalButtonColor()
	} else {
		self.parent.replaceButtonColor()
	}
	simra.GetInstance().AddSprite("ball.png",
		image.Rect(0, 0, int(self.parent.ball.W), int(self.parent.ball.H)),
		&self.parent.ball)
	simra.LogDebug("OUT")
}
func (self *ButtonRedTouchListener) OnTouchMove(x, y float32) {
	// nop
}
func (self *ButtonRedTouchListener) OnTouchEnd(x, y float32) {
	// nop
}

func (self *CtrlTrial) initButtonRed() {
	// set size of button red
	self.buttonRed.W = float32(80)
	self.buttonRed.H = float32(80)

	// put button red on right bottom
	self.buttonRed.X = config.SCREEN_WIDTH - BUTTON_MARGIN_RIGHT - self.buttonBlue.W -
		BUTTON_MARGIN_BETWEEN - self.buttonRed.W/2
	self.buttonRed.Y = BUTTON_MARGIN_BOTTOM + (self.buttonRed.H / 2)

	// add sprite to glpeer
	simra.GetInstance().AddSprite("red_circle.png",
		image.Rect(0, 0, int(self.buttonRed.W), int(self.buttonRed.H)),
		&self.buttonRed)

	// add touch listener for sprite
	listener := &ButtonRedTouchListener{}
	self.buttonRed.AddTouchListener(listener)
	listener.parent = self
}

var degree float32 = 0

func (self *CtrlTrial) Drive() {
	degree += 1
	if degree >= 360 {
		degree = 0
	}

	switch self.buttonState {
	case CTRL_UP:
		self.ball.Y += 1
	case CTRL_DOWN:
		self.ball.Y -= 1
	}

	self.ball.R = float32(degree) * math.Pi / 180
}
