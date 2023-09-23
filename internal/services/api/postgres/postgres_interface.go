package postgresinterface

import (
	"database/sql"
	"services/db/postgresql/controller"
	"services/db/postgresql/model"
)

type DBHandler interface {
	RetrieveProfile(db *sql.DB, id string, uid bool) (model.Profile, error)
	CreateProfile(db *sql.DB, user string, password string, randomUID string) error
	CreateExpense(db *sql.DB, m model.Expense) error
	UpdateExpense(db *sql.DB, setM model.Expense, whereM model.Expense) error
	RetrieveExpense(db *sql.DB, m model.Expense) ([]model.Expense, error)
	RetrieveTransaction(db *sql.DB, m model.Transaction) ([]model.Transaction, error)
	RetrieveToken(db *sql.DB, profileId int) (model.Token, error)
	RetrieveAccount(db *sql.DB, profileId int) ([]model.Account, error)
}

type DB struct{}
type MockDB struct {
	Expense []model.Expense
	Transaction []model.Transaction
	Profile model.Profile
	Token model.Token
	Account []model.Account
	ProfileErr error
	ExpenseErr error
	TransactionErr error
	TokenErr error
	AccountErr error
}

func (t MockDB) RetrieveProfile(db *sql.DB, id string, uid bool) (model.Profile, error) {
	return t.Profile, t.ProfileErr
}
func (t MockDB) CreateProfile(db *sql.DB, user string, password string, randomUID string) error {
	return t.ProfileErr
}
func (t MockDB) CreateExpense(db *sql.DB, m model.Expense) error {
	return t.ExpenseErr
}
func (t MockDB) RetrieveExpense(db *sql.DB, m model.Expense) ([]model.Expense, error) {
	return t.Expense, t.ExpenseErr
}
func (t MockDB) UpdateExpense(db *sql.DB, setM model.Expense, whereM model.Expense) error {
	return t.ExpenseErr
}
func (t MockDB) RetrieveTransaction(db *sql.DB, m model.Transaction) ([]model.Transaction, error) {
	return t.Transaction, t.TransactionErr
}
func (t MockDB) RetrieveToken(db *sql.DB, profileId int) (model.Token, error) {
	return t.Token, t.TokenErr
}
func (t MockDB) RetrieveAccount(db *sql.DB, profileId int) ([]model.Account, error) {
	return t.Account, t.AccountErr
}
func (t DB) RetrieveProfile(db *sql.DB, id string, uid bool) (model.Profile, error) {
	var profile model.Profile;
	var err error;
	if uid {
		profile, err = controller.RetrieveProfile(db, model.Profile{ RandomUID: id })
	} else {
		profile, err = controller.RetrieveProfile(db, model.Profile{ Name: id })
	}
	return profile, err
}
func (t DB) CreateProfile(db *sql.DB, user string, password string, randomUID string) error {
	err := controller.CreateProfile(db, model.Profile{
		Name: user,
		Password: password,
		RandomUID: randomUID,
	})
	return err
}
func (t DB) CreateExpense(db *sql.DB, m model.Expense) error {
	err := controller.CreateExpense(db, m)
	return err
}
func (t DB) RetrieveExpense(db *sql.DB, m model.Expense) ([]model.Expense, error) {
	expenses, err := controller.RetrieveExpense(db, m)
	return expenses, err
}
func (t DB) UpdateExpense(db *sql.DB, setM model.Expense, whereM model.Expense) error {
	err := controller.UpdateExpense(db, setM, whereM)
	return err
}
func (t DB) RetrieveTransaction(db *sql.DB, m model.Transaction) ([]model.Transaction, error) {
	transactions, err := controller.RetrieveTransaction(db, m)
	return transactions, err
}
func (t DB) RetrieveToken(db *sql.DB, profileId int) (model.Token, error) {
	token, err := controller.RetrieveToken(db, model.Token{ ProfileID: profileId })
	return token, err
}
func (t DB) RetrieveAccount(db *sql.DB, profileId int) ([]model.Account, error) {
	account, err := controller.RetrieveAccount(db, model.Account{ ProfileID: profileId })
	return account, err
}