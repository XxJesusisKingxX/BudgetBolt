package controller

import (
	"database/sql"

	q "budgetbolt/src/services/databases/postgresql/controller/querybuilder"
	"budgetbolt/src/services/databases/postgresql/model"
	"budgetbolt/src/services/databases/postgresql/view"
)

func CreateProfile(db *sql.DB, m model.Profile) error {
	query, err := q.BuildCreateQuery("profile", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateProfile(db *sql.DB, m model.Profile) error {
	query, err := q.BuildUpdateQuery("profile", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveProfile(db *sql.DB, m model.Profile) (model.Profile, error) {
	query, err := q.BuildRetrieveQuery("profile", m)
	if err == nil {
		rows, err := db.Query(query)
		if err != nil {
			return model.Profile{}, err
		}
		return view.ViewProfile(rows), nil
	}
	return model.Profile{}, err
}

func DeleteProfile(db *sql.DB, m model.Profile) error {
	query, err := q.BuildDeleteQuery("profile", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}