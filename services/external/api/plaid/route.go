package api

import (
    "services/internal/api/sql"

	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v12/plaid"
)

func SetupPlaidRoutes(router *gin.RouterGroup, ps Plaid, db map[string]*sql.DB, dbs api.DBHandler, plaid *plaid.APIClient) {
    router.POST("create_link_token", func(c *gin.Context) {
        CreateLinkToken(c, ps, dbs, db, plaid)
    })
    router.POST("create_access_token", func(c *gin.Context) {
        CreateAccessToken(c, ps, dbs, db, plaid, false)
    })
    router.POST("accounts/create", func(c *gin.Context) {
        CreateAccounts(c, ps, dbs, db, plaid, false)
    })
    router.POST("transactions/create", func(c *gin.Context) {
        CreateTransactions(c, ps, dbs, db, plaid, false)
    })

    // Investment and Holding if we add

    // router.POST("investment_transactions/create", func(c *gin.Context) {
    //     CreateInvestmentTransactions(c, ps, dbs, db, plaid, false)
    // })
    // router.POST("holdings/create", func(c *gin.Context) {
    //     CreateHoldings(c, ps, dbs, db, plaid, false)
    // })
}
