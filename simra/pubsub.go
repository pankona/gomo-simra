package simra

import "sync"

// Publisher is publisher interface for pubsub system
type Publisher interface {
	Publish(i interface{})
	Subscribe(id string, s Subscriber) error
	Unsubscribe(id string)
}

// Subscriber is subscriber interface for pubsub system
type Subscriber interface {
	OnEvent(i interface{})
}

// PubSub is an object to control publish and subscribe
type PubSub struct {
	m           sync.Mutex
	subscribers map[string]Subscriber
}

// NewPubSub returns new PubSub instance
func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string]Subscriber, 0),
	}
}

// Publish publishes specified object to all subscriber
func (p *PubSub) Publish(i interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	for _, v := range p.subscribers {
		v.OnEvent(i)
	}
}

// Subscribe adds a subscriber to publisher
func (p *PubSub) Subscribe(id string, s Subscriber) error {
	p.m.Lock()
	defer p.m.Unlock()

	p.subscribers[id] = s
	return nil
}

// Unsubscribe remove a subscriber
func (p *PubSub) Unsubscribe(id string) {
	delete(p.subscribers, id)
}
