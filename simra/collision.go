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
	// TODO: exclusive controll
	LogDebug("IN")
	comap = append(comap, &collisionMap{c1, c2, listener})
	LogDebug("OUT")
}

func (simra *Simra) removeCollisionMap(c *collisionMap) {
	result := []*collisionMap{}

	for _, v := range comap {
		if c.c1 != v.c1 && c.c2 != v.c2 && v != c {
			result = append(result, v)
		}
	}

	comap = result
}

// RemoveAllCollisionListener removes all registered listeners
func (simra *Simra) RemoveAllCollisionListener() {
	LogDebug("IN")
	comap = nil
	LogDebug("OUT")
}

func (simra *Simra) collisionCheckAndNotify() {
	//LogDebug("IN")

	// check collision
	for _, v := range comap {
		// TODO: refactor around here...
		x1, y1, w1, h1 := v.c1.GetXYWH()
		x2, y2, w2, h2 := v.c2.GetXYWH()

		p1 := &point{x1 - w1/2, y1 + h1/2}
		p2 := &point{x1 + w1/2, y1 + h1/2}
		p3 := &point{x1 - w1/2, y1 - h1/2}
		p4 := &point{x1 + w1/2, y1 - h1/2}

		if p1.x >= (x2-w2/2) && p1.x <= (x2+w2/2) &&
			p1.y >= (y2-h2/2) && p1.y <= (y2+h2/2) {
			v.listener.OnCollision(v.c1, v.c2)
			return
		}
		if p2.x >= (x2-w2/2) && p2.x <= (x2+w2/2) &&
			p2.y >= (y2-h2/2) && p2.y <= (y2+h2/2) {
			v.listener.OnCollision(v.c1, v.c2)
			return
		}
		if p3.x >= (x2-w2/2) && p3.x <= (x2+w2/2) &&
			p3.y >= (y2-h2/2) && p3.y <= (y2+h2/2) {
			v.listener.OnCollision(v.c1, v.c2)
			return
		}
		if p4.x >= (x2-w2/2) && p4.x <= (x2+w2/2) &&
			p4.y >= (y2-h2/2) && p4.y <= (y2+h2/2) {
			v.listener.OnCollision(v.c1, v.c2)
			return
		}
	}
	//LogDebug("OUT")
}

// RemoveCollisionListener removes a collision map by specified collider instance.
func (simra *Simra) RemoveCollisionListener(c1, c2 Collider) {
	// TODO: exclusive controll
	LogDebug("IN")
	simra.removeCollisionMap(&collisionMap{c1, c2, nil})
	LogDebug("OUT")
}

func (simra *Simra) comapLength() int {
	return len(comap)
}
