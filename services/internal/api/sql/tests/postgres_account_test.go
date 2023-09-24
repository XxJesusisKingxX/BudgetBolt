package test

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

	"services/internal/api/sql"
	transaction "services/internal/transaction_history/db/model"
	user "services/internal/user_managment/db/model"
	"services/internal/utils/testing"
)
func TestGetAccounts(t *testing.T) {
    testCases := []struct {
        TestName       string
        ExpectedCode   int
        ExpectedBody   string
		ProfileErr     error
		AccountErr     error
		Accounts        []transaction.Account
    }{
        {
            TestName:       "AccountsFound",
			Accounts: []transaction.Account{
				{
					Name:    "Test1",
					Balance: 11.11,
				},
			},
            ExpectedCode:   http.StatusOK,
        },
        {
            TestName:         "ProfileNotFound",
			ProfileErr: errors.New(""),
            ExpectedCode:   http.StatusNotFound,
            ExpectedBody: "PROFILE NOT FOUND",
        },
        {
            TestName:         "AccountsNotFound",
			AccountErr:      errors.New(""),
            ExpectedCode:   http.StatusNotFound,
            ExpectedBody: "ACCOUNT NOT FOUND",
        },
        {
            TestName:         "AccountsEmpty",
			Accounts: []transaction.Account{},
            ExpectedCode:   http.StatusNotFound,
            ExpectedBody: "ACCOUNT NOT FOUND",
        },
    }


    for _, tc := range testCases {
        t.Run(tc.TestName, func(t *testing.T) {
			// Create mock engine
			gin.SetMode(gin.TestMode)
            r := gin.Default()
			// Handle mock route
			r.GET("/get-accounts", func(c *gin.Context) {
				api.RetrieveAccounts(c,
					api.MockDB{
						Profile:    user.Profile{ID: 1},
						Account:    tc.Accounts,
						ProfileErr: tc.ProfileErr,
						AccountErr: tc.AccountErr,
					},
					nil,
					true,
				)
			})
			// Create request
			req, _ := http.NewRequest("GET", "/get-accounts", nil)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			// Receive response
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

func TestGetAccounts_AccountsRecieved(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/get-accounts", func(c *gin.Context) {
		api.RetrieveAccounts(c,
			api.MockDB{
				Profile: user.Profile{ID: 1},
				Account: []transaction.Account{
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
		api.RetrieveAccounts(c,
			api.MockDB{
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
		api.RetrieveAccounts(c, 
			api.MockDB{
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