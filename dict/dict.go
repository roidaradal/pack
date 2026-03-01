// Package dict contains Map functions
package dict

import (
	"cmp"
	"maps"
	"slices"

	"github.com/roidaradal/pack/ds"
)

// HasKey checks if Map has given key
func HasKey[K comparable, V any](items map[K]V, key K) bool {
	_, hasKey := items[key]
	return hasKey
}

// HasValue checks if Map has given value
func HasValue[K, V comparable](items map[K]V, value V) bool {
	for _, v := range items {
		if v == value {
			return true
		}
	}
	return false
}

// NoKey checks if Map does not have the given key
func NoKey[K comparable, V any](items map[K]V, key K) bool {
	return !HasKey(items, key)
}

// NoValue checks if Map does not have the given value
func NoValue[K, V comparable](items map[K]V, value V) bool {
	return !HasValue(items, value)
}

// SortedKeys returns the Map keys in sorted order
func SortedKeys[K cmp.Ordered, V any](items map[K]V) []K {
	keys := slices.Collect(maps.Keys(items))
	slices.Sort(keys)
	return keys
}

// SortedValues returns the Map values in sorted order
func SortedValues[K comparable, V cmp.Ordered](items map[K]V) []V {
	values := slices.Collect(maps.Values(items))
	slices.Sort(values)
	return values
}

// SortedEntries returns the Map entries in sorted key order
func SortedEntries[K cmp.Ordered, V any](items map[K]V) []ds.Entry[K, V] {
	keys := SortedKeys(items)
	entries := make([]ds.Entry[K, V], len(keys))
	for i, k := range keys {
		entries[i] = ds.Entry[K, V]{Key: k, Value: items[k]}
	}
	return entries
}

// SortValueLists sorts the list of values in place, for each key in the Map
func SortValueLists[K comparable, V cmp.Ordered](items map[K][]V) {
	for k, values := range items {
		slices.Sort(values)
		items[k] = values
	}
}
