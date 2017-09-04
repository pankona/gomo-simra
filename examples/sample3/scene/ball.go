package scene

import (
	"math"

	"github.com/pankona/gomo-simra/examples/sample3/scene/config"
	"github.com/pankona/gomo-simra/simra"
)

// Ball represents a ball
type Ball struct {
	simra.Spriter
	// direction is radian.
	direction float64
	speed     float64
}

/**
 * Ball implementation for Model interface
 */
func (ball *Ball) setPosition(x, y float32) {
	ball.SetPositionX((int)(x))
	ball.SetPositionY((int)(y))
}

func (ball *Ball) getPosition() (x, y float32) {
	p := ball.GetPosition()
	return (float32)(p.X), (float32)(p.Y)
}

func (ball *Ball) setRotate(r float32) {
	ball.SetRotate(r)
}

func (ball *Ball) getRotate() float32 {
	return (float32)(ball.GetRotate())
}

func (ball *Ball) setDirection(d float64) {
	ball.direction = d
}

func (ball *Ball) getDirection() float64 {
	return ball.direction
}

func (ball *Ball) setSpeed(s float64) {
	ball.speed = s
}

func (ball *Ball) getSpeed() float64 {
	return ball.speed
}

func (ball *Ball) move() {
	dx := ball.speed * math.Cos(ball.direction*math.Pi/180)
	dy := ball.speed * math.Sin(ball.direction*math.Pi/180)
	dy -= 9.8 / 60
	ball.speed = math.Sqrt(dx*dx + dy*dy)
	ball.direction = math.Atan2(dy, dx) * 180 / math.Pi

	p := ball.GetPosition()
	ball.SetPositionX(p.X + int(dx))
	ball.SetPositionY(p.Y + int(dy))

	p = ball.GetPosition()
	if p.Y < 0 {
		ball.SetPositionY(0)
		ball.speed = 0
	}

	if p.Y > config.ScreenHeight {
		ball.SetPositionY(config.ScreenHeight)
		ball.speed = 0
	}
}

// GetXYWH returns x, y w, h of receiver
func (ball *Ball) GetXYWH() (x, y, w, h int) {
	p := ball.GetPosition()
	s := ball.GetScale()
	return p.X, p.Y, s.W, s.H
}
