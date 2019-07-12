package ssql

import "database/sql"

type SelectResult struct {
	*sql.Rows
	error
}

func (r *SelectResult) Value(v interface{}) error {
	if r.error != nil {
		return r.error
	}
	err := scanStruct(r.Rows, v)
	if err != nil {
		return err
	}
	return r.Close()
}

func (r *SelectResult) Values(v interface{}) error {
	if r.error != nil {
		return r.error
	}
	err := scanStructSlice(r.Rows, v)
	if err != nil {
		return err
	}
	return r.Close()
}

func (r *SelectResult) Int64() (int64, error) {
	if r.error != nil {
		return 0, r.error
	}
	if !r.Next() {
		return 0, sql.ErrNoRows
	}

	var intVal int64
	err := r.Rows.Scan(&intVal)
	if err != nil {
		return 0, err
	}
	return intVal, r.Rows.Close()
}

func (r *SelectResult) Int() (int, error) {
	if r.error != nil {
		return 0, r.error
	}
	if !r.Next() {
		return 0, sql.ErrNoRows
	}

	var intVal int
	err := r.Rows.Scan(&intVal)
	if err != nil {
		return 0, err
	}
	return intVal, r.Rows.Close()
}

func (r *SelectResult) Bool() (bool, error) {
	v, err := r.Int()
	return v > 0, err
}

func (r *SelectResult) String() (string, error) {
	if r.error != nil {
		return "", r.error
	}
	if !r.Next() {
		return "", sql.ErrNoRows
	}

	var v string
	err := r.Rows.Scan(&v)
	if err != nil {
		return "", err
	}
	return v, r.Rows.Close()
}
