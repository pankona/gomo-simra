package peer

import (
	"fmt"

	"golang.org/x/mobile/event/size"
)

type Touch struct {
	touchListeners []*TouchListener
	sz             size.Event
}

var touchPeer *Touch

type TouchListener interface {
	OnTouchBegin(x, y float32)
	OnTouchMove(x, y float32)
	OnTouchEnd(x, y float32)
}

func GetTouchPeer() *Touch {
	if touchPeer == nil {
		touchPeer = &Touch{}
	}
	return touchPeer
}

func (self *Touch) AddTouchListener(listener TouchListener) {
	self.touchListeners = append(self.touchListeners, &listener)
}

func (self *Touch) RemoveAllTouchListener() {
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

func (self *Touch) OnTouchBegin(pxx, pxy float32) {

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

func (self *Touch) OnTouchMove(pxx, pxy float32) {

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

func (self *Touch) OnTouchEnd(pxx, pxy float32) {

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
