package controller

import (
	"database/sql"
	"fmt"

	q "budgetbolt/services/databases/postgresql/controller/querybuilder"
	table "budgetbolt/services/databases/postgresql/model"
)

func CreateInvestment(db *sql.DB, table table.Investment) error {
	query, err := q.BuildCreateQuery("investment", table)
	if err == nil {
		db.Exec(query)
	}
	return err
}

func UpdateInvestment(db *sql.DB, table table.Investment) error {
	query, err := q.BuildUpdateQuery("investment", table)
	if err == nil {
		db.Exec(query)
	}
	return err
}

func RetrieveInvestment(db *sql.DB, table table.Investment) error {
	query, err := q.BuildRetrieveQuery("investment", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func DeleteInvestment(db *sql.DB, table table.Investment) error {
	query, err := q.BuildDeleteQuery("investment", table)
	if err == nil {
		db.Exec(query)
	}
	return err
}