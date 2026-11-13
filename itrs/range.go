package itrs

import (
	"iter"
)

type Range iter.Seq[int]

func NewRange(stop int) Range {
	return func(yield func(n int) bool) {
		for i := range stop {
			if !yield(i) {
				return
			}
		}
	}
}
