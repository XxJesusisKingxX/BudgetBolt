package controller

import (
	"database/sql"

	q "services/internal/utils/sql/querybuilder"
	"services/internal/user_managment/db/view"
	"services/internal/user_managment/db/model"
)

func CreateProfile(db *sql.DB, m model.Profile) error {
	query, err := q.BuildCreateQuery("profiles", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func UpdateProfile(db *sql.DB, setM model.Profile, whereM model.Profile) error {
	query, err := q.BuildUpdateQuery("profiles", setM, whereM)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}

func RetrieveProfile(db *sql.DB, m model.Profile) (model.Profile, error) {
	query, err := q.BuildRetrieveQuery("profiles", m)
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
	query, err := q.BuildDeleteQuery("profiles", m)
	if err == nil {
		_, err := db.Exec(query)
		return err
	}
	return err
}