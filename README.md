# ssql
ssql provide some simple methods to make it easier to work with SQL in GO.

## Example 

### Open connection

```go
db, err := ssql.Open("mysql", "user:pass@tcp(localhost:3306)/dbname?parseTime=true")
``` 

### Select one row
```go
var item Item
err := db.Select("select * from items where id=?", 50).Value(&item)
``` 

### Select rows
```go
var items []Item
err := db.Select("select * from items where id=?", 50).Values(&items)
``` 

### Select scalar value
```go
v, err := db.Select("select count(*) from items").Int()
``` 
