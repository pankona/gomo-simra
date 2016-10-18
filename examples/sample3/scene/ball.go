package scene

import (
	"github.com/pankona/gomo-simra/examples/sample3/scene/config"
	"github.com/pankona/gomo-simra/simra"
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
func (ball *Ball) getPosition() (x, y float32) {
	x = 0
	y = 0
	return x, y
}

func (ball *Ball) setPosition(x, y float32) {
	ball.Sprite.X = x
	ball.Sprite.Y = y
}

func (ball *Ball) getRotate() float32 {
	return ball.Sprite.R
}

func (ball *Ball) setRotate(r float32) {
	ball.Sprite.R = r
}

func (ball *Ball) setDirection(d float64) {
	ball.direction = d
}

func (ball *Ball) move() {
	if ball.direction > 0 && ball.direction < 180 {
		ball.speed += 9.8 / 60
	} else {
		ball.speed -= 9.8 / 60

	}

	ball.Sprite.Y += float32(ball.speed)
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
