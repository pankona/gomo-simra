package scene

import (
	"math"

	"github.com/pankona/gomo-simra/simra"
)

// Viewer represents a view
type viewer interface {
	getPosition() (x float32, y float32)
	setPosition(x float32, y float32)
	setSpeed(s float64)
	getSpeed() float64
	setDirection(radian float64)
	getDirection() float64
	move()
}

// Views represents a collection of view for a scene
type views struct {
	ball             viewer
	listeners        []viewEventListener
	isDead           bool
	elapsedDeadFrame int
}

const (
	waitFrameAfterDead = 90
)

func (views *views) restart() {
	views.isDead = false
}

type viewEventListener interface {
	onFinishDead()
}

func (views *views) registerBall(ball viewer) {
	views.ball = ball
}

func (views *views) addEventListener(listener viewEventListener) {
	views.listeners = append(views.listeners, listener)
}

// Progress progresses the time of views 1 frame
func (views *views) Progress(isKeyTouching bool) {
	if !views.isDead {
		return
	}

	views.move()

	_, py := views.ball.getPosition()
	if py == 0 {
		if views.elapsedDeadFrame > waitFrameAfterDead {
			for _, v := range views.listeners {
				v.onFinishDead()
				views.isDead = false
				views.elapsedDeadFrame = 0
			}
		} else {
			views.elapsedDeadFrame++
		}
	}
}

func (views *views) move() {
	ball := views.ball
	ball.move()
}

// event notification from view
func (views *views) onDead() {
	simra.LogDebug("onDead!!")
	views.isDead = true

	ball := views.ball

	dx := ball.getSpeed() * math.Cos(ball.getDirection()*math.Pi/180)
	dy := ball.getSpeed() * math.Sin(ball.getDirection()*math.Pi/180)
	dy = 0
	dx -= 3
	ball.setSpeed(math.Sqrt(dx*dx + dy*dy))
	ball.setDirection(math.Atan2(dy, dx) * 180 / math.Pi)
}
