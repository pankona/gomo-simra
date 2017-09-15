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
	Initialize(onStart func(glc *GLContext), onStop func(), updateCallback func())
	// Start starts gomobile's main loop.
	// Most of events handled by peer is fired by this function.
	Start()
}

// Gomo represents gomobile instance.
// Singleton.
type Gomo struct {
	app            app.App
	screensize     ScreenSizer
	touch          Toucher
	onStart        func(glc *GLContext)
	onStop         func()
	updateCallback func()
}

// GetGomo returns a Gomo instance.
func GetGomo() Gomoer {
	return &Gomo{
		screensize: GetScreenSizePeer(),
		touch:      GetTouchPeer(),
	}
}

// Initialize initializes Gomo.
func (g *Gomo) Initialize(onStart func(glc *GLContext), onStop func(), updateCallback func()) {
	LogDebug("IN")
	g.onStart = onStart
	g.onStop = onStop
	g.updateCallback = updateCallback
	g.screensize = screensize
	LogDebug("OUT")
}

func (g *Gomo) handleLifeCycle(e lifecycle.Event) {
	switch e.Crosses(lifecycle.StageVisible) {
	case lifecycle.CrossOn:
		g.onStart(&GLContext{
			glcontext: e.DrawContext.(gl.Context),
			publish:   g.app.Publish,
		})
		g.app.Send(paint.Event{})
	case lifecycle.CrossOff:
		g.onStop()
	}

}

func (g *Gomo) handleSize(e size.Event) {
	g.screensize.SetScreenSize(e)
}

func (g *Gomo) handlePaint(e paint.Event) {
	if e.External {
		return
	}
	// update notify for simra
	g.updateCallback()
	// update notify for gl peer
	g.app.Send(paint.Event{})
}

func (g *Gomo) handleTouch(e touch.Event) {
	switch e.Type {
	case touch.TypeBegin:
		g.touch.OnTouchBegin(e.X, e.Y)
	case touch.TypeMove:
		g.touch.OnTouchMove(e.X, e.Y)
	case touch.TypeEnd:
		g.touch.OnTouchEnd(e.X, e.Y)
	}
}

func (g *Gomo) handleEvent(e interface{}) {
	switch e := g.app.Filter(e).(type) {
	case lifecycle.Event:
		g.handleLifeCycle(e)
	case size.Event:
		g.handleSize(e)
	case paint.Event:
		g.handlePaint(e)
	case touch.Event:
		g.handleTouch(e)
	}
}

// Start starts gomobile's main loop.
// Most of events handled by peer is fired by this function.
func (g *Gomo) Start() {
	LogDebug("IN")
	app.Main(func(a app.App) {
		g.app = a
		for e := range a.Events() {
			g.handleEvent(e)
		}
	})
	LogDebug("OUT")
}
