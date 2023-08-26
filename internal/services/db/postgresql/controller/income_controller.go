package controller

import (
	"database/sql"
	"fmt"

	q "services/db/postgresql/controller/querybuilder"
	"services/db/postgresql/model"
)

func CreateIncome(db *sql.DB, m model.Income) error {
	query, err := q.BuildCreateQuery("income", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateIncome(db *sql.DB, m model.Income) error {
	query, err := q.BuildUpdateQuery("income", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveIncome(db *sql.DB, m model.Income) error {
	query, err := q.BuildRetrieveQuery("income", m)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func DeleteIncome(db *sql.DB, m model.Income) error {
	query, err := q.BuildDeleteQuery("income", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}