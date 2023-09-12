package view

import (
	"services/db/postgresql/model"

	"database/sql"
)

func ViewAccount(rows *sql.Rows) []model.Account {
	var view []model.Account
	defer rows.Close()
	for rows.Next() {
		var id string
		var name string
		var balance float64
		var profileId int
		rows.Scan(&name, &balance, &id, &profileId)
		view = append(view, model.Account{
			ID: id,
			Name: name,
			Balance: balance,
			ProfileID: profileId,
		})
	}
	return view
}