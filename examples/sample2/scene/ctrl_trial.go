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
}

const (
	CTRL_NOP = iota
	CTRL_UP
	CTRL_DOWN
)

func (self *CtrlTrial) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.SCREEN_WIDTH, config.SCREEN_HEIGHT)

	// TODO: add global touch listener to catch touch end event

	// initialize sprites
	self.initSprites()

	simra.LogDebug("[OUT]")
}

func (self *CtrlTrial) initSprites() {
	self.initBall()
	self.initCtrlDown()
	self.initCtrlUp()
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
	CTRL_MARGIN_LEFT    = 10
	CTRL_MARGIN_BOTTOM  = 10
	CTRL_MARGIN_BETWEEN = 10
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
