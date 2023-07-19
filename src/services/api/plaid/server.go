package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	resp "budgetbolt/src/services/api/plaid/response"
	driver "budgetbolt/src/services/databases/postgresql/driver"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	plaid "github.com/plaid/plaid-go/v12/plaid"
)


var (
	PLAID_CLIENT_ID                      = ""
	PLAID_SECRET                         = ""
	PLAID_ENV                            = ""
	PLAID_PRODUCTS                       = ""
	PLAID_COUNTRY_CODES                  = ""
	PLAID_REDIRECT_URI                   = ""
	APP_PORT                             = ""
	client              *plaid.APIClient = nil
	db					*sql.DB	         = nil
)

var environments = map[string]plaid.Environment{
	"sandbox":     plaid.Sandbox,
	"development": plaid.Development,
	"production":  plaid.Production,
}

func init() {
	// load env vars from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error when loading environment variables from .env file %w", err)
	}

	// set constants from env
	PLAID_CLIENT_ID = os.Getenv("PLAID_CLIENT_ID")
	PLAID_SECRET = os.Getenv("PLAID_SECRET")

	if PLAID_CLIENT_ID == "" || PLAID_SECRET == "" {
		log.Fatal("Error: PLAID_SECRET or PLAID_CLIENT_ID is not set. Did you copy .env.example to .env and fill it out?")
	}

	PLAID_ENV = os.Getenv("PLAID_ENV")
	PLAID_PRODUCTS = os.Getenv("PLAID_PRODUCTS")
	PLAID_COUNTRY_CODES = os.Getenv("PLAID_COUNTRY_CODES")
	PLAID_REDIRECT_URI = os.Getenv("PLAID_REDIRECT_URI")
	APP_PORT = os.Getenv("APP_PORT")

	// set defaults
	if PLAID_PRODUCTS == "" {
		PLAID_PRODUCTS = "transactions"
	}
	if PLAID_COUNTRY_CODES == "" {
		PLAID_COUNTRY_CODES = "US"
	}
	if PLAID_ENV == "" {
		PLAID_ENV = "sandbox"
	}
	if APP_PORT == "" {
		APP_PORT = "8000"
	}
	if PLAID_CLIENT_ID == "" {
		log.Fatal("PLAID_CLIENT_ID is not set. Make sure to fill out the .env file")
	}
	if PLAID_SECRET == "" {
		log.Fatal("PLAID_SECRET is not set. Make sure to fill out the .env file")
	}

	// create Plaid client
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", PLAID_CLIENT_ID)
	configuration.AddDefaultHeader("PLAID-SECRET", PLAID_SECRET)
	configuration.UseEnvironment(environments[PLAID_ENV])
	client = plaid.NewAPIClient(configuration)

	// create database connection
	db, _ = driver.LogonDB(driver.CREDENTIALS{User:"postgres", Pass: `P-S$\\/M1n3!`}, "budgetbolt", driver.DB{}, false)
}

func maint() {
	r := gin.Default()
	r.POST("/api/create_link_token", createLinkToken)
	r.POST("/api/set_access_token", func(c *gin.Context){ getAccessToken(c, PlaidClient{}) })
	r.POST("/api/info", info)
	r.GET("/api/accounts", func(c *gin.Context){ accounts(c, PlaidClient{}, false) })
	r.GET("/api/transactions", transactions)
	r.GET("/api/investments_transactions", func(c *gin.Context){ investmentTransactions(c, PlaidClient{}, false) })
	r.GET("/api/holdings", func(c *gin.Context){ holdings(c, PlaidClient{}, false) })
	err := r.Run(":" + APP_PORT)
	if err != nil {
		panic("unable to start server")
	}
}

// We store the access_token in memory - in production, store it in a secure
// persistent data store.
var accessToken string
var itemID string

func renderError(c *gin.Context, originalErr error, plaid Plaid) {
	plaidError, err := plaid.ToPlaidError(originalErr)
	if err == nil {
		// Return 200 and allow the front end to render the error.
		c.JSON(http.StatusOK, gin.H{"error": plaidError})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": originalErr.Error()})
}

func getAccessToken(c *gin.Context, plaid Plaid) {
	publicToken := c.PostForm("public_token")
	ctx := context.Background()
	// exchange the public_token for an access_token
	exchangePublicTokenResp, err := plaid.ItemPublicTokenExchange(ctx, publicToken)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	accessToken = exchangePublicTokenResp.GetAccessToken()
	itemID = exchangePublicTokenResp.GetItemId()
	fmt.Println("public token: " + publicToken)
	fmt.Println("access token: " + accessToken)
	fmt.Println("item ID: " + itemID)
	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"item_id":      itemID,
	})
}

