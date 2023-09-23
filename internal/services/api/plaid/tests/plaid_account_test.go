package main

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	plaidinterface "services/api/plaid"
	plaidc "services/api/plaid/controller"
	"services/api/postgres"
	"services/db/postgresql/model"
	"services/utils/testing"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v12/plaid"
)

func TestCreateAccounts_AccountsCreated(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-accounts", func(c *gin.Context) {
		plaidc.CreateAccounts(c,
			plaidinterface.MockPlaidClient{
				Accounts: []plaid.AccountBase{
					{
						Name: "Test Account",
					},
				}},
			postgresinterface.MockDB{Profile: model.Profile{ID: 1}, Token: model.Token{Token: "access-sandbox-11-222-33"}},
			nil,
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	form.Set("username", "test_user")
	req, _ := http.NewRequest("POST", "/create-accounts", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Receive response
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isAccounts := strings.Contains(string(responseBody), "")
	// Assert
	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, isAccounts)
}
func TestCreateAccounts_ProfileNotReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-transactions", func(c *gin.Context) {
		plaidc.CreateAccounts(c,
			plaidinterface.MockPlaidClient{},
			postgresinterface.MockDB{ProfileErr: errors.New("Failed to get profile id")},
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
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isProfileId := !strings.Contains(string(responseBody), "\"Failed to get profile id\"")
	// Assert
	tests.Equals(t, http.StatusInternalServerError, w.Code)
	tests.Equals(t, false, isProfileId)
}
func TestCreateAccounts_TokenNotReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-transactions", func(c *gin.Context) {
		plaidc.CreateAccounts(c,
			plaidinterface.MockPlaidClient{},
			postgresinterface.MockDB{TokenErr: errors.New("Failed to get access token")},
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
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isAccessToken := !strings.Contains(string(responseBody), "\"Failed to get access token\"")
	// Assert
	tests.Equals(t, http.StatusInternalServerError, w.Code)
	tests.Equals(t, false, isAccessToken)
}
func TestCreateAccounts_AccountsNotCreated(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-transactions", func(c *gin.Context) {
		plaidc.CreateAccounts(c,
			plaidinterface.MockPlaidClient{Err: errors.New("Failed to get accounts")},
			postgresinterface.MockDB{},
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
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isAccounts := !strings.Contains(string(responseBody), "\"Failed to get accounts\"")
	// Assert
	tests.Equals(t, http.StatusInternalServerError, w.Code)
	tests.Equals(t, false, isAccounts)
}