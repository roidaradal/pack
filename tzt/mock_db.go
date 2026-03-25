package tzt

import (
	"database/sql"
	"errors"
	"slices"
)

var errNotFound = errors.New("not found")

type Conn[T any] struct {
	items  []T
	err    error
	testFn func(T) bool
	rowFn  func([]T) ([]any, error)
}

func NewConn[T any](items ...T) *Conn[T] {
	return new(Conn[T]{items: items, err: nil})
}

func (c *Conn[T]) Begin() (*Tx, error) {
	// TODO: Update for Tx
	return nil, c.err
}

func (c *Conn[T]) Exec(query string, args ...any) (sql.Result, error) {
	// TODO: Implement
	return nil, c.err
}

func (c *Conn[T]) Query(query string, args ...any) (*Rows, error) {
	// TODO: Implement
	return nil, c.err
}

func (c *Conn[T]) QueryRow(query string, args ...any) *Row {
	if c.testFn == nil || c.rowFn == nil || c.err != nil {
		return NewRow()
	}
	validItems := make([]T, 0, len(c.items))
	for _, item := range c.items {
		if c.testFn(item) {
			validItems = append(validItems, item)
		}
	}
	items, err := c.rowFn(validItems)
	if err != nil {
		return NewRow()
	}
	return NewRow(items...)
}

func (c *Conn[T]) SetError(err error) {
	c.err = err
}

func (c *Conn[T]) PrepareRow(testFn func(T) bool, rowFn func([]T) ([]any, error)) {
	c.testFn = testFn
	c.rowFn = rowFn
}

func PrepFn[T any](conn *Conn[T], testFn func(T) bool, rowFn func([]T) ([]any, error)) func() {
	return func() {
		conn.SetError(nil)
		conn.PrepareRow(testFn, rowFn)
	}
}

func PrepOne[T any](conn *Conn[T], testFn func(T) bool, rowFn func(T) []any) func() {
	return func() {
		conn.SetError(nil)
		conn.PrepareRow(testFn, func(items []T) ([]any, error) {
			if len(items) == 0 {
				return nil, errNotFound
			}
			return rowFn(items[0]), nil
		})
	}
}

func PrepSortOne[T any](conn *Conn[T], testFn func(T) bool, sortFn func(T, T) int, rowFn func(T) []any) func() {
	return func() {
		conn.SetError(nil)
		conn.PrepareRow(testFn, func(items []T) ([]any, error) {
			if len(items) == 0 {
				return nil, errNotFound
			}
			slices.SortFunc(items, sortFn)
			return rowFn(items[0]), nil
		})
	}
}
