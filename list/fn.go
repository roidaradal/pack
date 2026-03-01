package list

import "github.com/roidaradal/pack/ds"

// Map applies the convert function for each List item
func Map[L ~[]T, T, V any](items L, convert func(T) V) ds.List[V] {
	results := make(ds.List[V], len(items))
	for i, item := range items {
		results[i] = convert(item)
	}
	return results
}

// MapWithIndex applies the convert function for each List item, with index
func MapWithIndex[L ~[]T, T, V any](items L, convert func(int, T) V) ds.List[V] {
	results := make(ds.List[V], len(items))
	for i, item := range items {
		results[i] = convert(i, item)
	}
	return results
}

// MapList maps the list indexes to given List.
// Can have zero values for invalid indexes
func MapList[I ~[]int, L ~[]T, T any](indexes I, items L) ds.List[T] {
	results := make(ds.List[T], len(items))
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
func MapLookup[I ~[]K, L ~map[K]V, K comparable, V any](keys I, lookup L) ds.List[V] {
	results := make(ds.List[V], len(keys))
	for i, key := range keys {
		results[i] = lookup[key]
	}
	return results
}

// MapIf combines Map and Filter: apply the convert function, and filter items based on the result flag
func MapIf[L ~[]T, T, V any](items L, convert func(T) (V, bool)) ds.List[V] {
	results := ds.NewEmptyList[V](len(items))
	for _, item := range items {
		if item2, ok := convert(item); ok {
			results = append(results, item2)
		}
	}
	return results
}

// Filter filters the list by only keeping items that pass the keep function
func Filter[L ~[]T, T any](items L, keep func(T) bool) L {
	results := make(L, 0, len(items))
	for _, item := range items {
		if keep(item) {
			results = append(results, item)
		}
	}
	return results
}

// FilterWithIndex filters the list by only keeping items that pass the keep function: (index, item)
func FilterWithIndex[L ~[]T, T any](items L, keep func(int, T) bool) L {
	results := make(L, 0, len(items))
	for i, item := range items {
		if keep(i, item) {
			results = append(results, item)
		}
	}
	return results
}

// Reduce applies the reduce function to each item to get the final result
// Reducer function has the signature (result, item) => result
func Reduce[L ~[]T, T any](items L, reducer func(T, T) T, initial T) T {
	current := initial
	for _, item := range items {
		current = reducer(current, item)
	}
	return current
}

// Apply applies the task function to each item
func Apply[L ~[]T, T any](items L, task func(T) T) L {
	results := make(L, len(items))
	for i, item := range items {
		results[i] = task(item)
	}
	return results
}
