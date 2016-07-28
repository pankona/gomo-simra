package scene

import (
	"image"

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

// Background represents a sprite for background
type Background struct {
	simra.Sprite
}

// Stage1 represents scene of Stage1.
type Stage1 struct {
	models     Models
	ball       Ball
	background [2]Background
	isTouching bool
}

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

	/*
		ball.Sprite.X += float32(math.Cos(ball.direction*math.Pi/180) * ball.speed)
		ball.Sprite.Y += float32(math.Sin(ball.direction*math.Pi/180) * ball.speed)
	*/
	ball.Sprite.Y += float32(ball.speed)
	if ball.Sprite.Y < 0 {
		ball.Sprite.Y = 0
		ball.speed = 0
	}

	if ball.Sprite.Y > config.ScreenHeight {
		ball.Sprite.Y = config.ScreenHeight
		ball.speed = 0
	}
	//	simra.LogDebug("cos90, sin90", math.Cos(90*math.Pi/180), math.Sin(90*math.Pi/180))
}

// Initialize initializes Stage1 scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (scene *Stage1) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// add global touch listener to catch touch end event
	simra.GetInstance().AddTouchListener(scene)

	// TODO: when goes to next scene, remove global touch listener
	// simra.GetInstance().RemoveTouchListener(Stage1)

	// initialize sprites
	scene.initSprites()

	scene.registerModels()

	simra.LogDebug("[OUT]")
}

// OnTouchBegin is called when Stage1 scene is Touched.
func (scene *Stage1) OnTouchBegin(x, y float32) {
	scene.isTouching = true
}

// OnTouchMove is called when Stage1 scene is Touched and moved.
func (scene *Stage1) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when Stage1 scene is Touched and it is released.
func (scene *Stage1) OnTouchEnd(x, y float32) {
	scene.isTouching = false
}

func (scene *Stage1) initSprites() {

	// set size of background
	scene.background[0].W = config.ScreenWidth
	scene.background[0].H = config.ScreenHeight

	// put center of screen
	scene.background[0].X = config.ScreenWidth / 2
	scene.background[0].Y = config.ScreenHeight / 2
	simra.GetInstance().AddSprite("bg.png",
		image.Rect(0, 0, config.ScreenWidth, config.ScreenHeight),
		&scene.background[0].Sprite)

	// set size of background
	scene.background[1].W = config.ScreenWidth
	scene.background[1].H = config.ScreenHeight

	// put out of screen
	scene.background[1].X = config.ScreenWidth/2 + config.ScreenWidth
	scene.background[1].Y = config.ScreenHeight / 2
	simra.GetInstance().AddSprite("bg.png",
		image.Rect(0, 0, config.ScreenWidth, config.ScreenHeight),
		&scene.background[1].Sprite)

	// set size of ball
	scene.ball.W = float32(48)
	scene.ball.H = float32(48)

	// put center of screen at start
	scene.ball.X = config.ScreenWidth / 2
	scene.ball.Y = config.ScreenHeight / 2

	simra.GetInstance().AddSprite("ball.png",
		image.Rect(0, 0, int(scene.ball.W), int(scene.ball.H)),
		&scene.ball.Sprite)
}

func (scene *Stage1) registerModels() {
	scene.ball.speed = 1
	scene.ball.direction = 90
	scene.models.RegisterBall(&scene.ball)
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (scene *Stage1) Drive() {
	scene.models.Progress(scene.isTouching)
}
