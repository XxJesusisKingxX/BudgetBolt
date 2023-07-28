package controller

import (
	"database/sql"
	
	q "budgetbolt/src/services/databases/postgresql/controller/querybuilder"
	"budgetbolt/src/services/databases/postgresql/model"
	"budgetbolt/src/services/databases/postgresql/view"
)

func CreateToken(db *sql.DB, m model.Token) error {
	query, err := q.BuildCreateQuery("token", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateToken(db *sql.DB, m model.Token) error {
	query, err := q.BuildUpdateQuery("token", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveToken(db *sql.DB, m model.Token) (model.Token, error) {
	query, err := q.BuildRetrieveQuery("token", m)
	if err == nil {
		rows, err := db.Query(query)
		if err != nil {
			return model.Token{}, err
		}
		return view.ViewToken(rows), nil
	}
	return model.Token{}, err
}

func DeleteToken(db *sql.DB, m model.Token) error {
	query, err := q.BuildDeleteQuery("token", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}