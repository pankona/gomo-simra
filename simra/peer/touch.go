package peer

import (
	"fmt"

	"golang.org/x/mobile/event/size"
)

type TouchPeer struct {
	touchListeners []*TouchListener
	sz             size.Event
}

var touchPeer *TouchPeer

type TouchListener interface {
	OnTouchBegin(x, y float32)
	OnTouchMove(x, y float32)
	OnTouchEnd(x, y float32)
}

func GetTouchPeer() *TouchPeer {
	LogDebug("IN")
	if touchPeer == nil {
		touchPeer = &TouchPeer{}
	}
	LogDebug("OUT")
	return touchPeer
}

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

func (touchpeer *TouchPeer) RemoveTouchListener(listener TouchListener) {
	LogDebug("IN")
	touchpeer.touchListeners = remove(touchpeer.touchListeners, &listener)
	LogDebug("OUT")
}

func (touchpeer *TouchPeer) RemoveAllTouchListener() {
	LogDebug("IN")
	touchpeer.touchListeners = nil
	LogDebug("OUT")
}

func calcTouchedPosition(pxx, pxy float32) (float32, float32) {
	ptx := pxx / sz.PixelsPerPt
	pty := pxy / sz.PixelsPerPt

	var scale float32
	if desiredScreenSize.fitTo == FIT_HEIGHT {
		scale = desiredScreenSize.height / float32(sz.HeightPt)
	} else {
		scale = desiredScreenSize.width / float32(sz.WidthPt)
	}

	return (ptx - desiredScreenSize.marginWidth/2) * scale,
		desiredScreenSize.height - (pty-desiredScreenSize.marginHeight/2)*scale
}

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
