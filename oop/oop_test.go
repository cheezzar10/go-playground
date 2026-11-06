package oop

import (
	"testing"
)

func TestRectangleDraw(t *testing.T) {
	rect := Rectangle{Point{2, 10}, Point{10, 2}}

	var shape Shape
	if shape != nil {
		t.Fatal("unexpected shape reference initial value")
	}
	// only pointer to rectangle implements Shape
	shape = &rect

	shape.Draw()
}
