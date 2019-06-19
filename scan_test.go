package ssql

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func query() (*sql.Rows, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/devnews?parseTime=true")
	if err != nil {
		return nil, err
	}

	return db.Query(`select title, item_id from items where id > 50`)
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
	err = scanStructSlice(rows, items)
	if err != nil {
		t.Fatal(err)
	}

	//t.Log(it.ItemID, it.Title)
}
