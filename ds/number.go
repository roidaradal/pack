package ds

import "fmt"

// Int extends the int type
type Int int

// Uint extends the uint type
type Uint uint

// Float extends the float64 type
type Float float64

// String representation of Int
func (i Int) String() String {
	return String("%d").Format(i)
}

// String representation of Uint
func (u Uint) String() String {
	return String("%d").Format(u)
}

// String representation of Float
func (f Float) String() String {
	return String("%f").Format(f)
}

// StringDecimal returns the float string, with N decimal places
func (f Float) StringDecimal(decimalPlaces int) String {
	format := fmt.Sprintf("%%.%df", decimalPlaces)
	return String(format).Format(f)
}

// Abs gets the absolute value of the Int
func (i Int) Abs() Int {
	if i < 0 {
		return -i
	}
	return i
}
