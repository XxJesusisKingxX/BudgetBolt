package response

import (
	"database/sql"

	"github.com/plaid/plaid-go/v12/plaid"

	"budgetbolt/src/services/databases/postgresql/controller"
	"budgetbolt/src/services/databases/postgresql/model"
)

func ParseAccountsToDB(db *sql.DB, profileId int, accounts []plaid.AccountBase){
	for _, v := range accounts {
		acc := model.Account {
			Name: v.Name,
			Balance: v.Balances.GetAvailable(),
			PlaidAccountID: v.AccountId,
			ProfileID: profileId,
		}
		controller.CreateAccount(db, acc)
	}
}
func ParseTransactionsToDB(db *sql.DB, profileId int, transactions []plaid.Transaction){
	for _, v := range transactions {
    	acc, _ := controller.RetrieveAccount(db, model.Account{ PlaidAccountID: v.AccountId })
		trans := model.Transaction {
			From: acc[0].Name, //should never be dups for account id matches
			Vendor: v.GetMerchantName(),
			Amount: v.Amount,
			Date: v.Date,
			Description: v.Name,
			ProfileID: profileId,
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
func ParseProfileToDB(db *sql.DB, username string, password string){
	profile := model.Profile {
		Name: username,
		Password: password,
	}
	controller.CreateProfile(db, profile)
}
func ParseTokenToDB(db *sql.DB, itemId string, accessToken string){
	token := model.Token {
		Item: itemId,
		Token: accessToken,
	}
	controller.CreateToken(db, token)
}