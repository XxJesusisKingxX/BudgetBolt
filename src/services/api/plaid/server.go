package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	resp "budgetbolt/src/services/api/plaid/response"
	"budgetbolt/src/services/databases/postgresql/controller"
	driver "budgetbolt/src/services/databases/postgresql/driver"
	"budgetbolt/src/services/databases/postgresql/model"

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

func main() {
	r := gin.Default()
	r.POST("/api/create_link_token", func(c *gin.Context){ createLinkToken(c, PlaidClient{}) })
	r.POST("/api/set_access_token", func(c *gin.Context){ getAccessToken(c, PlaidClient{}, false) })
	r.POST("/api/info", info)
	r.POST("/api/accounts/create", func(c *gin.Context){ createAccounts(c, PlaidClient{}, controller.DB{}, true) })
	r.GET("/api/accounts/get", func(c *gin.Context){ retrieveAccounts(c, controller.DB{}) })
	r.POST("/api/transactions/create", func(c *gin.Context){ createTransactions(c, PlaidClient{}, controller.DB{}, false) })
	r.GET("/api/transactions/get", func(c *gin.Context){ retrieveTransactions(c, controller.DB{}) })
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

func info(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"item_id":      itemID,
		"access_token": accessToken,
		"products":     strings.Split(PLAID_PRODUCTS, ","),
	})
}

func createLinkToken(c *gin.Context, p Plaid) {
	ctx := context.Background()
	countryCodes := convertCountryCodes(strings.Split(PLAID_COUNTRY_CODES, ","))
	products := convertProducts(strings.Split(PLAID_PRODUCTS, ","))
	request := p.NewLinkTokenCreateRequest("Test User", "TestUser", countryCodes,  products, PLAID_REDIRECT_URI)
	linkTokenCreateResp, err := p.CreateLinkToken(client, ctx, request)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	c.JSON(http.StatusOK, gin.H{"link_token": linkTokenCreateResp.GetLinkToken()})
}

func getAccessToken(c *gin.Context, p Plaid, testMode bool) {
	publicToken := c.PostForm("public_token")
	user := c.PostForm("username")
	ctx := context.Background()
	exchangePublicTokenResp, err := p.ItemPublicTokenExchange(client, ctx, publicToken)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	accessToken = exchangePublicTokenResp.GetAccessToken()
	itemID = exchangePublicTokenResp.GetItemId()
	if !testMode {
		profile, err := controller.RetrieveProfile(db, model.Profile{ Name: user })
		if err == nil {
			controller.CreateToken(db, model.Token{ ProfileID: profile.ID, Item: itemID, Token: accessToken })
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
		"item_id":      itemID,
	})
}

func createAccounts(c *gin.Context, p Plaid, dbhandler controller.DBHandler, testMode bool) {
	ctx := context.Background()
	user := c.PostForm("username")
	profile, err := dbhandler.RetrieveProfile(db, user)
	if err == nil {
		id := profile.ID
		token, err := dbhandler.RetrieveToken(db, id)
		if err == nil {
			accessToken := token.Token
			accounts, err := p.AccountsGet(client, ctx, accessToken)
			if err != nil {
				renderError(c, err, PlaidClient{})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"accounts": accounts,
			})
			if !testMode {
				resp.ParseAccountsToDB(db, id, accounts)
			}
		}
		renderError(c, err, PlaidClient{})
	}
	renderError(c, err, PlaidClient{})
}
func retrieveAccounts(c *gin.Context, dbhandler controller.DBHandler) {
	user := c.Query("username")
	profile, err := dbhandler.RetrieveProfile(db, user)
	if err == nil {
		accounts, err := dbhandler.RetrieveAccount(db, profile.ID)
		if err != nil {
			renderError(c, err, PlaidClient{})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"accounts": accounts,
		})
	}
	renderError(c, err, PlaidClient{})
}

func retrieveTransactions(c *gin.Context, dbhandler controller.DBHandler) {
	user := c.Query("username")
	profile, err := dbhandler.RetrieveProfile(db, user)
	var transactions []model.Transaction
	if err == nil {
		transactions, err = dbhandler.RetrieveTransaction(db, model.Transaction{ 
			ProfileID: profile.ID, 
			Query: model.Querys{ 
				Select: model.QueryParameters{
					Desc: true,
					OrderBy: "transaction_date",
			}}})
	}
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}

func createTransactions(c *gin.Context, p Plaid, dbhandler controller.DBHandler, testMode bool) {
	ctx := context.Background()
	user := c.PostForm("username")
	profile, err := dbhandler.RetrieveProfile(db, user)
	if err == nil {
		token, err := dbhandler.RetrieveToken(db, profile.ID)
		if err == nil {
			accessToken := token.Token
			var cursor *string
			var transactions []plaid.Transaction
			hasMore := true
			for hasMore {
				resp, err := p.NewTransactionsSyncRequest(client, ctx, accessToken, cursor)
				fmt.Println(resp)
				if err != nil {
					renderError(c, err, PlaidClient{})
					return
				}
				transactions = append(transactions, resp.GetAdded()...)
				hasMore = resp.GetHasMore()
				nextCursor := resp.GetNextCursor()
				cursor = &nextCursor
			}
			if !testMode {
				resp.ParseTransactionsToDB(db, profile.ID, transactions)
			}
		}
		renderError(c, err, PlaidClient{})
	}
	renderError(c, err, PlaidClient{})
}

func investmentTransactions(c *gin.Context, p Plaid, testMode bool) {
	ctx := context.Background()
	if testMode {
		accessToken = c.PostForm("access_token")
	}
	invTxResp, err := p.InvestmentsTransactionsGet(client, ctx, accessToken)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	invest := invTxResp.InvestmentTransactions
	// accounts := invTxResp.Accounts
	c.JSON(http.StatusOK, gin.H{
		"investments_transactions": invTxResp,
	})
	if !testMode {
		// resp.ParseAccountsToDB(db, accessToken, accounts)
		resp.ParseInvestmentsToDB(db, invest)
	}
}

func holdings(c *gin.Context, p Plaid, testMode bool) {
	ctx := context.Background()
	if testMode {
		accessToken = c.PostForm("access_token")
	}
	holdingsGetResp, err := p.InvestmentsHoldingsGet(client, ctx, accessToken)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	// accounts := holdingsGetResp.Accounts
	holdings := holdingsGetResp.Holdings
	c.JSON(http.StatusOK, gin.H{
		"holdings": holdings,
	})
	if !testMode {
		// resp.ParseAccountsToDB(db, accessToken, accounts)
		resp.ParseHoldingsToDB(db, holdings)
	}
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

func renderError(c *gin.Context, originalErr error, p Plaid) {
	plaidError, err := p.ToPlaidError(originalErr)
	if err == nil {
		// Return 200 and allow the front end to render the error.
		c.JSON(http.StatusOK, gin.H{"error": plaidError})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": originalErr.Error()})
}