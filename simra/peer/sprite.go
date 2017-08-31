package peer

// Spriter represents interface of Sprite
type Spriter interface {
	// GetWidth returns width of sprite
	GetWidth() float32
	// GetHeight returns height of sprite
	GetHeight() float32
	// GetX returns x position of sprite
	GetX() float32
	// GetY returns y position of sprite
	GetY() float32
	// AddTouchListener registers a listener to notify touch event.
	AddTouchListener(listener TouchListener)
	// RemoveAllTouchListener removes all registered listeners from sprite.
	RemoveAllTouchListener()
}

// NewSprite returns an instance of Spriter
func NewSprite() Spriter {
	return &privateSprite{}
}

// TODO: remove after replacing &Sprite{} with NewSprite()
type privateSprite = Sprite

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

// GetWidth returns width of sprite
func (sprite *Sprite) GetWidth() float32 {
	return sprite.W
}

// GetHeight returns height of sprite
func (sprite *Sprite) GetHeight() float32 {
	return sprite.H
}

// GetX returns x position of sprite
func (sprite *Sprite) GetX() float32 {
	return sprite.X
}

// GetY returns y position of sprite
func (sprite *Sprite) GetY() float32 {
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
