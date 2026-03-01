package list

// All checks if all List items pass the ok function
func All[T any](items []T, ok func(T) bool) bool {
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
func AllWithIndex[T any](items []T, ok func(int, T) bool) bool {
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
func Any[T any](items []T, ok func(T) bool) bool {
	for _, item := range items {
		if ok(item) {
			return true
		}
	}
	return false
}

// AnyWithIndex checks if any List item passes the ok function: (index, item)
func AnyWithIndex[T any](items []T, ok func(int, T) bool) bool {
	for i, item := range items {
		if ok(i, item) {
			return true
		}
	}
	return false
}
