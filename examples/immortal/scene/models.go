package scene

import (
	"math"

	"github.com/pankona/gomo-simra/simra"
)

// modeler represents a model
type modeler interface {
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
	onDead()
}

type models struct {
	ball        modeler
	obstacles   [2]modeler
	backgrounds [2]modeler
	listeners   []modelEventListener
	isDead      bool
}

func (models *models) restart() {
	models.isDead = false
}

func (models *models) registerBall(ball modeler) {
	ball.setSpeed(1)
	ball.setDirection(90)
	models.ball = ball
}

func (models *models) registerObstacle(obstacle modeler, index int) {
	obstacle.setSpeed(3)
	obstacle.setDirection(180)
	models.obstacles[index] = obstacle
}

func (models *models) registerBackground(bg modeler, index int) {
	bg.setSpeed(3)
	models.backgrounds[index] = bg
}

func (models *models) addEventListener(listener modelEventListener) {
	models.listeners = append(models.listeners, listener)
}

var degree float32

// Progress progresses the time of models 1 frame
func (models *models) Progress(isKeyTouching bool) {
	ball := models.ball

	if !models.isDead {
		degree++
		if degree >= 360 {
			degree = 0
		}

		ball.setRotate(float32(degree) * math.Pi / 180)

		if isKeyTouching {
			dx := ball.getSpeed() * math.Cos(ball.getDirection()*math.Pi/180)
			dy := ball.getSpeed() * math.Sin(ball.getDirection()*math.Pi/180)
			dy += 9.8 / 60 * 2
			ball.setSpeed(math.Sqrt(dx*dx + dy*dy))
			ball.setDirection(math.Atan2(dy, dx) * 180 / math.Pi)
		}
		models.move()
	}
}

func (models *models) move() {
	ball := models.ball
	ball.move()
	obstacles := models.obstacles
	obstacles[0].move()
	obstacles[1].move()
	backgrounds := models.backgrounds
	backgrounds[0].move()
	backgrounds[1].move()
}

// OnCollision is called at collision detected
func (models *models) OnCollision(c1, c2 simra.Collider) {
	if _, ok := c1.(*Ball); ok {
		if _, ok := c2.(*Obstacle); ok {
			// collision indicates a miss. this will be decrease a life.
			if !models.isDead {
				models.isDead = true
				for _, v := range models.listeners {
					models.isDead = true
					v.onDead()
				}
			}
		}
	}
}
