package scene

import "github.com/pankona/gomo-simra/simra/image"

// ctrlUp
type ctrlUpTouchListener struct {
	*sample
}

func (c *ctrlUpTouchListener) OnTouchBegin(x, y float32) {
	c.buttonState = ctrlUp
}

func (c *ctrlUpTouchListener) OnTouchMove(x, y float32) {
	c.buttonState = ctrlUp
}

func (c *ctrlUpTouchListener) OnTouchEnd(x, y float32) {
	c.buttonState = ctrlNop
}

// ctrlDown
type ctrlDownTouchListener struct {
	*sample
}

func (c *ctrlDownTouchListener) OnTouchBegin(x, y float32) {
	c.buttonState = ctrlDown
}

func (c *ctrlDownTouchListener) OnTouchMove(x, y float32) {
	c.buttonState = ctrlDown
}

func (c *ctrlDownTouchListener) OnTouchEnd(x, y float32) {
	c.buttonState = ctrlNop
}

// ButtonBlueTouchListener represents a listener object
// to notify touch event of Blue Button
type ButtonBlueTouchListener struct {
	*sample
}

// OnTouchBegin is called when Blue Button is Touched.
func (c *ButtonBlueTouchListener) OnTouchBegin(x, y float32) {
	if c.buttonReplaced {
		c.originalButtonColor()
	} else {
		c.replaceButtonColor()
	}
	c.simra.RemoveSprite(c.ball)
}

// OnTouchMove is called when Blue Button is Touched and moved.
func (c *ButtonBlueTouchListener) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when Blue Button is Touched and it is released.
func (c *ButtonBlueTouchListener) OnTouchEnd(x, y float32) {
	// nop
}

// ButtonRedTouchListener represents a listener object
// to notify touch event of Red Button
type ButtonRedTouchListener struct {
	*sample
}

// OnTouchBegin is called when Red Button is Touched.
func (c *ButtonRedTouchListener) OnTouchBegin(x, y float32) {
	if c.buttonReplaced {
		c.originalButtonColor()
	} else {
		c.replaceButtonColor()
	}
	c.simra.AddSprite(c.ball)
	tex := c.simra.NewImageTexture("ball.png",
		image.Rect(0, 0, c.ball.GetScale().W, c.ball.GetScale().H))
	c.ball.ReplaceTexture(tex)
}

// OnTouchMove is called when Red Button is Touched and moved.
func (c *ButtonRedTouchListener) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when Red Button is Touched and it is released.
func (c *ButtonRedTouchListener) OnTouchEnd(x, y float32) {
	// nop
}
