package scene

import (
	"fmt"
	"image"
	"math"

	"github.com/pankona/gomo-simra/peer"
)

type Stage1 struct {
	gopher peer.PeerSprite
	ball   peer.PeerSprite
}

func (self *Stage1) Initialize() {

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

var degree float32 = 0

func (self *Stage1) Drive() {
	degree += 1
	if degree >= 360 {
		degree = 0
	}
	self.gopher.R = float32(degree) * math.Pi / 180
}

func (self *Stage1) OnTouchBegin(x, y float32) {
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
