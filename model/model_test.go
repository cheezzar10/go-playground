package model

import (
	"testing"
)

func TestCircle(t *testing.T) {
	circle := Circle{
		Center: Point{1.0, 1.0},
		Radius: 2.0,
	}

	if circle.String() != "Circle(Center=(1.0, 1.0), Radius=2.0)" {
		t.Error("unexpected circle String() result")
	}
}

func TestCircles(t *testing.T) {
	circles := []Circle{
		Circle{Point{0.0, 0.0}, 3.0},
	}

	if len(circles) != 1 {
		t.Error("unexpected circles length")
	}
}
