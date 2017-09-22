package simra

import (
	"context"

	"github.com/pankona/gomo-simra/simra/fps"
	"github.com/pankona/gomo-simra/simra/internal/peer"
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
	// SetPosition sets sprite's position
	SetPosition(x, y int)
	// SetPositionX sets sprite's position X
	SetPositionX(x int)
	// SetPositionY sets sprite's position Y
	SetPositionY(y int)
	// SetScale sets sprite's size
	SetScale(w, h int)
	// SetScaleW sets sprite's size W
	SetScaleW(w int)
	// SetScaleH sets sprite's size H
	SetScaleH(h int)
	// GetPosition gets sprites position
	GetPosition() Position
	// GetScale gets sprites size
	GetScale() Scale
	// SetRotate sets sprite's rotation
	SetRotate(r float32)
	// getRotate gets sprite's rotation
	GetRotate() float32
}

// Position represents position of sprite
type Position struct {
	X, Y int
}

// Scale represents size of sprite
type Scale struct {
	W, H int
}

// Sprite represents a sprite object.
type sprite struct {
	peer.Sprite
	simra           *simra
	animationSets   map[string]*AnimationSet
	animationCancel func()
	texture         *Texture
}

// ReplaceTexture replaces sprite's texture with specified image resource.
func (sprite *sprite) ReplaceTexture(texture *Texture) {
	LogDebug("IN")
	// retain reference for texture to avoid to be discarded by GC
	sprite.texture = texture
	sc := sprite.simra.spritecontainer
	sc.ReplaceTexture(&sprite.Sprite, texture.texture)
	LogDebug("OUT")
}

// AddTouchListener registers a listener for touch event.
// Touch event will be notified when "sprite" is touched.
func (sprite *sprite) AddTouchListener(listener peer.TouchListener) {
	LogDebug("IN")
	sprite.Sprite.AddTouchListener(listener)
	LogDebug("OUT")
}

// RemoveAllTouchListener removes all listeners already registered.
func (sprite *sprite) RemoveAllTouchListener() {
	LogDebug("IN")
	sprite.Sprite.RemoveAllTouchListener()
	LogDebug("OUT")
}

// AddAnimationSet adds a specified AnimationSet to sprite
func (sprite *sprite) AddAnimationSet(animationName string, set *AnimationSet) {
	LogDebug("IN")
	sprite.animationSets[animationName] = set
	LogDebug("OUT")
}

// StartAnimation starts animation by specified animation name
func (sprite *sprite) StartAnimation(animationName string, shouldLoop bool, animationEndCallback func()) {
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

func (sprite *sprite) startAnimation(ctx context.Context, animationName string, shouldLoop bool, animationEndCallback func()) {
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
func (sprite *sprite) StopAnimation() {
	LogDebug("IN")
	if sprite == nil {
		return
	}

	if sprite.animationCancel != nil {
		sprite.animationCancel()
	}
	LogDebug("OUT")
}

func (sprite *sprite) SetPosition(x, y int) {
	sprite.X, sprite.Y = x, y
}

func (sprite *sprite) SetPositionX(x int) {
	sprite.X = x
}

func (sprite *sprite) SetPositionY(y int) {
	sprite.Y = y
}

func (sprite *sprite) SetScale(w, h int) {
	sprite.W, sprite.H = w, h
}

func (sprite *sprite) SetScaleW(w int) {
	sprite.W = w
}

func (sprite *sprite) SetScaleH(h int) {
	sprite.H = h
}

func (sprite *sprite) GetPosition() Position {
	return Position{sprite.X, sprite.Y}
}

func (sprite *sprite) GetScale() Scale {
	return Scale{sprite.W, sprite.H}
}

func (sprite *sprite) SetRotate(r float32) {
	sprite.R = r
}

func (sprite *sprite) GetRotate() float32 {
	return sprite.R
}
