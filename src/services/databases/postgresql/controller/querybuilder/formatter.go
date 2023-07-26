package querybuilder

import (
	"reflect"
	"strings"
	"fmt"
)

func formatColumnsAndValues(t interface{}) (string, string) {
	id := reflect.ValueOf(t).FieldByName("ID")
	if id.IsValid() {
		var column []string
		var value []string
		tValue := reflect.ValueOf(t)
		tType := reflect.TypeOf(t)
		for i := 0; i < tValue.NumField(); i++ {
			fieldValue := tValue.Field(i)
			fieldType := tType.Field(i)
			// Check if the field is empty
			isEmpty := reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface())
			// Add values that are not empty to array
			if !isEmpty {
				columnName := fieldType.Tag.Get("db")
				if columnName == "" {
					panic("Missing `db` tag")
				}
				column = append(column, columnName)
				if fieldValue.Kind() == reflect.Int {
					value = append(value, fmt.Sprintf("%v", fieldValue))
				} else if fieldValue.Kind() == reflect.Float64 {
					value = append(value, fmt.Sprintf("%v", fieldValue))
				} else if fieldValue.Kind() == reflect.Bool {
					value = append(value, fmt.Sprintf("%v", fieldValue))
				} else {
					value = append(value, fmt.Sprintf("'%v'", fieldValue))
				}
			}
		}
		columns := strings.Join(column,", ")
		values := strings.Join(value,", ")
		return columns, values
	}
	panic("Invalid parameter type")
}

func setColumnsAndValues(t interface{}) string {
	id := reflect.ValueOf(t).FieldByName("ID")
	if id.IsValid() {
		var sets []string
		tValue := reflect.ValueOf(t)
		tType := reflect.TypeOf(t)
		for i := 0; i < tValue.NumField(); i++ {
			fieldValue := tValue.Field(i)
			fieldType := tType.Field(i)
			// skip ID to avoid updating it
			if fieldType.Name != "ID" {
				// Check if the field is empty
				isEmpty := reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface())
				// Add values that are not empty to array
				if isEmpty {
					columnName := fieldType.Tag.Get("db")
					if columnName == "" {
						panic("Missing `db` tag")
					}
					if fieldValue.Kind() == reflect.Int {
						sets = append(sets, fmt.Sprintf("%v = %v", columnName, fieldValue))
					} else if fieldValue.Kind() == reflect.Float64 {
						sets = append(sets, fmt.Sprintf("%v = %v", columnName, fieldValue))
					} else if fieldValue.Kind() == reflect.Bool {
						sets = append(sets, fmt.Sprintf("%v = %v", columnName, fieldValue))
					} else {
						sets = append(sets, fmt.Sprintf("%v = '%v'", columnName, fieldValue))
					}
				}
			}
		}
		set := strings.Join(sets,", ")
		return set
	}
	panic("Invalid parameter type")
}

func createWhereCondition(t interface{}) string {
	id := reflect.ValueOf(t).FieldByName("ID")
	if id.IsValid() {
		var condition []string
		tValue := reflect.ValueOf(t)
		tType := reflect.TypeOf(t)
		for i := 0; i < tValue.NumField(); i++ {
			fieldValue := tValue.Field(i)
			fieldType := tType.Field(i)
			// Check if the field is empty
			isEmpty := reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface())
			// Add values that are not empty to array
			if !isEmpty {
				columnName := fieldType.Tag.Get("db")
				if columnName == "" {
					continue
				}
				if fieldValue.Kind() == reflect.Int {
					condition = append(condition, fmt.Sprintf("%v = %v", columnName, fieldValue))
				} else if fieldValue.Kind() == reflect.Float64 {
					condition = append(condition, fmt.Sprintf("%v = %v", columnName, fieldValue))
				} else if fieldValue.Kind() == reflect.Bool {
					condition = append(condition, fmt.Sprintf("%v = %v", columnName, fieldValue))
				} else {
					condition = append(condition, fmt.Sprintf("%v = '%v'", columnName, fieldValue))
				}
			}
		}
		conditions := strings.Join(condition," AND ")
		return conditions
	}
	panic("Invalid parameter type")
}