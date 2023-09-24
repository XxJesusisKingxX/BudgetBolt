package view

import (
	"database/sql"
	"strings"
	
	"services/internal/user_managment/db/model"
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

func ViewToken(rows *sql.Rows) model.Token {
	var id int
	var itemId string
	var accesstoken string
	var profileId int
	defer rows.Close()
	rows.Next()
	rows.Scan(&id, &itemId, &accesstoken, &profileId)
	view := model.Token{ ID: id, Item: strings.TrimSpace(itemId), Token: strings.TrimSpace(accesstoken) }
	return view
}