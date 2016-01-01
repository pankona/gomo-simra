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
	peer.LogDebug("IN")
	if simra == nil {
		simra = &Simra{}
	}
	peer.LogDebug("OUT")
	return simra
}

func (self *Simra) onUpdate() {
	if self.driver != nil {
		self.driver.Drive()
	}
}

func (self *Simra) Start(startedCallback func()) {
	peer.LogDebug("IN")
	gomo.GetInstance().Initialize(self.onUpdate)
	gomo.GetInstance().Start(startedCallback)
	peer.LogDebug("OUT")
}

func (self *Simra) SetScene(driver Driver) {
	peer.LogDebug("IN")
	peer.GetGLPeer().Reset()
	peer.GetTouchPeer().RemoveAllTouchListener()

	self.driver = driver
	driver.Initialize()
	peer.LogDebug("OUT")
}

func (self *Simra) Stop() {
	peer.LogDebug("IN")
	// TODO implement
	peer.LogDebug("OUT")
}
