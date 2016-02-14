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

func (sprite *Sprite) AddTouchListener(listener TouchListener) {
	LogDebug("IN")
	sprite.touchListeners = append(sprite.touchListeners, &listener)
	LogDebug("OUT")
}

func (sprite *Sprite) RemoveAllTouchListener() {
	LogDebug("IN")
	sprite.touchListeners = nil
	LogDebug("OUT")
}

type SpriteNodePair struct {
	sprite *Sprite
	node   *sprite.Node
	inuse  bool
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

func (spritecontainer *SpriteContainer) Initialize() {
	LogDebug("IN")
	GetTouchPeer().AddTouchListener(spritecontainer)
	LogDebug("OUT")
}

func (spritecontainer *SpriteContainer) AddSprite(s *Sprite, subTex sprite.SubTex) {
	LogDebug("IN")
	for _, snpair := range spritecontainer.spriteNodePairs {
		if s == snpair.sprite && snpair.inuse {
			LogDebug("this sprite is already added and currently still being available.")
			return
		}
	}

	var sn *SpriteNodePair
	for _, snpair := range spritecontainer.spriteNodePairs {
		if !snpair.inuse {
			sn = snpair
		}
	}

	if sn == nil {
		sn = &SpriteNodePair{}
	}

	sn.sprite = s
	if sn.node == nil {
		sn.node = GetGLPeer().newNode()
		spritecontainer.spriteNodePairs = append(spritecontainer.spriteNodePairs, sn)
	} else {
		GetGLPeer().appendChild(sn.node)
	}
	sn.inuse = true
	GetGLPeer().eng.SetSubTex(sn.node, subTex)
	LogDebug("OUT")
}

func (spritecontainer *SpriteContainer) RemoveSprite(remove *Sprite) {
	LogDebug("IN")
	for _, sn := range spritecontainer.spriteNodePairs {
		if sn.sprite == remove {
			if !sn.inuse {
				LogDebug("already removed.")
				return
			}
			sn.inuse = false
			GetGLPeer().removeChild(sn.node)
		}
	}
	LogDebug("OUT")
}

func (spritecontainer *SpriteContainer) RemoveSprites() {
	LogDebug("IN")
	spritecontainer.spriteNodePairs = nil
	LogDebug("OUT")
}

func (spritecontainer *SpriteContainer) ReplaceTexture(sprite *Sprite, subTex sprite.SubTex) {
	LogDebug("IN")
	for i := range spritecontainer.spriteNodePairs {
		if spritecontainer.spriteNodePairs[i].sprite == sprite {
			node := spritecontainer.spriteNodePairs[i].node
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

func (spritecontainer *SpriteContainer) OnTouchBegin(x, y float32) {
	LogDebug("IN")
	for i := range spritecontainer.spriteNodePairs {
		listeners := spritecontainer.spriteNodePairs[i].sprite.touchListeners
		if isContained(spritecontainer.spriteNodePairs[i].sprite, x, y) {
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
func (spritecontainer *SpriteContainer) OnTouchMove(x, y float32) {
	LogDebug("IN")
	for i := range spritecontainer.spriteNodePairs {
		listeners := spritecontainer.spriteNodePairs[i].sprite.touchListeners
		if isContained(spritecontainer.spriteNodePairs[i].sprite, x, y) {
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

func (spritecontainer *SpriteContainer) OnTouchEnd(x, y float32) {
	LogDebug("IN")
	for i := range spritecontainer.spriteNodePairs {
		listeners := spritecontainer.spriteNodePairs[i].sprite.touchListeners
		if isContained(spritecontainer.spriteNodePairs[i].sprite, x, y) {
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
