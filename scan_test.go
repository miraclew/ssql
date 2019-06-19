package ssql

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func query() (*sql.Rows, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/devnews?parseTime=true")
	if err != nil {
		return nil, err
	}

	return db.Query(`select title, item_id from items`)
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

func TestAA(t *testing.T) {
	var items []*item
	typ := reflect.TypeOf(items)
	t.Log(typ.Kind().String(), typ.String())
	t.Log(typ.Elem().Kind().String(), typ.Elem().String())
	t.Log(typ.Elem().Elem().Kind().String())
	t.Log(typ.Elem().Elem().FieldByName("Title"))

	et := typ.Elem()
	if et.Kind() == reflect.Ptr {
		et = et.Elem()
	}
	t.Log(et.Kind().String())

	n := reflect.New(et)
	t.Log(n.Elem().String())
	n.Elem().FieldByName("Title").SetString("Hello")
	t.Logf("%v", n.Interface())

	//v := reflect.ValueOf(items)
	//t.Log(v.Kind().String())



}