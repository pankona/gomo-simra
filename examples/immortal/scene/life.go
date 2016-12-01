package scene

import "github.com/pankona/gomo-simra/simra"

// Life represents view part of remaining life
type Life struct {
	simra.Sprite
}

func (life *Life) getPosition() (x float32, y float32) {
	x, y = life.X, life.Y
	return
}

func (life *Life) setPosition(x float32, y float32) {
	life.X, life.Y = x, y
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
