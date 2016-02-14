package simra

import (
	"image"

	"github.com/pankona/gomo-simra/simra/peer"
)

type Sprite struct {
	peer.Sprite
}

func (sprite *Sprite) ReplaceTexture(assetName string, rect image.Rectangle) {
	LogDebug("IN")
	tex := peer.GetGLPeer().LoadTexture(assetName, rect)
	peer.GetSpriteContainer().ReplaceTexture(&sprite.Sprite, tex)
	LogDebug("OUT")
}

func (sprite *Sprite) AddTouchListener(listener peer.TouchListener) {
	LogDebug("IN")
	sprite.Sprite.AddTouchListener(listener)
	LogDebug("OUT")
}

func (sprite *Sprite) RemoveAllTouchListener() {
	LogDebug("IN")
	sprite.Sprite.RemoveAllTouchListener()
	LogDebug("OUT")
}