func accounts(c *gin.Context, plaid Plaid, testMode bool) {
	ctx := context.Background()
	if testMode == true {
		accessToken = c.PostForm("access_token")
	}
	accountsGetResp, err := plaid.AccountsGet(ctx, accessToken)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	accounts := accountsGetResp.Accounts
	c.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
	})
	if testMode != true {
		resp.ParseAccountsToDB(db, accessToken, accounts)
	}
}

func transactions(c *gin.Context) {
	ctx := context.Background()
	// Set cursor to empty to receive all historical updates
	var cursor *string
	// New transaction updates since "cursor"
	var added []plaid.Transaction
	var modified []plaid.Transaction
	var removed []plaid.RemovedTransaction // Removed transaction ids
	hasMore := true
	// Iterate through each page of new transaction updates for item
	for hasMore {
		request := plaid.NewTransactionsSyncRequest(accessToken)
		if cursor != nil {
			request.SetCursor(*cursor)
		}
		resp, _, err := client.PlaidApi.TransactionsSync(
			ctx,
		).TransactionsSyncRequest(*request).Execute()
		if err != nil {
			renderError(c, err, PlaidClient{})
			return
		}
		// Add this page of results
		added = append(added, resp.GetAdded()...)
		modified = append(modified, resp.GetModified()...)
		removed = append(removed, resp.GetRemoved()...)
		hasMore = resp.GetHasMore()
		// Update cursor to the next cursor
		nextCursor := resp.GetNextCursor()
		cursor = &nextCursor
	}

	sort.Slice(added, func(i, j int) bool {
		return added[i].GetDate() < added[j].GetDate()
	})
	var latestTransactions []plaid.Transaction = nil
	if len(added) >= 9 {
		latestTransactions = added[len(added)-9:]

	}
	c.JSON(http.StatusOK, gin.H{
		"latest_transactions": latestTransactions,
	})
	resp.ParseTransactionsToDB(db, latestTransactions)
}

func investmentTransactions(c *gin.Context, plaid Plaid, testMode bool) {
	ctx := context.Background()
	if testMode == true {
		accessToken = c.PostForm("access_token")
	}
	invTxResp, err := plaid.InvestmentsTransactionsGet(ctx, accessToken)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	invest := invTxResp.InvestmentTransactions
	accounts := invTxResp.Accounts
	c.JSON(http.StatusOK, gin.H{
		"investments_transactions": invTxResp,
	})
	if testMode != true {
		resp.ParseAccountsToDB(db, accessToken, accounts)
		resp.ParseInvestmentsToDB(db, invest)
	}
}

func holdings(c *gin.Context, plaid Plaid, testMode bool) {
	ctx := context.Background()
	if testMode == true {
		accessToken = c.PostForm("access_token")
	}
	holdingsGetResp, err := plaid.InvestmentsHoldingsGet(ctx, accessToken)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	accounts := holdingsGetResp.Accounts
	holdings := holdingsGetResp.Holdings
	c.JSON(http.StatusOK, gin.H{
		"holdings": holdings,
	})
	if testMode != true {
		resp.ParseAccountsToDB(db, accessToken, accounts)
		resp.ParseHoldingsToDB(db, holdings)
	}
}

func info(context *gin.Context) {
	context.JSON(http.StatusOK, map[string]interface{}{
		"item_id":      itemID,
		"access_token": accessToken,
		"products":     strings.Split(PLAID_PRODUCTS, ","),
	})
}

func createLinkToken(c *gin.Context) {
	linkToken, err := linkTokenCreate()
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	c.JSON(http.StatusOK, gin.H{"link_token": linkToken})
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

// linkTokenCreate creates a link token using the specified parameters
func linkTokenCreate() (string, error) {
	ctx := context.Background()
	countryCodes := convertCountryCodes(strings.Split(PLAID_COUNTRY_CODES, ","))
	redirectURI := PLAID_REDIRECT_URI
	user := plaid.LinkTokenCreateRequestUser{
		ClientUserId: time.Now().String(),
	}
	request := plaid.NewLinkTokenCreateRequest(
		"Plaid Quickstart",
		"en",
		countryCodes,
		user,
	)
	products := convertProducts(strings.Split(PLAID_PRODUCTS, ","))
	request.SetProducts(products)
	if redirectURI != "" {
		request.SetRedirectUri(redirectURI)
	}
	linkTokenCreateResp, _, err := client.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()
	if err != nil {
		return "", err
	}
	return linkTokenCreateResp.GetLinkToken(), nil
}