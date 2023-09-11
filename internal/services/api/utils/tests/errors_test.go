package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	plaidinterface "services/api/plaid"
	api "services/api/utils"
	"services/utils/testing"
	"testing"
	"github.com/gin-gonic/gin"
	plaid "github.com/plaid/plaid-go/v12/plaid"
)

func TestRenderError200(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	err := errors.New("Failed to connect")
	plaidErr := plaid.PlaidError{
		ErrorCode:    "111",
		ErrorMessage: "No PLAID_CLIENT or PLAID_SECRET",
	}
	api.RenderError(c, err, plaidinterface.MockPlaidClient{PlaidError: plaidErr})

	// 200 error
	tests.Equals(t, http.StatusOK, c.Writer.Status())
}

func TestRenderError500(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	err := errors.New("Failed to connect")

	api.RenderError(c, err, plaidinterface.PlaidClient{})

	// 500 error
	tests.Equals(t, http.StatusInternalServerError, c.Writer.Status())
}