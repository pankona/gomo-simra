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

func (self *Simra) onStopped() {
	peer.LogDebug("IN")
	self.driver = nil
	peer.GetGLPeer().Finalize()
	peer.LogDebug("OUT")
}

func (self *Simra) Start(onStart, onStop chan bool) {
	peer.LogDebug("IN")
	gomo.GetInstance().Initialize(onStart, onStop, self.onUpdate)
	gomo.GetInstance().Start()
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
