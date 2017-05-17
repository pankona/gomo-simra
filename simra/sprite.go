package simra

import (
	"context"
	"image"
	"time"

	"github.com/pankona/gomo-simra/simra/peer"
)

// Sprite represents a sprite object.
type Sprite struct {
	peer.Sprite
	animationSets   map[string]*AnimationSet
	animationCancel func()
}

// NewSprite returns an instance of Sprite
func NewSprite() *Sprite {
	return &Sprite{animationSets: map[string]*AnimationSet{}}
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
// TODO: deprecate. use ReplaceTexture2 and remove this, then rename function name
func (sprite *Sprite) ReplaceTexture(assetName string, rect image.Rectangle) {
	LogDebug("IN")
	tex := peer.GetGLPeer().LoadTexture(assetName, rect)
	peer.GetSpriteContainer().ReplaceTexture(&sprite.Sprite, tex)
	LogDebug("OUT")
}

// ReplaceTexture2 replaces sprite's texture with specified image resource.
func (sprite *Sprite) ReplaceTexture2(texture *Texture) {
	LogDebug("IN")
	peer.GetSpriteContainer().ReplaceTexture2(&sprite.Sprite, texture.Texture)
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
	LogDebug("IN")
	sprite.animationSets[animationName] = set
	LogDebug("OUT")
}

// StartAnimation starts animation by specified animation name
func (sprite *Sprite) StartAnimation(animationName string) {
	LogDebug("IN")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	sprite.animationCancel = cancel
	go sprite.startAnimation(ctx, animationName)
	LogDebug("OUT")
}

func (sprite *Sprite) startAnimation(ctx context.Context, animationName string) {
	LogDebug("IN")
	animationSet := sprite.animationSets[animationName]
	if animationSet == nil {
		panic("specified animation is not set. animation name")
	}

	loopCount := 0
animation:
	for {
		select {
		case <-ctx.Done():
			break animation
		case <-time.After(animationSet.interval):
			sprite.ReplaceTexture2(animationSet.textures[loopCount])
			loopCount = (loopCount + 1) % len(animationSet.textures)
		}
	}
	LogDebug("OUT")
}

// StopAnimation stops animation
func (sprite *Sprite) StopAnimation() {
	LogDebug("IN")
	if sprite.animationCancel != nil {
		sprite.animationCancel()
	}
	LogDebug("OUT")
}
