// Package fail contains common errors
package fail

import "errors"

var (
	MissingParams = errors.New("public: Missing required parameters")
	NotFoundItem  = errors.New("public: Item not found")
)
