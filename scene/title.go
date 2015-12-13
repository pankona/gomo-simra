package scene

import (
	"fmt"
	"image"

	"github.com/pankona/gomobile_gamelib_test/peer"
)

type Title struct {
	peer.PeerSprite
	notifySceneEnd func(nextScene Driver)
}

func (self *Title) Initialize(sceneEndCallback func(nextScene Driver)) {
	self.notifySceneEnd = sceneEndCallback
	self.initTitleSprite()
}

func (self *Title) initTitleSprite() {
	fmt.Println("[IN] Title.initTitleSprite")
	sz := peer.GetInstance().GetScreenSize()
	self.W = float32(sz.WidthPx)
	self.H = float32(sz.HeightPx)
	tex_title := peer.GetInstance().LoadTexture("title.png", image.Rect(0, 0, int(self.H), int(self.H)))
	peer.GetInstance().AddSprite(&self.PeerSprite, tex_title)
	fmt.Println("[OUT] Title.initTitleSprite")
}

func (self *Title) Drive() {
	//fmt.Println("[IN] Title.Drive")
	peer.GetInstance().Update()
	//fmt.Println("[OUT] Title.Drive")
}
