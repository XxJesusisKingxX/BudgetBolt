package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v12/plaid"

	"services/external/api/plaid"
	apiSql "services/internal/api/sql"
	"services/internal/user_managment/db/model"
	"services/internal/utils/testing"
)
func TestCreateTransactions(t *testing.T) {
	// Define a slice of test cases.
	testCases := []struct {
		TestName       string
		Transactions   []plaid.Transaction
		ProfileErr     error
		TokenErr       error
		TransactionsErr error
		ExpectedCode   int
		ExpectedBody   string
	}{
		{
			TestName: "TransactionsCreated",
			Transactions: []plaid.Transaction{
				{AccountId: "1", Date: "2023/01/01", Amount: 11.11},
				{AccountId: "2", Date: "2023/02/02", Amount: 22.22},
			},
			ExpectedCode:   http.StatusOK,
			ExpectedBody:   "",
		},
		{
			TestName:       "ProfileNotReceived",
			ProfileErr:     errors.New("Failed to get profile id"),
			ExpectedCode:   http.StatusInternalServerError,
			ExpectedBody:   "Failed to get profile id",
		},
		{
			TestName:       "TokenNotReceived",
			TokenErr:       errors.New("Failed to get access token"),
			ExpectedCode:   http.StatusInternalServerError,
			ExpectedBody:   "Failed to get access token",
		},
		{
			TestName:       "TransactionsNotCreated",
			TransactionsErr: errors.New("Failed to get transactions"),
			ExpectedCode:   http.StatusInternalServerError,
			ExpectedBody:   "Failed to get transactions",
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			// Create mock engine
			gin.SetMode(gin.TestMode)
			r := gin.Default()

			// Handle mock route
			r.POST("/create-transactions", func(c *gin.Context) {
				api.CreateTransactions(c,
					api.MockPlaidClient{SyncResp: plaid.TransactionsSyncResponse{Added: tc.Transactions}, Err: tc.TransactionsErr},
					apiSql.MockDB{Profile: model.Profile{ID: 1}, Token: model.Token{Token: "access-sandbox-111-222-3333-4444"}, ProfileErr: tc.ProfileErr, TokenErr: tc.TokenErr},
					nil,
					nil,
					true,
				)
			})

			// Create request
			form := url.Values{}
			form.Set("username", "test_user")
			req, _ := http.NewRequest("POST", "/create-transactions", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// Receive response
			var body map[string]string
			responseBody, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(responseBody, &body)
			defer w.Result().Body.Close()

			// Assert
			tests.Equals(t, tc.ExpectedCode, w.Code)
			if value, exists := body["error"]; exists {
				tests.Equals(t, tc.ExpectedBody, value)
			} else {
				tests.Equals(t, tc.ExpectedBody, body["transactions"])
			}
		})
	}
}