package scene

import "github.com/pankona/gomo-simra/simra"

// Life represents view part of remaining life
type Life struct {
	simra.Spriter
}

func (life *Life) getPosition() (float32, float32) {
	p := life.GetPosition()
	return (float32)(p.X), (float32)(p.Y)
}

func (life *Life) setPosition(x float32, y float32) {
	life.SetPosition((int)(x), (int)(y))
}

func (life *Life) setSpeed(s float64) {
}

func (life *Life) getSpeed() float64 {
	return 0
}

func (life *Life) setDirection(radian float64) {
}
func (life *Life) getDirection() float64 {
	return 0
}

func (life *Life) move() {
}
