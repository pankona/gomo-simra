package simra

import (
	"context"

	"github.com/pankona/gomo-simra/simra/fps"
	"github.com/pankona/gomo-simra/simra/peer"
)

// Spriter represents an interface of Sprite
type Spriter interface {
	// ReplaceTexture replaces sprite's texture with specified image resource.
	ReplaceTexture(texture *Texture)
	// AddTouchListener registers a listener for touch event.
	// Touch event will be notified when "sprite" is touched.
	AddTouchListener(listener peer.TouchListener)
	// RemoveAllTouchListener removes all listeners already registered.
	RemoveAllTouchListener()
	// AddAnimationSet adds a specified AnimationSet to sprite
	AddAnimationSet(animationName string, set *AnimationSet)
	// StartAnimation starts animation by specified animation name
	StartAnimation(animationName string, shouldLoop bool, animationEndCallback func())
	// StopAnimation stops animation
	StopAnimation()
	// ProgressAnimation progresses registered animation
	ProgressAnimation()
}

// Sprite represents a sprite object.
type Sprite struct {
	peer.Sprite
	animationSets   map[string]*AnimationSet
	animationCancel func()
	texture         *Texture
}

// NewSprite returns an instance of Sprite
func NewSprite() Spriter {
	return &Sprite{animationSets: map[string]*AnimationSet{}}
}

// ReplaceTexture replaces sprite's texture with specified image resource.
func (sprite *Sprite) ReplaceTexture(texture *Texture) {
	LogDebug("IN")
	// retain reference for texture to avoid to be discarded by GC
	sprite.texture = texture
	peer.GetSpriteContainer().ReplaceTexture(&sprite.Sprite, texture.Texture)
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
func (sprite *Sprite) StartAnimation(animationName string, shouldLoop bool, animationEndCallback func()) {
	LogDebug("IN")
	if sprite.animationCancel != nil {
		// animation is already in progress. don't start.
		// TODO: should exlude control
		return
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	sprite.animationCancel = cancel
	go sprite.startAnimation(ctx, animationName, shouldLoop, animationEndCallback)
	LogDebug("OUT")
}

func (sprite *Sprite) startAnimation(ctx context.Context, animationName string, shouldLoop bool, animationEndCallback func()) {
	LogDebug("IN")
	animationSet := sprite.animationSets[animationName]
	if animationSet == nil {
		panic("specified animation is not set. animation name = " + animationName)
	}

	loopCount := 0
animation:
	for {
		select {
		case <-ctx.Done():
			break animation
		case <-fps.After(animationSet.interval):
			sprite.ReplaceTexture(animationSet.textures[loopCount])
			loopCount = (loopCount + 1) % len(animationSet.textures)
			if !shouldLoop && loopCount == 0 {
				break animation
			}
		}
	}
	sprite.animationCancel = nil
	animationEndCallback()
	LogDebug("OUT")
}

// StopAnimation stops animation
func (sprite *Sprite) StopAnimation() {
	LogDebug("IN")
	if sprite == nil {
		return
	}

	if sprite.animationCancel != nil {
		sprite.animationCancel()
	}
	LogDebug("OUT")
}

// ProgressAnimation progresses registered animation
func (sprite *Sprite) ProgressAnimation() {
	fps.Progress()
}
