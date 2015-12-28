package simra

import "github.com/pankona/gomo-simra/gomo"

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

func (self *Simra) Initialize() {
	gomo.GetInstance().Initialize()
}

func (self *Simra) Start(driver Driver) {
	self.driver = driver
	driver.Initialize()
	gomo.GetInstance().Start()
}

func (self *Simra) Stop() {
}
