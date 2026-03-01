package ds

import (
	"strconv"
	"strings"
)

type String string

// Len returns the String length
func (s String) Len() int {
	return len(s)
}

// IsEmpty checks if the String is empty
func (s String) IsEmpty() bool {
	return s == ""
}

// NotEmpty checks if the String is not empty
func (s String) NotEmpty() bool {
	return s != ""
}

// ToInt parses the string as an int, defaults to 0 if invalid Int
func (s String) ToInt() int {
	text := strings.TrimSpace(string(s))
	number, err := strconv.Atoi(text)
	if err != nil {
		return 0
	}
	return number
}

// ToUint parses the string as uint, defaults to 0 if invalid Uint
func (s String) ToUint() uint {
	text := strings.TrimSpace(string(s))
	number, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		return 0
	}
	return uint(number)
}

// ToFloat parses the string as a float64, defaults to 0 if invalid Float
func (s String) ToFloat() float64 {
	text := strings.TrimSpace(string(s))
	number, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0
	}
	return number
}
