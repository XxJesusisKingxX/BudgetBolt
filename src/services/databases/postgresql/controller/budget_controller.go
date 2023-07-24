package controller

import (
	"database/sql"
	"fmt"

	q "budgetbolt/src/services/databases/postgresql/controller/querybuilder"
	"budgetbolt/src/services/databases/postgresql/model"
)

func CreateBudget(db *sql.DB, m model.Budget) error {
	query, err := q.BuildCreateQuery("budget", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateBudget(db *sql.DB, m model.Budget) error {
	query, err := q.BuildUpdateQuery("budget", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveBudget(db *sql.DB, m model.Budget) error {
	query, err := q.BuildRetrieveQuery("budget", m)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func DeleteBudget(db *sql.DB, m model.Budget) error {
	query, err := q.BuildDeleteQuery("budget", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}