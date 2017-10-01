package peer

import "github.com/pankona/gomo-simra/simra/simlog"

// Spriter represents interface of Sprite
type Spriter interface {
	// AddTouchListener registers a listener to notify touch event.
	AddTouchListener(l TouchListener)
	// RemoveAllTouchListener removes all registered listeners from sprite.
	RemoveAllTouchListener()
}

// NewSprite returns an instance of Spriter
func NewSprite() Spriter {
	return &Sprite{}
}

// Sprite represents a sprite.
// Deprecated. this will be unexposed.
type Sprite struct {
	// W = width of sprite
	W float32
	// H = height of sprite
	H float32
	// X = x position of sprite
	X float32
	// Y = y position of sprite
	Y float32
	// R = radius of sprite (use for rotation)
	R float32
	// touchListeners is listeners to notify touch event
	touchListeners []*TouchListener
}

// AddTouchListener registers a listener to notify touch event.
func (s *Sprite) AddTouchListener(l TouchListener) {
	simlog.FuncIn()
	s.touchListeners = append(s.touchListeners, &l)
	simlog.FuncOut()
}

// RemoveAllTouchListener removes all registered listeners from sprite.
func (s *Sprite) RemoveAllTouchListener() {
	simlog.FuncIn()
	s.touchListeners = nil
	simlog.FuncOut()
}
