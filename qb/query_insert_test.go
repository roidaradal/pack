package qb

import (
	"testing"

	"github.com/roidaradal/pack/dict"
	"github.com/roidaradal/tst"
)

func TestInsertRowQuery(t *testing.T) {
	type User struct {
		Username string
		Password string
		Count    int
		IP       *string
		secret   string
		Code     string `col:"UUID"`
	}
	table := "users"
	u := new(User)
	this := testPrelude(t, u)

	// NewInsertRowQuery
	q0 := NewInsertRowQuery(this, table) // empty row
	q1 := NewInsertRowQuery(this, "")    // no table

	// InsertRowQuery.Row
	q2 := NewInsertRowQuery(this, table)
	q2.Row(this, dict.Object{"Username": "admin", "Password": "123", "": "blank"}) // one blank column
	q3 := NewInsertRowQuery(this, table)
	q3.Row(this, dict.Object{"Username": "john", "Password": "12345", "Count": 5})
	q6 := NewInsertRowQuery(this, table)
	q6.Row(this, dict.Object{"IP": nil})
	q7 := NewInsertRowQuery(this, table)
	q7.Row(this, dict.Object{"Username": "homer", "IP": nil})

	// Using ToRow
	ip := new("127.0.0.1")
	u1 := new(User{Username: "Jane", Password: "6767", Count: 10, Code: "eagle", IP: ip})
	u2 := new(User{Username: "Jack", Password: "6969", Count: 5, Code: "tiger", IP: ip})
	q4 := NewInsertRowQuery(this, table)
	q4.Row(this, ToRow(this, u1))
	q5 := NewInsertRowQuery(this, table)
	q5.Row(this, ToRow(this, u2))
	q8 := NewInsertRowQuery(this, table)
	q8.Row(this, ToRow(this, new(User)))

	// InsertRowQuery.BuildQuery
	emptyValues := make([]any, 0)
	testCases := []tst.P1W2[*InsertRowQuery, string, []any]{
		{q0, "", emptyValues},
		{q1, "", emptyValues},
		{q2, "INSERT INTO `users` (`Password`, `Username`) VALUES (?, ?)", []any{"123", "admin"}},
		{q3, "INSERT INTO `users` (`Count`, `Password`, `Username`) VALUES (?, ?, ?)", []any{5, "12345", "john"}},
		{q6, "INSERT INTO `users` (`IP`) VALUES (?)", []any{nil}},
		{q7, "INSERT INTO `users` (`IP`, `Username`) VALUES (?, ?)", []any{nil, "homer"}},
		{q4, "INSERT INTO `users` (`Count`, `IP`, `Password`, `UUID`, `Username`) VALUES (?, ?, ?, ?, ?)", []any{10, ip, "6767", "eagle", "Jane"}},
		{q5, "INSERT INTO `users` (`Count`, `IP`, `Password`, `UUID`, `Username`) VALUES (?, ?, ?, ?, ?)", []any{5, ip, "6969", "tiger", "Jack"}},
		{q8, "INSERT INTO `users` (`Count`, `IP`, `Password`, `UUID`, `Username`) VALUES (?, ?, ?, ?, ?)", []any{0, nil, "", "", ""}},
	}
	// Note: used ListMixedEqual here because of IP (*string) which has nil checking
	tst.AllP1W2(t, testCases, "InsertRowQuery.BuildQuery", (*InsertRowQuery).BuildQuery, tst.AssertEqual, tst.AssertListMixedEqual)
}

func TestInsertRowsQuery(t *testing.T) {
	type User struct {
		Username string
		Password string
		Count    int
		IP       *string
		secret   string
		Code     string `col:"UUID"`
	}
	table := "users"
	u := new(User)
	this := testPrelude(t, u)

	// NewInsertRowsQuery
	q0 := NewInsertRowsQuery(this, table) // no rows
	q1 := NewInsertRowsQuery(this, "")    // no table

	// InsertRowsQuery.Rows
	q2 := NewInsertRowsQuery(this, table)
	q2.Rows(this, dict.Object{"Username": "admin", "Password": "123"})
	q3 := NewInsertRowsQuery(this, table)
	q3.Rows(this, dict.Object{"Username": "admin", "Password": "123"}, dict.Object{"Username": "root", "Password": "456"})
	q4 := NewInsertRowsQuery(this, table) // blank column
	q4.Rows(this, dict.Object{"Username": "admin", "": "blank"})
	q5 := NewInsertRowsQuery(this, table) // empty row
	q5.Rows(this, dict.Object{})
	q6 := NewInsertRowsQuery(this, table) // inconsistent signatures
	q6.Rows(this, dict.Object{"Username": "admin"}, dict.Object{"Password": "123"})

	// Rows with ToRow
	ip1, ip2 := new("127.0.0.1"), new("localhost")
	john := new(User{Username: "John", Password: "1234", Count: 5})
	jack := new(User{Username: "Jack", Password: "6969", Count: 10})
	jane := new(User{Username: "Jane", Password: "6767", Count: 3, Code: "eagle", IP: ip1})
	juno := new(User{Username: "Juno", Password: "3435", Count: 7, Code: "tiger", IP: ip2})
	q7 := NewInsertRowsQuery(this, table)
	q7.Rows(this, ToRow(this, john), ToRow(this, jack))
	q8 := NewInsertRowsQuery(this, table)
	q8.Rows(this, ToRow(this, new(User)))
	q9 := NewInsertRowsQuery(this, table)
	users := []dict.Object{ToRow(this, jack), ToRow(this, jane), ToRow(this, juno)}
	q9.Rows(this, users...)

	// InsertRowsQuery.BuildQuery
	emptyValues := make([]any, 0)
	testCases := []tst.P1W2[*InsertRowsQuery, string, []any]{
		{q0, "", emptyValues},
		{q1, "", emptyValues},
		{q2, "INSERT INTO `users` (`Password`, `Username`) VALUES (?, ?)", []any{"123", "admin"}},
		{q3, "INSERT INTO `users` (`Password`, `Username`) VALUES (?, ?), (?, ?)", []any{"123", "admin", "456", "root"}},
		{q4, "INSERT INTO `users` (`Username`) VALUES (?)", []any{"admin"}},
		{q5, "", emptyValues},
		{q6, "", emptyValues},
		{q7, "INSERT INTO `users` (`Count`, `IP`, `Password`, `UUID`, `Username`) VALUES (?, ?, ?, ?, ?), (?, ?, ?, ?, ?)", []any{5, nil, "1234", "", "John", 10, nil, "6969", "", "Jack"}},
		{q8, "INSERT INTO `users` (`Count`, `IP`, `Password`, `UUID`, `Username`) VALUES (?, ?, ?, ?, ?)", []any{0, nil, "", "", ""}},
		{q9, "INSERT INTO `users` (`Count`, `IP`, `Password`, `UUID`, `Username`) VALUES (?, ?, ?, ?, ?), (?, ?, ?, ?, ?), (?, ?, ?, ?, ?)", []any{10, nil, "6969", "", "Jack", 3, ip1, "6767", "eagle", "Jane", 7, ip2, "3435", "tiger", "Juno"}},
	}
	tst.AllP1W2(t, testCases, "InsertRowsQuery.BuildQuery", (*InsertRowsQuery).BuildQuery, tst.AssertEqual, tst.AssertListMixedEqual)
}
