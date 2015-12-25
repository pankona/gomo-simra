// +build darwin linux

package main

import (
	"github.com/pankona/gomo-simra/gomo"
	"github.com/pankona/gomo-simra/scene"
)

func main() {
	gomo := gomo.GetInstance()
	gomo.SetScene(&scene.InitScene{})
	gomo.Start()
}
