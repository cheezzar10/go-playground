package model

import (
	"testing"
)

type Garage struct {
	Cars []*Car
	Tags []string
}

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
		{Point{0.0, 0.0}, 3.0},
	}

	if len(circles) != 1 {
		t.Error("unexpected circles length")
	}
}

func TestGarage(t *testing.T) {
	garage := Garage{
		Cars: []*Car{
			{"BMW", "320d"},
			{"Mercedes-Benz", "E220d"},
		},
		Tags: []string{"main"},
	}

	if len(garage.Tags) != 1 {
		t.Errorf("unexpected 1 tag but was %d", len(garage.Tags))
	}
}
