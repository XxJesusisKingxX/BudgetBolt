package controller

import (
	"fmt"

	table "budgetbolt/services/databases/postgresql/model"
	q "budgetbolt/services/databases/postgresql/controller/querybuilder"
)

func CreateBudget(table table.Budget) error {
	query, err := q.BuildCreateQuery("budget", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func UpdateBudget(table table.Budget) error {
	query, err := q.BuildUpdateQuery("budget", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func RetrieveBudget(table table.Budget) error {
	query, err := q.BuildRetrieveQuery("budget", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func DeleteBudget(table table.Budget) error {
	query, err := q.BuildDeleteQuery("budget", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}