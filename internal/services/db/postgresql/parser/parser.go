package controller

import (
	"database/sql"
	"services/db/postgresql/controller"
	"services/db/postgresql/model"

	"github.com/plaid/plaid-go/v12/plaid"
)

func ParseAccountsToDB(db *sql.DB, profileId int, accounts []plaid.AccountBase){
	for _, v := range accounts {
		acc := model.Account {
			ID: v.AccountId,
			Name: v.Name,
			Balance: v.Balances.GetAvailable(),
			ProfileID: profileId,
		}
		controller.CreateAccount(db, acc)
	}
}

func ParseTransactionsToDB(db *sql.DB, profileId int, transactions []plaid.Transaction){
	for _, v := range transactions {
    	acc, _ := controller.RetrieveAccount(db, model.Account{ ID: v.AccountId })
		trans := model.Transaction {
			ID: v.TransactionId,
			From: acc[0].Name,
			Vendor: v.GetMerchantName(),
			Method: v.PaymentChannel,
			Amount: v.Amount,
			Date: v.Date,
			Description: v.Name,
			ProfileID: profileId,
			PrimaryCategory: v.PersonalFinanceCategory.Get().Primary,
			DetailCategory: v.PersonalFinanceCategory.Get().Detailed,
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