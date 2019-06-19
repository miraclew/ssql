package ssql

import "database/sql"

type SelectResult struct {
	*sql.Rows
	error
}

func (r *SelectResult) Value(v interface{}) error {
	return scanStruct(r.Rows, v)
}

func (r *SelectResult) Values(v []interface{}) error {
	return scanStructSlice(r.Rows, v)
}




