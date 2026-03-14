package qb

import "github.com/roidaradal/pack/ds"

type columnValuePair = ds.Tuple2[string, any]
type columnValueListPair = ds.Tuple2[string, []any]

// Internal: create new Column-Value pair
func newColumnValue[T any](this *Instance, fieldRef *T, value T) ds.Option[columnValuePair] {
	column := this.Column(fieldRef)
	if column == "" {
		return ds.Nil[columnValuePair]()
	}
	return ds.NewOption(&columnValuePair{V1: column, V2: value})
}

// Internal: create new Column-ValueList pair
func newColumnValueList[T any](this *Instance, fieldRef *T, values ds.List[T]) ds.Option[columnValueListPair] {
	column := this.Column(fieldRef)
	if column == "" {
		return ds.Nil[columnValueListPair]()
	}
	return ds.NewOption(&columnValueListPair{V1: column, V2: values.ToAny()})
}

// Internal: Create new Column-Value pair, by getting the column name from type and field name
func newFieldColumnValue(this *Instance, typeName, fieldName string, value any) ds.Option[columnValuePair] {
	column := this.getFieldColumnName(typeName, fieldName)
	if column == "" {
		return ds.Nil[columnValuePair]()
	}
	return ds.NewOption(&columnValuePair{V1: column, V2: value})
}
