package api

import (
	"database/sql"

	budgetCtrl "services/internal/budgeting/db/controller"
	userCtrl "services/internal/user_managment/db/controller"
	transactionCtrl "services/internal/transaction_history/db/controller"
	budget "services/internal/budgeting/db/model"
	transaction "services/internal/transaction_history/db/model"
	user "services/internal/user_managment/db/model"
)

type DBHandler interface {
	RetrieveProfile(db *sql.DB, id string, uid bool) (user.Profile, error)
	CreateProfile(db *sql.DB, user string, password string, randomUID string) error
	CreateToken(db *sql.DB, m user.Token) error
	CreateExpense(db *sql.DB, m budget.Expense) error
	UpdateExpense(db *sql.DB, setM budget.Expense, whereM budget.Expense) error
	RetrieveExpense(db *sql.DB, m budget.Expense) ([]budget.Expense, error)
	RetrieveTransaction(db *sql.DB, m transaction.Transaction) ([]transaction.Transaction, error)
	RetrieveToken(db *sql.DB, profileId int) (user.Token, error)
	RetrieveAccount(db *sql.DB, profileId int) ([]transaction.Account, error)
}

type DB struct{}
type MockDB struct {
	Expense []budget.Expense
	Transaction []transaction.Transaction
	Profile user.Profile
	Token user.Token
	Account []transaction.Account
	ProfileErr error
	ExpenseErr error
	TransactionErr error
	TokenErr error
	AccountErr error
}

func (t MockDB) RetrieveProfile(db *sql.DB, id string, uid bool) (user.Profile, error) {
	return t.Profile, t.ProfileErr
}
func (t MockDB) CreateProfile(db *sql.DB, user string, password string, randomUID string) error {
	return t.ProfileErr
}
func (t MockDB) CreateToken(db *sql.DB, m user.Token) error{
	return t.ExpenseErr
}
func (t MockDB) CreateExpense(db *sql.DB, m budget.Expense) error {
	return t.ExpenseErr
}
func (t MockDB) RetrieveExpense(db *sql.DB, m budget.Expense) ([]budget.Expense, error) {
	return t.Expense, t.ExpenseErr
}
func (t MockDB) UpdateExpense(db *sql.DB, setM budget.Expense, whereM budget.Expense) error {
	return t.ExpenseErr
}
func (t MockDB) RetrieveTransaction(db *sql.DB, m transaction.Transaction) ([]transaction.Transaction, error) {
	return t.Transaction, t.TransactionErr
}
func (t MockDB) RetrieveToken(db *sql.DB, profileId int) (user.Token, error) {
	return t.Token, t.TokenErr
}
func (t MockDB) RetrieveAccount(db *sql.DB, profileId int) ([]transaction.Account, error) {
	return t.Account, t.AccountErr
}
func (t DB) RetrieveProfile(db *sql.DB, id string, uid bool) (user.Profile, error) {
	var profile user.Profile;
	var err error;
	if uid {
		profile, err = userCtrl.RetrieveProfile(db, user.Profile{ RandomUID: id })
	} else {
		profile, err = userCtrl.RetrieveProfile(db, user.Profile{ Name: id })
	}
	return profile, err
}
func (t DB) CreateProfile(db *sql.DB, name string, password string, randomUID string) error {
	err := userCtrl.CreateProfile(db, user.Profile{
		Name: name,
		Password: password,
		RandomUID: randomUID,
	})
	return err
}
func (t DB) CreateExpense(db *sql.DB, m budget.Expense) error {
	err := budgetCtrl.CreateExpense(db, m)
	return err
}
func (t DB) RetrieveExpense(db *sql.DB, m budget.Expense) ([]budget.Expense, error) {
	expenses, err := budgetCtrl.RetrieveExpense(db, m)
	return expenses, err
}
func (t DB) UpdateExpense(db *sql.DB, setM budget.Expense, whereM budget.Expense) error {
	err := budgetCtrl.UpdateExpense(db, setM, whereM)
	return err
}
func (t DB) RetrieveTransaction(db *sql.DB, m transaction.Transaction) ([]transaction.Transaction, error) {
	transactions, err := transactionCtrl.RetrieveTransaction(db, m)
	return transactions, err
}
func (t DB) CreateToken(db *sql.DB, m user.Token) error {
	err := userCtrl.CreateToken(db, m)
	return err
}
func (t DB) RetrieveToken(db *sql.DB, profileId int) (user.Token, error) {
	token, err := userCtrl.RetrieveToken(db, user.Token{ ProfileID: profileId })
	return token, err
}
func (t DB) RetrieveAccount(db *sql.DB, profileId int) ([]transaction.Account, error) {
	account, err := transactionCtrl.RetrieveAccount(db, transaction.Account{ ProfileID: profileId })
	return account, err
}