package controller

import (
	"database/sql"
	
	"services/internal/transaction_history/db/model"
	"services/internal/transaction_history/db/view"
	q "services/internal/utils/sql/querybuilder"
)

const (
	TRANSACTIONS string = "transactions"
)

func CreateTransaction(db *sql.DB, m model.Transaction) error {
	query, err := q.BuildCreateQuery(TRANSACTIONS, m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateTransaction(db *sql.DB, setM model.Transaction, whereM model.Transaction) error {
	query, err := q.BuildUpdateQuery(TRANSACTIONS, setM, whereM)
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
	query, err := q.BuildDeleteQuery(TRANSACTIONS, m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}
