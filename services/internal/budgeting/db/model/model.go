package model

type Expense struct {
	ID            int64    `db:"expense_id"`
	Name          string   `db:"expense_name"`
	Limit         *float64 `db:"expense_limit"`
	Spent         *float64 `db:"expense_spent"`
	Category      []string `db:"expense_categories"`
	TransactionID []string `db:"transaction_ids"`
	ProfileID     int	   `db:"profile_id"`
}