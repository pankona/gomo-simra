package scene

import (
	"github.com/pankona/gomo-simra/simra"
	"math"
)

// model represents a model
type model interface {
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

type modelEventListener interface {
	onEvent()
}

type models struct {
	ball       model
	background [2]model
	listeners  []modelEventListener
	isDead     bool
}

// RegisterBall registers ball as a model component
func (models *models) RegisterBall(ball model) {
	ball.setSpeed(1)
	ball.setDirection(90)
	models.ball = ball
}

// RegisterBackground registers a model to array of model
func (models *models) RegisterBackground(bg model, index int) {
	bg.setSpeed(3)
	models.background[index] = bg
}

func (models *models) addEventListener(listener modelEventListener) {
	models.listeners = append(models.listeners, listener)
}

var degree float32

// Progress progresses the time of models 1 frame
func (models *models) Progress(isKeyTouching bool) {
	if !models.isDead {
		degree++
		if degree >= 360 {
			degree = 0
		}

		ball := models.ball
		ball.setRotate(float32(degree) * math.Pi / 180)

		if isKeyTouching {
			dx := ball.getSpeed() * math.Cos(ball.getDirection()*math.Pi/180)
			dy := ball.getSpeed() * math.Sin(ball.getDirection()*math.Pi/180)
			dy += 9.8 / 60 * 2
			ball.setSpeed(math.Sqrt(dx*dx + dy*dy))
			ball.setDirection(math.Atan2(dy, dx) * 180 / math.Pi)
		}
	}

	models.move()
}

func (models *models) move() {
	ball := models.ball
	ball.move()
	background := models.background
	background[0].move()
	background[1].move()
}

// OnCollision is called at collision detected
func (models *models) OnCollision(c1, c2 simra.Collider) {
	if _, ok := c1.(*Ball); ok {
		if _, ok := c2.(*Obstacle); ok {
			// onDead
			if !models.isDead {
				models.isDead = true
				for _, v := range models.listeners {
					models.isDead = true
					v.onEvent()
				}
			}
		}
	}
}
