package scene

import (
	"image"

	"github.com/pankona/gomo-simra/examples/sample2/scene/config"
	"github.com/pankona/gomo-simra/peer"
	"github.com/pankona/gomo-simra/simra"
)

type Title struct {
	background peer.PeerSprite
}

func (self *Title) Initialize() {
	peer.LogDebug("[IN]")

	peer.SetDesiredScreenSize(config.SCREEN_WIDTH, config.SCREEN_HEIGHT)
	peer.GetTouchPeer().AddTouchListener(self)

	// initialize sprites
	self.initialize()

	peer.LogDebug("[OUT]")
}

func (self *Title) initialize() {
	// add background sprite
	self.background.W = float32(config.SCREEN_WIDTH)
	self.background.H = float32(config.SCREEN_HEIGHT)

	// put center of screen
	self.background.X = config.SCREEN_WIDTH / 2
	self.background.Y = config.SCREEN_HEIGHT / 2

	tex_background := peer.GetGLPeer().LoadTexture("title.png",
		image.Rect(0, 0, int(self.background.W), int(self.background.H)))
	peer.GetGLPeer().AddSprite(&self.background, tex_background)
}

func (self *Title) Drive() {
}

func (self *Title) OnTouchBegin(x, y float32) {
}

func (self *Title) OnTouchMove(x, y float32) {
}

func (self *Title) OnTouchEnd(x, y float32) {
	// scene end. go to next scene
	simra.GetInstance().SetScene(&CtrlTrial{})
}
