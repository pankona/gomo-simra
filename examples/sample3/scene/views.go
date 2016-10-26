package scene

import (
	"github.com/pankona/gomo-simra/simra"
	"math"
)

// View represents a view
type view interface {
	getPosition() (x float32, y float32)
	setPosition(x float32, y float32)
	setSpeed(s float64)
	getSpeed() float64
	setDirection(radian float64)
	getDirection() float64
}

// Views represents a collection of view for a scene
type views struct {
	ball view
}

func (views *views) registerBall(ball view) {
	views.ball = ball
}

// event notification from model
func (views *views) onEvent() {
	// TODO: dispatch according to event type
	simra.LogDebug("onDead!!")
	ball := views.ball

	dx := ball.getSpeed() * math.Cos(ball.getDirection()*math.Pi/180)
	dy := ball.getSpeed() * math.Sin(ball.getDirection()*math.Pi/180)
	dy = 0
	dx -= 3
	ball.setSpeed(math.Sqrt(dx*dx + dy*dy))
	ball.setDirection(math.Atan2(dy, dx) * 180 / math.Pi)
}
