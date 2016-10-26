package scene

import (
	"github.com/pankona/gomo-simra/examples/sample3/scene/config"
	"github.com/pankona/gomo-simra/simra"
	"math"
)

// Ball represents a ball
type Ball struct {
	simra.Sprite
	// direction is radian.
	direction float64
	speed     float64
}

/**
 * Ball implementation for Model interface
 */
func (ball *Ball) setPosition(x, y float32) {
	ball.Sprite.X = x
	ball.Sprite.Y = y
}

func (ball *Ball) getPosition() (x, y float32) {
	return ball.Sprite.X, ball.Sprite.Y
}

func (ball *Ball) setRotate(r float32) {
	ball.Sprite.R = r
}

func (ball *Ball) getRotate() float32 {
	return ball.Sprite.R
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

	ball.Sprite.X += float32(dx)
	ball.Sprite.Y += float32(dy)

	if ball.Sprite.Y < 0 {
		ball.Sprite.Y = 0
		ball.speed = 0
	}

	if ball.Sprite.Y > config.ScreenHeight {
		ball.Sprite.Y = config.ScreenHeight
		ball.speed = 0
	}
}

// GetXYWH returns x, y w, h of receiver
func (ball *Ball) GetXYWH() (x, y, w, h int) {
	return int(ball.Sprite.X), int(ball.Sprite.Y), int(ball.Sprite.W), int(ball.Sprite.H)
}
