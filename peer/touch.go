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
	if touchPeer == nil {
		touchPeer = &TouchPeer{}
	}
	return touchPeer
}

func (self *TouchPeer) AddTouchListener(listener TouchListener) {
	self.touchListeners = append(self.touchListeners, &listener)
}

func (self *TouchPeer) RemoveAllTouchListener() {
	self.touchListeners = nil
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

	return ptx * scale, pty * scale
}

func (self *TouchPeer) OnTouchBegin(pxx, pxy float32) {

	x, y := calcTouchedPosition(pxx, pxy)

	for i := range self.touchListeners {
		listener := self.touchListeners[i]
		if listener == nil {
			fmt.Println("listener is nil!")
			continue
		}

		(*listener).OnTouchBegin(x, y)
	}
}

func (self *TouchPeer) OnTouchMove(pxx, pxy float32) {

	x, y := calcTouchedPosition(pxx, pxy)

	for i := range self.touchListeners {
		listener := self.touchListeners[i]
		if listener == nil {
			fmt.Println("listener is nil!")
			continue
		}

		(*listener).OnTouchMove(x, y)
	}
}

func (self *TouchPeer) OnTouchEnd(pxx, pxy float32) {

	x, y := calcTouchedPosition(pxx, pxy)

	for i := range self.touchListeners {
		listener := self.touchListeners[i]
		if listener == nil {
			fmt.Println("listener is nil!")
			continue
		}

		(*listener).OnTouchEnd(x, y)
	}
}
