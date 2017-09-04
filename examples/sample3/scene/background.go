package scene

import (
	"github.com/pankona/gomo-simra/examples/sample2/scene/config"
	"github.com/pankona/gomo-simra/simra"
)

// Background represents a sprite for background
type Background struct {
	simra.Spriter
	speed float64
}

/**
 * Background implementation for Model interface
 */
func (bg *Background) setPosition(x, y float32) {
}

func (bg *Background) getPosition() (x, y float32) {
	x = 0
	y = 0
	return x, y
}

func (bg *Background) setRotate(r float32) {
}

func (bg *Background) getRotate() float32 {
	return 0
}

func (bg *Background) setDirection(d float64) {
}

func (bg *Background) getDirection() float64 {
	return 0
}

func (bg *Background) setSpeed(s float64) {
	bg.speed = s
}

func (bg *Background) getSpeed() float64 {
	return bg.speed
}

func (bg *Background) move() {
	p := bg.GetPosition()
	s := bg.GetScale()
	p.X -= (int)(bg.speed)
	bg.SetPositionX(p.X)
	if p.X < -1*s.W/2 {
		bg.SetPositionX((int)(config.ScreenWidth/2 + (config.ScreenWidth - float32(bg.speed))))
	}
}
