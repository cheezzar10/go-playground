package gens

type List[T comparable] struct {
	elems []T
}

func NewList[T comparable](cap int) *List[T] {
	return &List[T]{
		elems: make([]T, 0, cap),
	}
}

func (list *List[T]) Len() int {
	return len(list.elems)
}

func (list *List[T]) Cap() int {
	return cap(list.elems)
}

func (list *List[T]) Append(elem T) {
	list.elems = append(list.elems, elem)
}
