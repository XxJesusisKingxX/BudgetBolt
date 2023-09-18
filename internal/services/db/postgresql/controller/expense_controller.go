package controller

import (
	"database/sql"
	
	q "services/db/postgresql/controller/querybuilder"
	"services/db/postgresql/model"
	"services/db/postgresql/view"
)

func CreateExpense(db *sql.DB, m model.Expense) error {
	query, err := q.BuildCreateQuery("expense", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateExpense(db *sql.DB, setM model.Expense, whereM model.Expense) error {
	query, err := q.BuildUpdateQuery("expense", setM, whereM)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveExpense(db *sql.DB, m model.Expense) ([]model.Expense, error) {
	query, err := q.BuildRetrieveQuery("expense", m)
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
	query, err := q.BuildDeleteQuery("expense", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}