package peer

import "testing"

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
