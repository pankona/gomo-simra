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
func (b *Ball) setPosition(x, y float32) {
	b.SetPosition(x, y)
}

func (b *Ball) getPosition() (x, y float32) {
	p := b.GetPosition()
	return p.X, p.Y
}

func (b *Ball) setRotate(r float32) {
	b.SetRotate(r)
}

func (b *Ball) getRotate() float32 {
	return b.GetRotate()
}

func (b *Ball) setDirection(d float64) {
	b.direction = d
}

func (b *Ball) getDirection() float64 {
	return b.direction
}

func (b *Ball) setSpeed(s float64) {
	b.speed = s
}

func (b *Ball) getSpeed() float64 {
	return b.speed
}

func (b *Ball) move() {
	dx := b.speed * math.Cos(b.direction*math.Pi/180)
	dy := b.speed * math.Sin(b.direction*math.Pi/180)
	dy -= 9.8 / 60
	b.speed = math.Sqrt(dx*dx + dy*dy)
	b.direction = math.Atan2(dy, dx) * 180 / math.Pi

	p := b.GetPosition()
	b.SetPosition(p.X+float32(dx), p.Y+float32(dy))

	p = b.GetPosition()
	if p.Y < 0 {
		b.SetPositionY(0)
		b.speed = 0
	}

	if p.Y > config.ScreenHeight {
		b.SetPositionY(config.ScreenHeight)
		b.speed = 0
	}
}

// GetXYWH returns x, y w, h of receiver
func (b *Ball) GetXYWH() (x, y, w, h float32) {
	p := b.GetPosition()
	s := b.GetScale()
	return p.X, p.Y, s.W, s.H
}
