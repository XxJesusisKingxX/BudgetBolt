package main

import (
	"errors"
	"os/exec"
	"time"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"budgetbolt/src/services/tests"
	browser "budgetbolt/src/services/tests/browser"

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
func TestGetAccessToken(t *testing.T) {
	// Start up mock server
	exec.Command("bash", "-c", "cd ../../tests/server && ./start.sh").Start()
	isRunning := checkMockServer()
	for isRunning != true {
		time.Sleep(1000 * time.Millisecond)
		isRunning = checkMockServer()
	}
	// Start test environment
	browser.BrowserTestSetup("http://localhost:8080/", true, browser.TestPlaidWorkFlow)
	// Stop mock server
	exec.Command("bash", "-c", "cd ../../tests/server && ./stop.sh").Run()

	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/get-access-token", func(c *gin.Context) {
		getAccessToken(c, PlaidClient{}, true)
	})
	// Create request
	form := url.Values{}
	wd, _ := os.Getwd()
	path := strings.Replace(wd, "\\api\\plaid", "\\tests\\server\\public_token.txt", 1)
	content, _ := os.ReadFile(path)
	form.Set("public_token", string(content))
	req, _ := http.NewRequest("POST", "/get-access-token", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	accessToken := strings.Contains(string(responseBody), "\"access_token\":")
	itemId := strings.Contains(string(responseBody), "\"item_id\":")

	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, accessToken)
	tests.Equals(t, true, itemId)
}
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
func TestAccounts(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/get-accounts", func(c *gin.Context) {
		accounts(c, PlaidClient{}, true)
	})
	// Create request
	form := url.Values{}
	wd, _ := os.Getwd()
	path := strings.Replace(wd, "\\api\\plaid", "\\tests\\server\\access_token.txt", 1)
	content, _ := os.ReadFile(path)
	form.Set("access_token", string(content))
	req, _ := http.NewRequest("POST", "/get-accounts", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	accounts := strings.Contains(string(responseBody), "\"accounts\":")

	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, accounts)
}
func TestAccountsFails(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/get-accounts", func(c *gin.Context) {
		accounts(c, MockPlaidClient{Err: errors.New("Failed")}, true)
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("GET", "/get-accounts", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	tests.Equals(t, http.StatusInternalServerError, w.Code)
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
		accounts(c, MockPlaidClient{Err: errors.New("Failed")}, true)
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
func TestTransactions(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/get-transactions", func(c *gin.Context) {
		transactions(c, PlaidClient{}, true)
	})
	// Create request
	form := url.Values{}
	wd, _ := os.Getwd()
	path := strings.Replace(wd, "\\api\\plaid", "\\tests\\server\\access_token.txt", 1)
	content, _ := os.ReadFile(path)
	form.Set("access_token", string(content))
	req, _ := http.NewRequest("POST", "/get-transactions", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	trans := strings.Contains(string(responseBody), "\"transactions\":")

	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, trans)
}
func TestTransactionsFails(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/get-transactions", func(c *gin.Context) {
		transactions(c, MockPlaidClient{ Err: errors.New("Failed") }, true)
	})
	// Create request
	form := url.Values{}
	wd, _ := os.Getwd()
	path := strings.Replace(wd, "\\api\\plaid", "\\tests\\server\\access_token.txt", 1)
	content, _ := os.ReadFile(path)
	form.Set("access_token", string(content))
	req, _ := http.NewRequest("POST", "/get-transactions", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	tests.Equals(t, http.StatusInternalServerError, w.Code)
}
