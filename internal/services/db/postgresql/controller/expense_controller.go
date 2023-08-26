package controller

import (
	"database/sql"
	"fmt"

	q "services/db/postgresql/controller/querybuilder"
	"services/db/postgresql/model"
)

func CreateExpense(db *sql.DB, m model.Expense) error {
	query, err := q.BuildCreateQuery("expense", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateExpense(db *sql.DB, m model.Expense) error {
	query, err := q.BuildUpdateQuery("expense", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveExpense(db *sql.DB, m model.Expense) error {
	query, err := q.BuildRetrieveQuery("expense", m)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func DeleteExpense(db *sql.DB, m model.Expense) error {
	query, err := q.BuildDeleteQuery("expense", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}