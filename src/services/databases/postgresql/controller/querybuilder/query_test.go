package querybuilder

import (
	"budgetbolt/src/services/databases/postgresql/model"
	tests "budgetbolt/src/services/tests"
	"errors"
	"testing"
)

func TestFormatColumnsAndValues(t *testing.T) {
	type TestStruct struct{ Test string }
	type TestStructNoTag struct{ ID int }
	testOne := model.Budget{}
	testTwo := model.Transaction{ ID: 1, Vendor: "Test", IsRecurring: true, Amount: 123.45 }
	testThree := TestStructNoTag{ ID: 1 }
	testFour := TestStruct{ Test: "Test" }

	resOne, _ := formatColumnsAndValues(testOne)
	resTwo, _ := formatColumnsAndValues(testTwo)
	resThree, _ := formatColumnsAndValues(testThree)
	
	// Test an empty string returns
	tests.Equals(t, "", resOne)
	// Test if a string returns - int and string values
	tests.Equals(t, "transaction_id, net_amount, vendor, is_recurring", resTwo)
	// Test if we miss struct tag
	tests.Equals(t, "", resThree)
	// Test if struct not passed as arg
	defer func() {
		if resFour := recover(); resFour != nil {
			tests.Equals(t, "Invalid parameter type", resFour)
		}
		}()
		formatColumnsAndValues(testFour)
	}
func TestSetColumnsAndValues(t *testing.T) {
	type TestStruct struct{ Test string }
	type TestStructNoTag struct{ ID int; Name string }
	testOne := model.Budget{}
	testTwo := model.Transaction{ ID: 1, Vendor: "Test", IsRecurring: true, Amount: 123.45}
	testThree := TestStructNoTag{ID: 1, Name: "Test"}
	testFour := TestStruct{Test: "Test"}

	resOne := setColumnsAndValues(testOne)
	resTwo := setColumnsAndValues(testTwo)
	resThree := setColumnsAndValues(testThree)
	
	// Test an empty string returns
	tests.Equals(t, "", resOne)
	// Test if a string returns
	tests.Equals(t, "net_amount = 123.45, vendor = 'Test', is_recurring = true", resTwo)
	// Test if we miss struct tag
	tests.Equals(t, "", resThree)
	// Test if struct not passed as arg
	defer func() {
		if resFour := recover(); resFour != nil {
			tests.Equals(t, "Invalid parameter type", resFour)
		}
	}()
	setColumnsAndValues(testFour)
}
func TestCreateWhereCondition(t *testing.T) {
	type TestStruct struct{ Test string }
	type TestStructNoTag struct{ ID int }
	testOne := model.Budget{}
	testTwo := model.Transaction{ ID: 1, Vendor: "Test", IsRecurring: true, Amount: 123.45}
	testThree := TestStruct{Test: "Test"}
	testFour := TestStructNoTag{ID: 1}

	resOne := createWhereCondition(testOne)
	resTwo := createWhereCondition(testTwo)

	// Test an empty string returns
	tests.Equals(t, "", resOne)
	// Test if a string returns
	tests.Equals(t, "transaction_id = 1 AND net_amount = 123.45 AND vendor = 'Test' AND is_recurring = true", resTwo)
	// Test if struct not passed as arg
	defer func() {
		if resThree := recover(); resThree != nil {
			tests.Equals(t, "Invalid parameter type", resThree)
		}
	}()
	createWhereCondition(testThree)
	// Test if we miss struct tag
	defer func() {
		if resFour := recover(); resFour != nil {
			tests.Equals(t, "Missing `db` tag", resFour)
		}
	}()
	createWhereCondition(testFour)
}

