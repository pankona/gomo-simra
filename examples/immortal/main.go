// +build darwin linux

package main

import (
	"github.com/pankona/gomo-simra/examples/immortal/scene"
	"github.com/pankona/gomo-simra/simra"
)

func eventHandle(onStart, onStop chan bool) {
	for {
		select {
		case <-onStart:
			simra.LogDebug("receive chan. onStart")
			engine := simra.GetInstance()
			// TODO: this will be called on rotation.
			// to keep state on rotation, SetScene must not call
			// every onStart.
			engine.SetScene(&scene.Title{})
		case <-onStop:
			simra.LogDebug("receive chan. onStop")
		}
	}
}

func main() {
	simra.LogDebug("[IN]")
	engine := simra.GetInstance()

	onStart := make(chan bool)
	onStop := make(chan bool)
	go eventHandle(onStart, onStop)
	engine.Start(onStart, onStop)
	simra.LogDebug("[OUT]")
}
