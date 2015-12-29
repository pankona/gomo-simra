package simra

import (
	"github.com/pankona/gomo-simra/gomo"
	"github.com/pankona/gomo-simra/peer"
)

type Simra struct {
	driver Driver
}

var simra *Simra = nil

func GetInstance() *Simra {
	if simra == nil {
		simra = &Simra{}
	}
	return simra
}

func (self *Simra) onUpdate() {
	if self.driver != nil {
		self.driver.Drive()
	}
	peer.GetGLPeer().Update()
}

func (self *Simra) Start(startedCallback func()) {
	gomo.GetInstance().Initialize(self.onUpdate)
	gomo.GetInstance().Start(startedCallback)
}

func (self *Simra) SetScene(driver Driver) {
	self.driver = driver
	driver.Initialize()
}

func (self *Simra) Stop() {
}
