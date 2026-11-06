package lang

import (
	m "github.com/cheezzar10/go-playground/model"
	"testing"
	"unsafe"
)

func TestIterateNilSlice(t *testing.T) {
	IterateNilSlice()
}

func TestNamedReturns(t *testing.T) {
	NamedReturnsDemo()
}

func TestValuesSlice(t *testing.T) {
	origin := m.Point{X: 0.0, Y: 0.0}
	points := []m.Point{
		origin,
		{X: 1.0, Y: 2.0},
	}

	point := points[0]
	if point != origin {
		t.Error("unexpected origin coordinates")
	}

	point.X = 3.0
	if points[0].X != 0.0 {
		t.Error("slice element modified")
	}
}

func TestValuesArray(t *testing.T) {
	points := [...]m.Point{
		{X: 1.0, Y: 1.0},
		{X: 2.0, Y: 2.0},
	}

	pointsSize := unsafe.Sizeof(points)
	if pointsSize != 32 {
		t.Error("unexpected points size:", pointsSize)
	}

	var pointsSlice []m.Point
	pointsSlice = points[1:]
	pointsSliceLength := len(pointsSlice)
	if pointsSliceLength != 1 {
		t.Error("unexpected points slice length:", pointsSliceLength)
	}
}
