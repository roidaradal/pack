package qb

import "github.com/roidaradal/pack/ds"

type Instance struct {
	dbType
	addressColumns   ds.Map[string, string]                 // {FieldAddress => ColumnName}
	typeColumns      ds.Map[string, ds.List[string]]        // {TypeName => []ColumnNames}
	typeColumnFields ds.Map[string, ds.Map[string, string]] // {TypeName => {ColumnName => FieldName}}
	typeFieldColumns ds.Map[string, ds.Map[string, string]] // {TypeName => {FieldName => ColumnName}}
}

// NewInstance creates a new QueryBuilder Instance
func NewInstance(db dbType) *Instance {
	return &Instance{
		dbType:           db,
		addressColumns:   make(ds.Map[string, string]),
		typeColumns:      make(ds.Map[string, ds.List[string]]),
		typeColumnFields: make(ds.Map[string, ds.Map[string, string]]),
		typeFieldColumns: make(ds.Map[string, ds.Map[string, string]]),
	}
}
