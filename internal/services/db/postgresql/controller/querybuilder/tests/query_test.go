package main

import (
	"services/db/postgresql/model"
	"services/db/postgresql/controller/querybuilder"
	tests "services/utils/testing"
	"errors"
	"testing"
)

func TestFormatColumnsAndValues(t *testing.T) {
	type TestStruct struct{ Test string }
	type TestStructNoTag struct{ ID int }
	testOne := model.Budget{}
	testTwo := model.Transaction{ ID: "1", Vendor: "Test", IsRecurring: true, Amount: 123.45 }
	testThree := TestStructNoTag{ ID: 1 }
	testFour := TestStruct{ Test: "Test" }
	spent := 0.00
	testFive := model.Expense{ Spent: &spent}

	resOne, _ := querybuilder.FormatColumnsAndValues(testOne)
	resTwo, _ := querybuilder.FormatColumnsAndValues(testTwo)
	resThree, _ := querybuilder.FormatColumnsAndValues(testThree)
	resFive, _ := querybuilder.FormatColumnsAndValues(testFive)
	
	// Test an empty string returns
	tests.Equals(t, "", resOne)
	// Test if a string returns - int and string values
	tests.Equals(t, "transaction_id, net_amount, vendor, is_recurring", resTwo)
	// Test if we miss struct tag
	tests.Equals(t, "", resThree)
	// Test if a string returns - ptr values
	tests.Equals(t, "expense_spent", resFive)
	// Test if struct not passed as arg
	defer func() {
		if resFour := recover(); resFour != nil {
			tests.Equals(t, "Invalid parameter type", resFour)
		}
	}()
	querybuilder.FormatColumnsAndValues(testFour)
}
func TestSetColumnsAndValues(t *testing.T) {
	type TestStruct struct{ Test string }
	type TestStructNoTag struct{ ID string; Name string }
	testOne := model.Expense{}
	testTwo := model.Transaction{ ID: "1", Vendor: "Test", IsRecurring: true, Amount: 123.45}
	testThree := TestStructNoTag{ID: "1", Name: "Test"}
	testFour := TestStruct{Test: "Test"}
	spent := 0.00
	testFive := model.Expense{ Spent: &spent }

	resOne := querybuilder.SetColumnsAndValues(testOne)
	resTwo := querybuilder.SetColumnsAndValues(testTwo)
	resThree := querybuilder.SetColumnsAndValues(testThree)
	resFive := querybuilder.SetColumnsAndValues(testFive)
	
	// Test an empty string returns
	tests.Equals(t, "", resOne)
	// Test if only value defined set
	tests.Equals(t, "net_amount = 123.45, vendor = 'Test', is_recurring = true", resTwo)
	// Test if a pointer value is set
	tests.Equals(t, "expense_spent = 0", resFive)
	// Test if we miss struct tag
	tests.Equals(t, "", resThree)
	// Test if struct not passed as arg
	defer func() {
		if resFour := recover(); resFour != nil {
			tests.Equals(t, "Invalid parameter type", resFour)
		}
	}()
	querybuilder.SetColumnsAndValues(testFour)
}
func TestCreateWhereCondition(t *testing.T) {
	type TestStruct struct{ Test string }
	type TestStructNoTag struct{ ID string }
	testOne := model.Budget{}
	testTwo := model.Transaction{ ID: "1", Vendor: "Test", IsRecurring: true, Amount: 123.45}
	testThree := TestStruct{Test: "Test"}
	testFour := TestStructNoTag{ID: "1"}
	spent := 0.00
	testFive := model.Expense{ Spent: &spent}

	resOne := querybuilder.CreateWhereCondition(testOne)
	resTwo := querybuilder.CreateWhereCondition(testTwo)
	resFive := querybuilder.CreateWhereCondition(testFive)

	// Test an empty string returns
	tests.Equals(t, "", resOne)
	// Test if a string returns
	tests.Equals(t, "transaction_id = '1' AND net_amount = 123.45 AND vendor = 'Test' AND is_recurring = true", resTwo)
	// Test if a string returns
	tests.Equals(t, "expense_spent = 0", resFive)
	// Test if struct not passed as arg
	defer func() {
		if resThree := recover(); resThree != nil {
			tests.Equals(t, "Invalid parameter type", resThree)
		}
	}()
	querybuilder.CreateWhereCondition(testThree)
	// Test if we miss struct tag
	defer func() {
		if resFour := recover(); resFour != nil {
			tests.Equals(t, "Missing `db` tag", resFour)
		}
	}()
	querybuilder.CreateWhereCondition(testFour)
}

