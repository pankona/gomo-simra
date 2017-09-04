package peer

import (
	"fmt"
)

// TouchPeer represents a Touch object.
// Singleton.
type TouchPeer struct {
	touchListeners []*TouchListener
}

var touchPeer *TouchPeer

// TouchListener is interface to be notifed touch event.
type TouchListener interface {
	OnTouchBegin(x, y float32)
	OnTouchMove(x, y float32)
	OnTouchEnd(x, y float32)
}

// GetTouchPeer returns instance of TouchPeer.
// Since TouchPeer is singleton, it is necessary to
// call this function to get instance of TouchPeer.
func GetTouchPeer() *TouchPeer {
	LogDebug("IN")
	if touchPeer == nil {
		touchPeer = &TouchPeer{}
	}
	LogDebug("OUT")
	return touchPeer
}

// AddTouchListener registeres a listener to notify touch event.
func (touchpeer *TouchPeer) AddTouchListener(listener TouchListener) {
	LogDebug("IN")
	touchpeer.touchListeners = append(touchpeer.touchListeners, &listener)
	LogDebug("OUT")
}

func remove(listeners []*TouchListener, remove *TouchListener) []*TouchListener {
	result := []*TouchListener{}

	for _, listener := range listeners {
		if listener != remove {
			result = append(result, listener)
		}
	}
	return result
}

// RemoveTouchListener removes specified listener.
func (touchpeer *TouchPeer) RemoveTouchListener(listener TouchListener) {
	LogDebug("IN")
	touchpeer.touchListeners = remove(touchpeer.touchListeners, &listener)
	LogDebug("OUT")
}

// RemoveAllTouchListener removes all registered listeners.
func (touchpeer *TouchPeer) RemoveAllTouchListener() {
	LogDebug("IN")
	touchpeer.touchListeners = nil
	LogDebug("OUT")
}

func calcTouchedPosition(pxx, pxy float32) (float32, float32) {
	ptx := pxx / screensize.sz.PixelsPerPt
	pty := pxy / screensize.sz.PixelsPerPt

	var scale float32
	if screensize.fitTo == FitHeight {
		scale = screensize.height / float32(screensize.sz.HeightPt)
	} else {
		scale = screensize.width / float32(screensize.sz.WidthPt)
	}

	return (ptx - screensize.marginWidth/2) * scale,
		screensize.height - (pty-screensize.marginHeight/2)*scale
}

// OnTouchBegin is called when touch is started.
// This event is notified to all registered listeners despite of the touched position.
func (touchpeer *TouchPeer) OnTouchBegin(pxx, pxy float32) {
	LogDebug("IN")

	x, y := calcTouchedPosition(pxx, pxy)
	for i := range touchpeer.touchListeners {
		listener := touchpeer.touchListeners[i]
		if listener == nil {
			fmt.Println("listener is nil!")
			continue
		}

		(*listener).OnTouchBegin(x, y)
	}
	LogDebug("OUT")
}

// OnTouchMove is called when touch is moved (dragged).
// This event is notified to all registered listeners despite of the touched position.
func (touchpeer *TouchPeer) OnTouchMove(pxx, pxy float32) {
	LogDebug("IN")

	x, y := calcTouchedPosition(pxx, pxy)

	for i := range touchpeer.touchListeners {
		listener := touchpeer.touchListeners[i]
		if listener == nil {
			fmt.Println("listener is nil!")
			continue
		}

		(*listener).OnTouchMove(x, y)
	}
	LogDebug("OUT")
}

// OnTouchEnd is called when touch is ended (released).
// This event is notified to all registered listeners despite of the touched position.
func (touchpeer *TouchPeer) OnTouchEnd(pxx, pxy float32) {
	LogDebug("IN")

	x, y := calcTouchedPosition(pxx, pxy)

	for i := range touchpeer.touchListeners {
		listener := touchpeer.touchListeners[i]
		if listener == nil {
			fmt.Println("listener is nil!")
			continue
		}

		(*listener).OnTouchEnd(x, y)
	}
	LogDebug("OUT")
}
