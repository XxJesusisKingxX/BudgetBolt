package api

import (
	"database/sql"

	"services/internal/transaction_history/db/controller"
	"services/internal/transaction_history/db/model"
	user "services/internal/user_management/db/model"
)

type DBHandler interface {
	RetrieveTransaction(db *sql.DB, m model.Transaction) ([]model.Transaction, error)
}

type DB struct{}
type MockDB struct {
	Profile user.Profile
	Transaction []model.Transaction
	TransactionErr error
	ProfileErr error
}

func (t MockDB) RetrieveTransaction(db *sql.DB, m model.Transaction) ([]model.Transaction, error) {
	return t.Transaction, t.TransactionErr
}

func (t DB) RetrieveTransaction(db *sql.DB, m model.Transaction) ([]model.Transaction, error) {
	transactions, err := controller.RetrieveTransaction(db, m)
	return transactions, err
}