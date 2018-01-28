package simra

import (
	"sync"
	"testing"
)

type sub struct {
	f func()
}

func (s *sub) OnEvent(i interface{}) {
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

	// discards duplicated registrations
	err := p.Subscribe("subscriber1", s)
	err = p.Subscribe("subscriber1", s)
	wg.Add(1)
	err = p.Subscribe("subscriber2", s)
	err = p.Subscribe("subscriber2", s)
	wg.Add(1)
	err = p.Subscribe("subscriber3", s)
	err = p.Subscribe("subscriber3", s)
	wg.Add(1)

	if err != nil {
		t.Errorf("unexpected error. err = %s", err)
	}

	p.Publish(nil)
	wg.Wait()
}
