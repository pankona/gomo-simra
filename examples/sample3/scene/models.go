package scene

import (
	"github.com/pankona/gomo-simra/simra"
	"math"
)

// Model represents a model
type Model interface {
	getPosition() (x float32, y float32)
	setPosition(x float32, y float32)
	getRotate() float32
	setRotate(radian float32)
	setDirection(radian float64)
	getDirection() float64
	setSpeed(s float64)
	getSpeed() float64
	move()
}

// Models represents collection of model for a scene
type Models struct {
	ball       Model
	background [2]Model
}

// RegisterBall registers a model to array of model
func (models *Models) RegisterBall(ball Model) {
	ball.setSpeed(1)
	ball.setDirection(90)
	models.ball = ball
}

// RegisterBackground registers a model to array of model
func (models *Models) RegisterBackground(bg Model, index int) {
	bg.setSpeed(3)
	models.background[index] = bg
}

var degree float32

// Progress progresses the time of models 1 frame
func (models *Models) Progress(isKeyTouching bool) {
	degree++
	if degree >= 360 {
		degree = 0
	}

	ball := models.ball
	ball.setRotate(float32(degree) * math.Pi / 180)

	if isKeyTouching {
		ball.setDirection(90)
	} else {
		ball.setDirection(270)
	}
	ball.move()

	background := models.background
	background[0].move()
	background[1].move()
}

// OnCollision is called at collision detected
func (models *Models) OnCollision(c1, c2 simra.Collider) {

	if _, ok := c1.(*Ball); ok {
		if _, ok := c2.(*Obstacle); ok {
			simra.LogDebug("collision detected!")
		}
	}
}
