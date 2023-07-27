package main

import (
	"errors"
	"os/exec"

	// "time"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"budgetbolt/src/services/databases/postgresql/controller"
	"budgetbolt/src/services/databases/postgresql/model"
	"budgetbolt/src/services/tests"

	// browser "budgetbolt/src/services/tests/browser"

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

	renderError(c, err, MockPlaidClient{PlaidError: plaidErr})

	// 200 error
	tests.Equals(t, http.StatusOK, c.Writer.Status())
}

func TestRenderError500(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	err := errors.New("Failed to connect")

	renderError(c, err, PlaidClient{})

	// 500 error
	tests.Equals(t, http.StatusInternalServerError, c.Writer.Status())
}
func checkMockServer() bool {
	c, _ := exec.Command("tasklist", "/FI", "IMAGENAME eq plaid_oauth*").Output()
	isRunning := !strings.Contains(string(c), "INFO:")
	return isRunning
}
// func TestGetAccessToken(t *testing.T) {
// 	// Start up mock server
// 	exec.Command("bash", "-c", "cd ../../tests/server && ./start.sh").Start()
// 	isRunning := checkMockServer()
// 	for isRunning != true {
// 		time.Sleep(1000 * time.Millisecond)
// 		isRunning = checkMockServer()
// 	}
// 	// Start test environment
// 	browser.BrowserTestSetup("http://localhost:8080/", true, browser.TestPlaidWorkFlow)
// 	// Stop mock server
// 	exec.Command("bash", "-c", "cd ../../tests/server && ./stop.sh").Run()

// 	// Create mock engine
// 	gin.SetMode(gin.TestMode)
// 	r := gin.Default()
// 	// Handle mock route
// 	r.POST("/get-access-token", func(c *gin.Context) {
// 		getAccessToken(c, PlaidClient{}, true)
// 	})
// 	// Create request
// 	form := url.Values{}
// 	wd, _ := os.Getwd()
// 	path := strings.Replace(wd, "\\api\\plaid", "\\tests\\server\\public_token.txt", 1)
// 	content, _ := os.ReadFile(path)
// 	form.Set("public_token", string(content))
// 	req, _ := http.NewRequest("POST", "/get-access-token", strings.NewReader(form.Encode()))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	responseBody, _ := io.ReadAll(w.Result().Body)
// 	defer w.Result().Body.Close()
// 	accessToken := strings.Contains(string(responseBody), "\"access_token\":")
// 	itemId := strings.Contains(string(responseBody), "\"item_id\":")

