package peer

import (
	"fmt"

	"golang.org/x/mobile/exp/sprite"
)

// Texture represents a subtexture
type Texture struct {
	subTex sprite.SubTex
}

// NewTexture returns a instance of texture
func NewTexture(s sprite.SubTex) *Texture {
	return &Texture{subTex: s}
}

// Sprite represents a sprite.
type Sprite struct {
	// W = width of sprite
	W float32
	// H = height of sprite
	H float32
	// X = x position of sprite
	X float32
	// Y = y position of sprite
	Y float32
	// R = radius of sprite (use for rotation)
	R float32
	// touchListeners is listeners to notify touch event
	touchListeners []*TouchListener
}

// AddTouchListener registers a listener to notify touch event.
func (sprite *Sprite) AddTouchListener(listener TouchListener) {
	LogDebug("IN")
	sprite.touchListeners = append(sprite.touchListeners, &listener)
	LogDebug("OUT")
}

// RemoveAllTouchListener removes all registered listeners from sprite.
func (sprite *Sprite) RemoveAllTouchListener() {
	LogDebug("IN")
	sprite.touchListeners = nil
	LogDebug("OUT")
}

// SpriteNodePair represents pair of Sprite and sprite.Node.
// This is used for calculate affine transformation of Node
// from Sprite's property.
type SpriteNodePair struct {
	sprite *Sprite
	node   *sprite.Node
	inuse  bool
}

// SpriteContainer represents array of SpriteNodePair.
type SpriteContainer struct {
	spriteNodePairs []*SpriteNodePair
}

var spriteContainer *SpriteContainer

// GetSpriteContainer returns SpriteContainer.
// Since SpriteContainer is singleton, use this function
// to get instance of SpriteContainer.
func GetSpriteContainer() *SpriteContainer {
	if spriteContainer == nil {
		spriteContainer = &SpriteContainer{}
	}
	return spriteContainer
}

// Initialize initializes SpriteContainer object.
// This must be called to use all SpriteContainer's function in advance.
func (spritecontainer *SpriteContainer) Initialize() {
	LogDebug("IN")
	GetTouchPeer().AddTouchListener(spritecontainer)
	LogDebug("OUT")
}

// AddSprite adds a sprite to SpriteContainer.
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

// RemoveSprite removes a spcified sprite from SpriteContainer.
// Since Unregister of Node is not implemented by gomobile, this function just
// marks the specified sprite as "not in use".
// The sprite marked as "not in use" will be reused at AddSprite.
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

// RemoveSprites removes all registered sprites from SpriteContainer.
func (spritecontainer *SpriteContainer) RemoveSprites() {
	LogDebug("IN")
	spritecontainer.spriteNodePairs = nil
	LogDebug("OUT")
}

// ReplaceTexture replaces sprite's texture to specified one.
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

// ReplaceTexture2 replaces sprite's texture to specified one.
func (spritecontainer *SpriteContainer) ReplaceTexture2(sprite *Sprite, texture *Texture) {
	LogDebug("IN")
	for i := range spritecontainer.spriteNodePairs {
		if spritecontainer.spriteNodePairs[i].sprite == sprite {
			node := spritecontainer.spriteNodePairs[i].node
			GetGLPeer().eng.SetSubTex(node, texture.subTex)
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

// OnTouchBegin is called when screen is started to touch.
// This function calls listener's OnTouchBegin if the touched position is
// contained by sprite's rectangle.
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

// OnTouchMove is called when touch is moved (dragged).
// This function calls listener's OnTouchMove if the touched position is
// contained by sprite's rectangle.
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

// OnTouchEnd is called when touch is ended (released).
// This function calls listener's OnTouchEnd if the touched position is
// contained by sprite's rectangle.
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
