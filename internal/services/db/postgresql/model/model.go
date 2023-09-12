package model

type Budget struct {
	ID          int	   `db:"budget_id"`
	Name        string `db:"budget_name"`
	Description string `db:"short_description"`
	Frequency   string `db:"budget_frequency"`
	ProfileID   int    `db:"profile_id"`
}
type Expense struct {
	ID            int       `db:"expense_id"`
	DueDate       string    `db:"due_date"`// YYYYMMDD
	Name          string    `db:"expense_name"`
	Limit         float64   `db:"expense_limit"`
	BudgetID      int	    `db:"budget_id"`
	TransactionID int	    `db:"transaction_id"`
	Category      string    `db:"expense_category"`
}
type Income struct {
	ID            int       `db:"income_id"`
	Name          string    `db:"income_name"`
	Expected      float64   `db:"income_amount_expected"`
	BudgetID      int	    `db:"budget_id"`
	TransactionID int	    `db:"transaction_id"`
	Category      string    `db:"income_category"`
	DueDate       string    `db:"due_date"`// YYYYMMDD
}
type Transaction struct {
	ID          string     `db:"transaction_id"`
	Date        string    `db:"transaction_date"`// YYYYMMDD
	Amount      float64   `db:"net_amount"`
	Method      string    `db:"payment_method"`
	From        string    `db:"payment_account_from_to"`
	Vendor      string    `db:"vendor"`
	IsRecurring bool      `db:"is_recurring"`
	Description string    `db:"short_description"`
	ProfileID   int       `db:"profile_id"`
	Query       Querys

}
type Account struct {
	ID             int	   `db:"account_id"`
	Name           string  `db:"account_name"`
	Balance        float64 `db:"account_balance"`
	PlaidAccountID string  `db:"plaid_account_id"`
	ProfileID      int     `db:"profile_id"`
}
type Investment struct {
	ID          int     `db:"investment_id"`
	PlaidId     string  `db:"plaid_id"`
	Fees        float64 `db:"fees"`
	Amount      float64 `db:"purchase_amount"`
	Date        string  `db:"purchase_date"`
	Description string  `db:"short_description"`
	Price       float64 `db:"price_at"`
	Quantity    float64 `db:"purchase_quantity"`
}

type Holding struct {
	ID          int     `db:"holding_id"`
	PlaidId     string  `db:"plaid_id"`
	CostBasis   float64 `db:"cost_basis"`
	TotalValue  float64 `db:"total_value"`
	LastPrice   float64 `db:"last_price"`
	Quantity    float64 `db:"purchase_quantity"`
}
type Profile struct {
	ID          int  `db:"profile_id"`
	Name      string `db:"profile_name"`
	Password  string `db:"profile_password"`
	RandomUID string `db:"v"`
}
type Token struct {
	ID          int  `db:"token_id"`
	Item      string `db:"item_id"`
	Token     string `db:"access_token"`
	ProfileID    int    `db:"profile_id"`
}