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
	"golang.org/x/crypto/bcrypt"
)

func TestGetAccounts_AccountsRecieved(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/get-accounts", func(c *gin.Context) {
		postgresc.RetrieveAccounts(c,
			postgresinterface.MockDB{
				Profile: model.Profile{ID: 1},
				Account: []model.Account{
					{
						Name:    "Test1",
						Balance: 11.11,
					},
				}},
			nil,
			true,
		)
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
		postgresc.RetrieveAccounts(c,
			postgresinterface.MockDB{
				AccountErr: errors.New("Failed to get accounts"),
			},
			nil,
			true,
		)
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
		postgresc.RetrieveAccounts(c, 
			postgresinterface.MockDB{
				ProfileErr: errors.New("Failed to get profile id"),
			},
			nil,
			true,
		)
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





func TestGetTransactions_TransactionsReceived(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Create mock transactions
	var transactions []model.Transaction
	transactions = append(transactions, model.Transaction{
		ID:          1001,
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





func TestCreateProfile_ProfileNameTaken(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/api/profile/create", func(c *gin.Context) {
		postgresc.CreateProfile(c,
			postgresinterface.MockDB{
				Profile: model.Profile{ID: 1},
			},
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("POST", "/api/profile/create", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Assert
	tests.Equals(t, http.StatusConflict, w.Code)
}
func TestCreateProfile_PasswordTooLong(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/api/profile/create", func(c *gin.Context) {
		postgresc.CreateProfile(c,
			postgresinterface.MockDB{
				Profile: model.Profile{ID: 0},
			},
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	form.Set("username", "test_user")
	form.Set("password", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz")
	req, _ := http.NewRequest("POST", "/api/profile/create", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Receive response
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isLong := strings.Contains(string(responseBody), "password length exceeds 72 bytes")
	// Assert
	tests.Equals(t, http.StatusInternalServerError, w.Code)
	tests.Equals(t, true, isLong)
}
func TestCreateProfile_ProfileNotCreated(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/api/profile/create", func(c *gin.Context) {
		postgresc.CreateProfile(c,
			postgresinterface.MockDB{
				Profile:    model.Profile{ID: 0},
				ProfileErr: errors.New("Failed to create profile"),
			},
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	form.Set("username", "test_user")
	form.Set("password", "abcdefghijklmnopqrstuvwxyz")
	req, _ := http.NewRequest("POST", "/api/profile/create", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Receive response
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isCreated := !strings.Contains(string(responseBody), "\"Failed to create profile\"")
	// Assert
	tests.Equals(t, http.StatusInternalServerError, w.Code)
	tests.Equals(t, false, isCreated)
}
func TestCreateProfile_ProfileCreated(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/api/profile/create", func(c *gin.Context) {
		postgresc.CreateProfile(c,
			postgresinterface.MockDB{
				Profile: model.Profile{ID: 0},
			},
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	form.Set("username", "test_user")
	form.Set("password", "abcdefghijklmnopqrstuvwxyz")
	req, _ := http.NewRequest("POST", "/api/profile/create", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Assert
	tests.Equals(t, http.StatusOK, w.Code)
}
func TestRetrieveProfile_NotExist(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/api/profile/get", func(c *gin.Context) {
		postgresc.RetrieveProfile(c,
			postgresinterface.MockDB{
				Profile: model.Profile{
					ID: 0,
				},
			},
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("GET", "/api/profile/get", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Assert
	tests.Equals(t, http.StatusNotFound, w.Code)
}
func TestRetrieveProfile_AuthFailed(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/api/profile/get", func(c *gin.Context) {
		postgresc.RetrieveProfile(c,
			postgresinterface.MockDB{
				Profile: model.Profile{
					ID: 1,
					Password: "abcdefghijklmnopqrstuvwxy",
				},
			},
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	form.Set("username", "test_user")
	form.Set("password", "abcdefghijklmnopqrstuvwxyz")
	req, _ := http.NewRequest("GET", "/api/profile/get", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Assert
	tests.Equals(t, http.StatusUnauthorized, w.Code)
}
func TestRetrieveProfile_AuthSucceed(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	hashPass, _ := bcrypt.GenerateFromPassword([]byte("abcdefghijklmnopqrstuvwxyz"), 1)
	r.GET("/api/profile/get", func(c *gin.Context) {
		postgresc.RetrieveProfile(c,
			postgresinterface.MockDB{
				Profile: model.Profile{
					ID: 1,
					Password: string(hashPass),
				},
			},
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	form.Set("username", "test_user")
	form.Set("password", "abcdefghijklmnopqrstuvwxyz")
	req, _ := http.NewRequest("GET", "/api/profile/get?" + form.Encode(), nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Assert
	tests.Equals(t, http.StatusOK, w.Code)
}