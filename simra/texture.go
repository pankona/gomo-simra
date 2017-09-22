package simra

import (
	"github.com/pankona/gomo-simra/simra/internal/peer"
)

// Texture represents a texture.
type Texture struct {
	simra   *simra
	texture *peer.Texture
}

func (t *Texture) release() {
	LogDebug("IN")
	gl := t.simra.gl
	gl.ReleaseTexture(t.texture)
	LogDebug("OUT")
}
