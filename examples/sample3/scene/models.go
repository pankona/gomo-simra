package scene

import "math"

// Model represents a model
type Model interface {
	getPosition() (x float32, y float32)
	setPosition(x float32, y float32)
	getRotate() float32
	setRotate(radian float32)
	setDirection(radian float64)
	move()
}

// Models represents collection of model for a scene
type Models struct {
	ball       Model
	background [2]Model
}

// RegisterBall registers a model to array of model
func (models *Models) RegisterBall(model Model) {
	models.ball = model
}

// RegisterBackground registers a model to array of model
func (models *Models) RegisterBackground(model Model, index int) {
	models.background[index] = model
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
