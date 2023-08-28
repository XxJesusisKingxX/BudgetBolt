package controller

import (
	"database/sql"

	q "services/db/postgresql/controller/querybuilder"
	"services/db/postgresql/model"
	"services/db/postgresql/view"
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