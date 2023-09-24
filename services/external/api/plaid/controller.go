package api

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v12/plaid"

	"services/internal/api/sql"
	transaction "services/internal/transaction_history/db/utils"
	user "services/internal/user_managment/db/model"
)

func CreateLinkToken(c *gin.Context, ps Plaid, dbs api.DBHandler, db map[string]*sql.DB, plaidapi *plaid.APIClient) {
	ctx := context.Background()
	uid, _ := c.Cookie("UID")
	profile, err := dbs.RetrieveProfile(db["user"], uid, true)
	if err != nil {
		RenderError(c, err, PlaidClient{})
		return
	}
	countryCodes := convertCountryCodes(strings.Split("US", ","))
	products := convertProducts(strings.Split("transactions", ","))
	request := ps.NewLinkTokenCreateRequest(uid, strconv.Itoa(profile.ID), countryCodes, products, "")
	linkTokenCreateResp, err := ps.CreateLinkToken(plaidapi, ctx, request)
	if err != nil {
		RenderError(c, err, PlaidClient{})
		return
	}
	c.JSON(http.StatusOK, gin.H{"link_token": linkTokenCreateResp.GetLinkToken()})
}

func CreateAccessToken(c *gin.Context, ps Plaid, dbs api.DBHandler, db map[string]*sql.DB, plaidapi *plaid.APIClient, debug bool) {
	ctx := context.Background()
	publicToken := c.PostForm("public_token")
	uid, _ := c.Cookie("UID")
	exchangePublicTokenResp, err := ps.ItemPublicTokenExchange(plaidapi, ctx, publicToken)
	if err != nil {
		RenderError(c, err, PlaidClient{})
		return
	}
	var id int
	accessToken := exchangePublicTokenResp.GetAccessToken()
	itemID := exchangePublicTokenResp.GetItemId()
	profile, err := dbs.RetrieveProfile(db["user"], uid, true)
	if err == nil {
		id = profile.ID
		err := dbs.CreateToken(db["user"], user.Token{ ProfileID: id, Item: itemID, Token: accessToken })
		if err != nil {
			RenderError(c, err, PlaidClient{})
			return
		}
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
}

func CreateAccounts(c *gin.Context, ps Plaid, dbs api.DBHandler, db map[string]*sql.DB, plaidapi *plaid.APIClient, debug bool) {
	ctx := context.Background()
	uid, _ := c.Cookie("UID")
	profile, err := dbs.RetrieveProfile(db["user"], uid, true)
	if err != nil {
		RenderError(c, err, PlaidClient{})
		return
	}
	id := profile.ID
	token, err := dbs.RetrieveToken(db["user"], id)
	if err != nil {
		RenderError(c, err, PlaidClient{})
		return
	}
	accessToken := token.Token
	accounts, err := ps.AccountsGet(plaidapi, ctx, accessToken)
	if err != nil {
		RenderError(c, err, PlaidClient{})
		return
	}
	if !debug {
		transaction.AccountsToDB(db["transaction"], id, accounts)
	}
	c.JSON(http.StatusOK, gin.H{})
}

func CreateTransactions(c *gin.Context, ps Plaid, dbs api.DBHandler, db map[string]*sql.DB, plaidapi *plaid.APIClient, debug bool) {
	ctx := context.Background()
	uid, _ := c.Cookie("UID")
	profile, err := dbs.RetrieveProfile(db["user"], uid, true)
	if err != nil {
		RenderError(c, err, PlaidClient{})
		return
	}
	token, err := dbs.RetrieveToken(db["user"], profile.ID)
	if err != nil {
		RenderError(c, err, PlaidClient{})
		return
	}
	accessToken := token.Token
	var cursor *string
	var transactions []plaid.Transaction
	hasMore := true
	for hasMore {
		resp, err := ps.NewTransactionsSyncRequest(plaidapi, ctx, accessToken, cursor)
		if err != nil {
			RenderError(c, err, PlaidClient{})
			return
		}
		transactions = append(transactions, resp.GetAdded()...)
		hasMore = resp.GetHasMore()
		nextCursor := resp.GetNextCursor()
		cursor = &nextCursor
	}
	if !debug {
		transaction.TransactionsToDB(db["transaction"], profile.ID, transactions)
	}
	c.JSON(http.StatusOK, gin.H{})
}

// Investment and Holding if we ever decide to add

// func CreateInvestmentTransactions(c *gin.Context, ps Plaid, dbs api.DBHandler, db map[string]*sql.DB, plaidapi *plaid.APIClient, debug bool) {
// 	ctx := context.Background()
// 	uid, _ := c.Cookie("UID")
// 	profile, err := dbs.RetrieveProfile(db["user"], uid, true)
// 	if err != nil {
// 		RenderError(c, err, PlaidClient{})
// 		return
// 	}
// 	token, err := dbs.RetrieveToken(db["user"], profile.ID)
// 	if err != nil {
// 		RenderError(c, err, PlaidClient{})
// 		return
// 	}
// 	accessToken := token.Token
// 	invTxResp, err := ps.InvestmentsTransactionsGet(plaidapi, ctx, accessToken)
// 	if err != nil {
// 		RenderError(c, err, PlaidClient{})
// 		return
// 	}
// 	invest := invTxResp.InvestmentTransactions
// 	// accounts := invTxResp.Accounts
// 	if !debug {
// 		// resp.ParseAccountsToDB(db, accessToken, accounts)
// 		postgresparser.ParseInvestmentsToDB(db[""], invest)
// 	}
// 	c.JSON(http.StatusOK, gin.H{})
// }

// func CreateHoldings(c *gin.Context, ps Plaid, dbs api.DBHandler, db map[string]*sql.DB, plaidapi *plaid.APIClient, debug bool) {
// 	ctx := context.Background()
// 	uid, _ := c.Cookie("UID")
// 	profile, err := dbs.RetrieveProfile(db["user"], uid, true)
// 	if err != nil {
// 		RenderError(c, err, PlaidClient{})
// 		return
// 	}
// 	token, err := dbs.RetrieveToken(db["user"], profile.ID)
// 	if err != nil {
// 		RenderError(c, err, PlaidClient{})
// 		return
// 	}
// 	accessToken := token.Token
// 	holdingsGetResp, err := ps.InvestmentsHoldingsGet(plaidapi, ctx, accessToken)
// 	if err != nil {
// 		RenderError(c, err, PlaidClient{})
// 		return
// 	}
// 	// accounts := holdingsGetResp.Accounts
// 	holdings := holdingsGetResp.Holdings
// 	if !debug {
// 		// resp.ParseAccountsToDB(db, accessToken, accounts)
// 		postgresparser.ParseHoldingsToDB(db[""], holdings)
// 	}
// 	c.JSON(http.StatusOK, gin.H{})
// }
