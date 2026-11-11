package lang

import (
	"fmt"
	"testing"
	"unsafe"

	m "github.com/cheezzar10/go-playground/model"
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

func TestSliceToString(t *testing.T) {
	vendors := []string{"GHSA", "NVD"}
	fmt.Printf("vendors: %v\n", vendors)

	vendorsStr := fmt.Sprintf("%s", vendors)
	if vendorsStr != "[GHSA NVD]" {
		t.Error("unexpected:", vendorsStr)
	}
}

func TestUnicodeString(t *testing.T) {
	bts := []byte{0x3e, 0x3d, 0x32, 0x2e, 0x31, 0x39, 0x2e, 0x30, 0x2c, 0x20, 0x3c, 0x32, 0x2e, 0x32, 0x31, 0x2e, 0x37}
	fmt.Printf("bytes string: %s\n", string(bts))
	str := "\u003e=2.19.0, \u003c2.21.7"
	for _, b := range []byte(str) {
		fmt.Printf("0x%x, ", b)
	}
	fmt.Println()

	symbols := "\u4e16\u754c"
	fmt.Printf("symbols length: %d\n", len(symbols))
	for _, s := range symbols {
		fmt.Printf("%q ", s)
	}
	fmt.Println()
}
