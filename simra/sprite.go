package simra

import (
	"image"
	"time"

	"github.com/pankona/gomo-simra/simra/peer"
)

// Sprite represents a sprite object.
type Sprite struct {
	peer.Sprite
	animationSets map[string]*AnimationSet
}

// AnimationSet represents a set of image for animation
type AnimationSet struct {
	textures []*Texture
	interval time.Duration
}

// NewAnimationSet returns an instance of AnimationSet
func NewAnimationSet() *AnimationSet {
	LogDebug("IN")
	LogDebug("OUT")
	return &AnimationSet{}
}

// AddTexture adds a specified texture to AnimationSet
func (animation *AnimationSet) AddTexture(texture *Texture) {
	LogDebug("IN")
	animation.textures = append(animation.textures, texture)
	LogDebug("OUT")
}

// SetInterval sets interval of animation
func (animation *AnimationSet) SetInterval(interval time.Duration) {
	LogDebug("IN")
	animation.interval = interval
	LogDebug("OUT")
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

// AddAnimationSet adds a specified AnimationSet to sprite
func (sprite *Sprite) AddAnimationSet(animationName string, set *AnimationSet) {
	sprite.animationSets[animationName] = set
}

// StartAnimation starts animation by specified animation name
func (sprite *Sprite) StartAnimation(animationName string) error {
	// TODO: implement
	return nil
}

// StopAnimation stops animation
func (sprite *Sprite) StopAnimation() {
	// TODO: implement
}
