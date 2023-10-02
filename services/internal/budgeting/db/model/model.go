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

type Income struct {
	ID            int64     `json:"income_id" db:"income_id"`
	Name          string    `json:"income_name" db:"income_name"`
	Amount       *float64   `json:"income_amount" db:"income_amount"`
	ProfileID     int64     `json:"profile_id" db:"profile_id"`
}