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
	Initialize(onStart func(glc *GLContext), onStop func(), updateCallback func(interface{}))
	// Start starts gomobile's main loop.
	// Most of events handled by peer is fired by this function.
	Start()
}

// Gomo represents gomobile instance.
// Singleton.
type Gomo struct {
	screensize     ScreenSizer
	touch          Toucher
	onStart        func(glc *GLContext)
	onStop         func()
	updateCallback func(interface{})
}

// GetGomo returns a Gomo instance.
func GetGomo() Gomoer {
	return &Gomo{
		screensize: GetScreenSizePeer(),
		touch:      GetTouchPeer(),
	}
}

// Initialize initializes Gomo.
func (g *Gomo) Initialize(onStart func(glc *GLContext), onStop func(), updateCallback func(i interface{})) {
	LogDebug("IN")
	g.onStart = onStart
	g.onStop = onStop
	g.updateCallback = updateCallback
	g.screensize = screensize
	LogDebug("OUT")
}

func (g *Gomo) handleLifeCycle(a app.App, e lifecycle.Event) {
	switch e.Crosses(lifecycle.StageVisible) {
	case lifecycle.CrossOn:
		// time to set first scene
		g.onStart(&GLContext{
			glcontext: e.DrawContext.(gl.Context),
		})
		a.Send(paint.Event{})
	case lifecycle.CrossOff:
		// time to stop application
		g.onStop()
	}

}

func (g *Gomo) handleSize(a app.App, e size.Event) {
	g.screensize.SetScreenSize(e)
}

func (g *Gomo) handlePaint(a app.App, e paint.Event) {
	if e.External {
		return
	}
	// update notify for simra
	g.updateCallback(a.Publish)
	// update notify for gl peer
	a.Send(paint.Event{})
}

func (g *Gomo) handleTouch(a app.App, e touch.Event) {
	switch e.Type {
	case touch.TypeBegin:
		g.touch.OnTouchBegin(e.X, e.Y)
	case touch.TypeMove:
		g.touch.OnTouchMove(e.X, e.Y)
	case touch.TypeEnd:
		g.touch.OnTouchEnd(e.X, e.Y)
	}
}

func (g *Gomo) handleEvent(a app.App, e interface{}) {
	switch e := a.Filter(e).(type) {
	case lifecycle.Event:
		g.handleLifeCycle(a, e)
	case size.Event:
		g.handleSize(a, e)
	case paint.Event:
		g.handlePaint(a, e)
	case touch.Event:
		g.handleTouch(a, e)
	}
}

// Start starts gomobile's main loop.
// Most of events handled by peer is fired by this function.
func (g *Gomo) Start() {
	LogDebug("IN")
	app.Main(func(a app.App) {
		for e := range a.Events() {
			g.handleEvent(a, e)
		}
	})
	LogDebug("OUT")
}
