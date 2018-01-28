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

type errPubSub struct {
	pubsub *PubSub
	err    error
}

func (e *errPubSub) Subscribe(id string, s Subscriber) {
	if e.err != nil {
		return
	}
	e.err = e.pubsub.Subscribe(id, s)
}

func (e *errPubSub) Publish(i interface{}) {
	e.pubsub.Publish(i)
}

func TestPubSub(t *testing.T) {
	var wg sync.WaitGroup
	s := &sub{
		f: func() {
			wg.Done()
		},
	}

	p := &errPubSub{pubsub: NewPubSub()}
	// discards duplicated registrations
	p.Subscribe("subscriber1", s)
	p.Subscribe("subscriber1", s)
	wg.Add(1)
	p.Subscribe("subscriber2", s)
	p.Subscribe("subscriber2", s)
	wg.Add(1)
	p.Subscribe("subscriber3", s)
	p.Subscribe("subscriber3", s)
	wg.Add(1)

	if p.err != nil {
		t.Errorf("unexpected error. err = %s", p.err)
	}

	p.Publish(nil)
	wg.Wait()
}