func TestBuildCreateQuery(t *testing.T) {
	type TestStruct struct{ Test string }
	err := errors.New("Empty model")
	testOne := model.Budget{}
	testTwo := model.Budget{Name: "Test"}
	testThree := TestStruct{Test: "Test"}

	_, resOne := querybuilder.BuildCreateQuery("budget", testOne)
	resTwo, _ := querybuilder.BuildCreateQuery("budget", testTwo)

	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne)
	// Test if columns and values are not empty
	tests.Equals(t, "INSERT INTO budget (budget_name) VALUES ('Test')", resTwo)
	// Test if pass wrong args
	defer func() {
		if resThree := recover(); resThree != nil {
			tests.Equals(t, "Invalid parameter type", resThree)
		}
	}()
	querybuilder.BuildCreateQuery("budget", testThree)
}
// func TestBuildUpdateQuery(t *testing.T) {
// 	type TestStruct struct{ Test string }
// 	err := errors.New("Empty model")
// 	testOne := model.Budget{}
// 	testTwo := model.Budget{Name: "Test"}
// 	testThree := TestStruct{Test: "Test"}

// 	_, resOne := querybuilder.BuildUpdateQuery("budget", testOne)
// 	resTwo, _ := querybuilder.BuildUpdateQuery("budget", testTwo)

