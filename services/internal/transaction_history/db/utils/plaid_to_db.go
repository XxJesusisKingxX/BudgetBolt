package utils

import (
	ctrl "services/internal/transaction_history/db/controller"
	"services/internal/transaction_history/db/model"
	
	"database/sql"
	"github.com/plaid/plaid-go/v12/plaid"
)

func TransactionsToDB(db *sql.DB, profileId int, transactions []plaid.Transaction){
	for _, v := range transactions {
    	acc, _ := ctrl.RetrieveAccount(db, model.Account{ ID: v.AccountId })
		trans := model.Transaction {
			ID: v.TransactionId,
			Vendor: v.GetMerchantName(),
			Method: v.PaymentChannel,
			Amount: v.Amount,
			Date: v.Date,
			Description: v.Name,
			ProfileID: profileId,
			PrimaryCategory: v.PersonalFinanceCategory.Get().Primary,
			DetailCategory: v.PersonalFinanceCategory.Get().Detailed,
			AccountName: acc[0].Name,
		}
		ctrl.CreateTransaction(db, trans)
	}
}

func AccountsToDB(db *sql.DB, profileId int, accounts []plaid.AccountBase){
	for _, v := range accounts {
		acc := model.Account {
			ID: v.AccountId,
			Name: v.Name,
			Balance: v.Balances.GetAvailable(),
			ProfileID: profileId,
		}
		ctrl.CreateAccount(db, acc)
	}
}