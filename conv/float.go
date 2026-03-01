package conv

import (
	"fmt"

	"github.com/roidaradal/pack/types"
)

// FloatToString converts a Float to string
func FloatToString[F types.Float](f F) string {
	return fmt.Sprintf("%f", f)
}

// FloatToStringDecimal converts a Float to string, with N decimal places
func FloatToStringDecimal[F types.Float](f F, decimalPlaces int) string {
	format := fmt.Sprintf("%%.%df", decimalPlaces)
	return fmt.Sprintf(format, f)
}

// FloatToInt converts Float to int
func FloatToInt[F types.Float](f F) int {
	return int(f)
}

// FloatToUint converts Float to uint, clips to 0 if negative float
func FloatToUint[F types.Float](f F) uint {
	if f < 0 {
		return 0
	}
	return uint(f)
}
