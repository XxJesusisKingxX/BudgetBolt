package querybuilder

import (
	table "budgetbolt/services/databases/postgresql/model"
	tests "budgetbolt/tests"
	"errors"
	"fmt"
	"testing"
)

func TestFormatColumnsAndValues(t *testing.T) {
	type TestStruct struct{ Test string }
	type TestStructNoTag struct{ ID int }
	testOne := table.Budget{}
	testTwo := table.Budget{ ID: 2, Name: "Test" }
	testThree := TestStruct{ Test: "Test"}
	testFour := TestStructNoTag{ ID: 1 }

	resOne, _ := formatColumnsAndValues(testOne)
	resTwo, _ := formatColumnsAndValues(testTwo)

	// Test an empty string returns
	tests.Equals(t, "", resOne, fmt.Sprintf("Expected %v but got %v", "", resOne))
	// Test if a string returns
	tests.Equals(t, "budget_id, budget_name", resTwo, fmt.Sprintf("Expected %v but got %v", "budget_id, budget_name", resTwo))
	// Test if struct not passed as arg
	defer func() {
		if resThree := recover(); resThree != nil {
			tests.Equals(t, "Invalid parameter type", resThree, fmt.Sprintf("Expected %v but got %v", "Invalid parameter type", resThree))
		}
	}()
	formatColumnsAndValues(testThree)
	// Test if we miss struct tag
	defer func() {
		if resFour := recover(); resFour != nil {
			tests.Equals(t, "Missing `db` tag", resFour, fmt.Sprintf("Expected %v but got %v", "Missing `db` tag", resFour))
		}
	}()
	formatColumnsAndValues(testFour)
	
}
func TestSetColumnsAndValues(t *testing.T) {
	type TestStruct struct{ Test string }
	type TestStructNoTag struct{ ID int }
	testOne := table.Budget{}
	testTwo := table.Budget{ ID: 2, Name: "Test" }
	testThree := TestStruct{ Test: "Test"}
	testFour := TestStructNoTag{ ID: 1 }

	resOne := setColumnsAndValues(testOne)
	resTwo := setColumnsAndValues(testTwo)

	// Test an empty string returns
	tests.Equals(t, "", resOne, fmt.Sprintf("Expected %v but got %v", "", resOne))
	// Test if a string returns
	tests.Equals(t, "budget_name = 'Test'", resTwo, fmt.Sprintf("Expected %v but got %v", "budget_name = 'Test'", resTwo))
	// Test if struct not passed as arg
	defer func() {
		if resThree := recover(); resThree != nil {
			tests.Equals(t, "Invalid parameter type", resThree, fmt.Sprintf("Expected %v but got %v", "Invalid parameter type", resThree))
		}
	}()
	setColumnsAndValues(testThree)
	// Test if we miss struct tag
	defer func() {
		if resFour := recover(); resFour != nil {
			tests.Equals(t, "Missing `db` tag", resFour, fmt.Sprintf("Expected %v but got %v", "Missing `db` tag", resFour))
		}
	}()
	setColumnsAndValues(testFour)
}
func TestCreateWhereCondition(t *testing.T) {
	type TestStruct struct{ Test string }
	type TestStructNoTag struct{ ID int }
	testOne := table.Budget{}
	testTwo := table.Budget{ ID: 2, Name: "Test" }
	testThree := TestStruct{ Test: "Test"}
	testFour := TestStructNoTag{ ID: 1 }

	resOne := createWhereCondition(testOne)
	resTwo := createWhereCondition(testTwo)

	// Test an empty string returns
	tests.Equals(t, "", resOne, fmt.Sprintf("Expected %v but got %v", "", resOne))
	// Test if a string returns
	tests.Equals(t, "budget_id = 2 AND budget_name = 'Test'", resTwo, fmt.Sprintf("Expected %v but got %v", "budget_id = 2 AND budget_name = 'Test'", resTwo))
	// Test if struct not passed as arg
	defer func() {
		if resThree := recover(); resThree != nil {
			tests.Equals(t, "Invalid parameter type", resThree, fmt.Sprintf("Expected %v but got %v", "Invalid parameter type", resThree))
		}
	}()
	createWhereCondition(testThree)
	// Test if we miss struct tag
	defer func() {
		if resFour := recover(); resFour != nil {
			tests.Equals(t, "Missing `db` tag", resFour, fmt.Sprintf("Expected %v but got %v", "Missing `db` tag", resFour))
		}
	}()
	createWhereCondition(testFour)
}

