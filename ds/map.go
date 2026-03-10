package ds

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Map[K comparable, V any] map[K]V

// String returns the string representation of Map, where keys are sorted
func (m Map[K, V]) String() string {
	out := make([]string, 0, len(m))
	for k, v := range m {
		out = append(out, fmt.Sprintf("%v: %v", k, v))
	}
	slices.Sort(out)
	return "{" + strings.Join(out, ", ") + "}"
}

// Len returns the Map size
func (m Map[K, V]) Len() int {
	return len(m)
}

// IsEmpty checks if Map is empty
func (m Map[K, V]) IsEmpty() bool {
	return len(m) == 0
}

// NotEmpty checks if Map is not empty
func (m Map[K, V]) NotEmpty() bool {
	return len(m) > 0
}

// Clear removes all Map entries
func (m Map[K, V]) Clear() {
	clear(m)
}

// Copy creates a new Map with copied entries
func (m Map[K, V]) Copy() Map[K, V] {
	m2 := make(Map[K, V], len(m))
	maps.Copy(m2, m)
	return m2
}
