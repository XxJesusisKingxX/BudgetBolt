package controller

import (
	"fmt"
	"reflect"
	"strings"

	table "budgetbolt/services/databases/postgresql/model"
)

func formatColumnsAndValues(t interface{}) (string, string) {
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
		if isEmpty != true {
			column = append(column, fieldType.Tag.Get("db"))
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

func setColumnsAndValues(t interface{}) string {
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
			if isEmpty != true {
				columnName := fieldType.Tag.Get("db")
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

func createWhereCondition(t interface{}) string {
	var condition []string
	tValue := reflect.ValueOf(t)
	tType := reflect.TypeOf(t)
	for i := 0; i < tValue.NumField(); i++ {
		fieldValue := tValue.Field(i)
		fieldType := tType.Field(i)
		// Check if the field is empty
		isEmpty := reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface())
		// Add values that are not empty to array
		if isEmpty != true {
			columnName := fieldType.Tag.Get("db")
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

func buildTransactionInsertQuery(model table.Transaction) string {
	query := "INSERT INTO transaction (%v) VALUES (%v)"
	columns, values := formatColumnsAndValues(model)
	preparedQuery := fmt.Sprintf(query, columns, values)
	return preparedQuery
}

func buildTransactionUpdateQuery(model table.Transaction) string {
	query := "UPDATE transaction SET %v WHERE transaction_id=%v" // TODO have the ability to update multiple transactions
	set := setColumnsAndValues(model)
	preparedQuery := fmt.Sprintf(query, set, model.ID)
	return preparedQuery
}

func buildTransactionGetQuery(model table.Transaction) string {
	query := "SELECT * FROM transaction WHERE %v" // TODO have the ability to make more complex where conditons sunch as nesting and other operators: >,<,IS NULL,etc
	conditions := createWhereCondition(model)
	preparedQuery := fmt.Sprintf(query, conditions)
	return preparedQuery
}