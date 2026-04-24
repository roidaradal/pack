package model

import (
	"github.com/zeroibot/pack/my"
	"github.com/zeroibot/pack/qb"
)

type AbstractTxManager interface {
	// Step runs a given function in a transaction, rolling back if any error occurs
	Step(rqtx *my.Request, fn func() error) error
}

type TxManager struct{}

// Step runs a given function in a transaction, rolling back if any error occurs
func (m *TxManager) Step(rqtx *my.Request, fn func() error) error {
	err := fn()
	if err != nil {
		return qb.Rollback(rqtx.Tx, err)
	}
	return nil
}
