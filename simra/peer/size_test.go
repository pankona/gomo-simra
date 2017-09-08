package peer

import (
	"testing"

	"golang.org/x/mobile/event/size"
)

func TestGetScreenSizePeer(t *testing.T) {
	s1 := GetScreenSizePeer()
	s2 := GetScreenSizePeer()
	if s1 != s2 {
		t.Errorf("GetScreenSizePeer didn't return same address. s1 = %v, s2 = %v", s1, s2)
	}
}

func TestSetScreenSize(t *testing.T) {
	s := &screenSize{}
	s.SetScreenSize(
		size.Event{
			WidthPt:  10,
			HeightPt: 20,
		},
	)

	s.SetScreenSize(
		size.Event{
			WidthPt:  20,
			HeightPt: 10,
		},
	)
}

func TestSetDesiredScreenSize(t *testing.T) {
	s := &screenSize{}
	s.SetDesiredScreenSize(1980, 1080)
}

func TestFitToHeight(t *testing.T) {
	s := &screenSize{}
	s.SetScreenSize(
		size.Event{
			WidthPt:  20,
			HeightPt: 10,
		},
	)
	s.SetDesiredScreenSize(1980, 1080)
}
