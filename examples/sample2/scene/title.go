package scene

import (
	"image"

	"github.com/pankona/gomo-simra/examples/sample2/scene/config"
	"github.com/pankona/gomo-simra/simra"
)

type Title struct {
	background simra.Sprite
}

func (self *Title) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.SCREEN_WIDTH, config.SCREEN_HEIGHT)
	simra.GetInstance().AddTouchListener(self)

	// initialize sprites
	self.initialize()

	simra.LogDebug("[OUT]")
}

func (self *Title) initialize() {
	// add background sprite
	self.background.W = float32(config.SCREEN_WIDTH)
	self.background.H = float32(config.SCREEN_HEIGHT)

	// put center of screen
	self.background.X = config.SCREEN_WIDTH / 2
	self.background.Y = config.SCREEN_HEIGHT / 2

	simra.GetInstance().AddSprite("title.png",
		image.Rect(0, 0, int(self.background.W), int(self.background.H)),
		&self.background)
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
