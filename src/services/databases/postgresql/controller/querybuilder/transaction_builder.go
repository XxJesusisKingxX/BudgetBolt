package querybuilder

import (
	"budgetbolt/src/services/databases/postgresql/model"
	"errors"
	"fmt"
)

func BuildTransactionCreateQuery(m model.Transaction) (string, error) {
	query := "INSERT INTO transaction (%v) VALUES (%v)"
	columns, values := formatColumnsAndValues(m)
	if columns == "" {
		err := errors.New("Empty model")
		return "", err
	}
	query = fmt.Sprintf(query, columns, values)
	return query, nil
}

func BuildTransactionUpdateQuery(m model.Transaction) (string, error) {
	query := "UPDATE transaction SET %v WHERE transaction_id=%v" // TODO have the ability to update multiple transactions
	set := setColumnsAndValues(m)
	if set == "" {
		err := errors.New("Empty model")
		return "", err
	}
	query = fmt.Sprintf(query, set, m.ID)
	return query, nil

}

func BuildTransactionRetrieveQuery(m model.Transaction) (string, error) {
	query := "SELECT * FROM transaction WHERE %v " // TODO have the ability to make more complex where conditons sunch as nesting and other operators: >,<,IS NULL,etc
	if m.Query.Select.Asc {
		query = query + fmt.Sprintf("ORDER BY %v ASC", m.Query.Select.OrderBy)
	} else if m.Query.Select.Desc {
		query = query + fmt.Sprintf("ORDER BY %v DESC", m.Query.Select.OrderBy)
	}
	conditions := createWhereCondition(m)
	if conditions == "" {
		err := errors.New("Empty model")
		return "", err
	}
	query = fmt.Sprintf(query, conditions)
	return query, nil
}

func BuildTransactionDeleteQuery(m model.Transaction) (string, error) {
	query := "DELETE FROM transaction WHERE %v"
	conditions := createWhereCondition(m)
	if conditions == "" {
		err := errors.New("Empty model")
		return "", err
	}
	query = fmt.Sprintf(query, conditions)
	return query, nil
}