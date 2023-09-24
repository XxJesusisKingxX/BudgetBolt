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

func TestCreateAccounts(t *testing.T) {
	// Define a slice of test cases.
	testCases := []struct {
		TestName     string
		PlaidClient  api.MockPlaidClient
		MockDB       apiSql.MockDB
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
			MockDB: apiSql.MockDB{
				Profile: model.Profile{ID: 1}, 
				Token: model.Token{Token: "access-sandbox-11-222-33"},
			},
			ExpectedCode: http.StatusOK,
		},
		{
			TestName:   "ProfileNotReceived",
			PlaidClient: api.MockPlaidClient{},
			MockDB: apiSql.MockDB{
				ProfileErr: errors.New("Failed to get profile id"), 
			},
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Failed to get profile id",
		},
		{
			TestName:     "TokenNotReceived",
			PlaidClient:  api.MockPlaidClient{},
			MockDB: apiSql.MockDB{
				TokenErr: errors.New("Failed to get token"), 
			},
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Failed to get token",
		},
		{
			TestName:     "AccountsNotCreated",
			PlaidClient:  api.MockPlaidClient{Err:errors.New("Failed to get accounts")},
			MockDB: apiSql.MockDB{},
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Failed to get accounts",
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			// Create mock engine
			gin.SetMode(gin.TestMode)
			r := gin.Default()

			// Handle mock route
			r.POST("/create-accounts", func(c *gin.Context) {
				api.CreateAccounts(c, tc.PlaidClient, tc.MockDB, nil, nil, true)
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