func TestBuildCreateQuery(t *testing.T) {
	type TestStruct struct{ Test string }
	err := errors.New("Empty model")
	testOne := table.Budget{}
	testTwo := table.Budget{ Name: "Test" }
	testThree := TestStruct{ Test: "Test"}

	_, resOne := BuildCreateQuery("budget", testOne)
	resTwo, _ := BuildCreateQuery("budget", testTwo)


	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne, fmt.Sprintf("Expected %v but got %v", err, resOne))
	// Test if columns and values are not empty
	tests.Equals(t, "INSERT INTO budget (budget_name) VALUES ('Test')", resTwo, fmt.Sprintf("Expected %v but got %v", "INSERT INTO budget (budget_name) VALUES ('Test')", resTwo))
	// Test if pass wrong args
	defer func() {
		if resThree := recover(); resThree != nil {
			tests.Equals(t, "Invalid parameter type", resThree, fmt.Sprintf("Expected %v but got %v", "Invalid parameter type", resThree))
		}
	}()
	BuildCreateQuery("budget", testThree)
}
func TestBuildUpdateQuery(t *testing.T) {
	type TestStruct struct{ Test string }
	err := errors.New("Empty model")
	testOne := table.Budget{}
	testTwo := table.Budget{ Name: "Test" }
	testThree := TestStruct{ Test: "Test"}

	_, resOne := BuildUpdateQuery("budget", testOne)
	resTwo, _ := BuildUpdateQuery("budget", testTwo)


	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne, fmt.Sprintf("Expected %v but got %v", err, resOne))
	// Test if columns and values are not empty
	tests.Equals(t, "UPDATE transaction SET budget_name = 'Test' WHERE transaction_id=0", resTwo, fmt.Sprintf("Expected %v but got %v", "UPDATE transaction SET budget_name = 'Test' WHERE transaction_id=0", resTwo))
	// Test if pass wrong args
	defer func() {
		if resThree := recover(); resThree != nil {
			tests.Equals(t, "Invalid parameter type", resThree, fmt.Sprintf("Expected %v but got %v", "Invalid parameter type", resThree))
		}
	}()
	BuildUpdateQuery("budget", testThree)
}
func TestBuildRetrieveQuery(t *testing.T) {
	type TestStruct struct{ Test string }
	err := errors.New("Empty model")
	testOne := table.Budget{}
	testTwo := table.Budget{ Name: "Test" }
	testThree := TestStruct{ Test: "Test"}

	_, resOne := BuildRetrieveQuery("budget", testOne)
	resTwo, _ := BuildRetrieveQuery("budget", testTwo)


	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne, fmt.Sprintf("Expected %v but got %v", err, resOne))
	// Test if columns and values are not empty
	tests.Equals(t, "SELECT * FROM transaction WHERE budget_name = 'Test'", resTwo, fmt.Sprintf("Expected %v but got %v", "SELECT * FROM transaction WHERE budget_name = 'Test'", resTwo))
	// Test if pass wrong args
	defer func() {
		if resThree := recover(); resThree != nil {
			tests.Equals(t, "Invalid parameter type", resThree, fmt.Sprintf("Expected %v but got %v", "Invalid parameter type", resThree))
		}
	}()
	BuildRetrieveQuery("budget", testThree)
}
func TestBuildDeleteQuery(t *testing.T) {
	type TestStruct struct{ Test string }
	err := errors.New("Empty model")
	testOne := table.Budget{}
	testTwo := table.Budget{ Name: "Test" }
	testThree := TestStruct{ Test: "Test"}

	_, resOne := BuildDeleteQuery("budget", testOne)
	resTwo, _ := BuildDeleteQuery("budget", testTwo)


	// Test if columns and values are empty
	tests.EqualsErr(t, err, resOne, fmt.Sprintf("Expected %v but got %v", err, resOne))
	// Test if columns and values are not empty
	tests.Equals(t, "DELETE FROM transaction WHERE budget_name = 'Test'", resTwo, fmt.Sprintf("Expected %v but got %v", "DELETE FROM transaction WHERE budget_name = 'Test'", resTwo))
	// Test if pass wrong args
	defer func() {
		if resThree := recover(); resThree != nil {
			tests.Equals(t, "Invalid parameter type", resThree, fmt.Sprintf("Expected %v but got %v", "Invalid parameter type", resThree))
		}
	}()
	BuildDeleteQuery("budget", testThree)
}