package api

import (
	"net/http"
	plaidinterface "services/api/plaid"
	"github.com/gin-gonic/gin"
)

func RenderError(c *gin.Context, originalErr error, ps plaidinterface.Plaid) {
	plaidError, err := ps.ToPlaidError(originalErr)
	if err == nil {
		// Return 200 and allow the front end to render the error.
		c.JSON(http.StatusOK, gin.H{"error": plaidError})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": originalErr.Error()})
}