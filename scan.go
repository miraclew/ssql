package ssql

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func scanStruct(rows *sql.Rows, dest interface{}) error {
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr {
		return errors.New("dest should be a pointer to struct")
	}

	v = reflect.Indirect(v)
	cols, _ := rows.Columns()

	if !rows.Next() {
		return sql.ErrNoRows
	}

	pointers := make([]interface{}, len(cols))
	for i, col := range cols {
		fieldName := covertColumnName(col)
		field := v.FieldByName(fieldName)
		if !field.IsValid() {
			return errors.New(fmt.Sprintf("field %s not valid", col))
		}
		pointers[i] = field.Addr().Interface()
	}

	return rows.Scan(pointers...)
}

func scanStructSlice(rows *sql.Rows, dest interface{}) error {
	t := reflect.TypeOf(dest)
	if t.Kind() != reflect.Slice {
		return errors.New("dest must be a slice")
	}
	v := reflect.ValueOf(dest)

	cols, _ := rows.Columns()
	var fields []string
	for _, c := range cols {
		fields = append(fields, covertColumnName(c))
	}

	for rows.Next() {
		et := t.Elem()
		if et.Kind() == reflect.Ptr {
			et = et.Elem()
		}
		fmt.Println(et.String())
		ev := reflect.New(et)

		pointers := make([]interface{}, len(cols))
		for i, col := range cols {
			fieldName := fields[i]
			fv := ev.Elem().FieldByName(fieldName)
			if !fv.IsValid() {
				return errors.New(fmt.Sprintf("fv %s not valid", col))
			}
			fmt.Println(fv.String())
			pointers[i] = fv.Addr().Interface()
		}

		err := rows.Scan(pointers...)
		if err != nil {
			return err
		}
		fmt.Println(ev.String())
		fmt.Println(ev.Elem().String())
		if t.Elem().Kind() != reflect.Ptr {
			ev = ev.Elem()
		}
		v = reflect.Append(v, ev)
	}

	fmt.Printf("%v", v.Interface())
	return nil
}

func covertColumnName(column string) string {
	  s := strings.Replace(column, "_", " ", -1)
	  s = strings.Replace(s, "id", "ID", -1)
	  s = strings.Title(s)
	  return strings.Replace(s, " ", "", -1)
}