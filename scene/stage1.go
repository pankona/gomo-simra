package scene

import (
	"fmt"
	"image"

	"github.com/pankona/gomobile_gamelib_test/peer"
)

type Stage1 struct {
	gopher         peer.PeerSprite
	ball           peer.PeerSprite
	notifySceneEnd func(nextScene Driver)
}

func (self *Stage1) Initialize(sceneEndCallback func(nextScene Driver)) {
	self.notifySceneEnd = sceneEndCallback

	peer.SetDesiredScreenSize(1080/2, 1920/2)
	peer.GetTouchPeer().AddTouchListener(self)

	// initialize sprites
	self.initSprite()
}

func (self *Stage1) initSprite() {
	fmt.Println("[IN] Stage1.initSprite")

	self.initGopher()

	fmt.Println("[OUT] Title.initTitleSprite")
}

func (self *Stage1) initGopher() {
	// add gopher sprite
	self.gopher.W = float32(140)
	self.gopher.H = float32(90)

	// put center of screen at start
	self.gopher.X = 1080 / 2 / 2
	self.gopher.Y = 1920 / 2 / 2

	tex_gopher := peer.GetGLPeer().LoadTexture("waza-gophers.jpeg",
		image.Rect(152, 10, 152+int(self.gopher.W), 10+int(self.gopher.H)))
	peer.GetGLPeer().AddSprite(&self.gopher, tex_gopher)
}

func (self *Stage1) Drive() {
	//peer.GetInstance().Update()
}

func (self *Stage1) OnTouchBegin(x, y float32) {
	fmt.Println("OnTouch = ", x, y)
	self.gopher.X = x
	self.gopher.Y = y
}

func (self *Stage1) OnTouchMove(x, y float32) {
	self.gopher.X = x
	self.gopher.Y = y
}

func (self *Stage1) OnTouchEnd(x, y float32) {
	self.gopher.X = x
	self.gopher.Y = y
}
