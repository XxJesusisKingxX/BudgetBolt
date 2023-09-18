package controller

import (
	"database/sql"
	"fmt"

	q "services/db/postgresql/controller/querybuilder"
	"services/db/postgresql/model"
)

func CreateInvestment(db *sql.DB, m model.Investment) error {
	query, err := q.BuildCreateQuery("investment", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateInvestment(db *sql.DB, setM model.Investment, whereM model.Investment) error {
	query, err := q.BuildUpdateQuery("investment", setM, whereM)
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