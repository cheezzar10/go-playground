package model

import (
	"fmt"
)

type Car struct {
	Maker string `json:"maker"`
	Model string `json:"model"`
}

type Point struct {
	X, Y float64
}

type Circle struct {
	Center Point
	Radius float64
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(Center=(%.1f, %.1f), Radius=%.1f)", c.Center.X, c.Center.X, c.Radius)

}

func New(maker string, model string) *Car {
	return &Car{maker, model}
}

type Ctx interface {
	Do()
}
