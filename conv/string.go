package conv

import (
	"strconv"
	"strings"
)

// StringToInt parses the string as int, defaults to 0 if invalid int
func StringToInt(s string) int {
	number, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return 0
	}
	return number
}

// StringToUint parses the string as uint, defaults to 0 if invalid uint
func StringToUint(s string) uint {
	number, err := strconv.ParseUint(strings.TrimSpace(s), 10, 64)
	if err != nil {
		return 0
	}
	return uint(number)
}

// StringToFloat parses the string as float64, defaults to 0 if invalid float
func StringToFloat(s string) float64 {
	number, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		return 0
	}
	return number
}
