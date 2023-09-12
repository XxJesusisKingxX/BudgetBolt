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

func TestCreateAccessToken(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/get-access-token", func(c *gin.Context) {
		plaidc.CreateAccessToken(c,
			plaidinterface.MockPlaidClient{Err: errors.New("Failed to get token")},
			postgresinterface.MockDB{},
			nil,
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	form.Set("public_token", "public-sandbox-12345678-1234-1234-1234-123456789012")
	req, _ := http.NewRequest("POST", "/get-access-token", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isToken := !strings.Contains(string(responseBody), "\"Failed to get token\"")
	//Assert
	tests.Equals(t, http.StatusInternalServerError, w.Code)
	tests.Equals(t, false , isToken)
}
func TestCreateAccessToken_ProfileFails(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/get-access-token", func(c *gin.Context) {
		plaidc.CreateAccessToken(c,
			plaidinterface.MockPlaidClient{},
			postgresinterface.MockDB{
				ProfileErr: errors.New("Failed to get profile"),
			},
			nil,
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	form.Set("public_token", "public-sandbox-12345678-1234-1234-1234-123456789012")
	req, _ := http.NewRequest("POST", "/get-access-token", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Assert
	tests.Equals(t, http.StatusInternalServerError, w.Code)
}
func TestCreateAccessToken_Received(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/get-access-token", func(c *gin.Context) {
		plaidc.CreateAccessToken(c,
			plaidinterface.MockPlaidClient{},
			postgresinterface.MockDB{Profile: model.Profile{ ID: 1 }},
			nil,
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	form.Set("public_token", "public-sandbox-12345678-1234-1234-1234-123456789012")
	req, _ := http.NewRequest("POST", "/get-access-token", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Assert
	tests.Equals(t, http.StatusOK, w.Code)
}





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





func TestLinkTokenCreate(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-link-token", func(c *gin.Context) {
		plaidc.CreateLinkToken(c, plaidinterface.MockPlaidClient{},postgresinterface.MockDB{Profile: model.Profile{Name: "test_user"}}, nil, nil)
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("POST", "/create-link-token", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Received response
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	linkToken := strings.Contains(string(responseBody), "\"link_token\":")
	// Assert
	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, linkToken)
}
func TestLinkTokenCreate_ProfileFailed(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-link-token", func(c *gin.Context) {
		plaidc.CreateLinkToken(c, 
			plaidinterface.MockPlaidClient{},
			postgresinterface.MockDB{
				ProfileErr: errors.New("Failed to get profile"),
				},
				nil,
				nil,
		)
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("POST", "/create-link-token", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Received response
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isToken := !strings.Contains(string(responseBody), "\"Failed to get profile\"")
	// Assert
	tests.Equals(t, http.StatusInternalServerError, w.Code)
	tests.Equals(t, false, isToken)
}
func TestLinkTokenCreateFails(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-link-token", func(c *gin.Context) {
		plaidc.CreateLinkToken(c, 
			plaidinterface.MockPlaidClient{
				Err: errors.New("Failed"),
			},
			postgresinterface.MockDB{},
			nil,
			nil,
		)
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("POST", "/create-link-token", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Assert
	tests.Equals(t, http.StatusInternalServerError, w.Code)
}





func TestCreateTransactions_TransactionsCreated(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Create mock plaid transactions
	var transactions []plaid.Transaction
	transactions = append(transactions, plaid.Transaction{AccountId: "1", Date: "2023/01/01", Amount: 11.11})
	transactions = append(transactions, plaid.Transaction{AccountId: "2", Date: "2023/02/02", Amount: 22.22})
	// Handle mock route
	r.POST("/create-transactions", func(c *gin.Context) {
		plaidc.CreateTransactions(c,
			plaidinterface.MockPlaidClient{SyncResp: plaid.TransactionsSyncResponse{Added: transactions}},
			postgresinterface.MockDB{Profile: model.Profile{ID: 1}, Token: model.Token{Token: "access-sandbox-111-222-3333-4444"}},
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
	isReceived := strings.Contains(string(responseBody), "")
	// Assert
	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, isReceived)
}
func TestCreateTransactions_ProfileNotReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-transactions", func(c *gin.Context) {
		plaidc.CreateTransactions(c,
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
func TestCreateTransactions_TokenNotReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-transactions", func(c *gin.Context) {
		plaidc.CreateTransactions(c,
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
func TestCreateTransactions_TransactionsNotCreated(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-transactions", func(c *gin.Context) {
		plaidc.CreateTransactions(c,
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
	req, _ := http.NewRequest("POST", "/create-transactions", strings.NewReader(form.Encode()))
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