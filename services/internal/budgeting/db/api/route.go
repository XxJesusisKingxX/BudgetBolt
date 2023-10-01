package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"

	"services/internal/utils/http"
)

func SetupBudgetRoutes(router *gin.RouterGroup, db *sql.DB) {
    httpClient := request.HTTPClient{}
    client := httpClient.NewHTTPClient("http://localhost:8000/api/")

    router.POST("expenses/create", func(c *gin.Context) {
        CreateExpenses(c, DB{}, db, client, false)
    })
    router.GET("expenses/get", func(c *gin.Context) {
        RetrieveExpenses(c, DB{}, db, client, false)
    })
    router.POST("expenses/update", func(c *gin.Context) {
        UpdateExpenses(c, DB{}, db, client, false)
    })
    router.POST("expenses/update/all", func(c *gin.Context) {
        UpdateAllExpenses(c, DB{}, db, client, false)
    })
    router.GET("incomes/get", func(c *gin.Context) {
        RetrieveIncome(c, DB{}, db, client, false)
    })
    router.POST("incomes/upsert", func(c *gin.Context) {
        UpsertIncome(c, DB{}, db, client, false)
    })
}
