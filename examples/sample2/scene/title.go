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

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// initialize sprites
	self.initialize()

	simra.LogDebug("[OUT]")
}

func (self *Title) initialize() {
	// add background sprite
	self.background.W = float32(config.ScreenWidth)
	self.background.H = float32(config.ScreenHeight)

	// put center of screen
	self.background.X = config.ScreenWidth / 2
	self.background.Y = config.ScreenHeight / 2

	simra.GetInstance().AddSprite("title.png",
		image.Rect(0, 0, int(self.background.W), int(self.background.H)),
		&self.background)

	self.background.AddTouchListener(self)
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
