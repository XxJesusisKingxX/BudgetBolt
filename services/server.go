package main

import (
	// api "services/external/api/plaid"
	api "services/external/api/plaid"
	budgetApi "services/internal/budgeting/db/api"
	transApi "services/internal/transaction_history/db/api"
	userApi "services/internal/user_management/db/api"
	"services/internal/utils/sql/driver"

	"database/sql"
	"fmt"
	"os"

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
)

var db = map[string]*sql.DB {
	"user":        nil,
	"transaction": nil,
	"budget":      nil,
}
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
	// create Plaid client
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", PLAID_CLIENT_ID)
	configuration.AddDefaultHeader("PLAID-SECRET", PLAID_SECRET)
	configuration.UseEnvironment(environments[PLAID_ENV])
	client = plaid.NewAPIClient(configuration)
}

func main() {
	r := gin.Default()
	router := r.Group("/api")
	api.SetupPlaidRoutes(router, api.PlaidClient{}, client)

	dbUser, _ := driver.LogonDB(driver.CREDENTIALS{ User: PG_USER, Pass: PG_SECRET }, "user", driver.DB{}, false)
	dbUser.SetMaxOpenConns(10)
    dbUser.SetMaxIdleConns(5)
	userApi.SetupUserRoutes(router, dbUser)

	dbTransaction, _ := driver.LogonDB(driver.CREDENTIALS{ User: PG_USER, Pass: PG_SECRET }, "transaction", driver.DB{}, false)
	dbTransaction.SetMaxOpenConns(10)
    dbTransaction.SetMaxIdleConns(5)
	transApi.SetupTransactionRoutes(router, dbTransaction)

	dbBudget, _ := driver.LogonDB(driver.CREDENTIALS{ User: PG_USER, Pass: PG_SECRET }, "budget", driver.DB{}, false)
	dbBudget.SetMaxOpenConns(10)
    dbBudget.SetMaxIdleConns(5)
	budgetApi.SetupBudgetRoutes(router, dbBudget)
	
	APP_PORT := "8000"
	err := r.Run(":" + APP_PORT)
	if err != nil {
		panic("unable to start server")
	}
}