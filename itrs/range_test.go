package itrs

import (
	"slices"
	"testing"
)

func TestRangeFull(t *testing.T) {
	got := []int{}
	for i := range NewRange(5) {
		got = append(got, i)
	}

	if want := []int{0, 1, 2, 3, 4}; !slices.Equal(want, got) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestRangeExit(t *testing.T) {
	got := []int{}
	for i := range NewRange(3) {
		got = append(got, i)
		if i == 1 {
			break
		}
	}

	if want := []int{0, 1}; !slices.Equal(want, got) {
		t.Errorf("expected %v, got %v", want, got)
	}
}
