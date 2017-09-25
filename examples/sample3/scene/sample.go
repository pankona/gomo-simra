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

// sample represents scene of sample.
type sample struct {
	simra         simra.Simraer
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

// Initialize initializes sample scene
// This is called from simra.
// simra.SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (s *sample) Initialize(sim simra.Simraer) {
	s.simra = sim
	s.simra.SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// add global touch listener to catch touch end event
	s.simra.AddTouchListener(s)

	s.initialize()
	s.resetPosition()
	s.setupSprites()
	s.registerViews()
	s.registerModels()
	s.remainingLife = remainingLifeAtStart

	s.simra.AddCollisionListener(&s.ball, &s.obstacle[0], &s.models)
	s.simra.AddCollisionListener(&s.ball, &s.obstacle[1], &s.models)

	s.showReadyText()
	s.gamestate = readyToStart
}

func (s *sample) initialize() {
	s.readytext[0] = s.simra.NewSprite()
	s.readytext[1] = s.simra.NewSprite()
	s.background[0].Spriter = s.simra.NewSprite()
	s.background[1].Spriter = s.simra.NewSprite()
	s.ball.Spriter = s.simra.NewSprite()
	s.obstacle[0].Spriter = s.simra.NewSprite()
	s.obstacle[1].Spriter = s.simra.NewSprite()
	s.life[0].Spriter = s.simra.NewSprite()
	s.life[1].Spriter = s.simra.NewSprite()
	s.life[2].Spriter = s.simra.NewSprite()
	s.gameovertext[0] = s.simra.NewSprite()
	s.gameovertext[1] = s.simra.NewSprite()
}

// OnTouchBegin is called when sample scene is Touched.
func (s *sample) OnTouchBegin(x, y float32) {
	s.isTouching = true

}

// OnTouchMove is called when sample scene is Touched and moved.
func (s *sample) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when sample scene is Touched and it is released.
func (s *sample) OnTouchEnd(x, y float32) {
	s.isTouching = false

	if s.gamestate == readyToStart {
		s.gamestate = started
		s.removeReadyText()
	} else if s.gamestate == readyToRestart {
		// TODO: methodize
		s.resetPosition()
		s.views.restart()
		s.models.restart()

		tex := s.simra.NewImageTexture("heart.png", image.Rect(0, 0, 384, 384))

		for i := 0; i < 3; i++ {
			s.simra.AddSprite(s.life[i].Spriter)
			s.life[i].ReplaceTexture(tex)
		}

		s.simra.RemoveSprite(s.gameovertext[0])
		s.simra.RemoveSprite(s.gameovertext[1])

		s.remainingLife = remainingLifeAtStart

		s.showReadyText()
		s.gamestate = readyToStart
	}
}

func (s *sample) showReadyText() {
	// ready text. will be removed after game start
	s.readytext[0].SetPosition(config.ScreenWidth/2, config.ScreenHeight/6*4-65/2)
	s.readytext[0].SetScale(config.ScreenWidth, 65)
	s.simra.AddSprite(s.readytext[0])

	s.readytext[1].SetPosition(config.ScreenWidth/2, config.ScreenHeight/6*3-65/2)
	s.readytext[1].SetScale(config.ScreenWidth, 65)
	s.simra.AddSprite(s.readytext[1])

	var tex *simra.Texture
	tex = s.simra.NewTextTexture("GET READY", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, config.ScreenWidth, 65))
	s.readytext[0].ReplaceTexture(tex)
	tex = s.simra.NewTextTexture("TAP TO GO", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, config.ScreenWidth, 65))
	s.readytext[1].ReplaceTexture(tex)

}

func (s *sample) removeReadyText() {
	s.simra.RemoveSprite(s.readytext[0])
	s.simra.RemoveSprite(s.readytext[1])
}

