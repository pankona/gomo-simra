package peer

import "golang.org/x/mobile/event/size"

var sz size.Event

const (
	FIT_HEIGHT = iota
	FIT_WIDTH
)

type screenSize struct {
	width        float32
	height       float32
	scale        float32
	fitTo        int
	marginWidth  float32
	marginHeight float32
}

var desiredScreenSize screenSize

func SetScreenSize(in_sz size.Event) {
	LogDebug("IN")
	sz = in_sz
	calcScale()
	LogDebug("OUT")
}

func GetScreenSize() size.Event {
	LogDebug("IN")
	LogDebug("OUT")
	return sz
}

func SetDesiredScreenSize(w, h float32) {
	LogDebug("IN")
	desiredScreenSize.height = h
	desiredScreenSize.width = w
	calcScale()
	LogDebug("OUT")
}

func calcScale() {
	h := desiredScreenSize.height
	w := desiredScreenSize.width

	if h/float32(sz.HeightPt) > w/float32(sz.WidthPt) {
		desiredScreenSize.scale = float32(sz.HeightPt) / h
		desiredScreenSize.fitTo = FIT_HEIGHT
		desiredScreenSize.marginWidth = float32(sz.WidthPt) - w*desiredScreenSize.scale
		desiredScreenSize.marginHeight = 0

	} else {
		desiredScreenSize.scale = float32(sz.WidthPt) / w
		desiredScreenSize.fitTo = FIT_WIDTH
		desiredScreenSize.marginWidth = 0
		desiredScreenSize.marginHeight = float32(sz.HeightPt) - h*desiredScreenSize.scale
	}
	LogDebug("scale = %f", desiredScreenSize.scale)
}
