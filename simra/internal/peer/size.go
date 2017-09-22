package peer

import (
	"github.com/pankona/gomo-simra/simra/simlog"
	"golang.org/x/mobile/event/size"
)

// ScreenSizer represents interface for configurating screen size
type ScreenSizer interface {
	// SetScreenSize sets device's actual screen size.
	// currently this function is called only from gomobile.go,
	// at updating window size, or orientating of screen.
	SetScreenSize(s size.Event)
	// SetDesiredScreenSize sets virtual screen size.
	// Any positive value can be specified to arguments.
	// like, w=1920, h=1080
	SetDesiredScreenSize(w, h float32)
}

const (
	fitHeight = iota
	fitWidth
)

type screenSize struct {
	width        float32
	height       float32
	scale        float32
	fitTo        int
	marginWidth  float32
	marginHeight float32
	sz           size.Event
}

var screensize = &screenSize{}

// GetScreenSizePeer returns an instance of ScreenSizer
func GetScreenSizePeer() ScreenSizer {
	return screensize
}

func (ss *screenSize) SetScreenSize(s size.Event) {
	simlog.FuncIn()
	ss.sz = s
	ss.calcScale()
	simlog.FuncOut()
}

// SetDesiredScreenSize sets virtual screen size.
// Any positive value can be specified to arguments.
// like, w=1920, h=1080
func (ss *screenSize) SetDesiredScreenSize(w, h float32) {
	simlog.FuncIn()
	ss.height = h
	ss.width = w
	ss.calcScale()
	simlog.FuncOut()
}

func (ss *screenSize) calcScale() {
	h := ss.height
	w := ss.width
	if h/float32(ss.sz.HeightPt) > w/float32(ss.sz.WidthPt) {
		ss.scale = float32(ss.sz.HeightPt) / h
		ss.fitTo = fitHeight
		ss.marginWidth = float32(ss.sz.WidthPt) - w*ss.scale
		ss.marginHeight = 0
	} else {
		ss.scale = float32(ss.sz.WidthPt) / w
		ss.fitTo = fitWidth
		ss.marginWidth = 0
		ss.marginHeight = float32(ss.sz.HeightPt) - h*ss.scale
	}
}
