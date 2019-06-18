package ssql

type Database interface {
	Select(query string, args ...interface{}) SelectQuery
	SelectOne(query string, args ...interface{}) SelectOneQuery
	Update(query string, args ...interface{})
	Count(query string, args ...interface{}) (int64, error)
	Exists(query string, args ...interface{}) (bool, error)
}

type SelectQuery interface {
}

type SelectOneQuery interface {
}

type UpdateQuery interface {
}

type DeleteQuery interface {
}