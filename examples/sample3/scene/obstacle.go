package scene

import "github.com/pankona/gomo-simra/simra"

// Obstacle represetnts a sprite for obstacle
type Obstacle struct {
	simra.Sprite
}

// GetXYWH returns x, y w, h of receiver
func (obstacle *Obstacle) GetXYWH() (x, y, w, h int) {
	return int(obstacle.Sprite.X), int(obstacle.Sprite.Y), int(obstacle.Sprite.W), int(obstacle.Sprite.H)
}
