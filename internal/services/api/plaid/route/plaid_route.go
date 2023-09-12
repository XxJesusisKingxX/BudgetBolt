package routes

import (
	"database/sql"
	plaidinterface "services/api/plaid"
	"services/api/plaid/controller"
    "services/api/postgres"
	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v12/plaid"
)

func SetupPlaidRoutes(router *gin.RouterGroup, ps plaidinterface.Plaid, db *sql.DB, dbs postgresinterface.DBHandler, plaid *plaid.APIClient) {
    router.POST("create_link_token", func(c *gin.Context) {
        controller.CreateLinkToken(c, ps, dbs, db, plaid)
    })
    router.POST("create_access_token", func(c *gin.Context) {
        controller.CreateAccessToken(c, ps, dbs, db, plaid, false)
    })
    router.POST("accounts/create", func(c *gin.Context) {
        controller.CreateAccounts(c, ps, dbs, db, plaid, false)
    })
    router.POST("transactions/create", func(c *gin.Context) {
        controller.CreateTransactions(c, ps, dbs, db, plaid, false)
    })
    router.POST("investment_transactions/create", func(c *gin.Context) {
        controller.CreateInvestmentTransactions(c, ps, dbs, db, plaid, false)
    })
    router.POST("holdings/create", func(c *gin.Context) {
        controller.CreateHoldings(c, ps, dbs, db, plaid, false)
    })
}
