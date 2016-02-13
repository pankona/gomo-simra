package simra

import (
	"image"

	"github.com/pankona/gomo-simra/simra/gomo"
	"github.com/pankona/gomo-simra/simra/peer"
)

type Simra struct {
	driver Driver
}

var simra *Simra

func GetInstance() *Simra {
	peer.LogDebug("IN")
	if simra == nil {
		simra = &Simra{}
	}
	peer.LogDebug("OUT")
	return simra
}

func (simra *Simra) onUpdate() {
	if simra.driver != nil {
		simra.driver.Drive()
	}
}

func (simra *Simra) onStopped() {
	peer.LogDebug("IN")
	simra.driver = nil
	peer.GetGLPeer().Finalize()
	peer.LogDebug("OUT")
}

func (simra *Simra) Start(onStart, onStop chan bool) {
	peer.LogDebug("IN")
	gomo.GetInstance().Initialize(onStart, onStop, simra.onUpdate)
	peer.GetSpriteContainer().Initialize()

	gomo.GetInstance().Start()
	peer.LogDebug("OUT")
}

func (simra *Simra) SetScene(driver Driver) {
	peer.LogDebug("IN")
	peer.GetGLPeer().Reset()
	peer.GetSpriteContainer().RemoveSprites()

	simra.driver = driver
	driver.Initialize()
	peer.LogDebug("OUT")
}

func (simra *Simra) AddSprite(assetName string, rect image.Rectangle, s *Sprite) {
	tex := peer.GetGLPeer().LoadTexture(assetName, rect)
	peer.GetSpriteContainer().AddSprite(&s.Sprite, tex)
}

func (simra *Simra) RemoveSprite(s *Sprite) {
	peer.GetSpriteContainer().RemoveSprite(&s.Sprite)
}

func (simra *Simra) SetDesiredScreenSize(w, h float32) {
	peer.SetDesiredScreenSize(w, h)
}

func (simra *Simra) AddTouchListener(listener peer.TouchListener) {
	peer.GetTouchPeer().AddTouchListener(listener)
}

func (simra *Simra) RemoveTouchListener(listener peer.TouchListener) {
	peer.GetTouchPeer().RemoveTouchListener(listener)
}

func LogDebug(format string, a ...interface{}) {
	peer.LogDebug(format, a...)
}

func LogError(format string, a ...interface{}) {
	peer.LogError(format, a...)
}
