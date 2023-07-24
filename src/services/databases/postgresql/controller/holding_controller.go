package controller

import (
	"database/sql"
	"fmt"

	q "budgetbolt/src/services/databases/postgresql/controller/querybuilder"
	"budgetbolt/src/services/databases/postgresql/model"
)

func CreateHolding(db *sql.DB, m model.Holding) error {
	query, err := q.BuildCreateQuery("holding", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateHolding(db *sql.DB, m model.Holding) error {
	query, err := q.BuildUpdateQuery("holding", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveHolding(db *sql.DB, m model.Holding) error {
	query, err := q.BuildRetrieveQuery("holding", m)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func DeleteHolding(db *sql.DB, m model.Holding) error {
	query, err := q.BuildDeleteQuery("holding", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}