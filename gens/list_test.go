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

	if got, want := list.Len(), 0; got != want {
		t.Errorf("initial length after construction = %d, want %d", got, want)
	}

	if got, want := list.Cap(), initialCap; got != want {
		t.Errorf("initial capacity after construction = %d, want %d", got, want)
	}

	wantElems := []int{1, 2}
	for _, e := range wantElems {
		list.Append(e)
	}

	if got, want := list.Len(), len(wantElems); got != want {
		t.Errorf("length after append = %d, want %d", got, want)
	}

	for i, e := range list.Iter() {
		if got, want := e, wantElems[i]; got != want {
			t.Errorf("unexpected list element at index %d: got = %d, want %d", i, got, want)
		}
	}

	desiredLen := 5
	for i := wantElems[len(wantElems)-1] + 1; list.Len() < desiredLen; i++ {
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

func TestListClip(t *testing.T) {
	const initialCap int = 10
	list := NewList[int](initialCap)
	wantElems := []int{1, 2, 3}
	for _, e := range wantElems {
		list.Append(e)
	}

	if got, want := list.Cap(), initialCap; got != want {
		t.Errorf("initial capacity = %d, want %d", got, want)
	}

	list.Clip()
	if got, want := list.Cap(), len(wantElems); got != want {
		t.Errorf("capacity after clip = %d, want %d", got, want)
	}
	if got, want := list.Len(), len(wantElems); got != want {
		t.Errorf("length after clip = %d, want %d", got, want)
	}
	for i, e := range list.Iter() {
		if got, want := e, wantElems[i]; got != want {
			t.Errorf("unexpected list element after clip at index %d: got = %d, want %d", i, got, want)
		}
	}
}
