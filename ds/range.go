package ds

import (
	"github.com/roidaradal/pack/list"
	"github.com/roidaradal/pack/types"
)

// Range represents a range of Integers from [start, end)
type Range[T types.Integer] struct {
	Start T
	End   T
}

// NewRange creates a new Range from [start, end)
func NewRange[T types.Integer](start, end T) *Range[T] {
	return &Range[T]{start, end}
}

// NewInclusiveRange creates a new Range from [first, last]
func NewInclusiveRange[T types.Integer](first, last T) *Range[T] {
	return &Range[T]{first, last + 1}
}

// ToList expands the Range into a List of Integers
func (r *Range[T]) ToList() []T {
	return list.Range(r.Start, r.End)
}

// Len returns the size of the Range
func (r *Range[T]) Len() int {
	return int(r.End - r.Start)
}

// Copy creates a new Range copy
func (r *Range[T]) Copy() *Range[T] {
	return NewRange[T](r.Start, r.End)
}

// Has checks if number is included in the Range
func (r *Range[T]) Has(number T) bool {
	return r.Start <= number && number < r.End
}

// Sum computes the sum of the Range
func (r *Range[T]) Sum() T {
	var total T = 0
	for number := r.Start; number < r.End; number++ {
		total += number
	}
	return total
}

// Product computes the product of the Range
func (r *Range[T]) Product() T {
	var product T = 1
	for number := r.Start; number < r.End; number++ {
		product *= number
	}
	return product
}
