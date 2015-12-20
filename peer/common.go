package peer

import (
	"fmt"

	"golang.org/x/mobile/event/size"
)

var sz size.Event

const (
	FIT_HEIGHT = iota
	FIT_WIDTH
)

type screenSize struct {
	width  float32
	height float32
	scale  float32
	fitTo  int
}

var desiredScreenSize screenSize

func SetScreenSize(in_sz size.Event) {
	sz = in_sz
	calcScale()
}

func GetScreenSize() size.Event {
	return sz
}

func SetDesiredScreenSize(w, h float32) {
	desiredScreenSize.height = h
	desiredScreenSize.width = w
	calcScale()
}

func calcScale() {
	h := desiredScreenSize.height
	w := desiredScreenSize.width

	if h/float32(sz.HeightPt) > w/float32(sz.WidthPt) {
		desiredScreenSize.scale = float32(sz.HeightPt) / h
		desiredScreenSize.fitTo = FIT_HEIGHT
		fmt.Println("scale = ", desiredScreenSize.scale)
	} else {
		desiredScreenSize.scale = float32(sz.WidthPt) / w
		desiredScreenSize.fitTo = FIT_WIDTH
		fmt.Println("scale = ", desiredScreenSize.scale)
	}
}
