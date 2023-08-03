package controller

import (
	"database/sql"

	"budgetbolt/src/services/databases/postgresql/model"
)

type DBHandler interface {
	RetrieveProfile(db *sql.DB, user string) (model.Profile, error)
	CreateProfile(db *sql.DB, user string, password string) error
	RetrieveTransaction(db *sql.DB, m model.Transaction) ([]model.Transaction, error)
	RetrieveToken(db *sql.DB, profileId int) (model.Token, error)
	RetrieveAccount(db *sql.DB, profileId int) ([]model.Account, error)
}

type DB struct{}
type MockDB struct {
	Transaction []model.Transaction
	Profile model.Profile
	Token model.Token
	Account []model.Account
	ProfileErr error
	TransactionErr error
	TokenErr error
	AccountErr error
}

func (t MockDB) RetrieveProfile(db *sql.DB, user string) (model.Profile, error) {
	return t.Profile, t.ProfileErr
}
func (t MockDB) CreateProfile(db *sql.DB, user string, password string) error {
	return t.ProfileErr
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
func (t DB) RetrieveProfile(db *sql.DB, user string) (model.Profile, error) {
	profile, err := RetrieveProfile(db, model.Profile{ Name: user })
	return profile, err
}
func (t DB) CreateProfile(db *sql.DB, user string, password string) error {
	err := CreateProfile(db, model.Profile{
		Name: user,
		Password: password,
	})
	return err
}
func (t DB) RetrieveTransaction(db *sql.DB, m model.Transaction) ([]model.Transaction, error) {
	transactions, err := RetrieveTransaction(db, m)
	return transactions, err
}
func (t DB) RetrieveToken(db *sql.DB, profileId int) (model.Token, error) {
	token, err := RetrieveToken(db, model.Token{ ProfileID: profileId })
	return token, err
}
func (t DB) RetrieveAccount(db *sql.DB, profileId int) ([]model.Account, error) {
	account, err := RetrieveAccount(db, model.Account{ ProfileID: profileId })
	return account, err
}