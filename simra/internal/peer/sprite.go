package peer

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
	W int
	// H = height of sprite
	H int
	// X = x position of sprite
	X int
	// Y = y position of sprite
	Y int
	// R = radius of sprite (use for rotation)
	R float32
	// touchListeners is listeners to notify touch event
	touchListeners []*TouchListener
}

// AddTouchListener registers a listener to notify touch event.
func (s *Sprite) AddTouchListener(l TouchListener) {
	LogDebug("IN")
	s.touchListeners = append(s.touchListeners, &l)
	LogDebug("OUT")
}

// RemoveAllTouchListener removes all registered listeners from sprite.
func (s *Sprite) RemoveAllTouchListener() {
	LogDebug("IN")
	s.touchListeners = nil
	LogDebug("OUT")
}
