// +build darwin linux

package peer

import (
	_ "image/jpeg" // must be imported here to treat jpeg
	_ "image/png"  // must be imported here to treat transparent of png

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/gl"
)

// Gomoer represents an interface of gomobile
type Gomoer interface {
	// Initialize initializes Gomo.
	Initialize(onStart, onStop func(), updateCallback func())
	// Start starts gomobile's main loop.
	// Most of events handled by peer is fired by this function.
	Start()
}

// Gomo represents gomobile instance.
// Singleton.
type Gomo struct {
	onStart        func()
	onStop         func()
	updateCallback func()
	screenSize     ScreenSizer
}

var gomo *Gomo

// GetGomo returns a Gomo instance.
// Since Gomo is singleton, it is necessary to
// call this function to get Gomo instance.
func GetGomo() Gomoer {
	LogDebug("IN")
	if gomo == nil {
		gomo = &Gomo{}
	}
	LogDebug("OUT")
	return gomo
}

// Initialize initializes Gomo.
func (gomo *Gomo) Initialize(onStart, onStop func(), updateCallback func()) {
	LogDebug("IN")
	gomo.onStart = onStart
	gomo.onStop = onStop
	gomo.updateCallback = updateCallback
	gomo.screenSize = screensize
	LogDebug("OUT")
}

func handleLifeCycle(a app.App, e lifecycle.Event) {
	switch e.Crosses(lifecycle.StageVisible) {
	case lifecycle.CrossOn:

		// initialize gl peer
		glctx, _ := e.DrawContext.(gl.Context)
		glPeer.Initialize(glctx)

		// time to set first scene
		gomo.onStart()
		a.Send(paint.Event{})
	case lifecycle.CrossOff:

		// time to stop application
		gomo.onStop()

		// finalize gl peer
		glPeer.Finalize()
	}

}

func handleSize(a app.App, e size.Event) {
	screensize.SetScreenSize(e)
}

func handlePaint(a app.App, e paint.Event) {
	if e.External {
		return
	}
	// update notify for simra
	gomo.updateCallback()
	// update notify for gl peer
	glPeer.Update(a.Publish)
	a.Send(paint.Event{})
}

func handleTouch(a app.App, e touch.Event) {
	switch e.Type {
	case touch.TypeBegin:
		touchPeer.OnTouchBegin(e.X, e.Y)
	case touch.TypeMove:
		touchPeer.OnTouchMove(e.X, e.Y)
	case touch.TypeEnd:
		touchPeer.OnTouchEnd(e.X, e.Y)
	}
}

func handleEvent(a app.App, e interface{}) {
	switch e := a.Filter(e).(type) {
	case lifecycle.Event:
		handleLifeCycle(a, e)
	case size.Event:
		handleSize(a, e)
	case paint.Event:
		handlePaint(a, e)
	case touch.Event:
		handleTouch(a, e)
	}
}

// Start starts gomobile's main loop.
// Most of events handled by peer is fired by this function.
func (gomo *Gomo) Start() {
	LogDebug("IN")
	app.Main(func(a app.App) {
		for e := range a.Events() {
			handleEvent(a, e)
		}
	})
	LogDebug("OUT")
}
