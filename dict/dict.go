// Package dict contains Map functions
package dict

import "maps"

// Len returns the map size
func Len[K comparable, V any](items map[K]V) int {
	return len(items)
}

// IsEmpty checks if map is empty
func IsEmpty[K comparable, V any](items map[K]V) bool {
	return len(items) == 0
}

// NotEmpty checks if map is not empty
func NotEmpty[K comparable, V any](items map[K]V) bool {
	return len(items) > 0
}

// Clear removes all map entries
func Clear[K comparable, V any](items map[K]V) {
	clear(items)
}

// Copy creates a new map with copied entries
func Copy[K comparable, V any](items map[K]V) map[K]V {
	items2 := make(map[K]V, len(items))
	maps.Copy(items2, items)
	return items2
}
