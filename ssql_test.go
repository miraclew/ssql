package ssql

//func TestName(t *testing.T) {
//	db, err := sql.Open("sqlite3", "/tmp/post_db.bin")
//	checkErr(err, "sql.Open failed")
//
//	// construct a gorp DbMap
//	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
//	dbmap.Select()
//}
//
//func checkErr(err error, msg string) {
//	if err != nil {
//		log.Fatalln(msg, err)
//	}
//}