package simra

import "github.com/pankona/gomo-simra/simra/simlog"

// AnimationSet represents a set of image for animation
type AnimationSet struct {
	textures []*Texture
	interval int64 // frames
}

// NewAnimationSet returns an instance of AnimationSet
func NewAnimationSet() *AnimationSet {
	simlog.FuncIn()
	defaultInterval := (int64)(6)
	simlog.FuncOut()
	return &AnimationSet{interval: defaultInterval}
}

// AddTexture adds a specified texture to AnimationSet
func (animation *AnimationSet) AddTexture(texture *Texture) {
	simlog.FuncIn()
	animation.textures = append(animation.textures, texture)
	simlog.FuncOut()
}

// SetInterval sets interval of animation
func (animation *AnimationSet) SetInterval(interval int64) {
	simlog.FuncIn()
	animation.interval = interval
	simlog.FuncOut()
}
