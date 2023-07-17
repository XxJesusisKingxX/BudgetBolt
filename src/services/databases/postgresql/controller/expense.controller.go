package controller

import (
	"fmt"

	table "budgetbolt/src/services/databases/postgresql/model"
	q "budgetbolt/src/services/databases/postgresql/controller/querybuilder"
)

func CreateExpense(table table.Expense) error {
	query, err := q.BuildCreateQuery("expense", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func UpdateExpense(table table.Expense) error {
	query, err := q.BuildUpdateQuery("expense", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func RetrieveExpense(table table.Expense) error {
	query, err := q.BuildRetrieveQuery("expense", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func DeleteExpense(table table.Expense) error {
	query, err := q.BuildDeleteQuery("expense", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}