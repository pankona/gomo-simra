package peer

import "testing"

func newTestTouchPeer() *TouchPeer {
	return &TouchPeer{
		screensize: &screenSize{},
	}
}

func TestGetTouchPeer(t *testing.T) {
	t1 := GetTouchPeer()
	t2 := GetTouchPeer()

	if t1 != t2 {
		t.Error("unexpected result .GetTouchPeer should return same address")
	}
}

func TestAddTouchListener(t *testing.T) {
	touch := newTestTouchPeer()
	for i := 0; i < 10; i++ {
		touch.AddTouchListener(&listener{})
		if len(touch.touchListeners) != i+1 {
			t.Errorf("unexpected result. [got] %d [want] %d", len(touch.touchListeners), i+1)
		}
	}
}

func TestRemoveAllTouchListeners(t *testing.T) {
	touch := newTestTouchPeer()
	for i := 0; i < 10; i++ {
		touch.AddTouchListener(&listener{})
	}
	if len(touch.touchListeners) != 10 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(touch.touchListeners), 10)
	}

	touch.RemoveAllTouchListeners()
	if len(touch.touchListeners) != 0 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(touch.touchListeners), 0)
	}
}

func TestRemoveTouchListner(t *testing.T) {
	touch := newTestTouchPeer()
	touches := make([]*listener, 10)
	for i := 0; i < 10; i++ {
		touches[i] = &listener{}
		touch.AddTouchListener(touches[i])
	}
	if len(touch.touchListeners) != 10 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(touch.touchListeners), 10)
	}
	for i := 0; i < 10; i++ {
		touch.RemoveTouchListener(touches[i])
		if len(touch.touchListeners) != 10-(i+1) {
			t.Errorf("unexpected result. [got] %d [want] %d", len(touch.touchListeners), 10-(i+1))
		}
	}
}

func TestTouchEvents(t *testing.T) {
	touch := newTestTouchPeer()
	l := &listener{
		touchBegin: func(x, y float32) {},
		touchMove:  func(x, y float32) {},
		touchEnd:   func(x, y float32) {},
	}
	touch.AddTouchListener(l)
	touch.OnTouchBegin(0, 0)
	touch.OnTouchMove(0, 0)
	touch.OnTouchEnd(0, 0)
}
