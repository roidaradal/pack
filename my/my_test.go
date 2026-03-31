package my

import (
	"testing"

	"github.com/zeroibot/tst"
)

func TestStatusCodes(t *testing.T) {
	pairs := [][2]int{
		{OK200, 200}, {OK201, 201}, {Err400, 400}, {Err401, 401},
		{Err403, 403}, {Err404, 404}, {Err429, 429}, {Err500, 500},
	}
	tst.All(t, pairs, "ErrorCodes", tst.AssertEqual)
}

func TestNewInstance(t *testing.T) {
	// TODO: NewInstance
}

func TestAddConnection(t *testing.T) {
	// TODO: Instance.AddConnection
}
