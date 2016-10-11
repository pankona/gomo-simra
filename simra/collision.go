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

// RemoveCollisionMap removes specified comap from list
func (simra *Simra) RemoveCollisionMap(c *collisionMap) {
	result := []*collisionMap{}

	for _, v := range comap {
		if v != c {
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
	LogDebug("IN")

CollisionDetection:
	// check collision
	for _, v := range comap {
		// TODO: refactor this Fxxkin' part

		if v.c1 == nil || v.c2 == nil {
			// remove and bailout...
			simra.RemoveCollisionMap(v)
			goto CollisionDetection
		}

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
	LogDebug("OUT")
}
