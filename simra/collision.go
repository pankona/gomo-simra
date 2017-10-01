package simra

// Collider represents an interface of collidables
type Collider interface {
	GetXYWH() (x, y, w, h float32)
}

// CollisionListener represents a interface of listener of collision
type CollisionListener interface {
	OnCollision(c1, c2 Collider)
}
