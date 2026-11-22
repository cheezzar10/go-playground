package gens

type List[T comparable] struct {
	elems []T
}

func NewList[T comparable](cap int) *List[T] {
	return &List[T]{
		elems: make([]T, 0, cap),
	}
}

func (list List[T]) Len() int {
	return len(list.elems)
}

func (list List[T]) Cap() int {
	return cap(list.elems)
}

func (list *List[T]) Append(elem T) {
	list.elems = append(list.elems, elem)
}

func (list *List[T]) Reserve(cap int) {
	if cap <= list.Cap() {
		return
	}

	newElems := make([]T, list.Len(), cap)
	copy(newElems, list.elems)
	list.elems = newElems
}

// add list iterators, to check actual list content after modifications

// Clip() - shrinks capacity to length
