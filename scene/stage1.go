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

	peer.GetInstance().SetDesiredScreenSize(1080/2, 1920/2)

	// initialize sprites
	self.initSprite()

	// add touch event listener
	peer.GetInstance().AddTouchListener(self)
}

func (self *Stage1) initSprite() {
	fmt.Println("[IN] Stage1.initSprite")

	// add gopher sprite
	self.gopher.W = float32(140)
	self.gopher.H = float32(90)
	tex_gopher := peer.GetInstance().LoadTexture("waza-gophers.jpeg",
		image.Rect(152, 10, 152+int(self.gopher.W), 10+int(self.gopher.H)))
	peer.GetInstance().AddSprite(&self.gopher, tex_gopher)

	fmt.Println("[OUT] Title.initTitleSprite")
}

func (self *Stage1) Drive() {
	//peer.GetInstance().Update()
}

func (self *Stage1) OnTouch(x, y float32) {
	fmt.Println("OnTouch = ", x, y)
}
