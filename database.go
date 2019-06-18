package ssql

type Database interface {
	Select(query string, args ...interface{}) SelectQuery
	SelectOne(query string, args ...interface{}) SelectOneQuery
	Update(query string, args ...interface{})
	Count(query string, args ...interface{}) (int64, error)
	Exists(query string, args ...interface{}) (bool, error)
}

type SelectQuery interface {
	Int64() (int64, error)
	Int() (int, error)
	Scan(v []interface{}, err error)
}

type SelectOneQuery interface {
	Int64() (int64, error)
	Int() (int, error)
	Scan(v interface{}, err error)
}

type UpdateQuery interface {
}

type DeleteQuery interface {
}