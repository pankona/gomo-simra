// +build darwin linux

package main

import (
	"fmt"
	"time"

	"github.com/pankona/gomo-simra/simra"
)

func onStarted() {
	fmt.Println("[IN] onStarted()")
	engine := simra.GetInstance()
	engine.SetScene(&simra.Title{})
	fmt.Println("[OUT] onStarted()")
}

func main() {
	fmt.Println("[IN] main()")
	engine := simra.GetInstance()
	engine.Start(onStarted)
	for {
		time.Sleep(10)
	}
	fmt.Println("[OUT] main()")
}
