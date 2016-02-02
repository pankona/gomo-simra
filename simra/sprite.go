package simra

import (
	"github.com/pankona/gomo-simra/simra/peer"
	"image"
)

type Sprite struct {
	peer.Sprite
}

func (self *Sprite) ReplaceTexture(assetName string, rect image.Rectangle) {
	LogDebug("IN")
	tex := peer.GetGLPeer().LoadTexture(assetName, rect)
	peer.GetSpriteContainer().ReplaceTexture(&self.Sprite, tex)
	LogDebug("OUT")
}

func (self *Sprite) AddTouchListener(listener peer.TouchListener) {
	LogDebug("IN")
	self.Sprite.AddTouchListener(listener)
	LogDebug("OUT")
}

func (self *Sprite) RemoveAllTouchListener() {
	LogDebug("IN")
	self.Sprite.RemoveAllTouchListener()
	LogDebug("OUT")
}
