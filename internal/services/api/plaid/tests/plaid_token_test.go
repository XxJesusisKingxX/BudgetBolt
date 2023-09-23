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