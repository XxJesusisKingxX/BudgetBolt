package main

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"services/api/postgres"
	postgresc "services/api/postgres/controller"
	"services/db/postgresql/model"
	"services/utils/testing"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
)

func TestGetTransactions_TransactionsReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Create mock transactions
	var transactions []model.Transaction
	transactions = append(transactions, model.Transaction{
		ID:          "1",
		Date:        "2023/01/01",
		Amount:      12.34,
		Method:      "",
		From:        "Test Account",
		Vendor:      "Test",
		Description: "A test case was received",
		ProfileID:   1,
	})
	// Handle mock route
	r.GET("/get-transactions", func(c *gin.Context) {
		postgresc.RetrieveTransactions(c,
			postgresinterface.MockDB{
				Profile: model.Profile{ID: 1}, 
				Transaction: transactions,
			},
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("GET", "/get-transactions", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Received response
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isReceived := strings.Contains(string(responseBody), "\"A test case was received\"")
	// Assert
	tests.Equals(t, http.StatusOK, w.Code)
	tests.Equals(t, true, isReceived)
}
func TestGetTransactions_ProfileIdNotReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/get-transactions", func(c *gin.Context) {
		postgresc.RetrieveTransactions(c,
			postgresinterface.MockDB{
				ProfileErr: errors.New("Failed to get profile id"),
			},
			nil,
			true,
		)
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
	// Assert
	tests.Equals(t, http.StatusInternalServerError, w.Code)
	tests.Equals(t, false, isProfileId)
}
func TestGetTransactions_TransactionsNotReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/get-transactions", func(c *gin.Context) {
		postgresc.RetrieveTransactions(c, 
			postgresinterface.MockDB{
				TransactionErr: errors.New("Failed to get transactions"),
			},
			nil,
			true,
		)
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