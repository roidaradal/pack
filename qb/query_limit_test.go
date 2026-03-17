package qb

import (
	"testing"

	"github.com/roidaradal/pack/dyn"
)

func TestOrderedLimit(t *testing.T) {
	this := NewInstance(MySQL)
	emptyQuery := new(orderedLimit)
	q0 := new(orderedLimit)
	q1 := new(orderedLimit)
	q2 := new(orderedLimit)
	q3 := new(orderedLimit)
	q4 := new(orderedLimit)
	q5 := new(orderedLimit)
	q6 := new(orderedLimit)
	q7 := new(orderedLimit)

	// OrderAsc, OrderDesc
	q1.OrderAsc(this, "")  // no column = no effect
	q2.OrderDesc(this, "") // no column = no effect
	if dyn.NotEqual(q1, emptyQuery) {
		t.Errorf("OrderAsc() = %v, want %v", q1, emptyQuery)
	}
	if dyn.NotEqual(q2, emptyQuery) {
		t.Errorf("OrderDesc() = %v, want %v", q2, emptyQuery)
	}

	q1.OrderAsc(this, "Name")
	q2.OrderDesc(this, "CreatedAt")
	q3.Limit(5)
	q4.OrderAsc(this, "Code").Limit(10)
	q5.OrderDesc(this, "UpdatedAt").Limit(5)
	q6.OrderAsc(this, "Code").OrderDesc(this, "UpdatedAt").Limit(5)
	q7.OrderDesc(this, "UpdatedAt").OrderAsc(this, "Code").Limit(10)

	// orderString
	type testCase struct {
		q         *orderedLimit
		wantOrder string
		wantLimit uint
	}
	testCases := []testCase{
		{q0, "", 0},
		{q1, "`Name` ASC", 0},
		{q2, "`CreatedAt` DESC", 0},
		{q3, "", 5},
		{q4, "`Code` ASC", 10},
		{q5, "`UpdatedAt` DESC", 5},
		{q6, "`Code` ASC, `UpdatedAt` DESC", 5},
		{q7, "`UpdatedAt` DESC, `Code` ASC", 10},
	}
	for _, x := range testCases {
		actualOrder := x.q.orderString()
		actualLimit := x.q.limit
		if actualOrder != x.wantOrder || actualLimit != x.wantLimit {
			t.Errorf("orderedLimit() = %q, %d, want %q, %d", actualOrder, actualLimit, x.wantOrder, x.wantLimit)
		}
	}

	// String
	type testCase2 struct {
		q          *orderedLimit
		wantString string
	}
	c0 := testCase2{q0, ""}
	c3 := testCase2{q3, "LIMIT 5"}
	c4 := testCase2{q4, "ORDER BY `Code` ASC LIMIT 10"}
	c5 := testCase2{q5, "ORDER BY `UpdatedAt` DESC LIMIT 5"}
	c6 := testCase2{q6, "ORDER BY `Code` ASC, `UpdatedAt` DESC LIMIT 5"}
	c7 := testCase2{q7, "ORDER BY `UpdatedAt` DESC, `Code` ASC LIMIT 10"}
	testCases2 := []testCase2{
		c0, c3, c4, c5, c6, c7,
		{q1, "ORDER BY `Name` ASC"},
		{q2, "ORDER BY `CreatedAt` DESC"},
	}
	for _, x := range testCases2 {
		actualString := x.q.String()
		if actualString != x.wantString {
			t.Errorf("orderedLimit.String() = %q, want %q", actualString, x.wantString)
		}
	}
	// mustLimitString
	testCases2 = []testCase2{
		c0, c3, c4, c5, c6, c7,
		{q1, ""},
		{q2, ""},
	}
	for _, x := range testCases2 {
		actualString := x.q.mustLimitString()
		if actualString != x.wantString {
			t.Errorf("orderedLimit.String() = %q, want %q", actualString, x.wantString)
		}
	}
}
