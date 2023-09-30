package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"

	"services/internal/transaction_history/db/api"
	"services/internal/utils/testing"
)

func TestStoreAccount(t *testing.T) {
    testCases := []struct {
        TestName       string
        Form           url.Values
        ExpectedCode   int
    }{
        {
            TestName: "IDInvalid",
            Form: url.Values{
                "id" : {"invalid"},
            },
            ExpectedCode:  http.StatusInternalServerError,
        },
        {
            TestName: "StoreFailed",
            Form: url.Values{
                "id" : {"12"},
                "accounts": {""},
            },
            ExpectedCode:  http.StatusInternalServerError,
        },
        {
            TestName: "StoreSuccess",
            Form: url.Values{
                "id" : {"12"},
                "accounts": {`[{"account_id": "account123"},{"account_id": "account456"}]`},
            },
            ExpectedCode:  http.StatusOK,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.TestName, func(t *testing.T) {
            // Create mock engine
            gin.SetMode(gin.TestMode)
            r := gin.Default()
            // Handle mock route
            r.POST("/store-accounts", func(c *gin.Context) {
                api.StoreAccounts(c,
                    api.MockDB{
                    },
                    nil,
                    true,
                )
            })
            // Create request
            req, _ := http.NewRequest("POST", "/store-accounts", strings.NewReader(tc.Form.Encode()))
            req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
            w := httptest.NewRecorder()
            r.ServeHTTP(w, req)

            // Assert
            tests.Equals(t, tc.ExpectedCode, w.Code)
        })
    }
}