package list

import "github.com/roidaradal/pack/types"

// Map applies the convert function for each List item
func Map[T, V any](items []T, convert func(T) V) []V {
	results := make([]V, len(items))
	for i, item := range items {
		results[i] = convert(item)
	}
	return results
}

// MapWithIndex applies the convert function for each List item, with index
func MapWithIndex[T, V any](items []T, convert func(int, T) V) []V {
	results := make([]V, len(items))
	for i, item := range items {
		results[i] = convert(i, item)
	}
	return results
}

// MapList maps the list indexes to given List.
// Can have zero values for invalid indexes
func MapList[T any](indexes []int, items []T) []T {
	results := make([]T, len(items))
	numItems := len(items)
	for i, idx := range indexes {
		if 0 <= idx && idx < numItems {
			results[i] = items[idx]
		}
	}
	return results
}

// MapLookup maps the keys to the given lookup Map.
// Can have zero values for invalid indexes
func MapLookup[K comparable, V any](keys []K, lookup map[K]V) []V {
	results := make([]V, len(keys))
	for i, key := range keys {
		results[i] = lookup[key]
	}
	return results
}

// MapIf combines Map and Filter: apply the convert function, and filter items based on the result flag
func MapIf[T, V any](items []T, convert func(T) (V, bool)) []V {
	results := NewEmpty[V](len(items))
	for _, item := range items {
		if item2, ok := convert(item); ok {
			results = append(results, item2)
		}
	}
	return results
}

// Filter filters the list by only keeping items that pass the keep function
func Filter[T any](items []T, keep func(T) bool) []T {
	results := make([]T, 0, len(items))
	for _, item := range items {
		if keep(item) {
			results = append(results, item)
		}
	}
	return results
}

// FilterWithIndex filters the list by only keeping items that pass the keep function: (index, item)
func FilterWithIndex[T any](items []T, keep func(int, T) bool) []T {
	results := make([]T, 0, len(items))
	for i, item := range items {
		if keep(i, item) {
			results = append(results, item)
		}
	}
	return results
}

// Reduce applies the reduce function to each item to get the final result
// Reducer function has the signature (result, item) => result
func Reduce[T any](items []T, reducer func(T, T) T, initial T) T {
	current := initial
	for _, item := range items {
		current = reducer(current, item)
	}
	return current
}

// Apply applies the task function to each item
func Apply[T any](items []T, task func(T) T) []T {
	results := make([]T, len(items))
	for i, item := range items {
		results[i] = task(item)
	}
	return results
}

// Sum computes the sum of number items
func Sum[T types.Number](numbers []T) T {
	var total T = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Product computes the product of number items
func Product[T types.Number](numbers []T) T {
	var product T = 1
	for _, number := range numbers {
		product *= number
	}
	return product
}
