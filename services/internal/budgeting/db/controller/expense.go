package controller

import (
	"database/sql"

	"services/internal/budgeting/db/model"
	"services/internal/budgeting/db/view"
	q "services/internal/utils/sql/querybuilder"
)

func CreateExpense(db *sql.DB, m model.Expense) error {
	query, err := q.BuildCreateQuery("expenses", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateExpense(db *sql.DB, setM model.Expense, whereM model.Expense) error {
	query, err := q.BuildUpdateQuery("expenses", setM, whereM)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveExpense(db *sql.DB, m model.Expense) ([]model.Expense, error) {
	query, err := q.BuildRetrieveQuery("expenses", m)
	if err == nil {
		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}
		return view.ViewExpense(rows), nil
	}
	return nil, err
}

func DeleteExpense(db *sql.DB, m model.Expense) error {
	query, err := q.BuildDeleteQuery("expenses", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}