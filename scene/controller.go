package scene

import (
	"fmt"

	"github.com/pankona/gomobile_gamelib_test/peer"
)

type Controller struct {
	current Driver
}

var controller *Controller = nil

func GetControllerInstance() *Controller {
	if controller == nil {
		controller = &Controller{}
	}
	return controller
}

func (self *Controller) Initialize() {
	self.current = GetInitScene()
}

func (self *Controller) onSceneEnd(nextScene Driver) {
	fmt.Println("[IN] callback function")
	self.current = nextScene
	self.current.Initialize(self.onSceneEnd)
}

func (self *Controller) Start() {
	if self.current != nil {
		self.current.Initialize(self.onSceneEnd)
	}
}

func (self *Controller) Update() {
	if self.current != nil {
		self.current.Drive()
		peer.GetInstance().Update()
	}
}

func (self *Controller) Stop() {
	self.current = nil
	peer.GetInstance().Finalize()
}
