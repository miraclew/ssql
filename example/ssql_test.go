package example

import (
	"github.com/miraclew/ssql"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func getDB() *ssql.DB {
	db, err := ssql.Open("mysql", "root@tcp(localhost:3306)/devnews?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func TestInt(t *testing.T) {
	db := getDB()
	v, err := db.Select("select count(*) from items where title=?", "华为鸿蒙之外还有后手？极光系统是什么").Int()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

func TestString(t *testing.T) {
	db := getDB()
	v, err := db.Select("select title from items where id=?", 50).String()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

func TestScan(t *testing.T) {
	db := getDB()
	var it item
	err := db.Select("select * from items where id=?", 50).Value(&it)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(it)
}


func TestSliceScan(t *testing.T) {
	db := getDB()
	var items []item
	err := db.Select("select * from items").Values(&items)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(items)
}

type item struct {
	Title string
	ItemID string
}