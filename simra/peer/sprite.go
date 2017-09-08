package peer

// Spriter represents interface of Sprite
type Spriter interface {
	// GetWidth returns width of sprite
	GetWidth() int
	// GetHeight returns height of sprite
	GetHeight() int
	// GetX returns x position of sprite
	GetX() int
	// GetY returns y position of sprite
	GetY() int
	// AddTouchListener registers a listener to notify touch event.
	AddTouchListener(listener TouchListener)
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

// GetWidth returns width of sprite
func (sprite *Sprite) GetWidth() int {
	return sprite.W
}

// GetHeight returns height of sprite
func (sprite *Sprite) GetHeight() int {
	return sprite.H
}

// GetX returns x position of sprite
func (sprite *Sprite) GetX() int {
	return sprite.X
}

// GetY returns y position of sprite
func (sprite *Sprite) GetY() int {
	return sprite.Y
}

// AddTouchListener registers a listener to notify touch event.
func (sprite *Sprite) AddTouchListener(listener TouchListener) {
	LogDebug("IN")
	sprite.touchListeners = append(sprite.touchListeners, &listener)
	LogDebug("OUT")
}

// RemoveAllTouchListener removes all registered listeners from sprite.
func (sprite *Sprite) RemoveAllTouchListener() {
	LogDebug("IN")
	sprite.touchListeners = nil
	LogDebug("OUT")
}
