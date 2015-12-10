// +build darwin linux

package main

import (
	"github.com/pankona/gomobile_gamelib_test/peer"

	_ "image/jpeg"
	_ "image/png"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/gl"
)

func main() {
	app.Main(func(a app.App) {
		peer := peer.GetInstance()

		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					glctx, _ := e.DrawContext.(gl.Context)
					peer.Initialize(glctx)
					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					peer.Stop()
				}
			case size.Event:
				peer.SetScreenSize(e)
			case paint.Event:
				if e.External {
					continue
				}

				peer.Update()

				a.Publish()
				a.Send(paint.Event{}) // keep animating
			case touch.Event:
			}
		}
	})
}
