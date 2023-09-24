package api

import (
	"database/sql"
	"math"

	budget "services/internal/budgeting/db/model"
	transaction "services/internal/transaction_history/db/model"
)

func UpdateAllExpenses(transactions []transaction.Transaction, dbs DBHandler, db *sql.DB, profileId int) {
	// Get the category totals for all possible expenses
	categoryTotals := make(map[string]float64)
	for _, transaction := range transactions {
		primary := transaction.PrimaryCategory
		detailed := transaction.DetailCategory
		categoryTotals[primary] += transaction.Amount
		categoryTotals[detailed] += transaction.Amount
	}
	// Update expense total spent
	expenses, _ := dbs.RetrieveExpense(db, budget.Expense{
		ProfileID: profileId,
	})
	for _, expense := range expenses {
		var total float64
		categories := expense.Category
		for _, category := range categories {
			total += categoryTotals[category]
		}
		total = math.Round(total*100) / 100 // round 2decimals
		dbs.UpdateExpense(db, budget.Expense{
			Spent: &total,
		}, budget.Expense{
			ID: expense.ID,
		})
	}
}