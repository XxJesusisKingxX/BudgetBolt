package response

import (
	"database/sql"

	controller "budgetbolt/services/databases/postgresql/controller"
	model "budgetbolt/services/databases/postgresql/model"

	plaid "github.com/plaid/plaid-go/v12/plaid"
)

func ParseAccountsToDB(db *sql.DB, accessToken string, accounts []plaid.AccountBase){
	for _, v := range accounts {
		acc := model.Account {
			PlaidID: v.AccountId,
			Name: v.Name,
			Balance: v.Balances.GetAvailable(),
			Token: accessToken,
		}
		controller.CreateAccount(db, acc)
	}
}
func ParseTransactionsToDB(db *sql.DB, transactions []plaid.Transaction){
	for _, v := range transactions {
		trans := model.Transaction {
			From: v.AccountId,
			Vendor: v.GetMerchantName(),
			Amount: v.Amount,
			Date: v.Date,
			Description: v.Name,
		}
		controller.CreateTransaction(db, trans)
	}
}
func ParseInvestmentsToDB(db *sql.DB, transactions []plaid.InvestmentTransaction){
	for _, v := range transactions {
		trans := model.Investment {
			PlaidId: v.AccountId,
			Fees: v.GetFees(),
			Amount: v.Amount,
			Date: v.Date,
			Description: v.Name,
			Price: v.Price,
			Quantity: v.Quantity,
		}
		controller.CreateInvestment(db, trans)
	}
}
func ParseHoldingsToDB(db *sql.DB, holdings []plaid.Holding){
	for _, v := range holdings {
		holdings := model.Holding {
			PlaidId: v.AccountId,
			CostBasis: v.GetCostBasis(),
			LastPrice: v.InstitutionPrice,
			TotalValue: v.InstitutionValue,
			Quantity: v.Quantity,
		}
		controller.CreateHolding(db, holdings)
	}
}