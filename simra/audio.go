package simra

import (
	"io"

	mp3 "github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"golang.org/x/mobile/asset"
)

type audio struct {
	isClosed chan bool
}

// Audioer is an interface for treating audio
type Audioer interface {
	Play(resource asset.File, loop bool, doneCallback func(err error)) error
	Stop() error
}

// NewAudio returns new audio instance that implements Audioer interface
func NewAudio() Audioer {
	return &audio{
		isClosed: make(chan bool, 1),
	}
}

func (a *audio) Play(resource asset.File, loop bool, doneCallback func(err error)) error {
	dec, err := mp3.NewDecoder(resource)
	if err != nil {
		return err
	}

	player, err := oto.NewPlayer(dec.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}

	go func() {
		err := a.doPlay(player, dec, resource, loop)
		if err != nil {
			LogError(err.Error())
		}
		err = dec.Close()
		if err != nil {
			LogError(err.Error())
		}
		err = player.Close()
		if err != nil {
			LogError(err.Error())
		}
		doneCallback(err)
	}()
	return nil
}

func (a *audio) doPlay(player *oto.Player, r io.ReadSeeker, f asset.File, loop bool) error {
	var written int64
	var err error
	readByte := (int64)(8192)
loop:
	for {
	playback:
		for {
			select {
			case <-a.isClosed:
				readByte = 0
				f.Seek(0, io.SeekEnd)
			default:
				// for non-blocking
			}
			written, err = io.CopyN(player, r, readByte)
			if err != nil || written == 0 {
				// error or EOF
				break playback
			}
		}
		if !loop || err != io.EOF {
			break loop
		}
		f.Seek(0, io.SeekStart)
	}
	return err
}

func (a *audio) Stop() error {
	a.isClosed <- true
	return nil
}
