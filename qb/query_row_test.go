package qb

import (
	"testing"

	"github.com/roidaradal/tst"
)

func TestCountQuery(t *testing.T) {
	type Person struct {
		Name string
		Age  int
		Job  string
	}
	p := new(Person)
	this := testPrelude(t, p)
	table := "persons"
	// NewCountQuery
	q0 := NewCountQuery[Person](this, "")    // blank table
	q1 := NewCountQuery[Person](this, table) // with condition
	q2 := NewCountQuery[Person](this, table) // no condition
	// CountQuery.Where
	q1.Where(GreaterEqual[Person](this, &p.Age, 18))
	// CountQuery.Test
	testCases := []tst.P2W1[*CountQuery[Person], Person, bool]{
		{q1, Person{"John", 20, "dev"}, true},
		{q1, Person{"Jane", 18, "student"}, true},
		{q1, Person{"Alice", 15, "student"}, false},
	}
	tst.AllP2W1(t, testCases, "CountQuery.Test", (*CountQuery[Person]).Test, tst.AssertEqual)
	// CountQuery.BuildQuery
	testCases2 := []tst.P1W2[*CountQuery[Person], string, []any]{
		{q0, "", []any{}},
		{q1, "SELECT COUNT(*) FROM `persons` WHERE `Age` >= ?", []any{18}},
		{q2, "SELECT COUNT(*) FROM `persons` WHERE false", []any{}},
	}
	tst.AllP1W2(t, testCases2, "CountQuery.BuildQuery", (*CountQuery[Person]).BuildQuery, tst.AssertEqual, tst.AssertListEqual)
}

func TestValueQuery(t *testing.T) {
	type User struct {
		ID     int
		Name   string
		Code   string
		Job    string
		Extra  string `col:"-"`
		secret string
	}
	u := new(User)
	this := testPrelude(t, u)
	table := "users"
	// NewValueQuery
	q0 := NewValueQuery[User](this, "", &u.Name)      // no table
	q1 := NewValueQuery[User](this, table, &u.Name)   // with condition
	q2 := NewValueQuery[User](this, table, &u.Code)   // with condition
	q3 := NewValueQuery[User](this, table, &u.Job)    //  no condition
	q4 := NewValueQuery[User](this, table, &u.Extra)  // blank column
	q5 := NewValueQuery[User](this, table, &u.secret) // private field
	// ValueQuery.Where
	q1.Where(Equal[User](this, &u.Code, "admin"))
	q2.Where(Equal[User](this, &u.ID, 2))
	// ValueQuery.Test
	u1 := User{1, "Admin", "admin", "dev", "", "123"}
	u2 := User{2, "Guest", "guest", "dev", "", "456"}
	testCases := []tst.P2W1[*ValueQuery[User, string], User, bool]{
		{q1, u1, true}, {q1, u2, false},
		{q2, u1, false}, {q2, u2, true},
		{q3, u1, false}, {q3, u2, false},
	}
	tst.AllP2W1(t, testCases, "ValueQuery.Test", (*ValueQuery[User, string]).Test, tst.AssertEqual)
	// ValueQuery.BuildQuery
	testCases2 := []tst.P1W2[*ValueQuery[User, string], string, []any]{
		{q0, "", []any{}},
		{q1, "SELECT `Name` FROM `users` WHERE `Code` = ?", []any{"admin"}},
		{q2, "SELECT `Code` FROM `users` WHERE `ID` = ?", []any{2}},
		{q3, "SELECT `Job` FROM `users` WHERE false", []any{}},
		{q4, "", []any{}},
		{q5, "", []any{}},
	}
	tst.AllP1W2(t, testCases2, "ValueQuery.BuildQuery", (*ValueQuery[User, string]).BuildQuery, tst.AssertEqual, tst.AssertListEqual)
}

func TestSelectRowQuery(t *testing.T) {
	// TODO: NewSelectRowQuery
	// TODO: NewFullSelectRowQuery
	// TODO: SelectRowQuery.Columns
	// TODO: SelectRowQuery.Where
	// TODO: SelectRowQuery.Test
	// TODO: SelectRowQuery.BuildQuery

}

func TestTopRowQuery(t *testing.T) {
	// TODO: NewTopRowQuery
	// TODO: TopRowQuery.Columns
	// TODO: TopRowQuery.OrderAsc, OrderDesc
	// TODO: TopRowQuery.Limit
	// TODO: TopRowQuery.Where
	// TODO: TopRowQuery.Test
	// TODO: TopRowQuery.BuildQuery
}

func TestTopValueQuery(t *testing.T) {
	// TODO: NewTopValueQuery
	// TODO: TopValueQuery.OrderAsc, OrderDesc
	// TODO: TopValueQuery.Limit
	// TODO: TopValueQuery.Where
	// TODO: TopValueQuery.Test
	// TODO: TopValueQuery.BuildQuery
}
