package ssql

import "database/sql"

func Open(driverName, dataSourceName string) (*DB, error) {
	rdb, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &DB{rdb}, nil
}

type DB struct {
	rdb *sql.DB
}

func (db *DB) Select(query string, args ...interface{}) *SelectResult {
	rows, err := db.rdb.Query(query, args...)
	return &SelectResult{rows, err}
}

func (db *DB) Update(query string, args ...interface{}) {
	panic("implement me")
}

