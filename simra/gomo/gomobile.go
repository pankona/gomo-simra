// +build darwin linux

package gomo

import (
	"github.com/pankona/gomo-simra/simra/peer"

	_ "image/jpeg" // must be imported here to treat jpeg
	_ "image/png"  // must be imported here to treat transparent of png

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/gl"
)

// Gomo represents gomobile instance.
// Singleton.
type Gomo struct {
	glPeer         *peer.GLPeer
	touchPeer      *peer.TouchPeer
	onStart        func()
	onStop         func()
	updateCallback func()
}

var gomo *Gomo

// GetInstance returns a Gomo instance.
// Since Gomo is singleton, it is necessary to
// call this function to get Gomo instance.
func GetInstance() *Gomo {
	peer.LogDebug("IN")
	if gomo == nil {
		gomo = &Gomo{}
	}
	peer.LogDebug("OUT")
	return gomo
}

// Initialize initializes Gomo.
func (gomo *Gomo) Initialize(onStart, onStop func(), updateCallback func()) {
	peer.LogDebug("IN")
	gomo.glPeer = peer.GetGLPeer()
	gomo.touchPeer = peer.GetTouchPeer()
	gomo.onStart = onStart
	gomo.onStop = onStop
	gomo.updateCallback = updateCallback
	peer.LogDebug("OUT")
}

// Start starts gomobile's main loop.
// Most of events handled by peer is fired by this function.
func (gomo *Gomo) Start() {
	peer.LogDebug("IN")
	app.Main(func(a app.App) {
		for e := range a.Events() {

			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:

					// initialize gl peer
					glctx, _ := e.DrawContext.(gl.Context)
					gomo.glPeer.Initialize(glctx)

					// time to set first scene
					gomo.onStart()
					a.Send(paint.Event{})
				case lifecycle.CrossOff:

					// time to stop application
					gomo.onStop()

					// finalize gl peer
					gomo.glPeer.Finalize()
				}
			case size.Event:
				peer.SetScreenSize(e)
			case paint.Event:
				if e.External {
					continue
				}

				// update notify for simra
				gomo.updateCallback()

				// update notify for gl peer
				gomo.glPeer.Update()

				a.Publish()
				a.Send(paint.Event{})
			case touch.Event:
				switch e.Type {
				case touch.TypeBegin:
					gomo.touchPeer.OnTouchBegin(e.X, e.Y)
				case touch.TypeMove:
					gomo.touchPeer.OnTouchMove(e.X, e.Y)
				case touch.TypeEnd:
					gomo.touchPeer.OnTouchEnd(e.X, e.Y)
				}
			}
		}
	})
	peer.LogDebug("OUT")
}
