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

func (self *Gomo) Initialize(updateCallback func()) {
	peer.LogDebug("IN")
	self.glPeer = peer.GetGLPeer()
	self.touchPeer = peer.GetTouchPeer()
	self.updateCallback = updateCallback
	peer.LogDebug("OUT")
}

func (self *Gomo) Start(startedCallback func()) {
	peer.LogDebug("IN")
	go app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					glctx, _ := e.DrawContext.(gl.Context)

					// initialize gl peer
					self.glPeer.Initialize(glctx)
					startedCallback()

					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					// TODO: notify to simra
				}
			case size.Event:
				peer.SetScreenSize(e)
			case paint.Event:
				if e.External {
					continue
				}

				self.updateCallback()
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
				default:
				}
			}
		}
	})
	peer.LogDebug("OUT")
}
