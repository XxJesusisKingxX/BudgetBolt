package view

import (
	"services/db/postgresql/model"

	"database/sql"
)

func ViewAccount(rows *sql.Rows) []model.Account {
	var view []model.Account
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var balance float64
		var plaidAccId string
		var profileId int
		rows.Scan(&id, &name, &balance, &plaidAccId, &profileId)
		view = append(view, model.Account{
			ID: id, 
			Name: name,
			Balance: balance,
			PlaidAccountID: plaidAccId,
			ProfileID: profileId,
		})
	}
	return view
}