// Package list contains List functions
package list

import "github.com/roidaradal/pack/ds"

// All checks if all List items pass the ok function
func All[L ~[]T, T any](items L, ok func(T) bool) bool {
	if len(items) == 0 {
		return false
	}
	for _, item := range items {
		if !ok(item) {
			return false
		}
	}
	return true
}

// AllWithIndex checks if all List items pass the ok function: (index, item)
func AllWithIndex[L ~[]T, T any](items L, ok func(int, T) bool) bool {
	if len(items) == 0 {
		return false
	}
	for i, item := range items {
		if !ok(i, item) {
			return false
		}
	}
	return true
}

// Any checks if any List item passes the ok function
func Any[L ~[]T, T any](items L, ok func(T) bool) bool {
	for _, item := range items {
		if ok(item) {
			return true
		}
	}
	return false
}

// AnyWithIndex checks if any List item passes the ok function: (index, item)
func AnyWithIndex[L ~[]T, T any](items L, ok func(int, T) bool) bool {
	for i, item := range items {
		if ok(i, item) {
			return true
		}
	}
	return false
}

// Sum computes the sum of number items
func Sum[L ~[]T, T ds.Number](numbers L) T {
	var total T = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Product computes the product of number items
func Product[L ~[]T, T ds.Number](numbers L) T {
	var product T = 1
	for _, number := range numbers {
		product *= number
	}
	return product
}
