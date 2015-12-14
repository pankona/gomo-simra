// +build darwin linux

package main

import (
	"github.com/pankona/gomobile_gamelib_test/peer"
	"github.com/pankona/gomobile_gamelib_test/scene"

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
		sceneCtrl := scene.GetControllerInstance()

		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
				case lifecycle.CrossOn:
					glctx, _ := e.DrawContext.(gl.Context)

					// initialize gl peer
					peer.Initialize(glctx)

					// initialize scene controller
					sceneCtrl.Initialize()

					// start scene controller
					sceneCtrl.Start()

					a.Send(paint.Event{})
				case lifecycle.CrossOff:
					sceneCtrl.Stop()
				}
			case size.Event:
				peer.SetScreenSize(e)
			case paint.Event:
				if e.External {
					continue
				}

				sceneCtrl.Update()

				a.Publish()
				a.Send(paint.Event{}) // keep animating
			case touch.Event:
				if e.Type == touch.TypeEnd {
					peer.OnTouch(e.X, e.Y)
				}
			}
		}
	})
}
