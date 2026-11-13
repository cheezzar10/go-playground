package gens

import (
	"testing"
)

func TestListNew(t *testing.T) {
	const capacity = 10
	list := NewList[int](capacity)

	if got := list.Cap(); got != capacity {
		t.Errorf("capacity = %d, want %d", got, capacity)
	}
	if got := list.Len(); got != 0 {
		t.Errorf("length = %d, want 0", got)
	}
}

func TestListAppend(t *testing.T) {
	list := NewList[int](1)
	initialLen := list.Len()
	initialCap := list.Cap()

	list.Append(1)
	if got, want := list.Len(), initialLen+1; got != want {
		t.Errorf("length after append = %d, want %d", got, want)
	}

	desiredLen := 4
	for i := 2; i <= desiredLen; i++ {
		list.Append(i)
	}
	if got := list.Len(); got != desiredLen {
		t.Errorf("length after consecutive appends = %d, want %d", got, desiredLen)
	}
	if got := list.Cap(); got <= initialCap {
		t.Errorf("capacity after consecutive appends = %d, want greater than %d", got, initialCap)
	}
}
