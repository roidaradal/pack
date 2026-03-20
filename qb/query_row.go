package qb

import "fmt"

// CountQuery counts the number of rows that satisfy the condition
type CountQuery[T any] struct {
	conditionQuery[T]
}

// NewCountQuery creates a new CountQuery
func NewCountQuery[T any](this *Instance, table string) *CountQuery[T] {
	q := new(CountQuery[T])
	q.initializeRequired(this, table)
	return q
}

// BuildQuery returns the query string and parameter values of CountQuery
func (q *CountQuery[T]) BuildQuery() (string, []any) {
	condition, values, err := q.conditionQuery.preBuildCheck()
	if err != nil {
		return emptyQueryValues()
	}
	query := "SELECT COUNT(*) FROM %s WHERE %s"
	query = fmt.Sprintf(query, q.table, condition)
	return query, values
}
