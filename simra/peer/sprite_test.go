package peer

import "testing"

func TestGetters(t *testing.T) {
	s := &Sprite{
		W: 10,
		H: 20,
		X: 30,
		Y: 40,
		R: 50,
	}
	if s.GetWidth() != 10 {
		t.Errorf("unexpected result. [got] %d [want] %d", s.GetWidth(), 10)
	}
	if s.GetHeight() != 20 {
		t.Errorf("unexpected result. [got] %d [want] %d", s.GetHeight(), 10)
	}
	if s.GetX() != 30 {
		t.Errorf("unexpected result. [got] %d [want] %d", s.GetX(), 10)
	}
	if s.GetY() != 40 {
		t.Errorf("unexpected result. [got] %d [want] %d", s.GetY(), 10)
	}
}

func TestAddRemoveTouchListener(t *testing.T) {
	s := &Sprite{}

	// Add listeners
	for i := 0; i < 5; i++ {
		s.AddTouchListener(&listener{})
		if len(s.touchListeners) != i+1 {
			t.Errorf("unexpected result. [got] %d [want] %d", len(s.touchListeners), i+1)
		}
	}

	// Remove all listeners
	s.RemoveAllTouchListener()
	if len(s.touchListeners) != 0 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(s.touchListeners), 0)
	}

	// Add listeners again after remove
	for i := 0; i < 5; i++ {
		s.AddTouchListener(&listener{})
		if len(s.touchListeners) != i+1 {
			t.Errorf("unexpected result. [got] %d [want] %d", len(s.touchListeners), i+1)
		}
	}
}
