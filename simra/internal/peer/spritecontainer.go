package peer

import (
	"fmt"
	"sync"

	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
)

// SpriteContainerer represetnts an interface of SpriteContaienr
type SpriteContainerer interface {
	// Initialize initializes SpriteContainer object.
	// This must be called to use all SpriteContainer's function in advance.
	Initialize(gl GLer)
	// AddSprite adds a sprite to SpriteContainer.
	AddSprite(s *Sprite, subTex *sprite.SubTex, arrangeCallback func()) error
	// RemoveSprite removes a spcified sprite from SpriteContainer.
	// Since Unregister of Node is not implemented by gomobile, this function just
	// marks the specified sprite as "not in use".
	// The sprite marked as "not in use" will be reused at AddSprite.
	RemoveSprite(remove *Sprite)
	// RemoveSprites removes all registered sprites from SpriteContainer.
	RemoveSprites()
	// GetSpriteNodePairs returns map represntation of sprite and node pair
	GetSpriteNodePairs() *sync.Map
	// ReplaceTexture replaces sprite's texture to specified one.
	ReplaceTexture(sprite *Sprite, texture *Texture)
	// OnTouchBegin is called when screen is started to touch.
	// This function calls listener's OnTouchBegin if the touched position is
	// contained by sprite's rectangle.
	OnTouchBegin(x, y float32)
	// OnTouchMove is called when touch is moved (dragged).
	// This function calls listener's OnTouchMove if the touched position is
	// contained by sprite's rectangle.
	OnTouchMove(x, y float32)
	// OnTouchEnd is called when touch is ended (released).
	// This function calls listener's OnTouchEnd if the touched position is
	// contained by sprite's rectangle.
	OnTouchEnd(x, y float32)
}

type spriteNodePair struct {
	sprite *Sprite
	node   *sprite.Node
	inuse  bool
}

// SpriteContainer represents array of SpriteNodePair.
type SpriteContainer struct {
	spriteNodePairs sync.Map // map[*Sprite]*spriteNodePair
	gl              GLer
}

// GetSpriteContainer returns SpriteContainer.
// Since SpriteContainer is singleton, use this function
// to get instance of SpriteContainer.
func GetSpriteContainer() SpriteContainerer {
	return &SpriteContainer{}
}

// Initialize initializes SpriteContainer object.
// This must be called to use all SpriteContainer's function in advance.
func (sc *SpriteContainer) Initialize(gl GLer) {
	LogDebug("IN")
	sc.gl = gl
	GetTouchPeer().AddTouchListener(sc)
	LogDebug("OUT")
}

// AddSprite adds a sprite to SpriteContainer.
func (sc *SpriteContainer) AddSprite(s *Sprite, subTex *sprite.SubTex, arrangeCallback func()) error {
	LogDebug("IN")
	var sn *spriteNodePair
	i, ok := sc.spriteNodePairs.Load(s)
	if !ok {
		sn = &spriteNodePair{}
	} else {
		sn = i.(*spriteNodePair)
		if sn.inuse {
			return fmt.Errorf("this sprite is already added and currently still being available")
		}
	}

	sn.sprite = s
	if sn.node == nil {
		sn.node = sc.gl.NewNode(func(eng sprite.Engine, n *sprite.Node, t clock.Time) {
			if arrangeCallback != nil {
				arrangeCallback()
			}
		})
		sc.spriteNodePairs.Store(s, sn)
	} else {
		sc.gl.AppendChild(sn.node)
	}
	sn.inuse = true
	if subTex != nil {
		sc.gl.SetSubTex(sn.node, subTex)
	}
	LogDebug("OUT")
	return nil
}

// RemoveSprite removes a spcified sprite from SpriteContainer.
// Since Unregister of Node is not implemented by gomobile, this function just
// marks the specified sprite as "not in use".
// The sprite marked as "not in use" will be reused at AddSprite.
func (sc *SpriteContainer) RemoveSprite(remove *Sprite) {
	LogDebug("IN")
	i, ok := sc.spriteNodePairs.Load(remove)
	if !ok {
		return
	}
	sn := i.(*spriteNodePair)
	if !sn.inuse {
		LogDebug("already removed.")
		return
	}
	sn.inuse = false
	sc.gl.RemoveChild(sn.node)
	LogDebug("OUT")
}

// RemoveSprites removes all registered sprites from SpriteContainer.
func (sc *SpriteContainer) RemoveSprites() {
	LogDebug("IN")
	sc.spriteNodePairs = sync.Map{}
	LogDebug("OUT")
}

// GetSpriteNodePairs returns map represntation of sprite and node pair
func (sc *SpriteContainer) GetSpriteNodePairs() *sync.Map {
	return &sc.spriteNodePairs
}

// ReplaceTexture replaces sprite's texture to specified one.
func (sc *SpriteContainer) ReplaceTexture(sprite *Sprite, texture *Texture) {
	LogDebug("IN")
	if i, ok := sc.spriteNodePairs.Load(sprite); ok {
		node := i.(*spriteNodePair).node
		sc.gl.SetSubTex(node, &texture.subTex)
	}
	LogDebug("OUT")
}

func isContained(sprite *Sprite, x, y float32) bool {
	LogDebug("IN")
	if x >= (float32)(sprite.X)-(float32)(sprite.W)/2 &&
		x <= (float32)(sprite.X)+(float32)(sprite.W)/2 &&
		y >= (float32)(sprite.Y)-(float32)(sprite.H)/2 &&
		y <= (float32)(sprite.Y)+(float32)(sprite.H)/2 {
		LogDebug("OUT true")
		return true
	}
	LogDebug("OUT false")
	return false
}

type event int

const (
	touchBegin event = iota
	touchMove
	touchEnd
)

func (sc *SpriteContainer) emitTouchEvent(x, y float32, e event) {
	LogDebug("IN")
	sc.spriteNodePairs.Range(func(k, v interface{}) bool {
		s := v.(*spriteNodePair).sprite
		listeners := s.touchListeners
		if isContained(s, x, y) {
			for i := range listeners {
				switch e {
				case touchBegin:
					(*listeners[i]).OnTouchBegin(x, y)
				case touchMove:
					(*listeners[i]).OnTouchMove(x, y)
				case touchEnd:
					(*listeners[i]).OnTouchEnd(x, y)
				default:
					panic("unknown touch event!")
				}
			}
		}
		return true
	})
	LogDebug("OUT")
}

// OnTouchBegin is called when screen is started to touch.
// This function calls listener's OnTouchBegin if the touched position is
// contained by sprite's rectangle.
func (sc *SpriteContainer) OnTouchBegin(x, y float32) {
	sc.emitTouchEvent(x, y, touchBegin)
}

// OnTouchMove is called when touch is moved (dragged).
// This function calls listener's OnTouchMove if the touched position is
// contained by sprite's rectangle.
func (sc *SpriteContainer) OnTouchMove(x, y float32) {
	sc.emitTouchEvent(x, y, touchMove)
}

// OnTouchEnd is called when touch is ended (released).
// This function calls listener's OnTouchEnd if the touched position is
// contained by sprite's rectangle.
func (sc *SpriteContainer) OnTouchEnd(x, y float32) {
	sc.emitTouchEvent(x, y, touchEnd)
}