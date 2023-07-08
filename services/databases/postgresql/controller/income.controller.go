package controller

import (
	"fmt"

	table "budgetbolt/services/databases/postgresql/model"
	q "budgetbolt/services/databases/postgresql/controller/querybuilder"
)

func CreateIncome(table table.Income) error {
	query, err := q.BuildCreateQuery("income", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func UpdateIncome(table table.Income) error {
	query, err := q.BuildUpdateQuery("income", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func RetrieveIncome(table table.Income) error {
	query, err := q.BuildRetrieveQuery("income", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func DeleteIncome(table table.Income) error {
	query, err := q.BuildDeleteQuery("income", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}