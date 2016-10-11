package simra

type collisionMap struct {
	c1       Collider
	c2       Collider
	listener CollisionListener
}

var comap []*collisionMap

// Collider represents an interface of collidables
type Collider interface {
	GetXYWH() (x, y, w, h int)
}

// CollisionListener represents a interface of listener of collision
type CollisionListener interface {
	OnCollision(c1, c2 Collider)
}

// AddCollisionListener add a callback function that is called on
// collision is detected between c1 and c2.
func (simra *Simra) AddCollisionListener(c1, c2 Collider, listener CollisionListener) {
	LogDebug("IN")
	comap = append(comap, &collisionMap{c1, c2, listener})
	LogDebug("OUT")
}

// RemoveAllCollisionListener removes all registered listeners
func (simra *Simra) RemoveAllCollisionListener() {
	LogDebug("IN")
	comap = nil
	LogDebug("OUT")
}
