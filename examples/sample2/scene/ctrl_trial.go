package scene

import (
	"image"
	"math"

	"github.com/pankona/gomo-simra/peer"
)

type CtrlTrial struct {
	ball peer.PeerSprite
}

func (self *CtrlTrial) Initialize() {
	peer.LogDebug("[IN]")

	peer.SetDesiredScreenSize(SCREEN_WIDTH, SCREEN_HEIGHT)
	peer.GetTouchPeer().AddTouchListener(self)

	// initialize sprites
	self.initBall()

	peer.LogDebug("[OUT]")
}

func (self *CtrlTrial) initBall() {
	// add ball sprite
	self.ball.W = float32(48)
	self.ball.H = float32(48)

	// put center of screen at start
	self.ball.X = SCREEN_WIDTH / 2
	self.ball.Y = SCREEN_HEIGHT / 2

	tex_ball := peer.GetGLPeer().LoadTexture("ball.png",
		image.Rect(0, 0, int(self.ball.W), int(self.ball.H)))
	peer.GetGLPeer().AddSprite(&self.ball, tex_ball)
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
