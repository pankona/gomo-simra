package simra

import (
	"image"

	"github.com/pankona/gomo-simra/simra/peer"
)

// Sprite represents a sprite object.
type Sprite struct {
	peer.Sprite
}

// ReplaceTexture replaces sprite's texture with specified image resource.
func (sprite *Sprite) ReplaceTexture(assetName string, rect image.Rectangle) {
	LogDebug("IN")
	tex := peer.GetGLPeer().LoadTexture(assetName, rect)
	peer.GetSpriteContainer().ReplaceTexture(&sprite.Sprite, tex)
	LogDebug("OUT")
}

// AddTouchListener registers a listener for touch event.
// Touch event will be notified when "sprite" is touched.
func (sprite *Sprite) AddTouchListener(listener peer.TouchListener) {
	LogDebug("IN")
	sprite.Sprite.AddTouchListener(listener)
	LogDebug("OUT")
}

// RemoveAllTouchListener removes all listeners already registered.
func (sprite *Sprite) RemoveAllTouchListener() {
	LogDebug("IN")
	sprite.Sprite.RemoveAllTouchListener()
	LogDebug("OUT")
}
