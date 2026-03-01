package list

import (
	"slices"

	"github.com/roidaradal/pack/ds"
)

// AllEqual checks if all List items are equal to given value
func AllEqual[L ~[]T, T comparable](items L, value T) bool {
	if len(items) == 0 {
		return false
	}
	for _, item := range items {
		if item != value {
			return false
		}
	}
	return true
}

// AllNotEqual checks if all List items are not equal to given value
func AllNotEqual[L ~[]T, T comparable](items L, value T) bool {
	return !slices.Contains(items, value)
}

// AnyEqual checks if any List items are equal to given value
func AnyEqual[L ~[]T, T comparable](items L, value T) bool {
	return slices.Contains(items, value)
}

// AllTrue checks if all List items are true
func AllTrue[L ~[]bool](items L) bool {
	return AllEqual(items, true)
}

// AllFalse checks if all List items are false
func AllFalse[L ~[]bool](items L) bool {
	return AllEqual(items, false)
}

// AnyTrue checks if any List item is true
func AnyTrue[L ~[]bool](items L) bool {
	return AnyEqual(items, true)
}

// AnyFalse checks if any List item is false
func AnyFalse[L ~[]bool](items L) bool {
	return AnyEqual(items, false)
}

// AllSame checks if all List items are the same
func AllSame[L ~[]T, T comparable](items L) bool {
	return len(Tally(items)) == 1
}

// AllUnique checks if all List items are unique
func AllUnique[L ~[]T, T comparable](items L) bool {
	return len(Tally(items)) == len(items)
}

// CountUnique counts the unique items in the List
func CountUnique[L ~[]T, T comparable](items L) int {
	return len(Tally(items))
}

// Deduplicate removes duplicates from the List, preserving the order of items
func Deduplicate[L ~[]T, T comparable](items L) L {
	unique := make(L, 0, len(items))
	done := make(map[T]bool)
	for _, item := range items {
		if done[item] {
			continue
		}
		unique = append(unique, item)
		done[item] = true
	}
	return unique
}

// Tally computes the number of occurrence of each item in the List
func Tally[L ~[]T, T comparable](items L) ds.Map[T, int] {
	count := make(ds.Map[T, int])
	for _, item := range items {
		count[item] += 1
	}
	return count
}

// Count counts the number of occurrence of given value in the List
func Count[L ~[]T, T comparable](items L, value T) int {
	count := 0
	for _, item := range items {
		if item == value {
			count += 1
		}
	}
	return count
}

// IndexLookup creates a lookup Map of { item : index } from the List.
// This loses data if items are not unique
func IndexLookup[L ~[]T, T comparable](items L) ds.Map[T, int] {
	lookup := make(ds.Map[T, int])
	for i, item := range items {
		lookup[item] = i
	}
	return lookup
}

// GroupBy groups List items using the key function
func GroupBy[L ~[]T, T any, K comparable](items L, keyFn func(T) K) ds.Map[K, L] {
	group := make(ds.Map[K, L])
	for _, item := range items {
		key := keyFn(item)
		group[key] = append(group[key], item)
	}
	return group
}
