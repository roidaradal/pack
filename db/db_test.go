package db

import (
	"database/sql"
	"testing"

	"github.com/roidaradal/tst"
)

func TestDBInterfaces(t *testing.T) {
	// Conn interface
	dbc1 := NewAdapter(new(sql.DB))
	dbc2 := NewMockAdapter(tst.NewConn[any]())
	var conn1 Conn = dbc1
	var conn2 Conn = dbc2
	tst.All(t, [][2]any{{conn1, dbc1}, {conn2, dbc2}}, "Conn", tst.AssertDeepEqual)

	// Tx interface
	tx1 := new(sql.Tx)
	tx2 := tst.NewTx()
	var tx3 Tx = tx1
	var tx4 Tx = tx2
	tst.All(t, [][2]any{{tx3, tx1}, {tx4, tx2}}, "Tx", tst.AssertDeepEqual)

	// Row interface
	row1 := new(sql.Row)
	row2 := tst.NewRow()
	var row3 Row = row1
	var row4 Row = row2
	tst.All(t, [][2]any{{row3, row1}, {row4, row2}}, "Row", tst.AssertDeepEqual)

	// Rows interface
	rows1 := new(sql.Rows)
	rows2 := tst.NewRows()
	var rows3 Rows = rows1
	var rows4 Rows = rows2
	tst.All(t, [][2]any{{rows3, rows1}, {rows4, rows2}}, "Rows", tst.AssertDeepEqual)

	// RowScanner interface
	var scan1 RowScanner = row1
	var scan2 RowScanner = row2
	var scan3 RowScanner = rows1
	var scan4 RowScanner = rows2
	pairs := [][2]any{{scan1, row1}, {scan2, row2}, {scan3, rows1}, {scan4, rows2}}
	tst.All(t, pairs, "RowScanner", tst.AssertDeepEqual)
}
