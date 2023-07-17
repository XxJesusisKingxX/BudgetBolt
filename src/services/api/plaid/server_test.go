package main

import (
	"errors"
	"os/exec"
	"time"

	// "fmt"
	"io/ioutil"
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
    	ErrorCode: "111",
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
	c, _ := exec.Command("tasklist", "/FI", "IMAGENAME eq plaid_oauth*" ).Output()
	isRunning := !strings.Contains(string(c), "INFO:")
	return isRunning
}
func TestGetAccessToken(t *testing.T) {
	// Start up mock server
	exec.Command("bash","-c", "cd ../../tests/server && ./start.sh").Start()
	isRunning := checkMockServer()
	for isRunning != true {
		time.Sleep(1000 * time.Millisecond)
		isRunning = checkMockServer()
	}
	// Start test environment
	browser.BrowserTestSetup("http://localhost:8080/", true, browser.TestPlaidWorkFlow)
	// Stop mock server
	exec.Command("bash","-c", "cd ../../tests/server && ./stop.sh").Run()

	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	wd, _ := os.Getwd()
	path := strings.Replace(wd, "\\api\\plaid", "\\tests\\server\\public_token.txt", 1)
	content, _ := os.ReadFile(path)
	// Handle mock route
	r.POST("/get-access-token", func(c *gin.Context) {
		getAccessToken(c, PlaidClient{})
	})
	// Create request
	form := url.Values{}
	form.Set("public_token", string(content))
	req, _ := http.NewRequest("POST", "/get-access-token", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseBody, _ := ioutil.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	accessToken := strings.Contains(string(responseBody), "\"access_token\":")
	itemId := strings.Contains(string(responseBody), "\"item_id\":")

	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, accessToken)
	tests.Equals(t, true, itemId)
}
