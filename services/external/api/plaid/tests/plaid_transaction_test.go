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

	"services/internal/utils/http"
	"services/external/api/plaid"
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
		Response        map[string]request.MockResponse
		ExpectedCode   int
		ExpectedBody   string
	}{
		{
			TestName: "TransactionsCreated",
			Transactions: []plaid.Transaction{
				{AccountId: "1", Date: "2023/01/01", Amount: 11.11},
				{AccountId: "2", Date: "2023/02/02", Amount: 22.22},
			},
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
				"token/get?uid=": {
					Code: http.StatusOK,
				},
			},
			ExpectedCode:   http.StatusOK,
		},
		{
			TestName:       "ProfileNotReceived",
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusInternalServerError,
				},
			},
			ExpectedCode:   http.StatusInternalServerError,
		},
		{
			TestName:       "TokenNotReceived",
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
				"token/get?uid=": {
					Code: http.StatusInternalServerError,
				},
			},
			ExpectedCode:   http.StatusInternalServerError,
		},
		{
			TestName:       "TransactionsNotCreated",
			TransactionsErr: errors.New("Failed to create transactions"),
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
				"token/get?uid=": {
					Code: http.StatusOK,
				},
			},
			ExpectedCode:   http.StatusInternalServerError,
			ExpectedBody:   "Failed to create transactions",
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			// Create mock engine
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			// Mock requests
			mockClient := request.MockHTTPClient{}
			mockClient.Responses = tc.Response
			// Handle mock route
			r.POST("/create-transactions", func(c *gin.Context) {
				api.CreateTransactions(c,
					api.MockPlaidClient{SyncResp: plaid.TransactionsSyncResponse{Added: tc.Transactions}, Err: tc.TransactionsErr},
					nil,
					mockClient,
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