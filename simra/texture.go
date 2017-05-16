package simra

import (
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/simra/peer"
)

// Texture represents a texture.
type Texture struct {
	*peer.Texture
}

// NewImageTexture allocates a texture from asset image
func NewImageTexture(assetName string, rect image.Rectangle) *Texture {
	LogDebug("IN")
	tex := peer.GetGLPeer().LoadTexture(assetName, rect)
	LogDebug("OUT")
	return &Texture{peer.NewTexture(tex)}
}

// NewTextTexture allocates a texture from specified text
func NewTextTexture(text string, fontsize float64, fontcolor color.RGBA, rect image.Rectangle) *Texture {
	LogDebug("IN")
	tex := peer.GetGLPeer().MakeTextureByText(text, fontsize, fontcolor, rect)
	return &Texture{peer.NewTexture(tex)}
}
