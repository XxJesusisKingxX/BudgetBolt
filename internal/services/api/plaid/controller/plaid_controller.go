package controller

import (
	"context"
	"database/sql"
	"net/http"
	plaidinterface "services/api/plaid"
	"services/api/postgres"
	"services/api/utils"
	"services/db/postgresql/controller"
	"services/db/postgresql/model"
	postgresparser "services/db/postgresql/parser"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v12/plaid"
)

func CreateLinkToken(c *gin.Context, ps plaidinterface.Plaid, dbs postgresinterface.DBHandler, db *sql.DB, plaidapi *plaid.APIClient) {
	ctx := context.Background()
	uid, _ := c.Cookie("UID")
	profile, err := dbs.RetrieveProfile(db, uid, true)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	countryCodes := convertCountryCodes(strings.Split("US", ","))
	products := convertProducts(strings.Split("transactions", ","))
	request := ps.NewLinkTokenCreateRequest(uid, strconv.Itoa(profile.ID), countryCodes, products, "")
	linkTokenCreateResp, err := ps.CreateLinkToken(plaidapi, ctx, request)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	c.JSON(http.StatusOK, gin.H{"link_token": linkTokenCreateResp.GetLinkToken()})
}

func CreateAccessToken(c *gin.Context, ps plaidinterface.Plaid, dbs postgresinterface.DBHandler, db *sql.DB, plaidapi *plaid.APIClient, debug bool) {
	ctx := context.Background()
	publicToken := c.PostForm("public_token")
	uid, _ := c.Cookie("UID")
	exchangePublicTokenResp, err := ps.ItemPublicTokenExchange(plaidapi, ctx, publicToken)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	var id int
	accessToken := exchangePublicTokenResp.GetAccessToken()
	itemID := exchangePublicTokenResp.GetItemId()
	profile, err := dbs.RetrieveProfile(db, uid, true)
	if err == nil {
		id = profile.ID
		if !debug {
			controller.CreateToken(db, model.Token{ ProfileID: id, Item: itemID, Token: accessToken })
		}
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
}

func CreateAccounts(c *gin.Context, ps plaidinterface.Plaid, dbs postgresinterface.DBHandler, db *sql.DB, plaidapi *plaid.APIClient, debug bool) {
	ctx := context.Background()
	uid, _ := c.Cookie("UID")
	profile, err := dbs.RetrieveProfile(db, uid, true)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	id := profile.ID
	token, err := dbs.RetrieveToken(db, id)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	accessToken := token.Token
	accounts, err := ps.AccountsGet(plaidapi, ctx, accessToken)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	if !debug {
		postgresparser.ParseAccountsToDB(db, id, accounts)
	}
	c.JSON(http.StatusOK, gin.H{})
}

func CreateTransactions(c *gin.Context, ps plaidinterface.Plaid, dbs postgresinterface.DBHandler, db *sql.DB, plaidapi *plaid.APIClient, debug bool) {
	ctx := context.Background()
	uid, _ := c.Cookie("UID")
	profile, err := dbs.RetrieveProfile(db, uid, true)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	token, err := dbs.RetrieveToken(db, profile.ID)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	accessToken := token.Token
	var cursor *string
	var transactions []plaid.Transaction
	hasMore := true
	for hasMore {
		resp, err := ps.NewTransactionsSyncRequest(plaidapi, ctx, accessToken, cursor)
		if err != nil {
			utils.RenderError(c, err, plaidinterface.PlaidClient{})
			return
		}
		transactions = append(transactions, resp.GetAdded()...)
		hasMore = resp.GetHasMore()
		nextCursor := resp.GetNextCursor()
		cursor = &nextCursor
	}
	if !debug {
		postgresparser.ParseTransactionsToDB(db, profile.ID, transactions)
	}
	c.JSON(http.StatusOK, gin.H{})
}

func CreateInvestmentTransactions(c *gin.Context, ps plaidinterface.Plaid, dbs postgresinterface.DBHandler, db *sql.DB, plaidapi *plaid.APIClient, debug bool) {
	ctx := context.Background()
	uid, _ := c.Cookie("UID")
	profile, err := dbs.RetrieveProfile(db, uid, true)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	token, err := dbs.RetrieveToken(db, profile.ID)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	accessToken := token.Token
	invTxResp, err := ps.InvestmentsTransactionsGet(plaidapi, ctx, accessToken)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	invest := invTxResp.InvestmentTransactions
	// accounts := invTxResp.Accounts
	if !debug {
		// resp.ParseAccountsToDB(db, accessToken, accounts)
		postgresparser.ParseInvestmentsToDB(db, invest)
	}
	c.JSON(http.StatusOK, gin.H{})
}

func CreateHoldings(c *gin.Context, ps plaidinterface.Plaid, dbs postgresinterface.DBHandler, db *sql.DB, plaidapi *plaid.APIClient, debug bool) {
	ctx := context.Background()
	uid, _ := c.Cookie("UID")
	profile, err := dbs.RetrieveProfile(db, uid, true)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	token, err := dbs.RetrieveToken(db, profile.ID)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	accessToken := token.Token
	holdingsGetResp, err := ps.InvestmentsHoldingsGet(plaidapi, ctx, accessToken)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	// accounts := holdingsGetResp.Accounts
	holdings := holdingsGetResp.Holdings
	if !debug {
		// resp.ParseAccountsToDB(db, accessToken, accounts)
		postgresparser.ParseHoldingsToDB(db, holdings)
	}
	c.JSON(http.StatusOK, gin.H{})
}

func convertCountryCodes(countryCodeStrs []string) []plaid.CountryCode {
	countryCodes := []plaid.CountryCode{}
	for _, countryCodeStr := range countryCodeStrs {
		countryCodes = append(countryCodes, plaid.CountryCode(countryCodeStr))
	}
	return countryCodes
}

func convertProducts(productStrs []string) []plaid.Products {
	products := []plaid.Products{}
	for _, productStr := range productStrs {
		products = append(products, plaid.Products(productStr))
	}
	return products
}

