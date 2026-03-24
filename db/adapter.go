package db

import (
	"database/sql"

	"github.com/roidaradal/tst"
)

// Adapter is an adapter for sql.DB so it follows the Conn interface
type Adapter struct {
	db *sql.DB
}

// MockAdapter is an adapter for tst.Conn so it follows the Conn interface
type MockAdapter struct {
	Conn *tst.Conn
}

// NewAdapter creates a new Adapter
func NewAdapter(db *sql.DB) *Adapter {
	return new(Adapter{db: db})
}

// NewMockAdapter creates a new MockAdapter
func NewMockAdapter(conn *tst.Conn) *MockAdapter {
	return new(MockAdapter{Conn: conn})
}

// QueryRow executes a query and returns a Row object
func (a *Adapter) QueryRow(query string, args ...any) Row {
	return a.db.QueryRow(query, args...)
}

// QueryRow executes a query and returns a Row object
func (a *MockAdapter) QueryRow(query string, args ...any) Row {
	return a.Conn.QueryRow(query, args...)
}

// Query executes a query and returns a Rows object
func (a *Adapter) Query(query string, args ...any) (Rows, error) {
	return a.db.Query(query, args...)
}

// Query executes a query and returns a Rows object
func (a *MockAdapter) Query(query string, args ...any) (Rows, error) {
	return a.Conn.Query(query, args...)
}

// Exec executes a query and returns a Result object
func (a *Adapter) Exec(query string, args ...any) (sql.Result, error) {
	return a.db.Exec(query, args...)
}

// Exec executes a query and returns a Result object
func (a *MockAdapter) Exec(query string, args ...any) (sql.Result, error) {
	return a.Conn.Exec(query, args...)
}

// Begin starts a transaction
func (a *Adapter) Begin() (Tx, error) {
	return a.db.Begin()
}

// Begin starts a transaction
func (a *MockAdapter) Begin() (Tx, error) {
	return a.Conn.Begin()
}