// 	tests.Equals(t, http.StatusOK, w.Code)
// 	tests.Equals(t, true, accessToken)
// 	tests.Equals(t, true, itemId)
// }
func TestGetAccessTokenFails(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/get-access-token", func(c *gin.Context) {
		getAccessToken(c, MockPlaidClient{Err: errors.New("Failed")}, true)
	})
	// Create request
	form := url.Values{}
	form.Set("public_token", "public-sandbox-12345678-1234-1234-1234-123456789012")
	req, _ := http.NewRequest("POST", "/get-access-token", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	tests.Equals(t, http.StatusInternalServerError, w.Code)
}
func TestGetAccounts_AccountsRecieved(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/get-accounts", func(c *gin.Context) {
		retrieveAccounts(c, 
			controller.MockDB{ 
				Profile: model.Profile{ ID: 1 }, 
				Account: []model.Account{
					{
						Name: "Test1",
						Balance: 11.11,
					},
				}})
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("GET", "/get-accounts", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Receive response
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isAccounts := strings.Contains(string(responseBody), "\"Test1\"")
	// Assert
	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, isAccounts)
}
func TestGetAccounts_AccountsNotReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/get-accounts", func(c *gin.Context) {
		retrieveAccounts(c, controller.MockDB{AccountErr: errors.New("Failed to get accounts")})
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("GET", "/get-accounts", strings.NewReader(form.Encode()))
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
func TestGetAccounts_ProfileNotReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/get-accounts", func(c *gin.Context) {
		retrieveAccounts(c, controller.MockDB{ ProfileErr: errors.New("Failed to get profile id") })
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("GET", "/get-accounts", strings.NewReader(form.Encode()))
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
func TestCreateAccounts_AccountsCreated(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-accounts", func(c *gin.Context) {
		createAccounts(c, 
			MockPlaidClient{ 
				Accounts: []plaid.AccountBase{
					{
						Name: "Test Account",
					},
			}},
			controller.MockDB{ Profile: model.Profile{ ID: 1 }, Token: model.Token{ Token: "access-sandbox-11-222-33"}},
			true)
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
	isAccounts := strings.Contains(string(responseBody), "\"Test Account\"")
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
		createAccounts(c, 
			MockPlaidClient{},
			controller.MockDB{ ProfileErr: errors.New("Failed to get profile id")},
			true,
		)})
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
		createAccounts(c, 
			MockPlaidClient{},
			controller.MockDB{ TokenErr: errors.New("Failed to get access token")},
			true,
		)})
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
		createAccounts(c, 
			MockPlaidClient{ Err: errors.New("Failed to get accounts") },
			controller.MockDB{},
			true,
		)})
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
func TestInvestments(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/get-investments", func(c *gin.Context) {
		investmentTransactions(c, PlaidClient{}, true)
	})
	// Create request
	form := url.Values{}
	wd, _ := os.Getwd()
	path := strings.Replace(wd, "\\api\\plaid", "\\tests\\server\\access_token.txt", 1)
	content, _ := os.ReadFile(path)
	form.Set("access_token", string(content))
	req, _ := http.NewRequest("POST", "/get-investments", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	invest := strings.Contains(string(responseBody), "\"NO_INVESTMENT_ACCOUNTS\"") //TODO get test accounts to models investments holdings

	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, invest)
}
func TestInvestmentsFails(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/get-investments", func(c *gin.Context) {
		investmentTransactions(c, MockPlaidClient{Err: errors.New("Failed")}, true)
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("GET", "/get-investments", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	tests.Equals(t, http.StatusInternalServerError, w.Code)
}
func TestHoldings(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/get-holdings", func(c *gin.Context) {
		holdings(c, PlaidClient{}, true)
	})
	// Create request
	form := url.Values{}
	wd, _ := os.Getwd()
	path := strings.Replace(wd, "\\api\\plaid", "\\tests\\server\\access_token.txt", 1)
	content, _ := os.ReadFile(path)
	form.Set("access_token", string(content))
	req, _ := http.NewRequest("POST", "/get-holdings", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	holdings := strings.Contains(string(responseBody), "\"NO_INVESTMENT_ACCOUNTS\"") //TODO get test accounts to models investments holdings

	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, holdings)
}
func TestHoldingsFails(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/get-holdings", func(c *gin.Context) {
		holdings(c, MockPlaidClient{Err: errors.New("Failed")}, true)
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("GET", "/get-holdings", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	tests.Equals(t, http.StatusInternalServerError, w.Code)
}

func TestInfo(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/get-info", func(c *gin.Context) {
		info(c)
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("GET", "/get-info", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	itemId := strings.Contains(string(responseBody), "\"item_id\"")
	accessToken := strings.Contains(string(responseBody), "\"access_token\"")
	products := strings.Contains(string(responseBody), "\"products\"")

	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, itemId)
	tests.Equals(t, true, accessToken)
	tests.Equals(t, true, products)
}
func TestLinkTokenCreate(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-link-token", func(c *gin.Context) {
		createLinkToken(c, PlaidClient{})
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("POST", "/create-link-token", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	linkToken := strings.Contains(string(responseBody), "\"link_token\":")

	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, linkToken)
}
func TestLinkTokenCreateFails(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-link-token", func(c *gin.Context) {
		createLinkToken(c, MockPlaidClient{ Err: errors.New("Failed") })
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("POST", "/create-link-token", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	tests.Equals(t, http.StatusInternalServerError, w.Code)
}
func TestGetTransactions_TransactionsReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Create mock transactions
	var transactions []model.Transaction
	transactions = append(transactions, model.Transaction{
		ID: 1001,
		Date: "2023/01/01",
		Amount: 12.34,
		Method: "",
		From: "Test Account",
		Vendor: "Test",
		Description: "A test case was received",
		ProfileID: 1,
	})
	// Handle mock route
	r.GET("/get-transactions", func(c *gin.Context) {
		retrieveTransactions(c, controller.MockDB{ Profile: model.Profile{ ID: 1 }, Transaction: transactions })
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("GET", "/get-transactions", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isReceived := strings.Contains(string(responseBody), "\"A test case was received\"")

	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, isReceived)
}
func TestGetTransactions_ProfileIdNotReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/get-transactions", func(c *gin.Context) {
		retrieveTransactions(c, controller.MockDB{ ProfileErr: errors.New("Failed to get profile id")})
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("GET", "/get-transactions", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Receive response
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isProfileId := !strings.Contains(string(responseBody), "\"Failed to get profile id\"")

	tests.Equals(t, http.StatusInternalServerError, w.Code)
	tests.Equals(t, false, isProfileId)
}
func TestGetTransactions_TransactionsNotReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/get-transactions", func(c *gin.Context) {
		retrieveTransactions(c, controller.MockDB{ TransactionErr: errors.New("Failed to get transactions")})
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("GET", "/get-transactions", strings.NewReader(form.Encode()))
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
		createTransactions(c, 
			MockPlaidClient{ SyncResp: plaid.TransactionsSyncResponse{ Added: transactions }},
			controller.MockDB{ Profile: model.Profile{ ID: 1 }, Token: model.Token{ Token:"access-sandbox-111-222-3333-4444" }},
			true,
		)})
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

	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, isReceived)
}
func TestCreateTransactions_ProfileNotReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/create-transactions", func(c *gin.Context) {
		createTransactions(c, 
			MockPlaidClient{},
			controller.MockDB{ ProfileErr: errors.New("Failed to get profile id")},
			true,
		)})
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
		createTransactions(c, 
			MockPlaidClient{},
			controller.MockDB{ TokenErr: errors.New("Failed to get access token")},
			true,
		)})
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
		createTransactions(c, 
			MockPlaidClient{ Err: errors.New("Failed to get transactions") },
			controller.MockDB{},
			true,
		)})
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