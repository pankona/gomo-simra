package scene

import (
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/examples/sample3/scene/config"
	"github.com/pankona/gomo-simra/simra"
)

type gameState int

const (
	readyToStart gameState = iota
	started
	readyToRestart
)

// Stage1 represents scene of Stage1.
type Stage1 struct {
	models        models
	views         views
	ball          Ball
	obstacle      [2]Obstacle
	background    [2]Background
	isTouching    bool
	remainingLife int
	life          [3]Life
	readytext     [2]simra.Spriter
	gameovertext  [2]simra.Spriter
	gamestate     gameState
}

const (
	remainingLifeAtStart = 3
)

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

	scene.initialize()

	scene.resetPosition()
	scene.setupSprites()
	scene.registerViews()
	scene.registerModels()
	scene.remainingLife = remainingLifeAtStart

	simra.GetInstance().AddCollisionListener(&scene.ball, &scene.obstacle[0], &scene.models)
	simra.GetInstance().AddCollisionListener(&scene.ball, &scene.obstacle[1], &scene.models)

	scene.showReadyText()
	scene.gamestate = readyToStart

	simra.LogDebug("[OUT]")
}

func (scene *Stage1) initialize() {
	scene.readytext[0] = simra.NewSprite()
	scene.readytext[1] = simra.NewSprite()
	scene.background[0].Spriter = simra.NewSprite()
	scene.background[1].Spriter = simra.NewSprite()
	scene.ball.Spriter = simra.NewSprite()
	scene.obstacle[0].Spriter = simra.NewSprite()
	scene.obstacle[1].Spriter = simra.NewSprite()
	scene.life[0].Spriter = simra.NewSprite()
	scene.life[1].Spriter = simra.NewSprite()
	scene.life[2].Spriter = simra.NewSprite()
	scene.gameovertext[0] = simra.NewSprite()
	scene.gameovertext[1] = simra.NewSprite()
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

	if scene.gamestate == readyToStart {
		scene.gamestate = started
		scene.removeReadyText()
	} else if scene.gamestate == readyToRestart {
		// TODO: methodize
		scene.resetPosition()
		scene.views.restart()
		scene.models.restart()

		tex := simra.NewImageTexture("heart.png", image.Rect(0, 0, 384, 384))

		for i := 0; i < 3; i++ {
			simra.GetInstance().AddSprite(scene.life[i].Spriter)
			scene.life[i].ReplaceTexture(tex)
		}

		simra.GetInstance().RemoveSprite(scene.gameovertext[0])
		simra.GetInstance().RemoveSprite(scene.gameovertext[1])

		scene.remainingLife = remainingLifeAtStart

		scene.showReadyText()
		scene.gamestate = readyToStart
	}
}

func (scene *Stage1) showReadyText() {
	// ready text. will be removed after game start
	scene.readytext[0].SetPosition(config.ScreenWidth/2, config.ScreenHeight/6*4-65/2)
	scene.readytext[0].SetScale(config.ScreenWidth, 65)
	simra.GetInstance().AddSprite(scene.readytext[0])

	scene.readytext[1].SetPosition(config.ScreenWidth/2, config.ScreenHeight/6*3-65/2)
	scene.readytext[1].SetScale(config.ScreenWidth, 65)
	simra.GetInstance().AddSprite(scene.readytext[1])

	var tex *simra.Texture
	tex = simra.NewTextTexture("GET READY", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, config.ScreenWidth, 65))
	scene.readytext[0].ReplaceTexture(tex)
	tex = simra.NewTextTexture("TAP TO GO", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, config.ScreenWidth, 65))
	scene.readytext[1].ReplaceTexture(tex)

}

func (scene *Stage1) removeReadyText() {
	simra.GetInstance().RemoveSprite(scene.readytext[0])
	simra.GetInstance().RemoveSprite(scene.readytext[1])
}

func (scene *Stage1) resetPosition() {
	simra.LogDebug("IN")
	scene.background[0].SetScale(config.ScreenWidth+1, config.ScreenHeight)
	scene.background[0].SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)

	scene.background[1].SetScale(config.ScreenWidth+1, config.ScreenHeight)
	scene.background[1].SetPosition(config.ScreenWidth/2+(config.ScreenWidth), config.ScreenHeight/2)

	scene.ball.SetScale(48, 48)
	scene.ball.SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)

	scene.obstacle[0].SetScale(50, 100)
	scene.obstacle[0].SetPosition(config.ScreenWidth+config.ScreenWidth/2, config.ScreenHeight/3*2)
	scene.obstacle[1].SetScale(50, 100)
	scene.obstacle[1].SetPosition(config.ScreenWidth+config.ScreenWidth/2, config.ScreenHeight/3*1)

	scene.life[0].SetPosition(48, 30)
	scene.life[0].SetScale(48, 48)
	scene.life[1].SetPosition(48*2, 30)
	scene.life[1].SetScale(48, 48)
	scene.life[2].SetPosition(48*3, 30)
	scene.life[2].SetScale(48, 48)
	simra.LogDebug("OUT")
}

