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
func (obstacle *Obstacle) GetXYWH() (x, y, w, h int) {
	p := obstacle.GetPosition()
	s := obstacle.GetScale()
	return p.X, p.Y, s.W, s.H
}

/**
 * Obstacle implementation for Model interface
 */
func (obstacle *Obstacle) setPosition(x, y float32) {
	obstacle.SetPosition((int)(x), (int)(y))
}

func (obstacle *Obstacle) getPosition() (float32, float32) {
	p := obstacle.GetPosition()
	return (float32)(p.X), (float32)(p.Y)
}

func (obstacle *Obstacle) setRotate(r float32) {
}

func (obstacle *Obstacle) getRotate() float32 {
	return 0
}

func (obstacle *Obstacle) setDirection(d float64) {
	obstacle.direction = d
}

func (obstacle *Obstacle) getDirection() float64 {
	return 0
}

func (obstacle *Obstacle) setSpeed(s float64) {
	obstacle.speed = s
}

func (obstacle *Obstacle) getSpeed() float64 {
	return obstacle.speed
}

func (obstacle *Obstacle) move() {
	dx := obstacle.speed * math.Cos(obstacle.direction*math.Pi/180)
	dy := obstacle.speed * math.Sin(obstacle.direction*math.Pi/180)
	obstacle.speed = math.Sqrt(dx*dx + dy*dy)
	obstacle.direction = math.Atan2(dy, dx) * 180 / math.Pi

	p := obstacle.GetPosition()
	obstacle.SetPositionX(p.X + int(dx))
	obstacle.SetPositionY(p.Y + int(dy))

	p = obstacle.GetPosition()
	s := obstacle.GetScale()
	if p.X < -1*s.W/2 {
		obstacle.SetPositionX(config.ScreenWidth + config.ScreenWidth/2)
	}
}
