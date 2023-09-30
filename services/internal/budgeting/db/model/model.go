package model

type Expense struct {
	ID            int64     `json:"expense_id" db:"expense_id"`
	Name          string    `json:"expense_name" db:"expense_name"`
	Limit         *float64  `json:"expense_limit" db:"expense_limit"`
	Spent         *float64  `json:"expense_spent" db:"expense_spent"`
	Category      []string  `json:"expense_categories" db:"expense_categories"`
	TransactionID []string  `json:"transaction_ids" db:"transaction_ids"`
	ProfileID     int64     `json:"profile_id" db:"profile_id"`
}