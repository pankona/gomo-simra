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
	LogDebug("IN")
	GetTouchPeer().AddTouchListener(self)
	LogDebug("OUT")
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

func (self *SpriteContainer) RemoveSprite(remove *Sprite) {
	result := []*SpriteNodePair{}
	for _, sn := range self.spriteNodePairs {
		if sn.sprite != remove {
			result = append(result, sn)
		} else {
			// eng.Unregister doesn't work
			// since it is not implemented by gomobile.
			// TODO: call this after gomobile's implement.
			//GetGLPeer().eng.Unregister(sn.node)
		}
	}
	self.spriteNodePairs = result
}

func (self *SpriteContainer) RemoveSprites() {
	LogDebug("IN")
	self.spriteNodePairs = nil
	LogDebug("OUT")
}

func (self *SpriteContainer) ReplaceTexture(sprite *Sprite, subTex sprite.SubTex) {
	LogDebug("IN")
	for i := range self.spriteNodePairs {
		if self.spriteNodePairs[i].sprite == sprite {
			node := self.spriteNodePairs[i].node
			GetGLPeer().eng.SetSubTex(node, subTex)
		}
	}
	LogDebug("OUT")
}

func isContained(sprite *Sprite, x, y float32) bool {
	LogDebug("IN")
	if x >= sprite.X-sprite.W/2 &&
		x <= sprite.X+sprite.W/2 &&
		y >= sprite.Y-sprite.H/2 &&
		y <= sprite.Y+sprite.H/2 {
		LogDebug("OUT true")
		return true
	}
	LogDebug("OUT false")
	return false
}

func (self *SpriteContainer) OnTouchBegin(x, y float32) {
	LogDebug("IN")
	for i := range self.spriteNodePairs {
		listeners := self.spriteNodePairs[i].sprite.touchListeners
		if isContained(self.spriteNodePairs[i].sprite, x, y) {
			for j := range listeners {
				listener := listeners[j]
				if listener == nil {
					fmt.Println("listener is nil!")
					continue
				}

				(*listener).OnTouchBegin(x, y)
			}
		}
	}
	LogDebug("OUT")
}
func (self *SpriteContainer) OnTouchMove(x, y float32) {
	LogDebug("IN")
	for i := range self.spriteNodePairs {
		listeners := self.spriteNodePairs[i].sprite.touchListeners
		if isContained(self.spriteNodePairs[i].sprite, x, y) {
			for j := range listeners {
				listener := listeners[j]
				if listener == nil {
					fmt.Println("listener is nil!")
					continue
				}

				(*listener).OnTouchMove(x, y)
			}
		}
	}
	LogDebug("OUT")
}

func (self *SpriteContainer) OnTouchEnd(x, y float32) {
	LogDebug("IN")
	for i := range self.spriteNodePairs {
		listeners := self.spriteNodePairs[i].sprite.touchListeners
		if isContained(self.spriteNodePairs[i].sprite, x, y) {
			for j := range listeners {
				listener := listeners[j]
				if listener == nil {
					fmt.Println("listener is nil!")
					continue
				}

				(*listener).OnTouchEnd(x, y)
			}
		}
	}
	LogDebug("OUT")
}
