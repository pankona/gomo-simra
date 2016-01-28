package simra

import (
	"image"

	"github.com/pankona/gomo-simra/simra/gomo"
	"github.com/pankona/gomo-simra/simra/peer"
)

type Simra struct {
	driver Driver
}

var simra *Simra = nil

func GetInstance() *Simra {
	peer.LogDebug("IN")
	if simra == nil {
		simra = &Simra{}
	}
	peer.LogDebug("OUT")
	return simra
}

func (self *Simra) onUpdate() {
	if self.driver != nil {
		self.driver.Drive()
	}
}

func (self *Simra) onStopped() {
	peer.LogDebug("IN")
	self.driver = nil
	peer.GetGLPeer().Finalize()
	peer.LogDebug("OUT")
}

func (self *Simra) Start(onStart, onStop chan bool) {
	peer.LogDebug("IN")
	gomo.GetInstance().Initialize(onStart, onStop, self.onUpdate)
	peer.GetSpriteContainer().Initialize()

	gomo.GetInstance().Start()
	peer.LogDebug("OUT")
}

func (self *Simra) SetScene(driver Driver) {
	peer.LogDebug("IN")
	peer.GetGLPeer().Reset()
	peer.GetSpriteContainer().RemoveSprites()

	self.driver = driver
	driver.Initialize()
	peer.LogDebug("OUT")
}

func (self *Simra) AddSprite(assetName string, rect image.Rectangle, s *Sprite) {
	tex := peer.GetGLPeer().LoadTexture(assetName, rect)
	peer.GetSpriteContainer().AddSprite(&s.Sprite, tex)
}

func (self *Simra) SetDesiredScreenSize(w, h float32) {
	peer.SetDesiredScreenSize(w, h)
}

func (self *Simra) AddTouchListener(listener peer.TouchListener) {
	peer.GetTouchPeer().AddTouchListener(listener)
}

func (self *Simra) RemoveTouchListener(listener peer.TouchListener) {
	peer.GetTouchPeer().RemoveTouchListener(listener)
}

func LogDebug(format string, a ...interface{}) {
	peer.LogDebug(format, a...)
}

func LogError(format string, a ...interface{}) {
	peer.LogError(format, a...)
}
