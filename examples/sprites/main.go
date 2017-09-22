// +build darwin linux

package main

import (
	"github.com/pankona/gomo-simra/examples/sprites/scene"
	"github.com/pankona/gomo-simra/simra"
)

func main() {
	simra.LogDebug("[IN]")
	engine := simra.GetInstance()
	engine.Start(&scene.Title{})
	simra.LogDebug("[OUT]")
}
