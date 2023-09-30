package querybuilder

import (
	"errors"
	"fmt"
	"reflect"

	"services/internal/transaction_history/db/model"
)

func BuildCreateQuery(tableName string, model interface{}) (string, error) {
	query := "INSERT INTO %v (%v) VALUES (%v)"
	columns, values := FormatColumnsAndValues(model)
	if columns == "" {
		err := errors.New("Empty model")
		return "", err
	}
	query = fmt.Sprintf(query, tableName, columns, values)
	return query, nil
}

func BuildUpdateQuery(tableName string, setModel interface{}, whereModel interface{}) (string, error) {
	query := "UPDATE %v SET %v WHERE %v" // TODO have the ability to update multiple transactions
	set := SetColumnsAndValues(setModel)
	conditions := CreateWhereCondition(whereModel)
	if set == "" {
		err := errors.New("Empty model")
		return "", err
	}
	query = fmt.Sprintf(query, tableName, set, conditions)
	return query, nil
}

func BuildRetrieveQuery(tableName string, model interface{}) (string, error) {
	query := "SELECT * FROM %v WHERE %v" // TODO have the ability to make more complex where conditons sunch as nesting and other operators: >,<,IS NULL,etc
	conditions := CreateWhereCondition(model)
	if conditions == "" {
		err := errors.New("Empty model")
		return "", err
	}
	query = fmt.Sprintf(query, tableName, conditions)
	return query, nil
}

func BuildTransactionRetrieveQuery(m model.Transaction) (string, error) {
	query := "SELECT * FROM transactions WHERE %v" // TODO have the ability to make more complex where conditons sunch as nesting and other operators: >,<,IS NULL,etc

	conditions := CreateWhereCondition(m)
	if conditions == "" {
		err := errors.New("Empty model")
		return "", err
	}
	query = fmt.Sprintf(query, conditions)

	// Additional conditionals
	if m.Query.Select.GreaterThanEq.Value != "" {
		if reflect.TypeOf(m.Query.Select.GreaterThanEq.Value).Kind() == reflect.String {
			query += fmt.Sprintf(" AND %v >= '%v'", m.Query.Select.GreaterThanEq.Column, m.Query.Select.GreaterThanEq.Value)
		} else {
			query += fmt.Sprintf(" AND %v >= %v", m.Query.Select.GreaterThanEq.Column, m.Query.Select.GreaterThanEq.Value)
		}
	}

	// Order by queries
	if m.Query.Select.OrderBy.Asc {
		query = query + fmt.Sprintf(" ORDER BY %v ASC", m.Query.Select.OrderBy.Column)
	} else if m.Query.Select.OrderBy.Desc {
		query = query + fmt.Sprintf(" ORDER BY %v DESC", m.Query.Select.OrderBy.Column)
	}
	
	return query, nil
}

func BuildDeleteQuery(tableName string, model interface{}) (string, error) {
	query := "DELETE FROM %v WHERE %v"
	conditions := CreateWhereCondition(model)
	if conditions == "" {
		err := errors.New("Empty model")
		return "", err
	}
	query = fmt.Sprintf(query, tableName, conditions)
	return query, nil
}