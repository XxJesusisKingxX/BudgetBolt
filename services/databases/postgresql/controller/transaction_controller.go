package controller

import (
	"database/sql"
	"fmt"

	q "budgetbolt/services/databases/postgresql/controller/querybuilder"
	table "budgetbolt/services/databases/postgresql/model"
)

func CreateTransaction(db *sql.DB, table table.Transaction) error {
	query, err := q.BuildCreateQuery("transaction", table)
	if err == nil {
		db.Exec(query)
	}
	return err
}

func UpdateTransaction(db *sql.DB, table table.Transaction) error {
	query, err := q.BuildUpdateQuery("transaction", table)
	if err == nil {
		db.Exec(query)
	}
	return err
}

func RetrieveTransaction(db *sql.DB, table table.Transaction) error {
	query, err := q.BuildRetrieveQuery("transaction", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func DeleteTransaction(db *sql.DB, table table.Transaction) error {
	query, err := q.BuildDeleteQuery("transaction", table)
	if err == nil {
		db.Exec(query)
	}
	return err
}