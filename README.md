[![Circle CI](https://circleci.com/gh/pankona/gomo-simra/tree/master.svg?style=svg)](https://circleci.com/gh/pankona/gomo-simra/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/pankona/gomo-simra)](https://goreportcard.com/report/github.com/pankona/gomo-simra)
[![GoDoc](https://godoc.org/github.com/pankona/gomo-simra?status.svg)](https://godoc.org/github.com/pankona/gomo-simra)  
<a href="https://app.zenhub.com/workspace/o/pankona/gomo-simra"><img src="https://raw.githubusercontent.com/ZenHubIO/support/master/zenhub-badge.png"></a>

# gomo-simra

GOMObile SIMple wRApper

## What is this

* simple wrapper for gomobile.
* purpose of this library is, let developers focus for scene creation using gomobile.
  * wraps gomobile's APIs for easy usage
  * wraps calculation of x, y coodinates using size.Event structure. 
  * wraps usage of f32.affine.
  * calculates scale to fit to any device's screen automatically.
  * provides easy scene transition

* see `example` directory to check sample applications.

## Build

* in each example directory, do `go build` or `gomobile build`.
  * `go build` generates an executable for PC (linux or mac)
  * `gomobile build` generates a package for android (APK).
  * for iPhone may be also available (but sorry, I've never confirmed)


## Basic usage

### Initialize

An entry point of an application using this library is like below.
(See example/sample1 to check whole of source codes of this example)

```go
// +build darwin linux

package main

import (
	"github.com/pankona/gomo-simra/examples/sample1/scene"
	"github.com/pankona/gomo-simra/simra"
)

func main() {
	sim := simra.NewSimra()   // create an instance of this library
	sim.Start(&scene.Title{}) // call Start method with an instance that implements simra.Driver
}
```

`simra.Driver` has two methods like below.

```go
package simra

// Driver represents a scene driver.
type Driver interface {
	// Initialize is called to initialize scene.
    // This method is called only once before starting Drive.
	Initialize(sim Simraer)

	// Drive is called about 60 times per 1 sec.
	// It is the chance to update sprite information like
	// position, appear/disappear, and change scene.
	Drive()
}
```

Application has chance to initialize sprites at `Initialize` callback.
After completion of `Initialize` method, `Drive` will be called repeatedly, about 60 times per 1 sec.

Here's example of `Initialize`.

```go
// sample represents a scene of sample
type sample struct {        
    simra  simra.Simraer
    gopher simra.Spriter
}

// Initialize initializes sample scene.
// This is called from simra.
// simra.SetDesiredScreenSize should be called to determine
// screen size of this scene.
// If SetDesiredScreenSize is already called in previous scene, this scene may not call the function.
func (s *sample) Initialize(sim simra.Simraer) { 
    // Retain simra instance passed via Initialize to call simra APIs later
    s.simra = sim

    // Screen size is assumed to be width=540, height=960.
    // Automatically fit to device's screen size.
    s.simra.SetDesiredScreenSize(540, 960)

    s.initGopher()
}

func (s *sample) initGopher() {
    s.gopher = s.simra.NewSprite() 

    // add gopher sprite.
    // specify size of this sprite
    s.gopher.SetScale(140, 90)

    // put center of screen at start
    s.gopher.SetPosition(540/2, 960/2)

    // AddSprite to show on screen
    s.simra.AddSprite(s.gopher)

    // obtain texture instance from jpeg image.
    // clip rectangle from specified image by x=152, y=10, width=140, height=90
    tex := s.simra.NewImageTexture("waza-gophers.jpeg",
        image.Rect(152, 10, 152+s.gopher.GetScale().W /*140*/, 10+s.gopher.GetScale().H /* 90 */))

    // anytime sprite's texture can be replaced with texture instance
    s.gopher.ReplaceTexture(tex)    

    // add touch listener to get touch event for this sprite
    // (need to implement simra.TouchListner interface to receive touch events)
    s.gopher.AddTouchListener(s) 
}
```

### Drive

`Drive` callback will be called 60 times per 1 sec.
Implementation example is like below.

```go
var degree float32

// Drive is called from simra.
// This is used to update sprites position.
// This function will be called 60 times per sec.
func (s *sample) Drive() {
    degree++
    if degree >= 360 {
        degree = 0
    }

    // continue rotating
    s.gopher.SetRotate(degree * math.Pi / 180)
}
```

### Touch event handling

`simra.TouchListner` represents an interface of touch event handler.
Implementation example is like below.

```go
// TouchListener is interface to receive touch events.
type TouchListener interface {
    OnTouchBegin(x, y float32)
    OnTouchMove(x, y float32)
    OnTouchEnd(x, y float32)
}

// OnTouchBegin is called when Title scene is Touched.
// It is caused by calling AddtouchListener for e.gopher sprite.
func (s *sample) OnTouchBegin(x, y float32) {
    // move position of sprite
    s.gopher.SetPosition(x, y)
}

// OnTouchMove is called when Title scene is Touched and moved.
// It is caused by calling AddtouchListener for e.gopher sprite.
func (s *sample) OnTouchMove(x, y float32) {
    // move position of sprite
    s.gopher.SetPosition(x, y)
}

// OnTouchEnd is called when Title scene is Touched and it is released.
// It is caused by calling AddtouchListener for e.gopher sprite.
func (s *sample) OnTouchEnd(x, y float32) {
    // move position of sprite
    s.gopher.SetPosition(x, y)
}
```

And then, call `AddTouchListener` method that is held by sprite instance.

```go
s.gopher.AddTouchListener(s) 
```
