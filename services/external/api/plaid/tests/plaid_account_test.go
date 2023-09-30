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
	"services/internal/utils/http"
	"services/internal/utils/testing"
)

func TestCreateAccounts(t *testing.T) {
	// Define a slice of test cases.
	testCases := []struct {
		TestName     string
		PlaidClient  api.MockPlaidClient
		Response     map[string]request.MockResponse
		ExpectedCode int
		ExpectedBody string
	}{
		{
			TestName: "AccountsCreated",
			PlaidClient: api.MockPlaidClient{
				Accounts: []plaid.AccountBase{
					{
						Name: "Test Account",
					},
				},
			},
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
				"token/get?uid=": {
					Code: http.StatusOK,
				},
			},
			ExpectedCode: http.StatusOK,
		},
		{
			TestName:   "ProfileNotReceived",
			PlaidClient: api.MockPlaidClient{},
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusInternalServerError,
				},
			},
			ExpectedCode: http.StatusInternalServerError,
		},
		{
			TestName:     "TokenNotReceived",
			PlaidClient:  api.MockPlaidClient{},
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
				"token/get": {
					Code: http.StatusInternalServerError,
				},
			},
			ExpectedCode: http.StatusInternalServerError,
		},
		{
			TestName:     "AccountsNotCreated",
			PlaidClient:  api.MockPlaidClient{Err:errors.New("Failed to get accounts")},
			ExpectedCode: http.StatusInternalServerError,
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
			r.POST("/create-accounts", func(c *gin.Context) {
				api.CreateAccounts(c, tc.PlaidClient, nil, mockClient, true)
			})

			// Create request
			form := url.Values{}
			form.Set("username", "test_user")
			req, _ := http.NewRequest("POST", "/create-accounts", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// Recieve response
			var body map[string]string
			responseBody, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(responseBody, &body)
			defer w.Result().Body.Close()

			// Assert
			tests.Equals(t, tc.ExpectedCode, w.Code)
			tests.Equals(t, tc.ExpectedBody, body["error"])
		})
	}
}