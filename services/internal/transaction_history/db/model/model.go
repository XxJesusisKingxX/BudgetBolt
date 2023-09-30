package model

type Transaction struct {
	ID              string  `json:"transaction_id" db:"transaction_id"`
    Date            string  `json:"transaction_date" db:"transaction_date"` // YYYYMMDD
    Amount          float64 `json:"net_amount" db:"net_amount"`
    Method          string  `json:"payment_method" db:"payment_method"`
    Vendor          string  `json:"vendor" db:"vendor"`
    IsRecurring     bool    `json:"is_recurring" db:"is_recurring"`
    Description     string  `json:"description" db:"description"`
    PrimaryCategory string  `json:"primary_category" db:"primary_category"`
    DetailCategory  string  `json:"secondary_category" db:"secondary_category"`
    ProfileID       int64   `json:"profile_id" db:"profile_id"`
    AccountName     string  `json:"from_account" db:"from_account"`
	Query            Querys

}
type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}
type Account struct {
	ID             string  `db:"account_id"`
	Name           string  `db:"account_name"`
	Balance        float64 `db:"account_balance"`
	ProfileID      int64   `db:"profile_id"`
}