package simra

import "github.com/pankona/gomo-simra/simra/peer"

type Sprite struct {
	peer.Sprite
}

func (self *Sprite) AddTouchListener(listener peer.TouchListener) {
	LogDebug("IN")
	self.Sprite.AddTouchListener(listener)
	LogDebug("OUT")
}

func (self *Sprite) RemoveAllTouchListener() {
	LogDebug("IN")
	self.Sprite.RemoveAllTouchListener()
	LogDebug("OUT")
}
