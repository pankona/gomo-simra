package scene

import (
	"fmt"
	"math"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/database"
	"github.com/pankona/gomo-simra/simra/fps"
	"github.com/pankona/gomo-simra/simra/image"
	"github.com/pankona/gomo-simra/simra/simlog"
	"github.com/pankona/gomo-simra/simra/storage"
)

// filestore represents a scene of filestore
type filestore struct {
	simra  simra.Simraer
	gopher simra.Spriter
	db     *simra.Database
}

// Initialize initializes filestore scene.
// This is called from simra.
// simra.SetDesiredScreenSize should be called to determine
// screen size of this scene.
// If SetDesiredScreenSize is already called in previous scene, this scene may not call the function.
func (f *filestore) Initialize(sim simra.Simraer) {
	f.simra = sim
	f.simra.SetDesiredScreenSize(1080/2, 1920/2)
	f.db = simra.OpenDB(&database.Boltdb{}, storage.NewStorage().DirectoryPath()) // TODO: when to call Close...?
	f.db.Close()
	f.initSprite()
	go func() {
		for {
			<-fps.After(60)
			f.storeCurrentPosition()
		}
	}()
}

func (f *filestore) storeCurrentPosition() {
	f.db = simra.OpenDB(&database.Boltdb{}, storage.NewStorage().DirectoryPath()) // TODO: when to call Close...?
	defer f.db.Close()

	p := f.gopher.GetPosition()
	pstr := fmt.Sprintf("%d,%d", p.X, p.Y)
	simlog.Debugf("store: %s", pstr)
	f.db.Put("position", pstr)
}

func (f *filestore) fetchCurrentPosition() {
	f.db = simra.OpenDB(&database.Boltdb{}, storage.NewStorage().DirectoryPath()) // TODO: when to call Close...?
	defer f.db.Close()

	pstr := (string)(f.db.Get("position").([]uint8))
	p := simra.Position{}
	_, err := fmt.Sscanf(pstr, "%f,%f", &p.X, &p.Y)
	if err != nil {
		simlog.Errorf("failed to fetch current position. %s", err.Error())
	}

	// restore last position from db
	simlog.Debugf("set: %s", pstr)
	f.gopher.SetPosition(p.X, p.Y)
}

func (f *filestore) initSprite() {
	f.initGopher()
	f.fetchCurrentPosition()
}

func (f *filestore) initGopher() {
	f.gopher = f.simra.NewSprite()
	// add gopher sprite
	f.gopher.SetScale(140, 90)

	// put center of screen at start
	f.gopher.SetPosition(1080/2/2, 1920/2/2)

	f.simra.AddSprite(f.gopher)
	tex := f.simra.NewImageTexture("waza-gophers.jpeg",
		image.Rect(152, 10, 152+f.gopher.GetScale().W, 10+f.gopher.GetScale().H))
	f.gopher.ReplaceTexture(tex)

	f.gopher.AddTouchListener(f)
}

var degree float32

// Drive is called from simra.
// This is used to update sprites position.
// This function will be called 60 times per sec.
func (f *filestore) Drive() {
	degree++
	if degree >= 360 {
		degree = 0
	}
	f.gopher.SetRotate(degree * math.Pi / 180)
}

// OnTouchBegin is called when Title scene is Touched.
// It is caused by calling AddtouchListener for e.gopher sprite.
func (f *filestore) OnTouchBegin(x, y float32) {
	f.gopher.SetPosition(x, y)
}

// OnTouchMove is called when Title scene is Touched and moved.
// It is caused by calling AddtouchListener for e.gopher sprite.
func (f *filestore) OnTouchMove(x, y float32) {
	f.gopher.SetPosition(x, y)
}

// OnTouchEnd is called when Title scene is Touched and it is released.
// It is caused by calling AddtouchListener for e.gopher sprite.
func (f *filestore) OnTouchEnd(x, y float32) {
	f.gopher.SetPosition(x, y)
}
