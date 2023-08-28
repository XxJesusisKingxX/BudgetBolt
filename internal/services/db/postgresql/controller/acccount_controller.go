package controller

import (
	"database/sql"

	q "services/db/postgresql/controller/querybuilder"
	"services/db/postgresql/model"
	"services/db/postgresql/view"
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