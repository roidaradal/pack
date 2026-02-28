package ds

import "fmt"

type String string

// Len returns the String length
func (s String) Len() Int {
	return Int(len(s))
}

// IsEmpty checks if the String is empty
func (s String) IsEmpty() Boolean {
	return s == ""
}

// NotEmpty checks if the String is not empty
func (s String) NotEmpty() Boolean {
	return s != ""
}

// Format uses Sprintf to build the string
func (s String) Format(args ...any) String {
	return String(fmt.Sprintf(string(s), args...))
}
