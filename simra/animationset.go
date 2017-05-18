package simra

import "time"

// AnimationSet represents a set of image for animation
type AnimationSet struct {
	textures []*Texture
	interval time.Duration
}

// NewAnimationSet returns an instance of AnimationSet
func NewAnimationSet() *AnimationSet {
	LogDebug("IN")
	LogDebug("OUT")
	defaultInterval := 100 * time.Millisecond
	return &AnimationSet{interval: defaultInterval}
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
