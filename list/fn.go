package list

import "github.com/roidaradal/pack/number"

// CountFunc counts the number of items that passes the ok function
func CountFunc[T any](items []T, ok func(T) bool) int {
	count := 0
	for _, item := range items {
		if ok(item) {
			count += 1
		}
	}
	return count
}

// Sum computes the sum of number items
func Sum[T number.Type](numbers []T) T {
	var total T = 0
	for _, x := range numbers {
		total += x
	}
	return total
}

// Product computes the product of number items
func Product[T number.Type](numbers []T) T {
	var product T = 1
	for _, x := range numbers {
		product *= x
	}
	return product
}
