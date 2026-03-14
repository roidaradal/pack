package qb

// Missing Combo: uses missingCondition
type missingCombo[T any] struct {
	missingCondition
}

func (c missingCombo[T]) Test(_ T) bool {
	return false
}

// MatchAll Combo: uses matchAllCondition
type matchAllCombo[T any] struct {
	matchAllCondition
}

func (c matchAllCombo[T]) Test(_ T) bool {
	return true
}

// Value Combo: uses valueCondition
type valueCombo[T any] struct {
	valueCondition
	test TestFn[T]
}

// newValueCombo creates a new valueCombo
func newValueCombo[T any](condition valueCondition, test TestFn[T]) valueCombo[T] {
	return valueCombo[T]{condition, test}
}

func (c valueCombo[T]) Test(item T) bool {
	return c.test(item)
}

// List Combo: uses listCondition
type listCombo[T any] struct {
	listCondition
	test TestFn[T]
}

// newListCombo creates a new listCombo
func newListCombo[T any](condition listCondition, test TestFn[T]) listCombo[T] {
	return listCombo[T]{condition, test}

}

func (c listCombo[T]) Test(item T) bool {
	return c.test(item)
}

// Multi Combo: uses multiCondition
type multiCombo[T any] struct {
	multiCondition
	test TestFn[T]
}

// newMultiCombo creates a new multiCombo
func newMultiCombo[T any](condition multiCondition, test TestFn[T]) multiCombo[T] {
	return multiCombo[T]{condition, test}
}

func (c multiCombo[T]) Test(item T) bool {
	return c.test(item)
}
