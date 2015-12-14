package scene

type InitScene struct {
	notifySceneEnd func(nextScene Driver)
}

func (self *InitScene) Initialize(sceneEndCallback func(nextScene Driver)) {
	self.notifySceneEnd = sceneEndCallback
}

func (self *InitScene) Drive() {
	self.notifySceneEnd(&Title{}) // TODO: specify first scene from other package
}

func (self *InitScene) OnTouch(x, y float32) {
}
