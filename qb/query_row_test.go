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
	// TODO: NewValueQuery
	// TODO: ValueQuery.Where
	// TODO: ValueQuery.Test
	// TODO: ValueQuery.BuildQuery
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
