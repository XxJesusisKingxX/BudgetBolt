package controller

import (
	"fmt"

	table "budgetbolt/services/databases/postgresql/model"
)

func AddTransaction(tableName string, table table.Transaction) error {
	query, err := buildAddQuery(tableName, table)
	if err == nil {
		fmt.Print(query)
	}
	return err
}

func UpdateTransaction(tableName string, table table.Transaction) error {
	query, err := buildUpdateQuery(tableName, table)
	if err == nil {
		fmt.Print(query)
	}
	return err
}

func GetTransaction(tableName string, table table.Transaction) error {
	query, err := buildGetQuery(tableName, table)
	if err == nil {
		fmt.Print(query)
	}
	return err
}

func DeleteTransaction(tableName string, table table.Transaction) error {
	query, err := buildDeleteQuery(tableName, table)
	if err == nil {
		fmt.Print(query)
	}
	return err
}