package qb

import (
	"cmp"
	"slices"
	"strings"

	"github.com/zeroibot/pack/ds"
	"github.com/zeroibot/pack/list"
)

// Condition interface unifies all Condition objects:
// BuildCondition() method outputs the condition string and parameter values
type Condition interface {
	BuildCondition() (string, []any) // Return (condition string, parameter values)
}

// DualCondition interface holds a Condition Builder and a struct Tester
type DualCondition[T any] interface {
	Condition
	Test(T) bool
}

// NoCondition creates a matchAllCondition
func NoCondition() Condition {
	return matchAllCondition{}
}

// NoCondition2 creates a matchAllCombo
func NoCondition2[T any]() DualCondition[T] {
	return matchAllCombo[T]{}
}

// Equal creates an Equal Condition
func Equal[T comparable](this *Instance, fieldRef *T, value T) Condition {
	return newValueCondition(this, fieldRef, value, opEqual)
}

// Equal2 creates an Equal Combo
func Equal2[T any, V comparable](this *Instance, fieldRef *V, value V) DualCondition[T] {
	fieldName := this.getFieldName(fieldRef)
	test := createFieldValueTest[T](fieldName, func(fieldValue V) bool {
		return fieldValue == value
	})
	condition := newValueCondition(this, fieldRef, value, opEqual)
	return newValueCombo(condition, test)
}

// NotEqual creates a NotEqual Condition
func NotEqual[T comparable](this *Instance, fieldRef *T, value T) Condition {
	return newValueCondition(this, fieldRef, value, opNotEqual)
}

// NotEqual2 creates a NotEqual Combo
func NotEqual2[T any, V comparable](this *Instance, fieldRef *V, value V) DualCondition[T] {
	fieldName := this.getFieldName(fieldRef)
	test := createFieldValueTest[T](fieldName, func(fieldValue V) bool {
		return fieldValue != value
	})
	condition := newValueCondition(this, fieldRef, value, opNotEqual)
	return newValueCombo(condition, test)
}

// Prefix creates a Prefix Condition
func Prefix(this *Instance, fieldRef *string, value string) Condition {
	return newValueCondition(this, fieldRef, value, opPrefix)
}

// Prefix2 creates a Prefix Combo
func Prefix2[T any](this *Instance, fieldRef *string, value string) DualCondition[T] {
	fieldName := this.getFieldName(fieldRef)
	test := createFieldValueTest[T](fieldName, func(fieldValue string) bool {
		return strings.HasPrefix(fieldValue, value)
	})
	condition := newValueCondition(this, fieldRef, value, opPrefix)
	return newValueCombo(condition, test)
}

// Suffix creates a Suffix Condition
func Suffix(this *Instance, fieldRef *string, value string) Condition {
	return newValueCondition(this, fieldRef, value, opSuffix)
}

// Suffix2 creates a Suffix Combo
func Suffix2[T any](this *Instance, fieldRef *string, value string) DualCondition[T] {
	fieldName := this.getFieldName(fieldRef)
	test := createFieldValueTest[T](fieldName, func(fieldValue string) bool {
		return strings.HasSuffix(fieldValue, value)
	})
	condition := newValueCondition(this, fieldRef, value, opSuffix)
	return newValueCombo(condition, test)
}

// Substring creates a Substring Condition
func Substring(this *Instance, fieldRef *string, value string) Condition {
	return newValueCondition(this, fieldRef, value, opSubstring)
}

// Substring2 creates a Substring Combo
func Substring2[T any](this *Instance, fieldRef *string, value string) DualCondition[T] {
	fieldName := this.getFieldName(fieldRef)
	test := createFieldValueTest[T](fieldName, func(fieldValue string) bool {
		return strings.Contains(fieldValue, value)
	})
	condition := newValueCondition(this, fieldRef, value, opSubstring)
	return newValueCombo(condition, test)
}

// Greater creates a GreaterThan Condition
func Greater[T cmp.Ordered](this *Instance, fieldRef *T, value T) Condition {
	return newValueCondition(this, fieldRef, value, opGreater)
}

// Greater2 creates a GreaterThan Combo
func Greater2[T any, V cmp.Ordered](this *Instance, fieldRef *V, value V) DualCondition[T] {
	fieldName := this.getFieldName(fieldRef)
	test := createFieldValueTest[T](fieldName, func(fieldValue V) bool {
		return fieldValue > value
	})
	condition := newValueCondition(this, fieldRef, value, opGreater)
	return newValueCombo(condition, test)
}

