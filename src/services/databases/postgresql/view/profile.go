package view

import (
	"database/sql"
	"strings"

	"budgetbolt/src/services/databases/postgresql/model"
)

func ViewProfile(rows *sql.Rows) model.Profile {
	var id int
	var user string
	var pass string
	defer rows.Close()
	rows.Next()
	rows.Scan(&id, &user, &pass)
	view := model.Profile{ID: id, Name: strings.TrimSpace(user), Password: strings.TrimSpace(pass)}
	return view
}