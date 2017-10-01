package image

import "image"

type Rectangle struct {
	image.Rectangle
}

func Rect(x0, y0, x1, y1 float32) Rectangle {
	return Rectangle{
		image.Rect(int(x0), int(y0), int(x1), int(y1)),
	}
}
