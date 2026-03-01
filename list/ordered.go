package list

import (
	"cmp"
)

// AllGreater checks if all List items are greater than given value
func AllGreater[T cmp.Ordered](items []T, value T) bool {
	return All(items, func(x T) bool {
		return x > value
	})
}

// AllGreaterEqual checks if all List items are greater or equal to given value
func AllGreaterEqual[T cmp.Ordered](items []T, value T) bool {
	return All(items, func(x T) bool {
		return x >= value
	})
}

// AllLess checks if all List items are lesser than given value
func AllLess[T cmp.Ordered](items []T, value T) bool {
	return All(items, func(x T) bool {
		return x < value
	})
}

// AllLessEqual checks if all List items are lesser or equal to given value
func AllLessEqual[T cmp.Ordered](items []T, value T) bool {
	return All(items, func(x T) bool {
		return x <= value
	})
}

// ArgMin finds the index of the minimum item of the List
func ArgMin[T cmp.Ordered](items []T) int {
	if len(items) == 0 {
		panic("empty list")
	}
	index, currMin := 0, items[0]
	for i := 1; i < len(items); i++ {
		if items[i] < currMin {
			index, currMin = i, items[i]
		}
	}
	return index
}

// ArgMax finds the index of the maximum item of the List
func ArgMax[T cmp.Ordered](items []T) int {
	if len(items) == 0 {
		panic("empty list")
	}
	index, currMax := 0, items[0]
	for i := 1; i < len(items); i++ {
		if items[i] > currMax {
			index, currMax = i, items[i]
		}
	}
	return index
}
