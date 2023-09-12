package view

import (
	"database/sql"
	"strings"

	"services/db/postgresql/model"
)

func ViewProfile(rows *sql.Rows) model.Profile {
	var id int
	var user string
	var pass string
	var randomUID string
	defer rows.Close()
	rows.Next()
	rows.Scan(&id, &user, &pass, &randomUID)
	view := model.Profile{ID: id, Name: strings.TrimSpace(user), Password: strings.TrimSpace(pass), RandomUID: randomUID}
	return view
}