// +build darwin linux

package gomo

import (
	"fmt"

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
	if gomo == nil {
		gomo = &Gomo{}
	}
	return gomo
}

func (self *Gomo) Initialize(updateCallback func()) {
	fmt.Println("[IN] gomo.initialize")
	self.glPeer = peer.GetGLPeer()
	self.touchPeer = peer.GetTouchPeer()
	self.updateCallback = updateCallback
	fmt.Println("[OUT] gomo.initialize")
}

func (self *Gomo) Start(startedCallback func()) {
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
					// sceneCtrl.Stop()
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
				a.Send(paint.Event{}) // keep animating
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
}
