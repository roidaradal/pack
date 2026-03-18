// Package tst contains unit testing functions and TestCase structs
package tst

import (
	"testing"
)

type P1W1[I, R any] struct {
	P1 I
	W1 R
}

// AllCompare1 tests 1 pair of [actual, want] from the given generic test cases
func AllCompare1[T, R any](t *testing.T, testCases []T, name string, testFn func(T) (R, R), assert assertFn[R]) {
	for _, x := range testCases {
		actual, want := testFn(x)
		assert(t, name, actual, want)
	}
}

// AllP1W1 tests all P1W1 test cases
func AllP1W1[I, R any](t *testing.T, testCases []P1W1[I, R], name string, testFn func(I) R, assert assertFn[R]) {
	for _, x := range testCases {
		actual := testFn(x.P1)
		assert(t, name, actual, x.W1)
	}
}
