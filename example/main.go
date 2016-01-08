// +build darwin linux

package main

import (
	"github.com/pankona/gomo-simra/example/scene"
	"github.com/pankona/gomo-simra/peer"
	"github.com/pankona/gomo-simra/simra"
)

func eventHandle(onStart, onStop chan bool) {
	for {
		select {
		case <-onStart:
			peer.LogDebug("receive chan. onStart")
			engine := simra.GetInstance()
			// TODO: this will be called on rotation.
			// to keep state on rotation, SetScene must not call
			// every onStart.
			engine.SetScene(&scene.Title{})
		case <-onStop:
			peer.LogDebug("receive chan. onStop")
		}
	}
}

func main() {
	peer.LogDebug("[IN]")
	engine := simra.GetInstance()

	onStart := make(chan bool)
	onStop := make(chan bool)
	go eventHandle(onStart, onStop)
	engine.Start(onStart, onStop)
	peer.LogDebug("[OUT]")
}
