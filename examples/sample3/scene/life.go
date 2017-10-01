package scene

import "github.com/pankona/gomo-simra/simra"

// Life represents view part of remaining life
type Life struct {
	simra.Spriter
}

func (l *Life) getPosition() (float32, float32) {
	p := l.GetPosition()
	return (float32)(p.X), (float32)(p.Y)
}

func (l *Life) setPosition(x float32, y float32) {
	l.SetPosition(x, y)
}

func (l *Life) setSpeed(s float64) {}

func (l *Life) getSpeed() float64 { return 0 }

func (l *Life) setDirection(radian float64) {}

func (l *Life) getDirection() float64 { return 0 }

func (l *Life) move() {}
