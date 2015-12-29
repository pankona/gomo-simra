package simra

import (
	"fmt"
	"image"

	"github.com/pankona/gomo-simra/peer"
)

type Title struct {
	background peer.PeerSprite
}

func (self *Title) Initialize() {
	fmt.Println("[IN] Title.Initialize()")
	peer.SetDesiredScreenSize(1080/2, 1920/2)
	peer.GetTouchPeer().AddTouchListener(self)

	// initialize sprites
	self.initialize()
}

func (self *Title) initialize() {
	fmt.Println("[IN] Title.initialize")

	// add background sprite
	self.background.W = float32(1080 / 2)
	self.background.H = float32(1920 / 2)

	// put center of screen
	self.background.X = 1080 / 2 / 2
	self.background.Y = 1920 / 2 / 2

	tex_background := peer.GetGLPeer().LoadTexture("title.png",
		image.Rect(0, 0, int(self.background.W), int(self.background.H)))
	peer.GetGLPeer().AddSprite(&self.background, tex_background)

	fmt.Println("[OUT] Title.initialize")
}

func (self *Title) Drive() {
}

func (self *Title) OnTouchBegin(x, y float32) {
}

func (self *Title) OnTouchMove(x, y float32) {
}

func (self *Title) OnTouchEnd(x, y float32) {
	fmt.Println("OnTouchBegin = ", x, y)
}
