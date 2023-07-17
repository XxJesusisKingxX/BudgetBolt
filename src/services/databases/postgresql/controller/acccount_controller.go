package controller

import (
	"database/sql"
	"fmt"

	q "budgetbolt/src/services/databases/postgresql/controller/querybuilder"
	table "budgetbolt/src/services/databases/postgresql/model"
)

func CreateAccount(db *sql.DB, table table.Account) error {
	query, err := q.BuildCreateQuery("account", table)
	if err == nil {
		db.Exec(query)
	}
	return err
}

func UpdateAccount(db *sql.DB, table table.Account) error {
	query, err := q.BuildUpdateQuery("account", table)
	if err == nil {
		db.Exec(query)
	}
	return err
}

func RetrieveAccount(db *sql.DB, table table.Account) error {
	query, err := q.BuildRetrieveQuery("account", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func DeleteAccount(db *sql.DB, table table.Account) error {
	query, err := q.BuildDeleteQuery("account", table)
	if err == nil {
		db.Exec(query)
	}
	return err
}