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

func TestCreateHoldings_HoldingsCreated(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	var holdings []plaid.Holding
	holdings = append(holdings, plaid.Holding{AccountId: "1", Quantity: 100, InstitutionValue: 11.11})
	holdings = append(holdings, plaid.Holding{AccountId: "2", Quantity: 200, InstitutionValue: 22.22})
	// Handle mock route
	r.POST("/create-holdings", func(c *gin.Context) {
		plaidc.CreateHoldings(c, plaidinterface.MockPlaidClient{
			InvestHoldResp: plaid.InvestmentsHoldingsGetResponse{
				Holdings: holdings,
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
	req, _ := http.NewRequest("POST", "/create-holdings", strings.NewReader(form.Encode()))
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
func TestCreateHoldings_ProfileNotReceived(t *testing.T) {
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
func TestCreateHoldings_TokenNotReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-holdings", func(c *gin.Context) {
		plaidc.CreateHoldings(c,
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
func TestCreateHoldings_HoldingsNotCreated(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-holdings", func(c *gin.Context) {
		plaidc.CreateHoldings(c,
			plaidinterface.MockPlaidClient{Err: errors.New("Failed to get holdings")},
			postgresinterface.MockDB{},
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
	isTransactions := !strings.Contains(string(responseBody), "\"Failed to get holdings\"")
	// Assert
	tests.Equals(t, http.StatusInternalServerError, w.Code)
	tests.Equals(t, false, isTransactions)
}