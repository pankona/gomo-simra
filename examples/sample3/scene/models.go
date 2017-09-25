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

func (m *models) restart() {
	m.isDead = false
}

func (m *models) registerBall(ball modeler) {
	ball.setSpeed(1)
	ball.setDirection(90)
	m.ball = ball
}

func (m *models) registerObstacle(obstacle modeler, index int) {
	obstacle.setSpeed(3)
	obstacle.setDirection(180)
	m.obstacles[index] = obstacle
}

func (m *models) registerBackground(bg modeler, index int) {
	bg.setSpeed(3)
	m.backgrounds[index] = bg
}

func (m *models) addEventListener(listener modelEventListener) {
	m.listeners = append(m.listeners, listener)
}

var degree float32

// Progress progresses the time of models 1 frame
func (m *models) Progress(isKeyTouching bool) {
	b := m.ball

	if !m.isDead {
		degree++
		if degree >= 360 {
			degree = 0
		}

		b.setRotate(float32(degree) * math.Pi / 180)

		if isKeyTouching {
			dx := b.getSpeed() * math.Cos(b.getDirection()*math.Pi/180)
			dy := b.getSpeed() * math.Sin(b.getDirection()*math.Pi/180)
			dy += 9.8 / 60 * 2
			b.setSpeed(math.Sqrt(dx*dx + dy*dy))
			b.setDirection(math.Atan2(dy, dx) * 180 / math.Pi)
		}
		m.move()
	}
}

func (m *models) move() {
	m.ball.move()
	m.obstacles[0].move()
	m.obstacles[1].move()
	m.backgrounds[0].move()
	m.backgrounds[1].move()
}

// OnCollision is called at collision detected
func (m *models) OnCollision(c1, c2 simra.Collider) {
	if _, ok := c1.(*Ball); ok {
		if _, ok := c2.(*Obstacle); ok {
			// collision indicates a miss. this will be decrease a life.
			if !m.isDead {
				m.isDead = true
				for _, v := range m.listeners {
					m.isDead = true
					v.onDead()
				}
			}
		}
	}
}
