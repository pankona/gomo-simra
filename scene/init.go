package scene

type InitScene struct {
	notifySceneEnd func(nextScene Driver)
}

var initScene *InitScene = nil

func GetInitScene() *InitScene {
	if initScene == nil {
		initScene = &InitScene{}
	}
	return initScene
}

func (self *InitScene) Initialize(sceneEndCallback func(nextScene Driver)) {
	//fmt.Println("[InitScene.Initialize] IN")
	self.notifySceneEnd = sceneEndCallback
}

func (self *InitScene) Drive() {
	//fmt.Println("[InitScene.Drive] IN")
	self.notifySceneEnd(&Title{}) // TODO: specify first scene
}
