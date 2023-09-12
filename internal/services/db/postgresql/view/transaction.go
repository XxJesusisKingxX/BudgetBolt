package view

import (
	"database/sql"
	"strings"

	"services/db/postgresql/model"
)

func ViewTransaction(rows *sql.Rows) []model.Transaction {
	var view []model.Transaction
	defer rows.Close()
	for rows.Next() {
		var id string
		var date string
		var amount float64
		var method *string
		var from string
		var vendor *string
		var isRecurring bool
		var description *string
		var profileId int
		rows.Scan(&date, &amount, &method, &from, &vendor, &isRecurring, &description, &profileId, &id)
		// If not pointing to empty string will lose information
		na := ""
		if description == nil {
			description = &na
		}
		if vendor == nil {
			vendor = description
		}
		if method == nil {
			method = &na
		}
		view = append(view, model.Transaction{
			ID: id, 
			Date: strings.Split(date,"T")[0],
			Amount: amount,
			Method: strings.TrimSpace(*method),
			From: strings.TrimSpace(from),
			Vendor: strings.TrimSpace(*vendor),
			IsRecurring: isRecurring,
			Description: strings.TrimSpace(*description),
			ProfileID: profileId,
		})
	}
	return view
}