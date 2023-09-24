package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func SetupPostgresRoutes(router *gin.RouterGroup, dbs DBHandler, db map[string]*sql.DB) {
    router.POST("profile/create", func(c *gin.Context) {
        CreateProfile(c, dbs, db, false)
    })
    router.POST("profile/get", func(c *gin.Context) {
        RetrieveProfile(c, dbs, db, false)
    })
    router.POST("expenses/create", func(c *gin.Context) {
        CreateExpenses(c, dbs, db, false)
    })
    router.GET("expenses/get", func(c *gin.Context) {
        RetrieveExpenses(c, dbs, db, false)
    })
    router.POST("expenses/update", func(c *gin.Context) {
        UpdateExpenses(c, dbs, db, false)
    })
    router.GET("transactions/get", func(c *gin.Context) {
        RetrieveTransactions(c, dbs, db, false)
    })
}
