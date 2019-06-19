package ssql

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func getDB() *DB {
	db, err := Open("mysql", "root@tcp(localhost:3306)/devnews?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func query() (*sql.Rows, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/devnews?parseTime=true")
	if err != nil {
		return nil, err
	}

	return db.Query(`select title, item_id from items`)
}

func TestCount(t *testing.T) {
	db := getDB()
	v, err := db.Select("select count(*) from items where title=?", "Hello").Int()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)

	var items []item
	err = db.Select("select * from items").Values(&items)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(items)
}

type item struct {
	Title string
	ItemID string
}

func TestScanStruct(t *testing.T) {
	rows, err := query()

	var it item
	err = scanStruct(rows, &it)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(it.ItemID, it.Title)
}

func TestScanStructSlice(t *testing.T) {
	rows, err := query()

	var items []item
	err = scanStructSlice(rows, &items)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%v", items)
	//t.Log(it.ItemID, it.Title)
}
