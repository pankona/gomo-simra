package scene

import "fmt"

type InitScene struct {
}

var initScene *InitScene = nil
var sceneEndCallback func(d Driver) = nil

func GetInitScene() *InitScene {
	if initScene == nil {
		initScene = &InitScene{}
	}
	return initScene
}

func (self *InitScene) Initialize(cb func(d Driver)) {
	fmt.Println("[InitScene.Initialize] IN")
	sceneEndCallback = cb
}

func (self *InitScene) Drive() {
	fmt.Println("[InitScene.Drive] IN")
	sceneEndCallback(nil) // TODO: specify first scene
}
