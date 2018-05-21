package simra

import (
	"image/color"
	"runtime"

	"github.com/pankona/gomo-simra/simra/fps"
	"github.com/pankona/gomo-simra/simra/image"
	"github.com/pankona/gomo-simra/simra/internal/peer"
	"github.com/pankona/gomo-simra/simra/simlog"
)

// Simraer represents an interface of simra instance
type Simraer interface {
	// Start needs to call to enable all function belong to simra package.
	Start(driver Driver)
	// SetScene sets a driver as a scene.
	// If a driver is already set, it is replaced with new one.
	SetScene(driver Driver)
	// NewSprite returns an instance of Spriter
	NewSprite() Spriter
	// AddSprite adds a sprite to current scene with empty texture.
	AddSprite(s Spriter)
	// RemoveSprite removes specified sprite from current scene.
	// Removed sprite will be disappeared.
	RemoveSprite(s Spriter)
	// SetZIndex sets specified zindex to specified Spriter.
	// All Spriter has 0 as zindex as default.
	// Spriters will be draw as the order of zindex in ascending.
	// Spriter that has lesser number of zindex will be draw latter.
	SetZIndex(s Spriter, z int) error
	// GetZIndex returns specified Spriter's zindex
	GetZIndex(s Spriter) (int, error)
	// SetDesiredScreenSize configures virtual screen size.
	// This function must be called at least once before calling Start.
	SetDesiredScreenSize(w, h float32)
	// AddTouchListener registers a listener for notifying touch event.
	// Event is notified when "screen" is touched.
	AddTouchListener(listener TouchListener)
	// RemoveTouchListener unregisters a listener for notifying touch event.
	RemoveTouchListener(listener TouchListener)
	// AddCollisionListener add a callback function that is called on
	// collision is detected between c1 and c2.
	AddCollisionListener(c1, c2 Collider, listener CollisionListener)
	// RemoveAllCollisionListener removes all registered listeners
	RemoveAllCollisionListener()
	// NewImageTexture returns a texture instance of image
	NewImageTexture(assetName string, rect image.Rectangle) *Texture
	// NewImageTexture returns a texture instance of text
	NewTextTexture(text string, fontsize float64, fontcolor color.RGBA, rect image.Rectangle) *Texture
	// SetOnStopCallback sets a callback function that will be called on application goes invisible
	SetOnStopCallback(f func())
}

// TouchListener is interface to receive touch event
type TouchListener peer.TouchListener

type collisionMap struct {
	c1       Collider
	c2       Collider
	listener CollisionListener
}

// Simra is a struct that provides API interface of simra
type simra struct {
	driver          Driver
	comap           []*collisionMap
	gl              peer.GLer
	spritecontainer peer.SpriteContainerer
	onStop          func()
}

// NewSimra returns an instance of Simraer
func NewSimra() Simraer {
	return &simra{
		comap: make([]*collisionMap, 0),
	}
}

func (sim *simra) onUpdate() {
	if sim.driver != nil {
		sim.driver.Drive()
	}
	sim.collisionCheckAndNotify()
	sim.gl.Update(sim.spritecontainer)
}

func (sim *simra) onGomoStart(glc *peer.GLContext) {
	sim.gl.Initialize(glc)
	sim.SetScene(sim.driver)
}

func (sim *simra) onGomoStop() {
	if sim.onStop != nil {
		sim.onStop()
	}
	sim.spritecontainer.Initialize(sim.gl)
	sim.gl.Finalize()
}

// Start starts to run gomobile and set specified scene as first driver
func (sim *simra) Start(driver Driver) {
	simlog.FuncIn()

	gl := peer.NewGLPeer()
	sc := peer.GetSpriteContainer()
	sc.Initialize(gl)
	sim.gl = gl
	sim.spritecontainer = sc
	sim.driver = driver
	gomo := peer.GetGomo()
	gomo.Initialize(sim.onGomoStart, sim.onGomoStop, sim.onUpdate)
	gomo.Start()

	simlog.FuncOut()
}

// SetScene sets a driver as a scene.
// If a driver is already set, it is replaced with new one.
func (sim *simra) SetScene(driver Driver) {
	simlog.FuncIn()

	sim.spritecontainer.RemoveSprites()
	sim.gl.Reset()
	sim.spritecontainer.Initialize(sim.gl)
	peer.GetTouchPeer().RemoveAllTouchListeners()
	sim.spritecontainer.RemoveSprites()

	sim.driver = driver
	sim.spritecontainer.Initialize(sim.gl)
	err := sim.spritecontainer.AddSprite(&peer.Sprite{}, nil, fps.Progress)
	if err != nil {
		simlog.Errorf("failed to add sprite. err: %s", err.Error())
		return
	}

	driver.Initialize(sim)

	simlog.FuncOut()
}

// NewSprite returns an instance of Sprite
func (sim *simra) NewSprite() Spriter {
	return &sprite{
		simra:         sim,
		animationSets: map[string]*AnimationSet{},
	}
}

