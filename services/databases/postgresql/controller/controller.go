package controller

import (
	"fmt"

	table "budgetbolt/services/databases/postgresql/model"
)

func AddTransaction(model table.Transaction) {
	query := buildTransactionInsertQuery(model)
	fmt.Println(query)
}

func UpdateTransaction(model table.Transaction) {
	query := buildTransactionUpdateQuery(model)
	fmt.Println(query)
}

func GetTransaction(model table.Transaction) {
	query := buildTransactionGetQuery(model)
	fmt.Println(query)
}