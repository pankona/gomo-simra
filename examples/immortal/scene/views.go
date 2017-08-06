package scene

import (
	"image"
	"image/color"
	"math"
	"strconv"

	"github.com/pankona/gomo-simra/examples/sample2/scene/config"
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
	score            simra.Sprite
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

// event notification from model
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

func (views *views) onScoreUpdate(score int) {
	views.score.X = config.ScreenWidth / 2
	views.score.Y = config.ScreenHeight / 2
	views.score.W = 65
	views.score.H = 65
	simra.GetInstance().RemoveSprite(&views.score)
	simra.GetInstance().AddTextSprite(strconv.Itoa(score),
		60, color.RGBA{255, 0, 0, 255},
		image.Rect(0, 0, 65, 65),
		&views.score)
}
