package qb

import (
	"fmt"
	"strings"
)

// orderedLimit is an abstractQuery with order column(s) and a limit.
// It does not implement the BuildQuery method; it is embedded by concrete Queries for method reuse
type orderedLimit struct {
	orders []string
	limit  uint
}

// OrderAsc adds a column with ascending order, returns orderedLimit for chaining
func (q *orderedLimit) OrderAsc(this *Instance, column string) *orderedLimit {
	if column == "" {
		return q
	}
	order := fmt.Sprintf("%s ASC", this.dbType.prepareIdentifier(column))
	q.orders = append(q.orders, order)
	return q
}

// OrderDesc adds a column with descending order, returns orderedLimit for chaining
func (q *orderedLimit) OrderDesc(this *Instance, column string) *orderedLimit {
	if column == "" {
		return q
	}
	order := fmt.Sprintf("%s DESC", this.dbType.prepareIdentifier(column))
	q.orders = append(q.orders, order)
	return q
}

// Limit sets the query limit, returns orderedLimit for chaining.
// Setting to 0 removes the limit
func (q *orderedLimit) Limit(limit uint) *orderedLimit {
	q.limit = limit
	return q
}

// String builds the orderString and limitString
func (q *orderedLimit) String() string {
	output := make([]string, 0, 2)
	orderString := q.orderString()
	if orderString != "" {
		output = append(output, fmt.Sprintf("ORDER BY %s", orderString))
	}
	if q.limit > 0 {
		output = append(output, fmt.Sprintf("LIMIT %d", q.limit))
	}
	return strings.Join(output, " ")
}

// orderString builds the list of orders into a string
func (q *orderedLimit) orderString() string {
	return strings.Join(q.orders, ", ")
}

// mustLimitString builds the orderString and limitString, but only includes the orderString if limitString is not empty
func (q *orderedLimit) mustLimitString() string {
	if q.limit == 0 {
		return ""
	}
	return q.String()
}
