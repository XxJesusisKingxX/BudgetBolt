package querybuilder

import (
	"errors"
	"fmt"
	"reflect"
)



func BuildCreateQuery(tableName string, model interface{}) (string, error) {
	var err error
	id := reflect.ValueOf(model).FieldByName("ID")
	if id.IsValid() {
		query := "INSERT INTO %v (%v) VALUES (%v)"
		columns, values := formatColumnsAndValues(model)
		if columns == "" {
			err = errors.New("Empty model")
			return "", err
		}
		query = fmt.Sprintf(query, tableName, columns, values)
		return query, err
	}
	panic("Invalid parameter type")
}

func BuildUpdateQuery(tableName string, model interface{}) (string, error) {
	var err error
	id := reflect.ValueOf(model).FieldByName("ID")
	if id.IsValid() {
		query := "UPDATE %v SET %v WHERE transaction_id=%v" // TODO have the ability to update multiple transactions
		set := setColumnsAndValues(model)
		if set == "" {
			err = errors.New("Empty model")
			return "", err
		}
		query = fmt.Sprintf(query, tableName, set, id)
		return query, err
	}
	panic("Invalid parameter type")
}

func BuildRetrieveQuery(tableName string, model interface{}) (string, error) {
	var err error
	id := reflect.ValueOf(model).FieldByName("ID")
	if id.IsValid() {
		query := "SELECT * FROM %v WHERE %v" // TODO have the ability to make more complex where conditons sunch as nesting and other operators: >,<,IS NULL,etc
		conditions := createWhereCondition(model)
		if conditions == "" {
			err = errors.New("Empty model")
			return "", err
		}
		query = fmt.Sprintf(query, tableName, conditions)
		return query, err
	}
	panic("Invalid parameter type")
}

func BuildDeleteQuery(tableName string, model interface{}) (string, error) {
	var err error
	id := reflect.ValueOf(model).FieldByName("ID")
	if id.IsValid() {
		query := "DELETE FROM %v WHERE %v"
		conditions := createWhereCondition(model)
		if conditions == "" {
			err = errors.New("Empty model")
			return "", err
		}
		query = fmt.Sprintf(query, tableName, conditions)
		return query, err
	}
	panic("Invalid parameter type")
}