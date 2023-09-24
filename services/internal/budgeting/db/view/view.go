package view

import (
	"database/sql"
	"strings"
	"github.com/lib/pq"
	
	"services/internal/budgeting/db/model"
)

func ViewExpense(rows *sql.Rows) []model.Expense {
	var view []model.Expense
	defer rows.Close()
	for rows.Next() {
		var id int64
		var name string
		var limit sql.NullFloat64
		var spent float64
		var category pq.StringArray
		var transactionIds pq.StringArray
		var profileId int
		rows.Scan(&id, &name, &limit, &spent, &category, &transactionIds, &profileId)
		view = append(view, model.Expense{
			ID: id,
			Name: strings.TrimSpace(name),
			Spent: &spent,
			Limit: &limit.Float64,
			Category: category,
			TransactionID: transactionIds,
			ProfileID: profileId,
		})
	}
	return view
}