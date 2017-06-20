package simra

import "sync"

type Commander interface{}

type Publisher interface {
	Publish(c Commander)
	Subscribe(id string, s Subscriber) error
	Unsubscribe(id string)
}

type Subscriber interface {
	OnEvent(c Commander)
}

type PubSub struct {
	m           sync.Mutex
	subscribers map[string]Subscriber
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string]Subscriber, 0),
	}
}

func (p *PubSub) Publish(c Commander) {
	p.m.Lock()
	defer p.m.Unlock()

	for _, v := range p.subscribers {
		v.OnEvent(c)
	}
}

func (p *PubSub) Subscribe(id string, s Subscriber) error {
	p.m.Lock()
	defer p.m.Unlock()

	p.subscribers[id] = s
	return nil
}

func (p *PubSub) Unsubscribe(id string) {
	delete(p.subscribers, id)
}
