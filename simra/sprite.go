package simra

import "github.com/pankona/gomo-simra/simra/peer"

type Sprite struct {
	peer.Sprite
}

func (self *Sprite) AddTouchListener(listener peer.TouchListener) {
	LogDebug("IN")
	// TODO: call func to add listener
	//self.Sprite.TouchListeners = append(self.Sprite.TouchListeners, &listener)
	LogDebug("OUT")
}

func (self *Sprite) RemoveAllTouchListener() {
	LogDebug("IN")
	// TODO: call func to remove listener
	//self.Sprite.TouchListeners = nil
	LogDebug("OUT")
}