func TestBuildCreateQuery(t *testing.T) {
	type TestStruct struct{ Test string }
	err := errors.New("Empty model")
	testOne := model.Budget{}
	testTwo := model.Budget{Name: "Test"}
	testThree := TestStruct{Test: "Test"}

	_, resOne := BuildCreateQuery("budget", testOne)
	resTwo, _ := BuildCreateQuery("budget", testTwo)

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
	BuildCreateQuery("budget", testThree)
}
func TestBuildUpdateQuery(t *testing.T) {
	type TestStruct struct{ Test string }
	err := errors.New("Empty model")
	testOne := model.Budget{}
	testTwo := model.Budget{Name: "Test"}
	testThree := TestStruct{Test: "Test"}

	_, resOne := BuildUpdateQuery("budget", testOne)
	resTwo, _ := BuildUpdateQuery("budget", testTwo)

	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne)
	// Test if columns and values are not empty
	tests.Equals(t, "UPDATE budget SET budget_name = 'Test' WHERE transaction_id=0", resTwo)
	// Test if pass wrong args
	defer func() {
		if resThree := recover(); resThree != nil {
			tests.Equals(t, "Invalid parameter type", resThree)
		}
	}()
	BuildUpdateQuery("budget", testThree)
}
func TestBuildRetrieveQuery(t *testing.T) {
	type TestStruct struct{ Test string }
	err := errors.New("Empty model")
	testOne := model.Budget{}
	testTwo := model.Budget{Name: "Test"}
	testThree := TestStruct{Test: "Test"}

	_, resOne := BuildRetrieveQuery("budget", testOne)
	resTwo, _ := BuildRetrieveQuery("budget", testTwo)

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
	BuildRetrieveQuery("budget", testThree)
}
func TestBuildDeleteQuery(t *testing.T) {
	type TestStruct struct{ Test string }
	err := errors.New("Empty model")
	testOne := model.Budget{}
	testTwo := model.Budget{Name: "Test"}
	testThree := TestStruct{Test: "Test"}

	_, resOne := BuildDeleteQuery("budget", testOne)
	resTwo, _ := BuildDeleteQuery("budget", testTwo)

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
	BuildDeleteQuery("budget", testThree)
}
func TestTransactionBuildQuery(t *testing.T) {
	type TestStruct struct{ Test string }
	err := errors.New("Empty model")
	testOne := model.Budget{}
	testTwo := model.Budget{Name: "Test"}
	testThree := TestStruct{Test: "Test"}

	_, resOne := BuildDeleteQuery("budget", testOne)
	resTwo, _ := BuildDeleteQuery("budget", testTwo)

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
	BuildDeleteQuery("budget", testThree)
}
func TestTransactionBuildCreateQuery(t *testing.T) {
	err := errors.New("Empty model")
	testOne := model.Transaction{}
	testTwo := model.Transaction{From: "Test"}

	_, resOne := BuildTransactionCreateQuery(testOne)
	resTwo, _ := BuildTransactionCreateQuery(testTwo)

	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne)
	// Test if columns and values are not empty
	tests.Equals(t, "INSERT INTO transaction (payment_account_from_to) VALUES ('Test')", resTwo)
}
func TestTransactionBuildUpdateQuery(t *testing.T) {
	err := errors.New("Empty model")
	testOne := model.Transaction{}
	testTwo := model.Transaction{From: "Test"}

	_, resOne := BuildTransactionUpdateQuery(testOne)
	resTwo, _ := BuildTransactionUpdateQuery(testTwo)

	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne)
	// Test if columns and values are not empty
	tests.Equals(t, "UPDATE transaction SET payment_account_from_to = 'Test' WHERE transaction_id=0", resTwo)
}
func TestTransactionBuildRetrieveQuery(t *testing.T) {
	err := errors.New("Empty model")
	testOne := model.Transaction{}
	testTwo := model.Transaction{From: "Test"}
	testThree := model.Transaction{From: "Test", Query: model.Querys{ Select: model.QueryParameters{ Asc: true, OrderBy: "transaction_date"}}}
	testFour := model.Transaction{From: "Test", Query: model.Querys{ Select: model.QueryParameters{ Desc: true, OrderBy: "transaction_date"}}}

	_, resOne := BuildTransactionRetrieveQuery(testOne)
	resTwo, _ := BuildTransactionRetrieveQuery(testTwo)
	resThree, _ := BuildTransactionRetrieveQuery(testThree)
	resFour, _ := BuildTransactionRetrieveQuery(testFour)

	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne)
	// Test if columns and values are not empty
	tests.Equals(t, "SELECT * FROM transaction WHERE payment_account_from_to = 'Test'", resTwo)
	// Test if ascending order is selected
	tests.Equals(t, "SELECT * FROM transaction WHERE payment_account_from_to = 'Test' ORDER BY transaction_date ASC", resThree)
	// Test if descding order is selected
	tests.Equals(t, "SELECT * FROM transaction WHERE payment_account_from_to = 'Test' ORDER BY transaction_date DESC", resFour)
}
func TestTransactionBuildDeleteQuery(t *testing.T) {
	err := errors.New("Empty model")
	testOne := model.Transaction{}
	testTwo := model.Transaction{From: "Test"}

	_, resOne := BuildTransactionDeleteQuery(testOne)
	resTwo, _ := BuildTransactionDeleteQuery(testTwo)

	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne)
	// Test if columns and values are not empty
	tests.Equals(t, "DELETE FROM transaction WHERE payment_account_from_to = 'Test'", resTwo)
}