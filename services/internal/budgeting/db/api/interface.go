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
	CreateIncome(db *sql.DB, m model.Income) error
	UpdateIncome(db *sql.DB, setM model.Income, whereM model.Income) error
	RetrieveIncome(db *sql.DB, m model.Income) ([]model.Income, error)

}

type DB struct{}
type MockDB struct {
	Profile user.Profile
	Expense []model.Expense
	Income []model.Income
	ExpenseErr error
	ProfileErr error
	IncomeErr error
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
func (t MockDB) CreateIncome(db *sql.DB, m model.Income) error {
	return t.IncomeErr
}
func (t MockDB) RetrieveIncome(db *sql.DB, m model.Income) ([]model.Income, error) {
	return t.Income, t.IncomeErr
}
func (t MockDB) UpdateIncome(db *sql.DB, setM model.Income, whereM model.Income) error {
	return t.IncomeErr
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
func (t DB) CreateIncome(db *sql.DB, m model.Income) error {
	err := controller.CreateIncome(db, m)
	return err
}
func (t DB) RetrieveIncome(db *sql.DB, m model.Income) ([]model.Income, error) {
	incomes, err := controller.RetrieveIncome(db, m)
	return incomes, err
}
func (t DB) UpdateIncome(db *sql.DB, setM model.Income, whereM model.Income) error {
	err := controller.UpdateIncome(db, setM, whereM)
	return err
}
