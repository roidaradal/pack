// Package ds contains data structures
package ds

import "fmt"

// Int extends the int type
type Int int

// Uint extends the uint type
type Uint uint

// Float extends the float64 type
type Float float64

// String representation of Int
func (i Int) String() string {
	return fmt.Sprintf("%d", i)
}

// String representation of Uint
func (u Uint) String() string {
	return fmt.Sprintf("%d", u)
}

// String representation of Float
func (f Float) String() string {
	return fmt.Sprintf("%f", f)
}

// StringDecimal returns the float string, with N decimal places
func (f Float) StringDecimal(decimalPlaces int) string {
	format := fmt.Sprintf("%%.%df", decimalPlaces)
	return fmt.Sprintf(format, f)
}
