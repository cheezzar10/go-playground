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
	const initialCap int = 1
	list := NewList[int](initialCap)

	initialLen := list.Len()
	if got, want := list.Len(), 0; got != want {
		t.Errorf("initial length after construction = %d, want %d", got, want)
	}

	if got, want := list.Cap(), initialCap; got != want {
		t.Errorf("initial capacity after construction = %d, want %d", got, want)
	}

	list.Append(1)
	if got, want := list.Len(), initialLen+1; got != want {
		t.Errorf("length after append = %d, want %d", got, want)
	}

	desiredLen := 5
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

func TestListReserve(t *testing.T) {
	list := NewList[int](2)
	list.Append(1)
	initialLen := list.Len()

	const increasedCap = 10
	list.Reserve(increasedCap)
	if got := list.Cap(); got != increasedCap {
		t.Errorf("capacity after reserve = %d, want %d", got, increasedCap)
	}
	if got := list.Len(); got != initialLen {
		t.Errorf("length after reserve = %d, want %d", got, initialLen)
	}

	list.Reserve(1)
	if got := list.Cap(); got != increasedCap {
		t.Errorf("capacity after reducing reserve = %d, want %d", got, increasedCap)
	}
	if got := list.Len(); got != initialLen {
		t.Errorf("length after reducing reserve = %d, want %d", got, initialLen)
	}
}
