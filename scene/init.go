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

func (self *InitScene) OnTouchBegin(x, y float32) {
}

func (self *InitScene) OnTouchMove(x, y float32) {
}

func (self *InitScene) OnTouchEnd(x, y float32) {
}
