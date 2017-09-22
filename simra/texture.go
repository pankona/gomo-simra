package simra

import (
	"github.com/pankona/gomo-simra/simra/internal/peer"
	"github.com/pankona/gomo-simra/simra/simlog"
)

// Texture represents a texture.
type Texture struct {
	simra   *simra
	texture *peer.Texture
}

func (t *Texture) release() {
	simlog.FuncIn()
	gl := t.simra.gl
	gl.ReleaseTexture(t.texture)
	simlog.FuncOut()
}
