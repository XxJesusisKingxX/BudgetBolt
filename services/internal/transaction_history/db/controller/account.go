package controller

import (
	"database/sql"

	"services/internal/transaction_history/db/view"
	"services/internal/transaction_history/db/model"
	q "services/internal/utils/sql/querybuilder"
)

func CreateAccount(db *sql.DB, m model.Account) error {
	query, err := q.BuildCreateQuery("accounts", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateAccount(db *sql.DB, setM model.Account, whereM model.Account) error {
	query, err := q.BuildUpdateQuery("accounts", setM, whereM)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveAccount(db *sql.DB, m model.Account) ([]model.Account , error) {
	query, err := q.BuildRetrieveQuery("accounts", m)
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
	query, err := q.BuildDeleteQuery("accounts", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}