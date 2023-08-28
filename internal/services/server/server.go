package main

import (
	"database/sql"
	"fmt"
	"os"
	plaidinterface "services/api/plaid"
	plaidroute "services/api/plaid/route"
	postgresinterface "services/api/postgres"
	postgresroute "services/api/postgres/route"
	driver "services/db/postgresql/driver"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	plaid "github.com/plaid/plaid-go/v12/plaid"
)

var (
	PG_USER             string
	PG_SECRET           string
	PG_DB               string
	PLAID_CLIENT_ID     string
	PLAID_SECRET        string
	PLAID_ENV           string
	PLAID_PRODUCTS      string
	PLAID_COUNTRY_CODES string
	PLAID_REDIRECT_URI  string
	APP_PORT            string
	client              *plaid.APIClient
	db                  *sql.DB
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
	PG_USER = os.Getenv("PG_USER")
	PG_SECRET = os.Getenv("PG_SECRET")
	PG_DB = os.Getenv("PG_DB")
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
	configuration.Scheme = "https"
	configuration.Host = "sandbox.plaid.com"
	configuration.UseEnvironment(environments[PLAID_ENV])
	client = plaid.NewAPIClient(configuration)

	// create database connection
	db, _ = driver.LogonDB(driver.CREDENTIALS{ User: PG_USER, Pass: PG_SECRET }, PG_DB, driver.DB{}, false)
	db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(5)
}

func main() {
	r := gin.Default()
	router := r.Group("/api")
	plaidroute.SetupPlaidRoutes(router, plaidinterface.PlaidClient{}, db, postgresinterface.DB{}, client)
	postgresroute.SetupPostgresRoutes(router, postgresinterface.DB{}, db )
	APP_PORT := "8000"
	err := r.Run(":" + APP_PORT)
	if err != nil {
		panic("unable to start server")
	}
}

// func renderError(c *gin.Context, originalErr error, ps plaidinterface.Plaid) {
// 	plaidError, err := ps.ToPlaidError(originalErr)
// 	if err == nil {
// 		// Return 200 and allow the front end to render the error.
// 		c.JSON(http.StatusOK, gin.H{"error": plaidError})
// 		return
// 	}
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": originalErr.Error()})
// }

// func retrieveProfile(c *gin.Context, dbhandler controller.DBHandler) {
// 	user := c.PostForm("username")
// 	pass := c.PostForm("password")
// 	userProfile, _ := dbhandler.RetrieveProfile(db, strings.ToLower(user))
// 	if userProfile.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{})
// 		return
// 	}
// 	auth := bcrypt.CompareHashAndPassword([]byte(userProfile.Password), []byte(pass))
// 	if auth == nil {
// 		c.JSON(http.StatusOK, gin.H{})
// 	} else {
// 		c.JSON(http.StatusUnauthorized, gin.H{})
// 	}
// }

// func createProfile(c *gin.Context, dbhandler controller.DBHandler, testMode bool) {
// 	user := c.PostForm("username")
// 	pass := c.PostForm("password")
// 	// Test if username is already taken
// 	profile, _ := dbhandler.RetrieveProfile(db, user)
// 	if profile.ID != 0 {
// 		c.JSON(http.StatusConflict, gin.H{})
// 		return
// 	}
// 	var err error
// 	var hashedPass []byte
// 	saltRounds := 17
// 	if testMode {
// 		saltRounds = 1
// 	}
// 	hashedPass, err = bcrypt.GenerateFromPassword([]byte(pass), saltRounds)
// 	if err != nil {
// 		renderError(c, err, plaidinterface.PlaidClient{})
// 		return
// 	}
// 	err = dbhandler.CreateProfile(db, strings.ToLower(user), string(hashedPass))
// 	if err != nil {
// 		renderError(c, err, plaidinterface.PlaidClient{})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{})
// }

// // func createLinkToken(c *gin.Context, p Plaid, dbhandler controller.DBHandler) {
// // 	ctx := context.Background()
// // 	user := c.PostForm("username")
// // 	profile, err := dbhandler.RetrieveProfile(db, user)
// // 	if err != nil {
// // 		renderError(c, err, plaidinterface.PlaidClient{})
// // 		return
// // 	}
// // 	countryCodes := convertCountryCodes(strings.Split(PLAID_COUNTRY_CODES, ","))
// // 	products := convertProducts(strings.Split(PLAID_PRODUCTS, ","))
// // 	request := p.NewLinkTokenCreateRequest(user, strconv.Itoa(profile.ID), countryCodes, products, PLAID_REDIRECT_URI)
// // 	linkTokenCreateResp, err := p.CreateLinkToken(client, ctx, request)
// // 	if err != nil {
// // 		renderError(c, err, plaidinterface.PlaidClient{})
// // 		return
// // 	}
// // 	c.JSON(http.StatusOK, gin.H{"link_token": linkTokenCreateResp.GetLinkToken()})
// // }

// func getAccessToken(c *gin.Context, p plaidinterface.Plaid, dbhandler controller.DBHandler, testMode bool) {
// 	ctx := context.Background()
// 	publicToken := c.PostForm("public_token")
// 	user := c.PostForm("username")
// 	exchangePublicTokenResp, err := p.ItemPublicTokenExchange(client, ctx, publicToken)
// 	if err != nil {
// 		renderError(c, err, plaidinterface.PlaidClient{})
// 		return
// 	}
// 	var id int
// 	accessToken := exchangePublicTokenResp.GetAccessToken()
// 	itemID := exchangePublicTokenResp.GetItemId()
// 	profile, err := dbhandler.RetrieveProfile(db, user)
// 	if err == nil {
// 		id = profile.ID
// 		if !testMode{
// 			controller.CreateToken(db, model.Token{ ProfileID: id, Item: itemID, Token: accessToken })
// 		}
// 		c.JSON(http.StatusOK, gin.H{})
// 	} else {
// 		c.JSON(http.StatusInternalServerError, gin.H{})
// 	}
// }

