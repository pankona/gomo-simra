package scene

import "fmt"

type InitScene struct {
}

var initScene *InitScene = nil

func GetInitScene() *InitScene {
	if initScene == nil {
		initScene = &InitScene{}
	}
	return initScene
}

func (self *InitScene) Initialize() {
	fmt.Println("[InitScene.Initialize] IN")
}

func (self *InitScene) Drive() {
	fmt.Println("[InitScene.Drive] IN")
}