// GreaterEqual creates a GreaterThanOrEqual Condition
func GreaterEqual[T cmp.Ordered](this *Instance, fieldRef *T, value T) Condition {
	return newValueCondition(this, fieldRef, value, opGreaterEqual)
}

// GreaterEqual2 creates a GreaterThanOrEqual Combo
func GreaterEqual2[T any, V cmp.Ordered](this *Instance, fieldRef *V, value V) DualCondition[T] {
	fieldName := this.getFieldName(fieldRef)
	test := createFieldValueTest[T](fieldName, func(fieldValue V) bool {
		return fieldValue >= value
	})
	condition := newValueCondition(this, fieldRef, value, opGreaterEqual)
	return newValueCombo(condition, test)
}

// Lesser creates a LesserThan Condition
func Lesser[T cmp.Ordered](this *Instance, fieldRef *T, value T) Condition {
	return newValueCondition(this, fieldRef, value, opLesser)
}

// Lesser2 creates a LesserThan Combo
func Lesser2[T any, V cmp.Ordered](this *Instance, fieldRef *V, value V) DualCondition[T] {
	fieldName := this.getFieldName(fieldRef)
	test := createFieldValueTest[T](fieldName, func(fieldValue V) bool {
		return fieldValue < value
	})
	condition := newValueCondition(this, fieldRef, value, opLesser)
	return newValueCombo(condition, test)
}

// LesserEqual creates a LesserThanOrEqual Condition
func LesserEqual[T cmp.Ordered](this *Instance, fieldRef *T, value T) Condition {
	return newValueCondition(this, fieldRef, value, opLesserEqual)
}

// LesserEqual2 creates a LesserThanOrEqual Combo
func LesserEqual2[T any, V cmp.Ordered](this *Instance, fieldRef *V, value V) DualCondition[T] {
	fieldName := this.getFieldName(fieldRef)
	test := createFieldValueTest[T](fieldName, func(fieldValue V) bool {
		return fieldValue <= value
	})
	condition := newValueCondition(this, fieldRef, value, opLesserEqual)
	return newValueCombo(condition, test)
}

// In creates an In Condition
func In[T comparable](this *Instance, fieldRef *T, values ds.List[T]) Condition {
	return newListCondition(this, fieldRef, values, opIn, opEqual)
}

// In2 creates an In Combo
func In2[T any, V comparable](this *Instance, fieldRef *V, values ds.List[V]) DualCondition[T] {
	fieldName := this.getFieldName(fieldRef)
	test := createFieldValueTest[T](fieldName, func(fieldValue V) bool {
		return slices.Contains(values, fieldValue)
	})
	condition := newListCondition(this, fieldRef, values, opIn, opEqual)
	return newListCombo(condition, test)
}

// NotIn creates a NotIn Condition
func NotIn[T comparable](this *Instance, fieldRef *T, values ds.List[T]) Condition {
	return newListCondition(this, fieldRef, values, opNotIn, opNotEqual)
}

// NotIn2 creates a NotIn Combo
func NotIn2[T any, V comparable](this *Instance, fieldRef *V, values ds.List[V]) DualCondition[T] {
	fieldName := this.getFieldName(fieldRef)
	test := createFieldValueTest[T](fieldName, func(fieldValue V) bool {
		return !slices.Contains(values, fieldValue)
	})
	condition := newListCondition(this, fieldRef, values, opNotIn, opNotEqual)
	return newListCombo(condition, test)
}

// And creates an And Condition
func And(conditions ...Condition) Condition {
	return newMultiCondition(opAnd, conditions...)
}

// And2 creates an And Combo
func And2[T any](conditions ...DualCondition[T]) DualCondition[T] {
	test := func(item T) bool {
		return list.All(conditions, func(c DualCondition[T]) bool {
			return c.Test(item)
		})
	}
	return newMultiCombo(conditions, opAnd, test)
}

// Or creates an Or Condition
func Or(conditions ...Condition) Condition {
	return newMultiCondition(opOr, conditions...)
}

// Or2 creates an Or Combo
func Or2[T any](conditions ...DualCondition[T]) DualCondition[T] {
	test := func(item T) bool {
		return list.Any(conditions, func(c DualCondition[T]) bool {
			return c.Test(item)
		})
	}
	return newMultiCombo(conditions, opOr, test)
}
