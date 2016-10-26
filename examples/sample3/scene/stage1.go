package scene

import (
	"image"

	"github.com/pankona/gomo-simra/examples/sample3/scene/config"
	"github.com/pankona/gomo-simra/simra"
)

// Stage1 represents scene of Stage1.
type Stage1 struct {
	models     models
	views      views
	ball       Ball
	obstacle   Obstacle
	background [2]Background
	isTouching bool
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
	scene.registerViews()
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
	scene.background[0].W = config.ScreenWidth + 1
	scene.background[0].H = config.ScreenHeight

	// put center of screen
	scene.background[0].X = config.ScreenWidth / 2
	scene.background[0].Y = config.ScreenHeight / 2
	simra.GetInstance().AddSprite("bg.png",
		image.Rect(0, 0, config.ScreenWidth, config.ScreenHeight),
		&scene.background[0].Sprite)

	// set size of background
	scene.background[1].W = config.ScreenWidth + 1
	scene.background[1].H = config.ScreenHeight

	// put out of screen
	scene.background[1].X = config.ScreenWidth/2 + (config.ScreenWidth)
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

	// set size of obstacle
	scene.obstacle.W = 50
	scene.obstacle.H = 100

	// put center/upper side of screen
	scene.obstacle.X = config.ScreenWidth / 2
	scene.obstacle.Y = config.ScreenHeight / 3 * 2
	simra.GetInstance().AddSprite("obstacle.png",
		image.Rect(0, 0, 100, 100),
		&scene.obstacle.Sprite)

	simra.GetInstance().AddCollisionListener(&scene.ball, &scene.obstacle, &scene.models)
}

func (scene *Stage1) registerViews() {
	scene.views.registerBall(&scene.ball)
}

func (scene *Stage1) registerModels() {
	scene.models.RegisterBall(&scene.ball)
	scene.models.RegisterBackground(&scene.background[0], 0)
	scene.models.RegisterBackground(&scene.background[1], 1)
	scene.models.addEventListener(&scene.views)
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (scene *Stage1) Drive() {
	scene.models.Progress(scene.isTouching)
}
