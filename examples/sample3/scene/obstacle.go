package scene

import (
	"math"

	"github.com/pankona/gomo-simra/examples/sample3/scene/config"
	"github.com/pankona/gomo-simra/simra"
)

// Obstacle represetnts a sprite for obstacle
type Obstacle struct {
	simra.Spriter
	// direction is radian.
	direction float64
	speed     float64
}

// GetXYWH returns x, y w, h of receiver
func (o *Obstacle) GetXYWH() (x, y, w, h float32) {
	p := o.GetPosition()
	s := o.GetScale()
	return p.X, p.Y, s.W, s.H
}

/**
 * Obstacle implementation for Model interface
 */
func (o *Obstacle) setPosition(x, y float32) {
	o.SetPosition(x, y)
}

func (o *Obstacle) getPosition() (float32, float32) {
	p := o.GetPosition()
	return p.X, p.Y
}

func (o *Obstacle) setRotate(r float32) {}

func (o *Obstacle) getRotate() float32 { return 0 }

func (o *Obstacle) setDirection(d float64) {
	o.direction = d
}

func (o *Obstacle) getDirection() float64 { return 0 }

func (o *Obstacle) setSpeed(s float64) {
	o.speed = s
}

func (o *Obstacle) getSpeed() float64 {
	return o.speed
}

func (o *Obstacle) move() {
	dx := o.speed * math.Cos(o.direction*math.Pi/180)
	dy := o.speed * math.Sin(o.direction*math.Pi/180)
	o.speed = math.Sqrt(dx*dx + dy*dy)
	o.direction = math.Atan2(dy, dx) * 180 / math.Pi

	p := o.GetPosition()
	o.SetPositionX(p.X + float32(dx))
	o.SetPositionY(p.Y + float32(dy))

	p = o.GetPosition()
	s := o.GetScale()
	if p.X < -1*s.W/2 {
		o.SetPositionX(config.ScreenWidth + config.ScreenWidth/2)
	}
}
