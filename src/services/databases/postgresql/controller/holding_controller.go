package controller

import (
	"database/sql"
	"fmt"

	q "budgetbolt/src/services/databases/postgresql/controller/querybuilder"
	table "budgetbolt/src/services/databases/postgresql/model"
)

func CreateHolding(db *sql.DB, table table.Holding) error {
	query, err := q.BuildCreateQuery("holding", table)
	if err == nil {
		db.Exec(query)
	}
	return err
}

func UpdateHolding(db *sql.DB, table table.Holding) error {
	query, err := q.BuildUpdateQuery("holding", table)
	if err == nil {
		db.Exec(query)
	}
	return err
}

func RetrieveHolding(db *sql.DB, table table.Holding) error {
	query, err := q.BuildRetrieveQuery("holding", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func DeleteHolding(db *sql.DB, table table.Holding) error {
	query, err := q.BuildDeleteQuery("holding", table)
	if err == nil {
		db.Exec(query)
	}
	return err
}