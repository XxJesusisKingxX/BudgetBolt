package controller

import (
	"database/sql"

	q "budgetbolt/src/services/databases/postgresql/controller/querybuilder"
	"budgetbolt/src/services/databases/postgresql/model"
	"budgetbolt/src/services/databases/postgresql/view"
)

func CreateTransaction(db *sql.DB, m model.Transaction) error {
	query, err := q.BuildTransactionCreateQuery(m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateTransaction(db *sql.DB, m model.Transaction) error {
	query, err := q.BuildTransactionUpdateQuery(m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveTransaction(db *sql.DB, m model.Transaction) ([]model.Transaction, error) {
	query, err := q.BuildTransactionRetrieveQuery(m)
	if err == nil {
		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}
		return view.ViewTransaction(rows), nil
	}
	return nil, err
}

func DeleteTransaction(db *sql.DB, m model.Transaction) error {
	query, err := q.BuildTransactionDeleteQuery(m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}
