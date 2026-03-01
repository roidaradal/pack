// Package types contains type functions
package types

// Number interface unifies the number types
type Number interface {
	Integer | Float
}

// Integer interface unifies the integer types
type Integer interface {
	~int | ~uint | ~int64
}

// Float interface unifies the float types
type Float interface {
	~float32 | ~float64
}
