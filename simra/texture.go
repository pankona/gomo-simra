package simra

import (
	"image"
	"image/color"
	"runtime"

	"github.com/pankona/gomo-simra/simra/peer"
)

// Texture represents a texture.
type Texture struct {
	texture *peer.Texture
}

// NewImageTexture allocates a texture from asset image
func NewImageTexture(assetName string, rect image.Rectangle) *Texture {
	LogDebug("IN")
	// FIXME:
	gl := GetInstance().(*simra).gl
	tex := gl.LoadTexture(assetName, rect)
	LogDebug("OUT")
	t := &Texture{
		texture: gl.NewTexture(tex),
	}
	runtime.SetFinalizer(t, (*Texture).release)
	return t
}

// NewTextTexture allocates a texture from specified text
func NewTextTexture(text string, fontsize float64, fontcolor color.RGBA, rect image.Rectangle) *Texture {
	LogDebug("IN")
	// FIXME:
	gl := GetInstance().(*simra).gl
	tex := gl.MakeTextureByText(text, fontsize, fontcolor, rect)
	t := &Texture{
		texture: gl.NewTexture(tex),
	}
	runtime.SetFinalizer(t, (*Texture).release)
	LogDebug("OUT")
	return t
}

func (t *Texture) release() {
	LogDebug("IN")
	// FIXME:
	gl := GetInstance().(*simra).gl
	gl.ReleaseTexture(t.texture)
	LogDebug("OUT")
}
