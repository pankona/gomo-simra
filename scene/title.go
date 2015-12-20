package scene

import (
	"fmt"
	"image"

	"github.com/pankona/gomobile_gamelib_test/peer"
)

type Title struct {
	background     peer.PeerSprite
	notifySceneEnd func(nextScene Driver)
}

func (self *Title) Initialize(sceneEndCallback func(nextScene Driver)) {
	self.notifySceneEnd = sceneEndCallback

	peer.SetDesiredScreenSize(1080/2, 1920/2)

	// initialize sprites
	self.initSprite()

	// add touch event listener
	peer.GetTouchPeer().AddTouchListener(self)
}

func (self *Title) initSprite() {
	fmt.Println("[IN] Title.initTitleSprite")

	// add background sprite
	self.background.W = float32(1080 / 2)
	self.background.H = float32(1920 / 2)

	// put center of screen
	self.background.X = 1080 / 2 / 2
	self.background.Y = 1920 / 2 / 2

	tex_background := peer.GetGLPeer().LoadTexture("title.png",
		image.Rect(0, 0, int(self.background.W), int(self.background.H)))
	peer.GetGLPeer().AddSprite(&self.background, tex_background)

	fmt.Println("[OUT] Title.initTitleSprite")
}

func (self *Title) Drive() {
}

func (self *Title) OnTouchBegin(x, y float32) {
}

func (self *Title) OnTouchMove(x, y float32) {
}

func (self *Title) OnTouchEnd(x, y float32) {
	fmt.Println("OnTouchBegin = ", x, y)
	self.notifySceneEnd(&Stage1{})
}
