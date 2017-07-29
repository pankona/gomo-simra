package simra

import (
	"io"

	"github.com/hajimehoshi/oto"
	"golang.org/x/mobile/asset"
)

type Audio struct {
	player *oto.Player
}

type Audioer interface {
	Play(resource asset.File) error
}

func NewAudio() Audioer {
	return &Audio{}
}

func (a *Audio) Play(resource asset.File) error {
	player, err := oto.NewPlayer(44100, 2, 2, 8192)
	if err != nil {
		return err
	}

	_, err = io.Copy(player, resource)
	if err != nil {
		return err
	}

	return player.Close()
}
