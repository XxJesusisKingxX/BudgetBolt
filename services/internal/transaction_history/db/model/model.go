package model

type Transaction struct {
	ID               string    `db:"transaction_id"`
	Date             string    `db:"transaction_date"`// YYYYMMDD
	Amount           float64   `db:"net_amount"`
	Method           string    `db:"payment_method"`
	Vendor           string    `db:"vendor"`
	IsRecurring      bool      `db:"is_recurring"`
	Description      string    `db:"description"`
	PrimaryCategory  string    `db:"primary_category"`
	DetailCategory   string    `db:"secondary_category"`
	ProfileID        int       `db:"profile_id"`
	AccountName      string    `db:"from_account"`
	Query            Querys

}
type Account struct {
	ID             string  `db:"account_id"`
	Name           string  `db:"account_name"`
	Balance        float64 `db:"account_balance"`
	ProfileID      int     `db:"profile_id"`
}