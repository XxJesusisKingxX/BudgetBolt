package api

import (
	"database/sql"

	"services/internal/transaction_history/db/controller"
	"services/internal/transaction_history/db/model"
	user "services/internal/user_management/db/model"
)

type DBHandler interface {
	RetrieveTransaction(db *sql.DB, m model.Transaction) ([]model.Transaction, error)
	DeleteTransaction(db *sql.DB, m model.Transaction) error
	RetrieveRecurringTransaction(db *sql.DB, m model.RecurringTransaction) ([]model.RecurringTransaction, error)
}

type DB struct{}
type MockDB struct {
	Profile user.Profile
	Transaction []model.Transaction
	RecurringTransaction []model.RecurringTransaction
	TransactionErr error
	ProfileErr error
}

func (t MockDB) RetrieveTransaction(db *sql.DB, m model.Transaction) ([]model.Transaction, error) {
	return t.Transaction, t.TransactionErr
}
func (t MockDB) DeleteTransaction(db *sql.DB, m model.Transaction) error {
	return t.TransactionErr
}
func (t MockDB) RetrieveRecurringTransaction(db *sql.DB, m model.RecurringTransaction) ([]model.RecurringTransaction, error) {
	return t.RecurringTransaction, t.TransactionErr
}

func (t DB) RetrieveTransaction(db *sql.DB, m model.Transaction) ([]model.Transaction, error) {
	transactions, err := controller.RetrieveTransaction(db, m)
	return transactions, err
}
func (t DB) DeleteTransaction(db *sql.DB, m model.Transaction) error {
	err := controller.DeleteTransaction(db, m)
	return err
}
func (t DB) RetrieveRecurringTransaction(db *sql.DB, m model.RecurringTransaction) ([]model.RecurringTransaction, error) {
	transactions, err := controller.RetrieveRecurringTransaction(db, m)
	return transactions, err
}