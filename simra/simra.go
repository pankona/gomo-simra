package simra

import (
	"github.com/pankona/gomo-simra/simra/fps"
	"github.com/pankona/gomo-simra/simra/internal/peer"
)

// Simraer represents an interface of simra instance
type Simraer interface {
	// Start needs to call to enable all function belong to simra package.
	Start(onStart, onStop func())
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
	// SetDesiredScreenSize configures virtual screen size.
	// This function must be called at least once before calling Start.
	SetDesiredScreenSize(w, h float32)
	// AddTouchListener registers a listener for notifying touch event.
	// Event is notified when "screen" is touched.
	AddTouchListener(listener peer.TouchListener)
	// RemoveTouchListener unregisters a listener for notifying touch event.
	RemoveTouchListener(listener peer.TouchListener)
	// AddCollisionListener add a callback function that is called on
	// collision is detected between c1 and c2.
	AddCollisionListener(c1, c2 Collider, listener CollisionListener)
	// RemoveAllCollisionListener removes all registered listeners
	RemoveAllCollisionListener()
}

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
	onStart         func()
	onStop          func()
}

// TODO: don't declare as package global.
var sim = &simra{
	comap: make([]*collisionMap, 0),
}

// GetInstance returns instance of Simra.
// It is necessary to call this function to get Simra instance
// since Simra is single instance.
func GetInstance() Simraer {
	return sim
}

type point struct {
	x, y int
}

func (sim *simra) onUpdate() {
	if sim.driver != nil {
		sim.driver.Drive()
	}
	sim.collisionCheckAndNotify()
	sim.gl.Update(sim.spritecontainer)
}

func (sim *simra) onStopped() {
	peer.LogDebug("IN")
	sim.driver = nil
	sim.gl.Finalize()
	peer.LogDebug("OUT")
}

func (sim *simra) onGomoStart(glc *peer.GLContext) {
	sim.gl.Initialize(glc)
	sim.onStart()
}

func (sim *simra) onGomoStop() {
	sim.spritecontainer.Initialize(sim.gl)
	sim.gl.Finalize()
	sim.onStop()
}

// Start needs to call to enable all function belong to simra package.
func (sim *simra) Start(onStart, onStop func()) {
	peer.LogDebug("IN")
	gl := peer.NewGLPeer()
	sc := peer.GetSpriteContainer()
	sc.Initialize(gl)
	sim.gl = gl
	sim.spritecontainer = sc
	sim.onStart = onStart
	sim.onStop = onStop
	gomo := peer.GetGomo()
	gomo.Initialize(sim.onGomoStart, sim.onGomoStop, sim.onUpdate)
	gomo.Start()
	peer.LogDebug("OUT")
}

// SetScene sets a driver as a scene.
// If a driver is already set, it is replaced with new one.
func (sim *simra) SetScene(driver Driver) {
	peer.LogDebug("IN")
	sim.spritecontainer.RemoveSprites()
	sim.gl.Reset()
	sim.spritecontainer.Initialize(sim.gl)
	peer.GetTouchPeer().RemoveAllTouchListeners()
	sim.spritecontainer.RemoveSprites()

	sim.driver = driver
	sim.spritecontainer.Initialize(sim.gl)
	sim.spritecontainer.AddSprite(&peer.Sprite{}, nil, fps.Progress)

	driver.Initialize()
	peer.LogDebug("OUT")
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
	sim.spritecontainer.AddSprite(&sp.Sprite, nil, nil)
}

// RemoveSprite removes specified sprite from current scene.
// Removed sprite will be disappeared.
func (sim *simra) RemoveSprite(s Spriter) {
	sp := s.(*sprite)
	sp.texture = nil
	sim.spritecontainer.RemoveSprite(&sp.Sprite)
}

// SetDesiredScreenSize configures virtual screen size.
// This function must be called at least once before calling Start.
func (sim *simra) SetDesiredScreenSize(w, h float32) {
	ss := peer.GetScreenSizePeer()
	ss.SetDesiredScreenSize(w, h)
}

// AddTouchListener registers a listener for notifying touch event.
// Event is notified when "screen" is touched.
func (sim *simra) AddTouchListener(listener peer.TouchListener) {
	peer.GetTouchPeer().AddTouchListener(listener)
}

// RemoveTouchListener unregisters a listener for notifying touch event.
func (sim *simra) RemoveTouchListener(listener peer.TouchListener) {
	peer.GetTouchPeer().RemoveTouchListener(listener)
}

// AddCollisionListener add a callback function that is called on
// collision is detected between c1 and c2.
func (sim *simra) AddCollisionListener(c1, c2 Collider, listener CollisionListener) {
	// TODO: exclusive controll
	LogDebug("IN")
	sim.comap = append(sim.comap, &collisionMap{c1, c2, listener})
	LogDebug("OUT")
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
	LogDebug("IN")
	sim.comap = nil
	LogDebug("OUT")
}

func (sim *simra) collisionCheckAndNotify() {
	//LogDebug("IN")

	// check collision
	for _, v := range sim.comap {
		// TODO: refactor around here...
		x1, y1, w1, h1 := v.c1.GetXYWH()
		x2, y2, w2, h2 := v.c2.GetXYWH()

		p1 := &point{x1 - w1/2, y1 + h1/2}
		p2 := &point{x1 + w1/2, y1 + h1/2}
		p3 := &point{x1 - w1/2, y1 - h1/2}
		p4 := &point{x1 + w1/2, y1 - h1/2}

		if p1.x >= (x2-w2/2) && p1.x <= (x2+w2/2) &&
			p1.y >= (y2-h2/2) && p1.y <= (y2+h2/2) {
			v.listener.OnCollision(v.c1, v.c2)
			return
		}
		if p2.x >= (x2-w2/2) && p2.x <= (x2+w2/2) &&
			p2.y >= (y2-h2/2) && p2.y <= (y2+h2/2) {
			v.listener.OnCollision(v.c1, v.c2)
			return
		}
		if p3.x >= (x2-w2/2) && p3.x <= (x2+w2/2) &&
			p3.y >= (y2-h2/2) && p3.y <= (y2+h2/2) {
			v.listener.OnCollision(v.c1, v.c2)
			return
		}
		if p4.x >= (x2-w2/2) && p4.x <= (x2+w2/2) &&
			p4.y >= (y2-h2/2) && p4.y <= (y2+h2/2) {
			v.listener.OnCollision(v.c1, v.c2)
			return
		}
	}
	//LogDebug("OUT")
}

// RemoveCollisionListener removes a collision map by specified collider instance.
func (sim *simra) RemoveCollisionListener(c1, c2 Collider) {
	// TODO: exclusive controll
	LogDebug("IN")
	sim.removeCollisionMap(&collisionMap{c1, c2, nil})
	LogDebug("OUT")
}

func (sim *simra) comapLength() int {
	return len(sim.comap)
}

// LogDebug prints logs.
// From simra, just call peer.LogDebug.
// This is disabled at Release Build.
func LogDebug(format string, a ...interface{}) {
	peer.LogDebug(format, a...)
}

// LogError prints logs.
// From simra, just call peer.LogError.
// This is never disabled even for Release build.
func LogError(format string, a ...interface{}) {
	peer.LogError(format, a...)
}
