package simra

import "github.com/pankona/gomo-simra/simra/peer"

type Sprite struct {
	peer.Sprite
}

func (self *Sprite) AddTouchListener(listener peer.TouchListener) {
	LogDebug("IN")
	LogDebug("OUT")
}

func (self *Sprite) RemoveAllTouchListener() {
	LogDebug("IN")
	LogDebug("OUT")
}
