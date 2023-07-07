package model

type Budget struct {
	ID          int	   `db:"budget_id"`
	Name        string `db:"budget_name"`
	Description string `db:"short_description"`
	Frequency   string `db:"budget_frequency"`
}
type Expense struct {
	ID            int     `db:"expense_id"`
	DueDate       string  `db:"due_date"`// YYYYMMDD
	Name          string  `db:"expense_name"`
	Limit         float64 `db:"expense_limit"`
	BudgetID      int	  `db:"budget_id"`
	TransactionID int	  `db:"transaction_id"`
	Category      string  `db:"expense_category"`
}
type Income struct {
	ID            int     `db:"income_id"`
	Name          string  `db:"income_name"`
	Expected      float64 `db:"income_amount_expected"`
	BudgetID      int	  `db:"budget_id"`
	TransactionID int	  `db:"transaction_id"`
	Category      string  `db:"income_category"`
	DueDate       string  `db:"due_date"`// YYYYMMDD
}
type Transaction struct {
	ID          int     `db:"transaction_id"`
	Date        string  `db:"transaction_date"`// YYYYMMDD
	Amount      float64 `db:"net_amount"`
	Method      string  `db:"payment_method"`
	From        string  `db:"payment_account_from_to"`
	Vendor      string  `db:"vendor"`
	IsRecurring bool    `db:"is_recurring"`
	Description string  `db:"short_description"`
}