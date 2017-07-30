package simra

import (
	"context"
	"fmt"
	"io"

	mp3 "github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"golang.org/x/mobile/asset"
)

type Audio struct {
	ctx        context.Context
	player     *oto.Player
	cancelFunc func()
}

type Audioer interface {
	Play(resource asset.File, loop bool, doneCallback func()) error
	Stop() error
}

func NewAudio() Audioer {
	return &Audio{}
}

func (a *Audio) Play(resource asset.File, loop bool, doneCallback func()) error {
	a.ctx, a.cancelFunc = context.WithCancel(context.Background())

	dec, err := mp3.NewDecoder(resource)
	if err != nil {
		return err
	}

	player, err := oto.NewPlayer(dec.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}

	doneChan := make(chan error)
	go func() {
		doneChan <- a.doPlay(player, dec, loop)
	}()

	go func() {
		defer dec.Close()
		defer player.Close()
		defer a.cancelFunc()

		select {
		case err := <-doneChan:
			LogDebug("playback completed. %s\n", err)
		case <-a.ctx.Done():
			LogDebug("playback canceled")
		}
		if doneCallback != nil {
			doneCallback()
		}
	}()

	return nil
}

func (a *Audio) doPlay(player *oto.Player, r io.ReadSeeker, loop bool) error {
	for {
		r.Seek(0, 0)
		_, err := io.Copy(player, r)
		if err != nil {
			return err
		}
		if !loop {
			break
		}
	}
	return nil
}

func (a *Audio) Stop() error {
	if a.cancelFunc == nil {
		return fmt.Errorf("stop didn't effect. not playing now")
	}
	a.cancelFunc()
	return nil
}
