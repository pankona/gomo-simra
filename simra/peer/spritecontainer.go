package peer

import (
	"golang.org/x/mobile/exp/sprite"
)

type Sprite struct {
	W float32
	H float32
	X float32
	Y float32
	R float32
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
	self.spriteNodePairs = nil
}
