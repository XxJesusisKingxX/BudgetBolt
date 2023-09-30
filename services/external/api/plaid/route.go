package api

import (
	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v12/plaid"

    "services/internal/utils/http"
)

func SetupPlaidRoutes(router *gin.RouterGroup, ps Plaid, plaid *plaid.APIClient) {
    httpClient := request.HTTPClient{} 
    client := httpClient.NewHTTPClient("http://localhost:8000/api/")

    router.POST("link_token/create", func(c *gin.Context) {
        CreateLinkToken(c, ps, client, plaid)
    })
    router.POST("access_token/create", func(c *gin.Context) {
        CreateAccessToken(c, ps, plaid, client, false)
    })
    router.POST("accounts/create", func(c *gin.Context) {
        CreateAccounts(c, ps, plaid, client, false)
    })
    router.POST("transactions/create", func(c *gin.Context) {
        CreateTransactions(c, ps, plaid, client, false)
    })
}
