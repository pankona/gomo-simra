package image

import "image"

// Rectangle is a wrapper of image.Rectangle
type Rectangle struct {
	image.Rectangle
}

// Rect returns a Rectangle instance
func Rect(x0, y0, x1, y1 float32) Rectangle {
	return Rectangle{
		image.Rect(int(x0), int(y0), int(x1), int(y1)),
	}
}
