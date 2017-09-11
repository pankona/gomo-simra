package peer

// Toucher reporesents an interface for touch controller
type Toucher interface {
	// AddTouchListener registeres a listener to notify touch event.
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

// AddTouchListener registeres a listener to notify touch event.
func (touchpeer *TouchPeer) AddTouchListener(listener TouchListener) {
	LogDebug("IN")
	touchpeer.touchListeners = append(touchpeer.touchListeners, listener)
	LogDebug("OUT")
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
func (touchpeer *TouchPeer) RemoveTouchListener(listener TouchListener) {
	LogDebug("IN")
	touchpeer.touchListeners = remove(touchpeer.touchListeners, listener)
	LogDebug("OUT")
}

// RemoveAllTouchListeners removes all registered listeners.
func (touchpeer *TouchPeer) RemoveAllTouchListeners() {
	LogDebug("IN")
	touchpeer.touchListeners = nil
	LogDebug("OUT")
}

func (touchpeer *TouchPeer) calcTouchedPosition(pxx, pxy float32) (float32, float32) {
	ptx := pxx / touchpeer.screensize.sz.PixelsPerPt
	pty := pxy / touchpeer.screensize.sz.PixelsPerPt

	var scale float32
	if touchpeer.screensize.fitTo == fitHeight {
		scale = touchpeer.screensize.height / float32(touchpeer.screensize.sz.HeightPt)
	} else {
		scale = touchpeer.screensize.width / float32(touchpeer.screensize.sz.WidthPt)
	}

	return (ptx - touchpeer.screensize.marginWidth/2) * scale,
		touchpeer.screensize.height - (pty-touchpeer.screensize.marginHeight/2)*scale
}

// OnTouchBegin is called when touch is started.
// This event is notified to all registered listeners despite of the touched position.
func (touchpeer *TouchPeer) OnTouchBegin(pxx, pxy float32) {
	LogDebug("IN")
	x, y := touchpeer.calcTouchedPosition(pxx, pxy)
	for i := range touchpeer.touchListeners {
		touchpeer.touchListeners[i].OnTouchBegin(x, y)
	}
	LogDebug("OUT")
}

// OnTouchMove is called when touch is moved (dragged).
// This event is notified to all registered listeners despite of the touched position.
func (touchpeer *TouchPeer) OnTouchMove(pxx, pxy float32) {
	LogDebug("IN")
	x, y := touchpeer.calcTouchedPosition(pxx, pxy)
	for i := range touchpeer.touchListeners {
		touchpeer.touchListeners[i].OnTouchMove(x, y)
	}
	LogDebug("OUT")
}

// OnTouchEnd is called when touch is ended (released).
// This event is notified to all registered listeners despite of the touched position.
func (touchpeer *TouchPeer) OnTouchEnd(pxx, pxy float32) {
	LogDebug("IN")
	x, y := touchpeer.calcTouchedPosition(pxx, pxy)
	for i := range touchpeer.touchListeners {
		touchpeer.touchListeners[i].OnTouchEnd(x, y)
	}
	LogDebug("OUT")
}
