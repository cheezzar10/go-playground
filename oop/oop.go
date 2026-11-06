package oop

import (
	"fmt"
)

type Shape interface {
	Draw()
}

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("Point(%d, %d)", p.X, p.Y)
}

type Rectangle struct {
	TopLeft     Point
	BottomRight Point
}

func (r *Rectangle) Draw() {
	fmt.Printf("Rectangle(topLeft: %s, bottomRight: %s)::Draw()", r.TopLeft, r.BottomRight)
}
