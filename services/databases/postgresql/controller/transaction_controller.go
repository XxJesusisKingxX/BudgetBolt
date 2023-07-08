package controller

import (
	"fmt"

	table "budgetbolt/services/databases/postgresql/model"
	q "budgetbolt/services/databases/postgresql/controller/querybuilder"
)

func CreateTransaction(table table.Transaction) error {
	query, err := q.BuildCreateQuery("transaction", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func UpdateTransaction(table table.Transaction) error {
	query, err := q.BuildUpdateQuery("transaction", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func RetrieveTransaction(table table.Transaction) error {
	query, err := q.BuildRetrieveQuery("transaction", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}

func DeleteTransaction(table table.Transaction) error {
	query, err := q.BuildDeleteQuery("transaction", table)
	if err == nil {
		fmt.Println(query)
	}
	return err
}