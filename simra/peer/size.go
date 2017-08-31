package peer

import "golang.org/x/mobile/event/size"

var sz size.Event

const (
	// FitHeight indicates screen should fit to height length
	FitHeight = iota
	// FitWidth indicates screen should fit to width length
	FitWidth
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

// SetScreenSize sets screen size of device.
// This will be called only from gomobile.go.
func SetScreenSize(s size.Event) {
	LogDebug("IN")
	sz = s
	calcScale()
	LogDebug("OUT")
}

// GetScreenSize returns screen size of device.
// This value is set by SetScreenSize in advance.
func GetScreenSize() size.Event {
	LogDebug("IN")
	LogDebug("OUT")
	return sz
}

// SetDesiredScreenSize sets virtual screen size.
// Any positive value can be specified to arguments.
// like, w=1920, h=1080
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
		desiredScreenSize.fitTo = FitHeight
		desiredScreenSize.marginWidth = float32(sz.WidthPt) - w*desiredScreenSize.scale
		desiredScreenSize.marginHeight = 0
	} else {
		desiredScreenSize.scale = float32(sz.WidthPt) / w
		desiredScreenSize.fitTo = FitWidth
		desiredScreenSize.marginWidth = 0
		desiredScreenSize.marginHeight = float32(sz.HeightPt) - h*desiredScreenSize.scale
	}
	LogDebug("scale = %f", desiredScreenSize.scale)
}
