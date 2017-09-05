package simra

import (
	"testing"
	"time"
)

type c struct{}
type l struct{}

var onCollision = make(chan bool)

func (c *c) GetXYWH() (x, y, w, h int) {
	x, y, w, h = 0, 0, 0, 0
	return
}

func (l *l) OnCollision(c1, c2 Collider) {
	go func() {
		//time.Sleep(time.Millisecond * 500)
		onCollision <- true
	}()
}

func waitOnCollision(t *testing.T, shouldCallback bool) {
	select {
	case <-onCollision:
		if !shouldCallback {
			t.Error("unexpected OnCollision.")
		}
	case <-time.After(time.Millisecond * 300):
		if shouldCallback {
			t.Error("expected OnCollision but not fired.")
		}
	}
}

func Test_AddCollisionListener(t *testing.T) {
	var c1, c2 c
	var l l

	simra := sim
	simra.RemoveAllCollisionListener()
	if simra.comapLength() != 0 {
		t.Error("unexpected comap length. comapLength() =", simra.comapLength())
	}
	simra.AddCollisionListener(&c1, &c2, &l)
	if simra.comapLength() != 1 {
		t.Error("unexpected comap length. comapLength() =", simra.comapLength())
	}
	simra.collisionCheckAndNotify()
	waitOnCollision(t, true)
	if simra.comapLength() != 1 {
		t.Error("unexpected comap length. comapLength() =", simra.comapLength())
	}
	simra.AddCollisionListener(&c1, &c2, &l)
	if simra.comapLength() != 2 {
		t.Error("unexpected comap length. comapLength() =", simra.comapLength())
	}
	simra.RemoveAllCollisionListener()
	if simra.comapLength() != 0 {
		t.Error("unexpected comap length. comapLength() =", simra.comapLength())
	}
}

func Test_RemoveCollisionListener(t *testing.T) {

	var c1, c2 c
	var l l

	simra := sim
	simra.RemoveAllCollisionListener()
	if simra.comapLength() != 0 {
		t.Error("unexpected comap length. comapLength() =", simra.comapLength())
	}

	simra.AddCollisionListener(&c1, &c2, &l)
	if simra.comapLength() != 1 {
		t.Error("unexpected comap length. comapLength() =", simra.comapLength())
	}

	simra.collisionCheckAndNotify()
	waitOnCollision(t, true)
	if simra.comapLength() != 1 {
		t.Error("unexpected comap length. comapLength() =", simra.comapLength())
	}

	simra.RemoveCollisionListener(&c1, nil)
	simra.collisionCheckAndNotify()
	waitOnCollision(t, false)
	if simra.comapLength() != 0 {
		t.Error("unexpected comap length. comapLength() =", simra.comapLength())
	}

	simra.AddCollisionListener(&c1, &c2, &l)
	if simra.comapLength() != 1 {
		t.Error("unexpected comap length. comapLength() =", simra.comapLength())
	}

	simra.RemoveCollisionListener(nil, &c2)
	simra.collisionCheckAndNotify()
	waitOnCollision(t, false)
	if simra.comapLength() != 0 {
		t.Error("unexpected comap length. comapLength() =", simra.comapLength())
	}

	simra.AddCollisionListener(&c1, &c2, &l)
	if simra.comapLength() != 1 {
		t.Error("unexpected comap length. comapLength() =", simra.comapLength())
	}

	simra.RemoveCollisionListener(&c1, &c2)
	simra.collisionCheckAndNotify()
	waitOnCollision(t, false)
	if simra.comapLength() != 0 {
		t.Error("unexpected comap length. comapLength() =", simra.comapLength())
	}
}
