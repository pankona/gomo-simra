package peer

import "github.com/pankona/gomo-simra/simra/simlog"

// Toucher reporesents an interface for touch controller
type Toucher interface {
	// AddTouchListener registers a listener to notify touch event.
	AddTouchListener(listener TouchListener)
	// RemoveTouchListener removes specified listener.
	RemoveTouchListener(listener TouchListener)
	// RemoveAllTouchListener removes all registered listeners.
	RemoveAllTouchListeners()
	// OnTouchBegin is called when touch is started.
	// This event is notified to all registered listeners despite of the touched position.
	OnTouchBegin(pxx, pxy float32)
	// OnTouchMove is called when touch is moved (dragged).
	// This event is notified to all registered listeners despite of the touched position.
	OnTouchMove(pxx, pxy float32)
	// OnTouchEnd is called when touch is ended (released).
	// This event is notified to all registered listeners despite of the touched position.
	OnTouchEnd(pxx, pxy float32)
}

// TouchPeer represents a Touch object.
// Singleton.
type TouchPeer struct {
	screensize     *screenSize
	touchListeners []TouchListener
}

var touchPeer = &TouchPeer{}

// TouchListener is interface to be notifed touch event.
type TouchListener interface {
	OnTouchBegin(x, y float32)
	OnTouchMove(x, y float32)
	OnTouchEnd(x, y float32)
}

// GetTouchPeer returns instance of TouchPeer.
// Since TouchPeer is singleton, it is necessary to
// call this function to get instance of TouchPeer.
func GetTouchPeer() Toucher {
	touchPeer.screensize = screensize
	return touchPeer
}

// AddTouchListener registers a listener to notify touch event.
func (tp *TouchPeer) AddTouchListener(listener TouchListener) {
	simlog.FuncIn()
	tp.touchListeners = append(tp.touchListeners, listener)
	simlog.FuncOut()
}

func remove(listeners []TouchListener, remove TouchListener) []TouchListener {
	result := []TouchListener{}
	for _, listener := range listeners {
		if listener != remove {
			result = append(result, listener)
		}
	}
	return result
}

// RemoveTouchListener removes specified listener.
func (tp *TouchPeer) RemoveTouchListener(listener TouchListener) {
	simlog.FuncIn()
	tp.touchListeners = remove(tp.touchListeners, listener)
	simlog.FuncOut()
}

// RemoveAllTouchListeners removes all registered listeners.
func (tp *TouchPeer) RemoveAllTouchListeners() {
	simlog.FuncIn()
	tp.touchListeners = nil
	simlog.FuncOut()
}

func (tp *TouchPeer) calcTouchedPosition(pxx, pxy float32) (float32, float32) {
	ptx := pxx / tp.screensize.sz.PixelsPerPt
	pty := pxy / tp.screensize.sz.PixelsPerPt

	var scale float32
	if tp.screensize.fitTo == fitHeight {
		scale = tp.screensize.height / float32(tp.screensize.sz.HeightPt)
	} else {
		scale = tp.screensize.width / float32(tp.screensize.sz.WidthPt)
	}

	return (ptx - tp.screensize.marginWidth/2) * scale,
		tp.screensize.height - (pty-tp.screensize.marginHeight/2)*scale
}

// OnTouchBegin is called when touch is started.
// This event is notified to all registered listeners despite of the touched position.
func (tp *TouchPeer) OnTouchBegin(pxx, pxy float32) {
	simlog.FuncIn()
	x, y := tp.calcTouchedPosition(pxx, pxy)
	for i := range tp.touchListeners {
		tp.touchListeners[i].OnTouchBegin(x, y)
	}
	simlog.FuncOut()
}

// OnTouchMove is called when touch is moved (dragged).
// This event is notified to all registered listeners despite of the touched position.
func (tp *TouchPeer) OnTouchMove(pxx, pxy float32) {
	simlog.FuncIn()
	x, y := tp.calcTouchedPosition(pxx, pxy)
	for i := range tp.touchListeners {
		tp.touchListeners[i].OnTouchMove(x, y)
	}
	simlog.FuncOut()
}

// OnTouchEnd is called when touch is ended (released).
// This event is notified to all registered listeners despite of the touched position.
func (tp *TouchPeer) OnTouchEnd(pxx, pxy float32) {
	simlog.FuncIn()
	x, y := tp.calcTouchedPosition(pxx, pxy)
	for i := range tp.touchListeners {
		tp.touchListeners[i].OnTouchEnd(x, y)
	}
	simlog.FuncOut()
}
