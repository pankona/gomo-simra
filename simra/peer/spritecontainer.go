package peer

import (
	"fmt"

	"golang.org/x/mobile/exp/sprite"
)

type Sprite struct {
	W              float32
	H              float32
	X              float32
	Y              float32
	R              float32
	touchListeners []*TouchListener
}

type SpriteNodePair struct {
	sprite *Sprite
	node   *sprite.Node
}

type SpriteContainer struct {
	spriteNodePairs []*SpriteNodePair
}

var spriteContainer *SpriteContainer

func GetSpriteContainer() *SpriteContainer {
	if spriteContainer == nil {
		spriteContainer = &SpriteContainer{}
	}
	return spriteContainer
}

func (self *SpriteContainer) Initialize() {
	GetTouchPeer().AddTouchListener(self)
}

func (self *SpriteContainer) AddSprite(s *Sprite, subTex sprite.SubTex) {
	LogDebug("IN")
	var sn SpriteNodePair
	sn.sprite = s
	sn.node = GetGLPeer().newNode()
	self.spriteNodePairs = append(self.spriteNodePairs, &sn)
	GetGLPeer().eng.SetSubTex(sn.node, subTex)
	LogDebug("OUT")
}

func (self *SpriteContainer) RemoveSprites() {
	LogDebug("IN")
	self.spriteNodePairs = nil
	LogDebug("OUT")
}

func (self *Sprite) AddTouchListener(listener TouchListener) {
	LogDebug("IN")
	self.touchListeners = append(self.touchListeners, &listener)
	LogDebug("OUT")
}

func (self *Sprite) RemoveAllTouchListener() {
	LogDebug("IN")
	self.touchListeners = nil
	LogDebug("OUT")
}

func (self *SpriteContainer) OnTouchBegin(x, y float32) {
	LogDebug("IN")
	for i := range self.spriteNodePairs {
		listeners := self.spriteNodePairs[i].sprite.touchListeners
		for j := range listeners {
			listener := listeners[j]
			if listener == nil {
				fmt.Println("listener is nil!")
				continue
			}
			(*listener).OnTouchBegin(x, y)
		}
	}
	LogDebug("OUT")
}
func (self *SpriteContainer) OnTouchMove(x, y float32) {
	LogDebug("IN")
	for i := range self.spriteNodePairs {
		listeners := self.spriteNodePairs[i].sprite.touchListeners
		for j := range listeners {
			listener := listeners[j]
			if listener == nil {
				fmt.Println("listener is nil!")
				continue
			}
			(*listener).OnTouchMove(x, y)
		}
	}
	LogDebug("OUT")
}

func (self *SpriteContainer) OnTouchEnd(x, y float32) {
	LogDebug("IN")
	for i := range self.spriteNodePairs {
		listeners := self.spriteNodePairs[i].sprite.touchListeners
		for j := range listeners {
			listener := listeners[j]
			if listener == nil {
				fmt.Println("listener is nil!")
				continue
			}
			(*listener).OnTouchEnd(x, y)
		}
	}
	LogDebug("OUT")
}
