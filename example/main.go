// +build darwin linux

package main

import (
	"github.com/pankona/gomo-simra/example/scene"
	"github.com/pankona/gomo-simra/peer"
	"github.com/pankona/gomo-simra/simra"
)

func main() {
	peer.LogDebug("[IN]")
	engine := simra.GetInstance()

	onStart := make(chan bool)
	onStop := make(chan bool)

	engine.Start(onStart, onStop)

	for {
	loop:
		select {
		case <-onStart:
			engine := simra.GetInstance()
			engine.SetScene(&scene.Title{})
		case <-onStop:
			break loop
		}
	}
	peer.LogDebug("[OUT]")
}
