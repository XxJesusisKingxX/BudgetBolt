package view

import (
	"database/sql"
	"strings"
	
	"services/internal/transaction_history/db/model"
)

func ViewAccount(rows *sql.Rows) []model.Account {
	var view []model.Account
	defer rows.Close()
	for rows.Next() {
		var id string
		var name string
		var balance float64
		var profileId int64
		rows.Scan(&id, &name, &balance, &profileId)
		view = append(view, model.Account{
			ID: id,
			Name: name,
			Balance: balance,
			ProfileID: profileId,
		})
	}
	return view
}

func ViewTransaction(rows *sql.Rows) []model.Transaction {
	var view []model.Transaction
	defer rows.Close()
	for rows.Next() {
		var id string
		var date string
		var amount float64
		var method sql.NullString
		var vendor sql.NullString
		var isRecurring bool
		var description string
		var primaryCat sql.NullString
		var secondaryCat sql.NullString
		var profileId int64
		var accName string
		rows.Scan(&id, &date, &amount, &method, &vendor, &isRecurring, &description, &primaryCat, &secondaryCat, &profileId, &accName)
		view = append(view, model.Transaction{
			ID: id,
			Date: strings.Split(date,"T")[0],
			Amount: amount,
			Method: strings.TrimSpace(method.String),
			Vendor: strings.TrimSpace(vendor.String),
			IsRecurring: isRecurring,
			Description: strings.TrimSpace(description),
			ProfileID: profileId,
			PrimaryCategory: strings.TrimSpace(primaryCat.String),
			DetailCategory: strings.TrimSpace(secondaryCat.String),
			AccountName: accName,
		})
	}
	return view
}