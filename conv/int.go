package conv

import (
	"fmt"

	"github.com/roidaradal/pack/types"
)

// Abs gets the absolute value of an Integer
func Abs[I types.Integer](i I) I {
	if i < 0 {
		return -i
	}
	return i
}

// IntToString converts an Integer to string
func IntToString[I types.Integer](i I) string {
	return fmt.Sprintf("%d", i)
}

// IntToUint converts Integer to uint, clips to 0 if negative int
func IntToUint[I types.Integer](i I) uint {
	if i < 0 {
		return 0
	}
	return uint(i)
}

// IntToFloat converts Integer to float64
func IntToFloat[I types.Integer](i I) float64 {
	return float64(i)
}

// IntToBool converts Integer to bool (0 is false, else true)
func IntToBool[I types.Integer](i I) bool {
	return i != 0
}

// UintToInt converts uint to int
func UintToInt(i uint) int {
	return int(i)
}
