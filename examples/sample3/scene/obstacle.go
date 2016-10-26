package scene

import (
	"github.com/pankona/gomo-simra/examples/sample3/scene/config"
	"github.com/pankona/gomo-simra/simra"
	"math"
)

// Obstacle represetnts a sprite for obstacle
type Obstacle struct {
	simra.Sprite
	// direction is radian.
	direction float64
	speed     float64
}

// GetXYWH returns x, y w, h of receiver
func (obstacle *Obstacle) GetXYWH() (x, y, w, h int) {
	return int(obstacle.Sprite.X), int(obstacle.Sprite.Y), int(obstacle.Sprite.W), int(obstacle.Sprite.H)
}

/**
 * Obstacle implementation for Model interface
 */
func (obstacle *Obstacle) setPosition(x, y float32) {
	obstacle.Sprite.X = x
	obstacle.Sprite.Y = y
}

func (obstacle *Obstacle) getPosition() (x, y float32) {
	return obstacle.Sprite.X, obstacle.Sprite.Y
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

	obstacle.Sprite.X += float32(dx)
	obstacle.Sprite.Y += float32(dy)

	if obstacle.Sprite.X < -1*obstacle.Sprite.W/2 {
		obstacle.Sprite.X = config.ScreenWidth + config.ScreenWidth/2
	}
}
