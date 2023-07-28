package view

import (
	"database/sql"
	"strings"

	"budgetbolt/src/services/databases/postgresql/model"
)


func ViewToken(rows *sql.Rows) model.Token {
	var id int
	var itemId string
	var accesstoken string
	var profileId int
	rows.Next() 
	rows.Scan(&id, &itemId, &accesstoken, &profileId)
	view := model.Token{ ID: id, Item: strings.TrimSpace(itemId), Token: strings.TrimSpace(accesstoken) }
	return view
}