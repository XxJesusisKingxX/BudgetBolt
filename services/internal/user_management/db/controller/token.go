package controller

import (
	"database/sql"

	"services/internal/user_management/db/view"
	"services/internal/user_management/db/model"
	q "services/internal/utils/sql/querybuilder"
)

func CreateToken(db *sql.DB, m model.Token) error {
	query, err := q.BuildCreateQuery("tokens", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateToken(db *sql.DB, setM model.Token, whereM model.Token) error {
	query, err := q.BuildUpdateQuery("tokens", setM, whereM)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveToken(db *sql.DB, m model.Token) ([]model.Token, error) {
	query, err := q.BuildRetrieveQuery("tokens", m)
	if err == nil {
		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}
		return view.ViewToken(rows), nil
	}
	return nil, err
}

func DeleteToken(db *sql.DB, m model.Token) error {
	query, err := q.BuildDeleteQuery("tokens", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}