// 	// Test if columns and values are empty
// 	tests.EqualsErr(t, err, resOne)
// 	// Test if columns and values are not empty
// 	tests.Equals(t, "UPDATE budget SET budget_name = 'Test' WHERE transaction_id=0", resTwo)
// 	// Test if pass wrong args
// 	defer func() {
// 		if resThree := recover(); resThree != nil {
// 			tests.Equals(t, "Invalid parameter type", resThree)
// 		}
// 	}()
// 	querybuilder.BuildUpdateQuery("budget", testThree)
// }
func TestBuildRetrieveQuery(t *testing.T) {
	type TestStruct struct{ Test string }
	err := errors.New("Empty model")
	testOne := model.Budget{}
	testTwo := model.Budget{Name: "Test"}
	testThree := TestStruct{Test: "Test"}

	_, resOne := querybuilder.BuildRetrieveQuery("budget", testOne)
	resTwo, _ := querybuilder.BuildRetrieveQuery("budget", testTwo)

	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne)
	// Test if columns and values are not empty
	tests.Equals(t, "SELECT * FROM budget WHERE budget_name = 'Test'", resTwo)
	// Test if pass wrong args
	defer func() {
		if resThree := recover(); resThree != nil {
			tests.Equals(t, "Invalid parameter type", resThree)
		}
	}()
	querybuilder.BuildRetrieveQuery("budget", testThree)
}
func TestBuildDeleteQuery(t *testing.T) {
	type TestStruct struct{ Test string }
	err := errors.New("Empty model")
	testOne := model.Budget{}
	testTwo := model.Budget{Name: "Test"}
	testThree := TestStruct{Test: "Test"}

	_, resOne := querybuilder.BuildDeleteQuery("budget", testOne)
	resTwo, _ := querybuilder.BuildDeleteQuery("budget", testTwo)

	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne)
	// Test if columns and values are not empty
	tests.Equals(t, "DELETE FROM budget WHERE budget_name = 'Test'", resTwo)
	// Test if pass wrong args
	defer func() {
		if resThree := recover(); resThree != nil {
			tests.Equals(t, "Invalid parameter type", resThree)
		}
	}()
	querybuilder.BuildDeleteQuery("budget", testThree)
}
func TestTransactionBuildQuery(t *testing.T) {
	type TestStruct struct{ Test string }
	err := errors.New("Empty model")
	testOne := model.Budget{}
	testTwo := model.Budget{Name: "Test"}
	testThree := TestStruct{Test: "Test"}

	_, resOne := querybuilder.BuildDeleteQuery("budget", testOne)
	resTwo, _ := querybuilder.BuildDeleteQuery("budget", testTwo)

	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne)
	// Test if columns and values are not empty
	tests.Equals(t, "DELETE FROM budget WHERE budget_name = 'Test'", resTwo)
	// Test if pass wrong args
	defer func() {
		if resThree := recover(); resThree != nil {
			tests.Equals(t, "Invalid parameter type", resThree)
		}
	}()
	querybuilder.BuildDeleteQuery("budget", testThree)
}
func TestTransactionBuildCreateQuery(t *testing.T) {
	err := errors.New("Empty model")
	testOne := model.Transaction{}
	testTwo := model.Transaction{From: "Test"}

	_, resOne := querybuilder.BuildTransactionCreateQuery(testOne)
	resTwo, _ := querybuilder.BuildTransactionCreateQuery(testTwo)

	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne)
	// Test if columns and values are not empty
	tests.Equals(t, "INSERT INTO transaction (payment_account_from_to) VALUES ('Test')", resTwo)
}
func TestTransactionBuildUpdateQuery(t *testing.T) {
	err := errors.New("Empty model")
	testOne := model.Transaction{}
	testTwo := model.Transaction{ID: "1", From: "Test"}

	_, resOne := querybuilder.BuildTransactionUpdateQuery(testOne)
	resTwo, _ := querybuilder.BuildTransactionUpdateQuery(testTwo)

	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne)
	// Test if columns and values are not empty
	tests.Equals(t, "UPDATE transaction SET payment_account_from_to = 'Test' WHERE transaction_id='1'", resTwo)
}
func TestTransactionBuildRetrieveQuery(t *testing.T) {
	err := errors.New("Empty model")
	testOne := model.Transaction{}
	testTwo := model.Transaction{From: "Test"}
	testThree := model.Transaction{From: "Test", Query: model.Querys{ Select: model.QueryParameters{ OrderBy: model.OrderBy{ Asc: true, Column: "transaction_date"}}}}
	testFour := model.Transaction{From: "Test", Query: model.Querys{ Select: model.QueryParameters{ OrderBy: model.OrderBy{ Desc: true, Column: "transaction_date"}}}}
	testFive := model.Transaction{From: "Test", Query: model.Querys{ Select: model.QueryParameters{ GreaterThanEq: model.GreaterThanEq{ Value: "2023-08-31", Column: "transaction_date"}}}}

	_, resOne := querybuilder.BuildTransactionRetrieveQuery(testOne)
	resTwo, _ := querybuilder.BuildTransactionRetrieveQuery(testTwo)
	resThree, _ := querybuilder.BuildTransactionRetrieveQuery(testThree)
	resFour, _ := querybuilder.BuildTransactionRetrieveQuery(testFour)
	resFive, _ := querybuilder.BuildTransactionRetrieveQuery(testFive)

	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne)
	// Test if columns and values are not empty
	tests.Equals(t, "SELECT * FROM transaction WHERE payment_account_from_to = 'Test'", resTwo)
	// Test if ascending order is selected
	tests.Equals(t, "SELECT * FROM transaction WHERE payment_account_from_to = 'Test' ORDER BY transaction_date ASC", resThree)
	// Test if descding order is selected
	tests.Equals(t, "SELECT * FROM transaction WHERE payment_account_from_to = 'Test' ORDER BY transaction_date DESC", resFour)
	// Test if greater than 
	tests.Equals(t, "SELECT * FROM transaction WHERE payment_account_from_to = 'Test' AND transaction_date > '2023-08-31'", resFive)
}
func TestTransactionBuildDeleteQuery(t *testing.T) {
	err := errors.New("Empty model")
	testOne := model.Transaction{}
	testTwo := model.Transaction{From: "Test"}

	_, resOne := querybuilder.BuildTransactionDeleteQuery(testOne)
	resTwo, _ := querybuilder.BuildTransactionDeleteQuery(testTwo)

	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne)
	// Test if columns and values are not empty
	tests.Equals(t, "DELETE FROM transaction WHERE payment_account_from_to = 'Test'", resTwo)
}