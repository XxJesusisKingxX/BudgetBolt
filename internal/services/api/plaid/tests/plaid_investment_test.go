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

func TestInvestments_TransactionsCreated(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	var transactions []plaid.InvestmentTransaction
	transactions = append(transactions, plaid.InvestmentTransaction{AccountId: "1", Date: "2023/01/01", Amount: 11.11})
	transactions = append(transactions, plaid.InvestmentTransaction{AccountId: "2", Date: "2023/02/02", Amount: 22.22})
	// Handle mock route
	r.POST("/get-investments", func(c *gin.Context) {
		plaidc.CreateInvestmentTransactions(c, plaidinterface.MockPlaidClient{
			InvestTransResp: plaid.InvestmentsTransactionsGetResponse{
				InvestmentTransactions: transactions,
			},
		}, 
		postgresinterface.MockDB{
			Profile: model.Profile{ID: 1},
			Token: model.Token{Token: "access-sandbox-111-222-3333-4444"},
		},
		nil,
		nil,
		true,
	)
	})
	// Create request
	form := url.Values{}
	form.Set("username", "test_user")
	req, _ := http.NewRequest("POST", "/get-investments", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Received response
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	invest := strings.Contains(string(responseBody), "") //TODO get test accounts to models investments holdings
	// Assert
	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, invest)
}
func TestCreateInvestments_ProfileNotReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-holdings", func(c *gin.Context) {
		plaidc.CreateInvestmentTransactions(c,
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
	req, _ := http.NewRequest("POST", "/create-holdings", strings.NewReader(form.Encode()))
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
func TestCreateInvestments_TokenNotReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-holdings", func(c *gin.Context) {
		plaidc.CreateInvestmentTransactions(c,
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
	req, _ := http.NewRequest("POST", "/create-holdings", strings.NewReader(form.Encode()))
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
func TestCreateInvestments_TransactionsNotCreated(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-investments_transactions", func(c *gin.Context) {
		plaidc.CreateInvestmentTransactions(c,
			plaidinterface.MockPlaidClient{Err: errors.New("Failed to get transactions")},
			postgresinterface.MockDB{},
			nil,
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	form.Set("username", "test_user")
	req, _ := http.NewRequest("POST", "/create-investments_transactions", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Receive response
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isTransactions := !strings.Contains(string(responseBody), "\"Failed to get transactions\"")
	// Assert
	tests.Equals(t, http.StatusInternalServerError, w.Code)
	tests.Equals(t, false, isTransactions)
}