// func createAccounts(c *gin.Context, p plaidinterface.Plaid, dbhandler controller.DBHandler, testMode bool) {
// 	ctx := context.Background()
// 	user := c.PostForm("profile")
// 	profile, err := dbhandler.RetrieveProfile(db, user)
// 	if err != nil {
// 		renderError(c, err, plaidinterface.PlaidClient{})
// 		return
// 	}
// 	id := profile.ID
// 	token, err := dbhandler.RetrieveToken(db, id)
// 	if err != nil {
// 		renderError(c, err, plaidinterface.PlaidClient{})
// 		return
// 	}
// 	accessToken := token.Token
// 	accounts, err := p.AccountsGet(client, ctx, accessToken)
// 	if err != nil {
// 		renderError(c, err, plaidinterface.PlaidClient{})
// 		return
// 	}
// 	if !testMode {
// 		resp.ParseAccountsToDB(db, id, accounts)
// 	}
// 	c.JSON(http.StatusOK, gin.H{})
// }

// func retrieveAccounts(c *gin.Context, dbhandler controller.DBHandler) {
// 	user := c.Query("username")
// 	profile, err := dbhandler.RetrieveProfile(db, user)
// 	if err != nil {
// 		renderError(c, err, plaidinterface.PlaidClient{})
// 		return
// 	}
// 	accounts, err := dbhandler.RetrieveAccount(db, profile.ID)
// 	if err != nil {
// 		renderError(c, err, plaidinterface.PlaidClient{})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"accounts": accounts,
// 	})
// }

// // func retrieveTransactions(c *gin.Context, dbhandler controller.DBHandler) {
// // 	user := c.Query("username")
// // 	profile, err := dbhandler.RetrieveProfile(db, user)
// // 	var transactions []model.Transaction
// // 	if err == nil {
// // 		transactions, err = dbhandler.RetrieveTransaction(db, model.Transaction{
// // 			ProfileID: profile.ID,
// // 			Query: model.Querys{
// // 				Select: model.QueryParameters{
// // 					Desc: true,
// // 					OrderBy: "transaction_date",
// // 			}}})
// // 	}
// // 	if err != nil {
// // 		renderError(c, err, plaidinterface.PlaidClient{})
// // 		return
// // 	}
// // 	c.JSON(http.StatusOK, gin.H{
// // 		"transactions": transactions,
// // 	})
// // }

// func createTransactions(c *gin.Context, p plaidinterface.Plaid, dbhandler controller.DBHandler, testMode bool) {
// 	ctx := context.Background()
// 	user := c.PostForm("username")
// 	profile, err := dbhandler.RetrieveProfile(db, user)
// 	if err != nil {
// 		renderError(c, err, plaidinterface.PlaidClient{})
// 		return
// 	}
// 	token, err := dbhandler.RetrieveToken(db, profile.ID)
// 	if err != nil {
// 		renderError(c, err, plaidinterface.PlaidClient{})
// 		return
// 	}
// 	accessToken := token.Token
// 	var cursor *string
// 	var transactions []plaid.Transaction
// 	hasMore := true
// 	for hasMore {
// 		resp, err := p.NewTransactionsSyncRequest(client, ctx, accessToken, cursor)
// 		if err != nil {
// 			renderError(c, err, plaidinterface.PlaidClient{})
// 			return
// 		}
// 		transactions = append(transactions, resp.GetAdded()...)
// 		hasMore = resp.GetHasMore()
// 		nextCursor := resp.GetNextCursor()
// 		cursor = &nextCursor
// 	}
// 	if !testMode {
// 		resp.ParseTransactionsToDB(db, profile.ID, transactions)
// 	}
// 	c.JSON(http.StatusOK, gin.H{})
// }

// func investmentTransactions(c *gin.Context, p plaidinterface.Plaid, testMode bool) {
// 	ctx := context.Background()
// 	var accessToken string
// 	if testMode {
// 		accessToken = c.PostForm("access_token")
// 	}
// 	invTxResp, err := p.InvestmentsTransactionsGet(client, ctx, accessToken)
// 	if err != nil {
// 		renderError(c, err, plaidinterface.PlaidClient{})
// 		return
// 	}
// 	invest := invTxResp.InvestmentTransactions
// 	// accounts := invTxResp.Accounts
// 	c.JSON(http.StatusOK, gin.H{
// 		"investments_transactions": invTxResp,
// 	})
// 	if !testMode {
// 		// resp.ParseAccountsToDB(db, accessToken, accounts)
// 		resp.ParseInvestmentsToDB(db, invest)
// 	}
// }

// func holdings(c *gin.Context, p plaidinterface.Plaid, testMode bool) {
// 	ctx := context.Background()
// 	var accessToken string
// 	if testMode {
// 		accessToken = c.PostForm("access_token")
// 	}
// 	holdingsGetResp, err := p.InvestmentsHoldingsGet(client, ctx, accessToken)
// 	if err != nil {
// 		renderError(c, err, plaidinterface.PlaidClient{})
// 		return
// 	}
// 	// accounts := holdingsGetResp.Accounts
// 	holdings := holdingsGetResp.Holdings
// 	c.JSON(http.StatusOK, gin.H{
// 		"holdings": holdings,
// 	})
// 	if !testMode {
// 		// resp.ParseAccountsToDB(db, accessToken, accounts)
// 		resp.ParseHoldingsToDB(db, holdings)
// 	}
// }