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
		var method sql.NullString
		var from string
		var vendor sql.NullString
		var isRecurring bool
		var description sql.NullString
		var profileId int
		var primary sql.NullString
		var detail sql.NullString
		rows.Scan(&date, &amount, &method, &from, &vendor, &isRecurring, &description, &profileId, &id, &primary, &detail)
		view = append(view, model.Transaction{
			ID: id,
			Date: strings.Split(date,"T")[0],
			Amount: amount,
			Method: strings.TrimSpace(method.String),
			From: strings.TrimSpace(from),
			Vendor: strings.TrimSpace(vendor.String),
			IsRecurring: isRecurring,
			Description: strings.TrimSpace(description.String),
			ProfileID: profileId,
			PrimaryCategory: strings.TrimSpace(primary.String),
			DetailCategory: strings.TrimSpace(detail.String),
		})
	}
	return view
}