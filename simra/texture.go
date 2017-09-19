package simra

import (
	"image"
	"image/color"
	"runtime"

	"github.com/pankona/gomo-simra/simra/internal/peer"
)

// Texture represents a texture.
type Texture struct {
	simra   *simra
	texture *peer.Texture
}

// NewImageTexture allocates a texture from asset image
// TODO: use simra as receiver for this function
func NewImageTexture(assetName string, rect image.Rectangle) *Texture {
	LogDebug("IN")
	simra := GetInstance().(*simra)
	gl := simra.gl
	tex := gl.LoadTexture(assetName, rect)
	LogDebug("OUT")
	t := &Texture{
		simra:   simra,
		texture: gl.NewTexture(tex),
	}
	runtime.SetFinalizer(t, (*Texture).release)
	return t
}

// NewTextTexture allocates a texture from specified text
// TODO: use simra as receiver for this function
func NewTextTexture(text string, fontsize float64, fontcolor color.RGBA, rect image.Rectangle) *Texture {
	LogDebug("IN")
	simra := GetInstance().(*simra)
	gl := simra.gl
	tex := gl.MakeTextureByText(text, fontsize, fontcolor, rect)
	t := &Texture{
		simra:   simra,
		texture: gl.NewTexture(tex),
	}
	runtime.SetFinalizer(t, (*Texture).release)
	LogDebug("OUT")
	return t
}

func (t *Texture) release() {
	LogDebug("IN")
	gl := t.simra.gl
	gl.ReleaseTexture(t.texture)
	LogDebug("OUT")
}
