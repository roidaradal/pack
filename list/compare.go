package list

import "slices"

// IndexOf gets the index of given item, returns -1 if not in list
func IndexOf[T comparable](items []T, item T) int {
	return slices.Index(items, item)
}

// AllIndexOf gets all indexes of given item
func AllIndexOf[T comparable](items []T, item T) []int {
	indexes := make([]int, 0, len(items))
	for i, x := range items {
		if x == item {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// Remove removes the first item from list that matches given item
func Remove[T comparable](items []T, item T) ([]T, bool) {
	index := IndexOf(items, item)
	if index < 0 {
		return items, false
	}
	result := slices.Delete(items, index, index+1)
	return result, true
}

// RemoveAll removes all items from list that matches given item
func RemoveAll[T comparable](items []T, item T) []T {
	return slices.DeleteFunc(items, func(x T) bool {
		return x == item
	})
}

// Has checks if any list items are equal to given value
func Has[T comparable](items []T, item T) bool {
	return slices.Contains(items, item)
}

// HasNo checks if no list items are equal to given value
func HasNo[T comparable](items []T, item T) bool {
	return !slices.Contains(items, item)
}

// GetOrDefault returns the first item that matches given item, or returns default value
func GetOrDefault[T comparable](items []T, item T, defaultValue T) T {
	index := IndexOf(items, item)
	if index < 0 {
		return defaultValue
	}
	return items[index]
}

// AllEqual checks if all list items are equal to given value
func AllEqual[T comparable](items []T, item T) bool {
	if len(items) == 0 {
		return false
	}
	for _, x := range items {
		if x != item {
			return false
		}
	}
	return true
}

// AllTrue checks if all list items are true
func AllTrue(items []bool) bool {
	return AllEqual(items, true)
}

// AllFalse checks if all list items are false
func AllFalse(items []bool) bool {
	return AllEqual(items, false)
}

// AnyTrue checks if any list item is true
func AnyTrue(items []bool) bool {
	return Has(items, true)
}

// AnyFalse checks if any list item is false
func AnyFalse(items []bool) bool {
	return Has(items, false)
}
