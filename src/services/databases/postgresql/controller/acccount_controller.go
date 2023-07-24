package controller

import (
	"database/sql"

	q "budgetbolt/src/services/databases/postgresql/controller/querybuilder"
	"budgetbolt/src/services/databases/postgresql/model"
	"budgetbolt/src/services/databases/postgresql/view"
)

func CreateAccount(db *sql.DB, m model.Account) error {
	query, err := q.BuildCreateQuery("account", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateAccount(db *sql.DB, m model.Account) error {
	query, err := q.BuildUpdateQuery("account", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveAccount(db *sql.DB, m model.Account) ([]model.Account , error) {
	query, err := q.BuildRetrieveQuery("account", m)
	if err == nil {
		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}
		return view.ViewAccount(rows), nil
	}
	return nil, err
}

func DeleteAccount(db *sql.DB, m model.Account) error {
	query, err := q.BuildDeleteQuery("account", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}