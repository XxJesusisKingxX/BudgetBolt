package querybuilder

import (
	"errors"
	"fmt"
	"reflect"
	"services/db/postgresql/model"
)

func BuildTransactionCreateQuery(m model.Transaction) (string, error) {
	query := "INSERT INTO transaction (%v) VALUES (%v)"
	columns, values := FormatColumnsAndValues(m)
	if columns == "" {
		err := errors.New("Empty model")
		return "", err
	}
	query = fmt.Sprintf(query, columns, values)
	return query, nil
}

func BuildTransactionUpdateQuery(m model.Transaction) (string, error) {
	query := "UPDATE transaction SET %v WHERE transaction_id='%v'" // TODO have the ability to update multiple transactions
	set := SetColumnsAndValues(m)
	if set == "" {
		err := errors.New("Empty model")
		return "", err
	}
	query = fmt.Sprintf(query, set, m.ID)
	return query, nil

}

func BuildTransactionRetrieveQuery(m model.Transaction) (string, error) {
	query := "SELECT * FROM transaction WHERE %v" // TODO have the ability to make more complex where conditons sunch as nesting and other operators: >,<,IS NULL,etc

	// Order by queries
	if m.Query.Select.OrderBy.Asc {
		query = query + fmt.Sprintf(" ORDER BY %v ASC", m.Query.Select.OrderBy.Column)
	} else if m.Query.Select.OrderBy.Desc {
		query = query + fmt.Sprintf(" ORDER BY %v DESC", m.Query.Select.OrderBy.Column)
	}

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
	return query, nil
}

func BuildTransactionDeleteQuery(m model.Transaction) (string, error) {
	query := "DELETE FROM transaction WHERE %v"
	conditions := CreateWhereCondition(m)
	if conditions == "" {
		err := errors.New("Empty model")
		return "", err
	}
	query = fmt.Sprintf(query, conditions)
	return query, nil
}