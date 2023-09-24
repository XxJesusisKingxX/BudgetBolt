package main

import (
	"errors"
	"testing"
	
	budget "services/internal/budgeting/db/model"
	"services/internal/utils/sql/querybuilder"
	transaction "services/internal/transaction_history/db/model"
	tests "services/internal/utils/testing"
)

func TestFormatColumnsAndValues(t *testing.T) {
	t.Run("Empty Model", func(t *testing.T) {
		testModel := budget.Expense{}
		result, _ := querybuilder.FormatColumnsAndValues(testModel)
		tests.Equals(t, "", result)
	})

	t.Run("Non-Empty Model with Tags", func(t *testing.T) {
		testModel := transaction.Transaction{ID: "1", Vendor: "Test", IsRecurring: true, Amount: 123.45}
		expectedColumns := "transaction_id, net_amount, vendor, is_recurring"
		result, _ := querybuilder.FormatColumnsAndValues(testModel)
		tests.Equals(t, expectedColumns, result)
	})

	t.Run("Model Missing `db` Tag", func(t *testing.T) {
		testModel := struct{ ID int }{ID: 1}
		result, _ := querybuilder.FormatColumnsAndValues(testModel)
		tests.Equals(t, "", result)
	})

	t.Run("Invalid Parameter Type", func(t *testing.T) {
		testModel := struct{ Test string }{Test: "Test"}
		defer func() {
			if result := recover(); result != nil {
				tests.Equals(t, "Invalid parameter type", result)
			}
		}()
		querybuilder.FormatColumnsAndValues(testModel)

	})
}

func TestSetColumnsAndValues(t *testing.T) {
	t.Run("Empty Model", func(t *testing.T) {
		testModel := budget.Expense{}
		result := querybuilder.SetColumnsAndValues(testModel)
		tests.Equals(t, "", result)
	})

	t.Run("Non-Empty Model with Tags", func(t *testing.T) {
		testModel := transaction.Transaction{ID: "1", Vendor: "Test", IsRecurring: true, Amount: 123.45}
		expectedQuery := "net_amount = 123.45, vendor = 'Test', is_recurring = true"
		result := querybuilder.SetColumnsAndValues(testModel)
		tests.Equals(t, expectedQuery, result)
	})

	t.Run("Model Missing `db` Tag", func(t *testing.T) {
		testModel := struct{ ID string; Name string }{ID: "1", Name: "Test"}
		result := querybuilder.SetColumnsAndValues(testModel)
		tests.Equals(t, "", result)
	})

	t.Run("Invalid Parameter Type", func(t *testing.T) {
		testModel := struct{ Test string }{Test: "Test"}
		defer func() {
			if result := recover(); result != nil {
				tests.Equals(t, "Invalid parameter type", result)
			}
		}()
		querybuilder.SetColumnsAndValues(testModel)
	})
}

func TestCreateWhereCondition(t *testing.T) {
	t.Run("Empty Model", func(t *testing.T) {
		testModel := budget.Expense{}
		result := querybuilder.CreateWhereCondition(testModel)
		tests.Equals(t, "", result)
	})

	t.Run("Non-Empty Model with Tags", func(t *testing.T) {
		testModel := transaction.Transaction{ID: "1", Vendor: "Test", IsRecurring: true, Amount: 123.45}
		expectedQuery := "transaction_id = '1' AND net_amount = 123.45 AND vendor = 'Test' AND is_recurring = true"
		result := querybuilder.CreateWhereCondition(testModel)
		tests.Equals(t, expectedQuery, result)
	})

	t.Run("Invalid Parameter Type", func(t *testing.T) {
		testModel := struct{ Test string }{Test: "Test"}
		defer func() {
			if result := recover(); result != nil {
				tests.Equals(t, "Invalid parameter type", result)
			}
		}()
		querybuilder.CreateWhereCondition(testModel)
	})

	t.Run("Missing `db` Tag", func(t *testing.T) {
		testModel := struct{ ID string }{ID: "1"}
		defer func() {
			if result := recover(); result != nil {
				tests.Equals(t, "Missing `db` tag", result)
			}
		}()
		querybuilder.CreateWhereCondition(testModel)
	})
}

func TestBuildCreateQuery(t *testing.T) {
	t.Run("Empty Model", func(t *testing.T) {
		err := errors.New("Empty model")
		testModel := budget.Expense{}
		_, result := querybuilder.BuildCreateQuery("expenses", testModel)
		tests.EqualsErr(t, err, result)
	})

	t.Run("Non-Empty Model", func(t *testing.T) {
		testModel := budget.Expense{Name: "Test"}
		expectedQuery := "INSERT INTO expenses (expense_name) VALUES ('Test')"
		result, _ := querybuilder.BuildCreateQuery("expenses", testModel)
		tests.Equals(t, expectedQuery, result)
	})
}

func TestBuildUpdateQuery(t *testing.T) {

	t.Run("Valid input with ID", func(t *testing.T) {
		setModel := budget.Expense{Name: "Test"}
		whereModel := budget.Expense{ID: 1}
		expectedQuery := "UPDATE tablename SET expense_name = 'Test' WHERE expense_id = '1'"
		query, _ := querybuilder.BuildUpdateQuery("tablename", setModel, whereModel)

		tests.Equals(t, expectedQuery, query)
	})

	t.Run("Empty model", func(t *testing.T) {
		setModel := budget.Expense{}
		whereModel := budget.Expense{}
		expectedErr := errors.New("Empty model")
		expectedQuery := ""
		query, err := querybuilder.BuildUpdateQuery("tablename", setModel, whereModel)

		tests.Equals(t, expectedErr.Error(), err.Error())
		tests.Equals(t, expectedQuery, query)
	})
}

func TestBuildRetrieveQuery(t *testing.T) {
	t.Run("Empty Model", func(t *testing.T) {
		err := errors.New("Empty model")
		testModel := budget.Expense{}
		_, result := querybuilder.BuildRetrieveQuery("expenses", testModel)
		tests.EqualsErr(t, err, result)
	})

	t.Run("Non-Empty Model", func(t *testing.T) {
		testModel := budget.Expense{Name: "Test"}
		expectedQuery := "SELECT * FROM expenses WHERE expense_name = 'Test'"
		result, _ := querybuilder.BuildRetrieveQuery("expenses", testModel)
		tests.Equals(t, expectedQuery, result)
	})
}
func TestBuildDeleteQuery(t *testing.T) {
	t.Run("Empty Model", func(t *testing.T) {
		err := errors.New("Empty model")
		testModel := budget.Expense{}
		_, result := querybuilder.BuildDeleteQuery("expenses", testModel)
		tests.EqualsErr(t, err, result)
	})

	t.Run("Non-Empty Model", func(t *testing.T) {
		testModel := budget.Expense{Name: "Test"}
		expectedQuery := "DELETE FROM expenses WHERE expense_name = 'Test'"
		result, _ := querybuilder.BuildDeleteQuery("expenses", testModel)
		tests.Equals(t, expectedQuery, result)
	})
}
