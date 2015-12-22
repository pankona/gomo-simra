package scene

import (
	"fmt"

	"github.com/pankona/gomo-simra/peer"
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
	self.current = &InitScene{}
}

func (self *Controller) onSceneEnd(nextScene Driver) {
	fmt.Println("[IN] callback function")
	peer.GetGLPeer().Reset()
	peer.GetTouchPeer().RemoveAllTouchListener()
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
		peer.GetGLPeer().Update()
	}
}

func (self *Controller) Stop() {
	self.current = nil
	peer.GetGLPeer().Finalize()
}
