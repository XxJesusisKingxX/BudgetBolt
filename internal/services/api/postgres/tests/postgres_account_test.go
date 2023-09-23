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