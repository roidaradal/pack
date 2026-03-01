// Package conv contains conversion function
package conv

import "fmt"

// BoolToInt converts bool to int (true = 1, false = 0)
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// BoolToUint converts bool to uint (true = 1, false = 0)
func BoolToUint(b bool) uint {
	if b {
		return 1
	}
	return 0
}

// BoolToString converts bool to string
func BoolToString(b bool) string {
	return fmt.Sprintf("%t", b)
}
