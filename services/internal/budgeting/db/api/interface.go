package api

import (
	"database/sql"

	"services/internal/budgeting/db/controller"
	"services/internal/budgeting/db/model"
	user "services/internal/user_management/db/model"
)

type DBHandler interface {
	CreateExpense(db *sql.DB, m model.Expense) error
	UpdateExpense(db *sql.DB, setM model.Expense, whereM model.Expense) error
	RetrieveExpense(db *sql.DB, m model.Expense) ([]model.Expense, error)

}

type DB struct{}
type MockDB struct {
	Profile user.Profile
	Expense []model.Expense
	ExpenseErr error
	ProfileErr error
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
