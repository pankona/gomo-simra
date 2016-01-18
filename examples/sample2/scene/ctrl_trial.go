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
}

func (self *CtrlTrial) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.SCREEN_WIDTH, config.SCREEN_HEIGHT)
	simra.GetInstance().AddTouchListener(self)

	// initialize sprites
	self.initSprites()

	simra.LogDebug("[OUT]")
}

func (self *CtrlTrial) initSprites() {
	self.initBall()
	self.initCtrlUp()
	self.initCtrlDown()
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

func (self *CtrlTrial) initCtrlUp() {
	// set size of CtrlUp
	self.ctrlup.W = float32(120)
	self.ctrlup.H = float32(120)

	// put CtrlUp on left bottom
	self.ctrlup.X = (self.ctrlup.W / 2) + 10
	self.ctrlup.Y =
		config.SCREEN_HEIGHT - (self.ctrlup.H / 2) -
			self.ctrlup.H - CTRL_MARGIN_BOTTOM - CTRL_MARGIN_BETWEEN

	// add sprite to glpeer
	simra.GetInstance().AddSprite("arrow.png",
		image.Rect(0, 0, int(self.ctrlup.W), int(self.ctrlup.H)),
		&self.ctrlup)
}

func (self *CtrlTrial) initCtrlDown() {
	// set size of CtrlDown
	self.ctrldown.W = float32(120)
	self.ctrldown.H = float32(120)

	// put CtrlDown on left bottom
	self.ctrldown.X = (self.ctrldown.W / 2) + 10
	self.ctrldown.Y =
		config.SCREEN_HEIGHT - (self.ctrldown.H / 2) - CTRL_MARGIN_BOTTOM

	// rotate arrow to indicate down control
	self.ctrldown.R = math.Pi

	// add sprite to glpeer
	simra.GetInstance().AddSprite("arrow.png",
		image.Rect(0, 0, int(self.ctrldown.W), int(self.ctrldown.H)),
		&self.ctrldown)
}

var degree float32 = 0

func (self *CtrlTrial) Drive() {
	degree += 1
	if degree >= 360 {
		degree = 0
	}
	self.ball.R = float32(degree) * math.Pi / 180
}

func (self *CtrlTrial) OnTouchBegin(x, y float32) {
}

func (self *CtrlTrial) OnTouchMove(x, y float32) {
}

func (self *CtrlTrial) OnTouchEnd(x, y float32) {
}
