// +build darwin linux

package main

import "github.com/pankona/gomo-simra/simra"

func main() {
	engine := simra.GetInstance()

	engine.Initialize()
	engine.Start(&simra.Title{})
}
