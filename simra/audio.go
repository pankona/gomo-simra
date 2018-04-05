package simra

import (
	"io"

	mp3 "github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"github.com/pankona/gomo-simra/simra/simlog"
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
		if err != nil && err != io.EOF {
			simlog.Error(err)
		}
		err = dec.Close()
		if err != nil {
			simlog.Error(err)
		}
		err = player.Close()
		if err != nil {
			simlog.Error(err)
		}
		doneCallback(err)
	}()
	return nil
}

func (a *audio) doPlay(player io.Writer, r io.Reader, f io.Seeker, loop bool) error {
	var written int64
	var err error
	readByte := (int64)(8192)
loop:
	for {
	playback:
		for {
			select {
			case <-a.isClosed:
				break loop
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
		_, err = f.Seek(0, io.SeekStart)
		if err != nil {
			return err
		}

	}
	return err
}

func (a *audio) Stop() error {
	a.isClosed <- true
	return nil
}
