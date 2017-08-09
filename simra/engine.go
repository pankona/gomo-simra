package simra

import (
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/simra/gomo"
	"github.com/pankona/gomo-simra/simra/peer"
)

// Simra is a struct that provides API interface of simra
type Simra struct {
	driver Driver
}

var simra *Simra

// GetInstance returns instance of Simra.
// It is necessary to call this function to get Simra instance
// since Simra is single instance.
func GetInstance() *Simra {
	//peer.LogDebug("IN")
	if simra == nil {
		simra = &Simra{}
	}
	//peer.LogDebug("OUT")
	return simra
}

type point struct {
	x, y int
}

func (simra *Simra) onUpdate() {
	if simra.driver != nil {
		simra.driver.Drive()
	}

	simra.collisionCheckAndNotify()
}

func (simra *Simra) onStopped() {
	peer.LogDebug("IN")
	simra.driver = nil
	peer.GetGLPeer().Finalize()
	peer.LogDebug("OUT")
}

// Start needs to call to enable all function belong to simra package.
func (simra *Simra) Start(onStart, onStop func()) {
	peer.LogDebug("IN")
	gomo.GetInstance().Initialize(onStart, onStop, simra.onUpdate)
	peer.GetSpriteContainer().Initialize()

	gomo.GetInstance().Start()
	peer.LogDebug("OUT")
}

// SetScene sets a driver as a scene.
// If a driver is already set, it is replaced with new one.
func (simra *Simra) SetScene(driver Driver) {
	peer.LogDebug("IN")
	peer.GetGLPeer().Reset()
	peer.GetTouchPeer().RemoveAllTouchListener()
	peer.GetSpriteContainer().RemoveSprites()

	simra.driver = driver
	peer.GetSpriteContainer().Initialize()
	driver.Initialize()
	peer.LogDebug("OUT")
}

// AddSprite adds a sprite to current scene.
// To call this function, SetScene must be called in advance.
// TODO: remove this function. this function is deprecated. use AddImageSprite instead.
func (simra *Simra) AddSprite(assetName string, rect image.Rectangle, s *Sprite) {
	simra.AddImageSprite(assetName, rect, s)
}

// AddSprite2 adds a sprite to current scene with empty texture.
func (simra *Simra) AddSprite2(s *Sprite) {
	tex := peer.GetGLPeer().MakeTextureByText("", 0, color.RGBA{0, 0, 0, 0}, image.Rect(0, 0, 1, 1))
	peer.GetSpriteContainer().AddSprite(&s.Sprite, tex, s.ProgressAnimation)
}

// AddImageSprite adds a sprite to current scene.
// To call this function, SetScene must be called in advance.
func (simra *Simra) AddImageSprite(assetName string, rect image.Rectangle, s *Sprite) {
	tex := peer.GetGLPeer().LoadTexture(assetName, rect)
	peer.GetSpriteContainer().AddSprite(&s.Sprite, tex, s.ProgressAnimation)
}

// AddTextSprite adds a sprite to current scene.
// To call this function, SetScene must be called in advance.
func (simra *Simra) AddTextSprite(text string, fontsize float64, fontcolor color.RGBA, rect image.Rectangle, s *Sprite) {
	tex := peer.GetGLPeer().MakeTextureByText(text, fontsize, fontcolor, rect)
	peer.GetSpriteContainer().AddSprite(&s.Sprite, tex, s.ProgressAnimation)
}

// RemoveSprite removes specified sprite from current scene.
// Removed sprite will be disappeared.
func (simra *Simra) RemoveSprite(s *Sprite) {
	peer.GetSpriteContainer().RemoveSprite(&s.Sprite)
}

// SetDesiredScreenSize configures virtual screen size.
// This function must be called at least once before calling Start.
func (simra *Simra) SetDesiredScreenSize(w, h float32) {
	peer.SetDesiredScreenSize(w, h)
}

// AddTouchListener registers a listener for notifying touch event.
// Event is notified when "screen" is touched.
func (simra *Simra) AddTouchListener(listener peer.TouchListener) {
	peer.GetTouchPeer().AddTouchListener(listener)
}

// RemoveTouchListener unregisters a listener for notifying touch event.
func (simra *Simra) RemoveTouchListener(listener peer.TouchListener) {
	peer.GetTouchPeer().RemoveTouchListener(listener)
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
