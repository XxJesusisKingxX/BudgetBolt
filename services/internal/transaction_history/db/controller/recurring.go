package controller

import (
	"database/sql"

	"services/internal/transaction_history/db/view"
	"services/internal/transaction_history/db/model"
	q "services/internal/utils/sql/querybuilder"
)

const (
	RECURRINGS string = "recurrings"
)

func CreateRecurringTransaction(db *sql.DB, m model.RecurringTransaction) error {
	query, err := q.BuildCreateQuery(RECURRINGS, m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateRecurringTransaction(db *sql.DB, setM model.RecurringTransaction, whereM model.RecurringTransaction) error {
	query, err := q.BuildUpdateQuery(RECURRINGS, setM, whereM)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveRecurringTransaction(db *sql.DB, m model.RecurringTransaction) ([]model.RecurringTransaction , error) {
	query, err := q.BuildRetrieveQuery(RECURRINGS, m)
	if err == nil {
		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}
		return view.ViewRecurringTransaction(rows), nil
	}
	return nil, err
}

func DeleteRecurringTransaction(db *sql.DB, m model.RecurringTransaction) error {
	query, err := q.BuildDeleteQuery(RECURRINGS, m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}