func (scene *Stage1) setupSprites() {
	simra.GetInstance().AddSprite(scene.background[0].Spriter)
	simra.GetInstance().AddSprite(scene.background[1].Spriter)
	simra.GetInstance().AddSprite(scene.ball.Spriter)
	simra.GetInstance().AddSprite(scene.obstacle[0].Spriter)
	simra.GetInstance().AddSprite(scene.obstacle[1].Spriter)
	simra.GetInstance().AddSprite(scene.life[0].Spriter)
	simra.GetInstance().AddSprite(scene.life[1].Spriter)
	simra.GetInstance().AddSprite(scene.life[2].Spriter)

	var tex *simra.Texture

	tex = simra.NewImageTexture("bg.png", image.Rect(0, 0, config.ScreenWidth, config.ScreenHeight))
	scene.background[0].ReplaceTexture(tex)

	tex = simra.NewImageTexture("bg.png", image.Rect(0, 0, config.ScreenWidth, config.ScreenHeight))
	scene.background[1].ReplaceTexture(tex)

	s := scene.ball.GetScale()
	tex = simra.NewImageTexture("ball.png", image.Rect(0, 0, s.W, s.H))
	scene.ball.ReplaceTexture(tex)

	tex = simra.NewImageTexture("obstacle.png", image.Rect(0, 0, 100, 100))
	scene.obstacle[0].ReplaceTexture(tex)
	scene.obstacle[1].ReplaceTexture(tex)

	tex = simra.NewImageTexture("heart.png", image.Rect(0, 0, 384, 384))
	scene.life[0].ReplaceTexture(tex)
	scene.life[1].ReplaceTexture(tex)
	scene.life[2].ReplaceTexture(tex)
}

func (scene *Stage1) registerViews() {
	scene.views.registerBall(&scene.ball)
	scene.views.addEventListener(scene)
}

func (scene *Stage1) showGameover() {
	scene.gameovertext[0].SetPosition(config.ScreenWidth/2, config.ScreenHeight/6*4-65/2)
	scene.gameovertext[0].SetScale(config.ScreenWidth, 65)
	simra.GetInstance().AddSprite(scene.gameovertext[0])

	scene.gameovertext[1].SetPosition(config.ScreenWidth/2, config.ScreenHeight/6*3-65/2)
	scene.gameovertext[1].SetScale(config.ScreenWidth, 65)
	simra.GetInstance().AddSprite(scene.gameovertext[1])

	var tex *simra.Texture
	tex = simra.NewTextTexture("GAME OVER", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, config.ScreenWidth, 65))
	scene.gameovertext[0].ReplaceTexture(tex)
	tex = simra.NewTextTexture("RESTART!!", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, config.ScreenWidth, 65))
	scene.gameovertext[1].ReplaceTexture(tex)
}

func (scene *Stage1) onFinishDead() {
	simra.LogDebug("IN")
	if scene.remainingLife == 0 {
		scene.showGameover()
		scene.gamestate = readyToRestart
		return
	}

	// life is still remaining. continue.
	scene.resetPosition()
	scene.views.restart()
	scene.models.restart()

	simra.GetInstance().RemoveSprite(scene.life[scene.remainingLife-1].Spriter)
	scene.remainingLife--
	simra.LogDebug("OUT")
}

func (scene *Stage1) registerModels() {
	scene.models.registerBall(&scene.ball)
	scene.models.registerObstacle(&scene.obstacle[0], 0)
	scene.models.registerObstacle(&scene.obstacle[1], 1)
	scene.models.registerBackground(&scene.background[0], 0)
	scene.models.registerBackground(&scene.background[1], 1)
	scene.models.addEventListener(&scene.views)
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (scene *Stage1) Drive() {
	if scene.gamestate == started {
		scene.models.Progress(scene.isTouching)
		scene.views.Progress(scene.isTouching)
	}
}
