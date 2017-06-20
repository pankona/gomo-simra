package simra

import (
	"sync"
	"testing"
)

type sub struct {
	f func()
}

func (s *sub) OnEvent(c Commander) {
	s.f()
}

func TestPubSub(t *testing.T) {
	p := NewPubSub()

	var wg sync.WaitGroup
	s := &sub{
		f: func() {
			wg.Done()
		},
	}

	// discards duplicated registerations
	p.Subscribe("subscriber1", s)
	p.Subscribe("subscriber1", s)
	wg.Add(1)
	p.Subscribe("subscriber2", s)
	p.Subscribe("subscriber2", s)
	wg.Add(1)
	p.Subscribe("subscriber3", s)
	p.Subscribe("subscriber3", s)
	wg.Add(1)

	p.Publish(nil)
	wg.Wait()
}