func (s *sample) resetPosition() {
	s.background[0].SetScale(config.ScreenWidth+1, config.ScreenHeight)
	s.background[0].SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)

	s.background[1].SetScale(config.ScreenWidth+1, config.ScreenHeight)
	s.background[1].SetPosition(config.ScreenWidth/2+(config.ScreenWidth), config.ScreenHeight/2)

	s.ball.SetScale(48, 48)
	s.ball.SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)

	s.obstacle[0].SetScale(50, 100)
	s.obstacle[0].SetPosition(config.ScreenWidth+config.ScreenWidth/2, config.ScreenHeight/3*2)
	s.obstacle[1].SetScale(50, 100)
	s.obstacle[1].SetPosition(config.ScreenWidth+config.ScreenWidth/2, config.ScreenHeight/3*1)

	s.life[0].SetPosition(48, 30)
	s.life[0].SetScale(48, 48)
	s.life[1].SetPosition(48*2, 30)
	s.life[1].SetScale(48, 48)
	s.life[2].SetPosition(48*3, 30)
	s.life[2].SetScale(48, 48)
}

func (s *sample) setupSprites() {
	s.simra.AddSprite(s.background[0].Spriter)
	s.simra.AddSprite(s.background[1].Spriter)
	s.simra.AddSprite(s.ball.Spriter)
	s.simra.AddSprite(s.obstacle[0].Spriter)
	s.simra.AddSprite(s.obstacle[1].Spriter)
	s.simra.AddSprite(s.life[0].Spriter)
	s.simra.AddSprite(s.life[1].Spriter)
	s.simra.AddSprite(s.life[2].Spriter)

	var tex *simra.Texture

	tex = s.simra.NewImageTexture("bg.png", image.Rect(0, 0, config.ScreenWidth, config.ScreenHeight))
	s.background[0].ReplaceTexture(tex)

	tex = s.simra.NewImageTexture("bg.png", image.Rect(0, 0, config.ScreenWidth, config.ScreenHeight))
	s.background[1].ReplaceTexture(tex)

	scale := s.ball.GetScale()
	tex = s.simra.NewImageTexture("ball.png", image.Rect(0, 0, scale.W, scale.H))
	s.ball.ReplaceTexture(tex)

	tex = s.simra.NewImageTexture("obstacle.png", image.Rect(0, 0, 100, 100))
	s.obstacle[0].ReplaceTexture(tex)
	s.obstacle[1].ReplaceTexture(tex)

	tex = s.simra.NewImageTexture("heart.png", image.Rect(0, 0, 384, 384))
	s.life[0].ReplaceTexture(tex)
	s.life[1].ReplaceTexture(tex)
	s.life[2].ReplaceTexture(tex)
}

func (s *sample) registerViews() {
	s.views.registerBall(&s.ball)
	s.views.addEventListener(s)
}

func (s *sample) showGameover() {
	s.gameovertext[0].SetPosition(config.ScreenWidth/2, config.ScreenHeight/6*4-65/2)
	s.gameovertext[0].SetScale(config.ScreenWidth, 65)
	s.simra.AddSprite(s.gameovertext[0])

	s.gameovertext[1].SetPosition(config.ScreenWidth/2, config.ScreenHeight/6*3-65/2)
	s.gameovertext[1].SetScale(config.ScreenWidth, 65)
	s.simra.AddSprite(s.gameovertext[1])

	var tex *simra.Texture
	tex = s.simra.NewTextTexture("GAME OVER", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, config.ScreenWidth, 65))
	s.gameovertext[0].ReplaceTexture(tex)
	tex = s.simra.NewTextTexture("RESTART!!", 60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, config.ScreenWidth, 65))
	s.gameovertext[1].ReplaceTexture(tex)
}

func (s *sample) onFinishDead() {
	if s.remainingLife == 0 {
		s.showGameover()
		s.gamestate = readyToRestart
		return
	}

	// life is still remaining. continue.
	s.resetPosition()
	s.views.restart()
	s.models.restart()

	s.simra.RemoveSprite(s.life[s.remainingLife-1].Spriter)
	s.remainingLife--
}

func (s *sample) registerModels() {
	s.models.registerBall(&s.ball)
	s.models.registerObstacle(&s.obstacle[0], 0)
	s.models.registerObstacle(&s.obstacle[1], 1)
	s.models.registerBackground(&s.background[0], 0)
	s.models.registerBackground(&s.background[1], 1)
	s.models.addEventListener(&s.views)
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (s *sample) Drive() {
	if s.gamestate == started {
		s.models.Progress(s.isTouching)
		s.views.Progress(s.isTouching)
	}
}
