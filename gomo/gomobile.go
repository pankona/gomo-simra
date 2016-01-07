// +build darwin linux

package gomo

import (
	"github.com/pankona/gomo-simra/peer"

	_ "image/jpeg"
	_ "image/png"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/gl"
)

type Gomo struct {
	glPeer         *peer.GLPeer
	touchPeer      *peer.TouchPeer
	onStart        chan bool
	onStop         chan bool
	updateCallback func()
}

var gomo *Gomo = nil

func GetInstance() *Gomo {
	peer.LogDebug("IN")
	if gomo == nil {
		gomo = &Gomo{}
	}
	peer.LogDebug("OUT")
	return gomo
}

func (self *Gomo) Initialize(onStart, onStop chan bool, updateCallback func()) {
	peer.LogDebug("IN")
	self.glPeer = peer.GetGLPeer()
	self.touchPeer = peer.GetTouchPeer()
	self.onStart = onStart
	self.onStop = onStop
	self.updateCallback = updateCallback
	peer.LogDebug("OUT")
}

func (self *Gomo) Start() {
	peer.LogDebug("IN")
	app.Main(func(a app.App) {
		for e := range a.Events() {

			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:

					// initialize gl peer
					glctx, _ := e.DrawContext.(gl.Context)
					self.glPeer.Initialize(glctx)

					// time to set first scene
					self.onStart <- true
					a.Send(paint.Event{})
				case lifecycle.CrossOff:

					// time to stop application
					self.onStop <- true

					// finalize gl peer
					self.glPeer.Finalize()
				}
			case size.Event:
				peer.SetScreenSize(e)
			case paint.Event:
				if e.External {
					continue
				}

				// update notify for simra
				self.updateCallback()

				// update notify for gl peer
				self.glPeer.Update()

				a.Publish()
				a.Send(paint.Event{})
			case touch.Event:
				switch e.Type {
				case touch.TypeBegin:
					self.touchPeer.OnTouchBegin(e.X, e.Y)
				case touch.TypeMove:
					self.touchPeer.OnTouchMove(e.X, e.Y)
				case touch.TypeEnd:
					self.touchPeer.OnTouchEnd(e.X, e.Y)
				}
			}
		}
	})
	peer.LogDebug("OUT")
}
