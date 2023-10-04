package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"

    "services/internal/utils/http"
)

func SetupTransactionRoutes(router *gin.RouterGroup, db *sql.DB ) {
    httpClient := request.HTTPClient{} 
    client := httpClient.NewHTTPClient("http://localhost:8000/api/")
    
    router.DELETE("transactions/pending/remove", func(c *gin.Context) {
        DeletePendingTransactions(c, DB{}, db, client, false)
    })
    router.GET("transactions/get", func(c *gin.Context) {
        RetrieveTransactions(c, DB{}, db, client, false)
    })
    router.POST("transactions/store", func(c *gin.Context) {
        StoreTransactions(c, DB{}, db, false)
    })
    router.POST("accounts/store", func(c *gin.Context) {
        StoreAccounts(c, DB{}, db, false)
    })
}
