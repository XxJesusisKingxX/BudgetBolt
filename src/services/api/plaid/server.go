package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"
	"golang.org/x/crypto/bcrypt"

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
	PLAID_ENV = os.Getenv("PLAID_ENV")
	PLAID_PRODUCTS = os.Getenv("PLAID_PRODUCTS")
	PLAID_COUNTRY_CODES = os.Getenv("PLAID_COUNTRY_CODES")
	PLAID_REDIRECT_URI = os.Getenv("PLAID_REDIRECT_URI")

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
	r.POST("/api/set_access_token", func(c *gin.Context){ getAccessToken(c, PlaidClient{}, controller.DB{}, false) })
	r.POST("/api/accounts/create", func(c *gin.Context){ createAccounts(c, PlaidClient{}, controller.DB{}, false) })
	r.GET("/api/accounts/get", func(c *gin.Context){ retrieveAccounts(c, controller.DB{}) })
	r.POST("/api/profile/get", func(c *gin.Context){ retrieveProfile(c, controller.DB{}) })
	r.POST("/api/profile/create", func(c *gin.Context){ createProfile(c, controller.DB{}, false) })
	r.POST("/api/transactions/create", func(c *gin.Context){ createTransactions(c, PlaidClient{}, controller.DB{}, false) })
	r.GET("/api/transactions/get", func(c *gin.Context){ retrieveTransactions(c, controller.DB{}) })
	r.GET("/api/investments_transactions", func(c *gin.Context){ investmentTransactions(c, PlaidClient{}, false) })
	r.GET("/api/holdings", func(c *gin.Context){ holdings(c, PlaidClient{}, false) })
	APP_PORT := "8000"
	err := r.Run(":" + APP_PORT)
	if err != nil {
		panic("unable to start server")
	}
}

func retrieveProfile(c *gin.Context, dbhandler controller.DBHandler) {
	user := c.PostForm("username")
	pass := c.PostForm("password")
	userProfile, err := dbhandler.RetrieveProfile(db, strings.ToLower(user))
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	auth := bcrypt.CompareHashAndPassword([]byte(userProfile.Password), []byte(pass))
	if auth == nil {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}
}

func createProfile(c *gin.Context, dbhandler controller.DBHandler, testMode bool) {
	user := c.PostForm("username")
	pass := c.PostForm("password")
	// Test if username is already taken
	profile, _ := dbhandler.RetrieveProfile(db, user)
	if profile.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{})
		return
	}
	var err error
	var hashedPass []byte
	saltRounds := 17
	if testMode {
		saltRounds = 1
	}
	hashedPass, err = bcrypt.GenerateFromPassword([]byte(pass), saltRounds)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	err = dbhandler.CreateProfile(db, strings.ToLower(user), string(hashedPass))
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
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

func getAccessToken(c *gin.Context, p Plaid, dbhandler controller.DBHandler, testMode bool) {
	ctx := context.Background()
	publicToken := c.PostForm("public_token")
	user := c.PostForm("profile")
	exchangePublicTokenResp, err := p.ItemPublicTokenExchange(client, ctx, publicToken)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	var id int
	accessToken := exchangePublicTokenResp.GetAccessToken()
	itemID := exchangePublicTokenResp.GetItemId()
	profile, err := dbhandler.RetrieveProfile(db, user)
	if err == nil {
		id = profile.ID
		if !testMode{
			controller.CreateToken(db, model.Token{ ProfileID: id, Item: itemID, Token: accessToken })
		}
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
}

func createAccounts(c *gin.Context, p Plaid, dbhandler controller.DBHandler, testMode bool) {
	ctx := context.Background()
	user := c.PostForm("profile")
	profile, err := dbhandler.RetrieveProfile(db, user)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	id := profile.ID
	token, err := dbhandler.RetrieveToken(db, id)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	accessToken := token.Token
	accounts, err := p.AccountsGet(client, ctx, accessToken)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	if !testMode {
		resp.ParseAccountsToDB(db, id, accounts)
	}
	c.JSON(http.StatusOK, gin.H{})
}

func retrieveAccounts(c *gin.Context, dbhandler controller.DBHandler) {
	user := c.Query("username")
	profile, err := dbhandler.RetrieveProfile(db, user)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	accounts, err := dbhandler.RetrieveAccount(db, profile.ID)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
	})
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
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	token, err := dbhandler.RetrieveToken(db, profile.ID)
	if err != nil {
		renderError(c, err, PlaidClient{})
		return
	}
	accessToken := token.Token
	var cursor *string
	var transactions []plaid.Transaction
	hasMore := true
	for hasMore {
		resp, err := p.NewTransactionsSyncRequest(client, ctx, accessToken, cursor)
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
	c.JSON(http.StatusOK, gin.H{})
}

func investmentTransactions(c *gin.Context, p Plaid, testMode bool) {
	ctx := context.Background()
	var accessToken string
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
	var accessToken string
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