package view

import (
	"services/db/postgresql/model"
	"strings"

	"database/sql"
)

func ViewExpense(rows *sql.Rows) []model.Expense {
	var view []model.Expense
	defer rows.Close()
	for rows.Next() {
		var id int
		var date sql.NullString
		var name string
		var limit sql.NullFloat64
		var transactionId sql.NullInt32
		var category sql.NullString
		var spent float64
		var profileId int
		rows.Scan(&id, &date, &name, &limit, &transactionId, &category, &spent, &profileId)

		view = append(view, model.Expense{
			ID: id,
			Name: strings.TrimSpace(name),
			Spent: spent,
			Limit: limit.Float64,
			Category: strings.TrimSpace(category.String),
			DueDate: date.String,
			TransactionID: transactionId.Int32,
			ProfileID: profileId,
		})
	}
	return view
}