// AddSprite adds a sprite to current scene with empty texture.
func (sim *simra) AddSprite(s Spriter) {
	sp := s.(*sprite)
	err := sim.spritecontainer.AddSprite(&sp.Sprite, nil, nil)
	if err != nil {
		simlog.Errorf("failed to add sprite. err: %s", err.Error())
	}
}

// RemoveSprite removes specified sprite from current scene.
// Removed sprite will be disappeared.
func (sim *simra) RemoveSprite(s Spriter) {
	sp := s.(*sprite)
	sp.texture = nil
	sim.spritecontainer.RemoveSprite(&sp.Sprite)
}

// SetZIndex sets specified zindex to specified Spriter.
// All Spriter has 0 as zindex as default.
// Spriters will be draw as the order of zindex in ascending.
// Spriter that has lesser number of zindex will be draw latter.
//
// If specified sprite is not found (removed or not added yet),
// this function returns non-nil error.
func (sim *simra) SetZIndex(s Spriter, z int) error {
	sp := s.(*sprite)
	return sim.spritecontainer.SetZIndex(&sp.Sprite, z)
}

// GetZIndex returns specified Spriter's zindex
//
// If specified sprite is not found (removed or not added yet),
// this function returns non-nil error.
func (sim *simra) GetZIndex(s Spriter) (int, error) {
	sp := s.(*sprite)
	return sim.spritecontainer.GetZIndex(&sp.Sprite)
}

// SetDesiredScreenSize configures virtual screen size.
// This function must be called at least once before calling Start.
func (sim *simra) SetDesiredScreenSize(w, h float32) {
	ss := peer.GetScreenSizePeer()
	ss.SetDesiredScreenSize(w, h)
}

// AddTouchListener registers a listener for notifying touch event.
// Event is notified when "screen" is touched.
func (sim *simra) AddTouchListener(listener TouchListener) {
	peer.GetTouchPeer().AddTouchListener(listener)
}

// RemoveTouchListener unregisters a listener for notifying touch event.
func (sim *simra) RemoveTouchListener(listener TouchListener) {
	peer.GetTouchPeer().RemoveTouchListener(listener)
}

// AddCollisionListener add a callback function that is called on
// collision is detected between c1 and c2.
func (sim *simra) AddCollisionListener(c1, c2 Collider, listener CollisionListener) {
	// TODO: exclusive control
	simlog.FuncIn()
	sim.comap = append(sim.comap, &collisionMap{c1, c2, listener})
	simlog.FuncOut()
}

func (sim *simra) removeCollisionMap(c *collisionMap) {
	result := []*collisionMap{}

	for _, v := range sim.comap {
		if c.c1 != v.c1 && c.c2 != v.c2 && v != c {
			result = append(result, v)
		}
	}

	sim.comap = result
}

// RemoveAllCollisionListener removes all registered listeners
func (sim *simra) RemoveAllCollisionListener() {
	simlog.FuncIn()
	sim.comap = nil
	simlog.FuncOut()
}

// NewImageTexture allocates a texture from asset image
func (sim *simra) NewImageTexture(assetName string, rect image.Rectangle) *Texture {
	simlog.FuncIn()

	gl := sim.gl
	tex := gl.LoadTexture(assetName, rect.Rectangle)
	t := &Texture{
		simra:   sim,
		texture: gl.NewTexture(tex),
	}
	runtime.SetFinalizer(t, (*Texture).release)

	simlog.FuncOut()
	return t
}

// NewTextTexture allocates a texture from specified text
func (sim *simra) NewTextTexture(text string, fontsize float64, fontcolor color.RGBA, rect image.Rectangle) *Texture {
	simlog.FuncIn()

	gl := sim.gl
	tex := gl.MakeTextureByText(text, fontsize, fontcolor, rect.Rectangle)
	t := &Texture{
		simra:   sim,
		texture: gl.NewTexture(tex),
	}
	runtime.SetFinalizer(t, (*Texture).release)

	simlog.FuncOut()
	return t
}

type point struct {
	x, y float32
}

func (sim *simra) collisionCheckAndNotify() {
	// check collision
	for _, v := range sim.comap {
		// TODO: refactoring around here...
		x1, y1, w1, h1 := v.c1.GetXYWH()
		x2, y2, w2, h2 := v.c2.GetXYWH()
		ps := []*point{
			&point{x1 - w1/2, y1 + h1/2},
			&point{x1 + w1/2, y1 + h1/2},
			&point{x1 - w1/2, y1 - h1/2},
			&point{x1 + w1/2, y1 - h1/2},
		}
		for _, p := range ps {
			if p.x >= (x2-w2/2) && p.x <= (x2+w2/2) &&
				p.y >= (y2-h2/2) && p.y <= (y2+h2/2) {
				v.listener.OnCollision(v.c1, v.c2)
				return
			}
		}
	}
}

// RemoveCollisionListener removes a collision map by specified collider instance.
func (sim *simra) RemoveCollisionListener(c1, c2 Collider) {
	// TODO: exclusive control
	simlog.FuncIn()
	sim.removeCollisionMap(&collisionMap{c1, c2, nil})
	simlog.FuncOut()
}

func (sim *simra) comapLength() int {
	return len(sim.comap)
}

func (sim *simra) SetOnStopCallback(f func()) {
	sim.onStop = f
}
