package ssql

import "database/sql"

type SelectOneResult struct {
	*sql.Row
}

func (r *SelectOneResult) Int64() (int64, error) {
	var intVal int64
	err := r.Row.Scan(&intVal)
	return intVal, err
}

func (r *SelectOneResult) Int() (int, error) {
	var intVal int
	err := r.Row.Scan(&intVal)
	return intVal, err
}

func (r *SelectOneResult) Float() (float32, error) {
	var floatVal float32
	err := r.Row.Scan(&floatVal)
	return floatVal, err
}

func (r *SelectOneResult) Float64() (float64, error) {
	var floatVal float64
	err := r.Row.Scan(&floatVal)
	return floatVal, err
}

func (r *SelectOneResult) Bool() (bool, error) {
	v, err := r.Int()
	return v > 0, err
}

func (r *SelectOneResult) Val(v interface{}) error {
	return nil
}
