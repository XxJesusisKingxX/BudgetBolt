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
	// create Plaid client
	configuration := plaid.NewConfiguration()
	configuration.AddDefaultHeader("PLAID-CLIENT-ID", PLAID_CLIENT_ID)
	configuration.AddDefaultHeader("PLAID-SECRET", PLAID_SECRET)
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