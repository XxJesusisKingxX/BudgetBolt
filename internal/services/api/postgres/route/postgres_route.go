package routes

import (
	"database/sql"
	"services/api/postgres/controller"
	"services/api/postgres"
	"github.com/gin-gonic/gin"
)

func SetupPostgresRoutes(router *gin.RouterGroup, dbs postgresinterface.DBHandler, db *sql.DB ) {
    router.POST("profile/create", func(c *gin.Context) {
        controller.CreateProfile(c, dbs, db, false)
    })
    router.POST("profile/get", func(c *gin.Context) {
        controller.RetrieveProfile(c, dbs, db, false)
    })
    router.POST("expenses/create", func(c *gin.Context) {
        controller.CreateExpenses(c, dbs, db, false)
    })
    router.GET("expenses/get", func(c *gin.Context) {
        controller.RetrieveExpenses(c, dbs, db, false)
    })
    router.GET("transactions/get", func(c *gin.Context) {
        controller.RetrieveTransactions(c, dbs, db, false)
    })
    router.GET("accounts/get", func(c *gin.Context) {
        controller.RetrieveAccounts(c, dbs, db, false)
    })
    //TODO investment transactions
    //TODO holdings
}
