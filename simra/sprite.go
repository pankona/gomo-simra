package simra

import (
	"context"

	"github.com/pankona/gomo-simra/simra/fps"
	"github.com/pankona/gomo-simra/simra/internal/peer"
	"github.com/pankona/gomo-simra/simra/simlog"
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
	SetPosition(x, y float32)
	// SetPositionX sets sprite's position X
	SetPositionX(x float32)
	// SetPositionY sets sprite's position Y
	SetPositionY(y float32)
	// SetScale sets sprite's size
	SetScale(w, h float32)
	// SetScaleW sets sprite's size W
	SetScaleW(w float32)
	// SetScaleH sets sprite's size H
	SetScaleH(h float32)
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
	X, Y float32
}

// Scale represents size of sprite
type Scale struct {
	W, H float32
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
	simlog.FuncIn()
	// retain reference for texture to avoid to be discarded by GC
	sprite.texture = texture
	sc := sprite.simra.spritecontainer
	sc.ReplaceTexture(&sprite.Sprite, texture.texture)
	simlog.FuncOut()
}

// AddTouchListener registers a listener for touch event.
// Touch event will be notified when "sprite" is touched.
func (sprite *sprite) AddTouchListener(listener peer.TouchListener) {
	simlog.FuncIn()
	sprite.Sprite.AddTouchListener(listener)
	simlog.FuncOut()
}

// RemoveAllTouchListener removes all listeners already registered.
func (sprite *sprite) RemoveAllTouchListener() {
	simlog.FuncIn()
	sprite.Sprite.RemoveAllTouchListener()
	simlog.FuncOut()
}

// AddAnimationSet adds a specified AnimationSet to sprite
func (sprite *sprite) AddAnimationSet(animationName string, set *AnimationSet) {
	simlog.FuncIn()
	sprite.animationSets[animationName] = set
	simlog.FuncOut()
}

// StartAnimation starts animation by specified animation name
func (sprite *sprite) StartAnimation(animationName string, shouldLoop bool, animationEndCallback func()) {
	simlog.FuncIn()
	if sprite.animationCancel != nil {
		// animation is already in progress. don't start.
		// TODO: should exlude control
		return
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	sprite.animationCancel = cancel
	go sprite.startAnimation(ctx, animationName, shouldLoop, animationEndCallback)
	simlog.FuncOut()
}

func (sprite *sprite) startAnimation(ctx context.Context, animationName string, shouldLoop bool, animationEndCallback func()) {
	simlog.FuncIn()
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
	simlog.FuncOut()
}

// StopAnimation stops animation
func (sprite *sprite) StopAnimation() {
	simlog.FuncIn()
	if sprite == nil {
		return
	}

	if sprite.animationCancel != nil {
		sprite.animationCancel()
	}
	simlog.FuncOut()
}

func (sprite *sprite) SetPosition(x, y float32) {
	sprite.X, sprite.Y = x, y
}

func (sprite *sprite) SetPositionX(x float32) {
	sprite.X = x
}

func (sprite *sprite) SetPositionY(y float32) {
	sprite.Y = y
}

func (sprite *sprite) SetScale(w, h float32) {
	sprite.W, sprite.H = w, h
}

func (sprite *sprite) SetScaleW(w float32) {
	sprite.W = w
}

func (sprite *sprite) SetScaleH(h float32) {
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
