package controller

import (
	"database/sql"
	"fmt"

	q "budgetbolt/src/services/databases/postgresql/controller/querybuilder"
	"budgetbolt/src/services/databases/postgresql/model"
)

func CreateInvestment(db *sql.DB, m model.Investment) error {
	query, err := q.BuildCreateQuery("investment", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateInvestment(db *sql.DB, m model.Investment) error {
	query, err := q.BuildUpdateQuery("investment", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveInvestment(db *sql.DB, m model.Investment) error {
	query, err := q.BuildRetrieveQuery("investment", m)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func DeleteInvestment(db *sql.DB, m model.Investment) error {
	query, err := q.BuildDeleteQuery("investment", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}