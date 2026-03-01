// Package list contains list functions
package list

import (
	"math/rand/v2"
	"slices"

	"github.com/roidaradal/pack/types"
)

// NewEmpty creates an empty list, with given capacity
func NewEmpty[T any](capacity int) []T {
	return make([]T, 0, capacity)
}

// Range creates a new list, containing numbers from [start, end)
func Range[T types.Integer](start, end T) []T {
	items := NewEmpty[T](int(end - start))
	for i := start; i < end; i++ {
		items = append(items, i)
	}
	return items
}

// InclusiveRange creates a new list, containing numbers from [first, last]
func InclusiveRange[T types.Integer](first, last T) []T {
	return Range(first, last+1)
}

// RepeatedItem creates a new list, with <value> repeated <count> times
func RepeatedItem[T any](value T, count int) []T {
	return slices.Repeat([]T{value}, count)
}

// Len returns the list length
func Len[T any](items []T) int {
	return len(items)
}

// Cap returns the list capacity
func Cap[T any](items []T) int {
	return cap(items)
}

// IsEmpty checks if list is empty
func IsEmpty[T any](items []T) bool {
	return len(items) == 0
}

// NotEmpty checks if list is not empty
func NotEmpty[T any](items []T) bool {
	return len(items) > 0
}

// Copy creates a new list with copied items
func Copy[T any](items []T) []T {
	return append([]T{}, items...)
}

// ToAny creates a list of <any> items from list
func ToAny[T any](items []T) []any {
	items2 := make([]any, len(items))
	for i, item := range items {
		items2[i] = item
	}
	return items2
}

// Last returns the nth item from the back of the list (starts at 1)
// Panics if rank is invalid
func Last[T any](items []T, rank int) T {
	numItems := len(items)
	if rank > numItems || rank <= 0 {
		panic("invalid rank")
	}
	return items[numItems-rank]
}

// Shuffle shuffles the list in place
func Shuffle[T any](items []T) {
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
}

// GetRandom gets a random item from list
// Panics if list is empty
func GetRandom[T any](items []T) T {
	if len(items) == 0 {
		panic("empty list")
	}
	return items[rand.IntN(len(items))]
}
