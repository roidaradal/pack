package qb

// DualCondition interface holds a Condition Builder and a struct Tester
type DualCondition[T any] interface {
	Condition
	Test(T) bool
}
