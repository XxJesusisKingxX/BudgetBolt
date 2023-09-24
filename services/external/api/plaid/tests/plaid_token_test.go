package main

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

	"services/external/api/plaid"
	apiSql "services/internal/api/sql"
	"services/internal/user_managment/db/model"
	"services/internal/utils/testing"
)

func TestCreateAccessToken(t *testing.T) {
	// Define a slice of test cases.
	testCases := []struct {
		TestName       string
		PublicToken    string
		TokenErr       error
		ProfileErr     error
		ExpectedCode   int
		ExpectedBody   string
	}{
		{
			TestName:     "TokenCreationFailed",
			TokenErr: errors.New("Failed to get token"),
			PublicToken:  "public-sandbox-12345678-1234-1234-1234-123456789012",
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Failed to get token",
		},
		{
			TestName:     "ProfileFails",
			PublicToken:  "public-sandbox-12345678-1234-1234-1234-123456789012",
			ProfileErr:   errors.New("Failed to get profile"),
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Failed to get profile",
		},
		{
			TestName:     "Received",
			PublicToken:  "public-sandbox-12345678-1234-1234-1234-123456789012",
			ExpectedCode: http.StatusOK,
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			// Create mock engine
			gin.SetMode(gin.TestMode)
			r := gin.Default()

			// Handle mock route
			r.POST("/get-access-token", func(c *gin.Context) {
				api.CreateAccessToken(c,
					api.MockPlaidClient{Err: tc.ProfileErr},
					apiSql.MockDB{ 
						Profile: model.Profile{ID: 1},
						TokenErr: tc.TokenErr,
					},
					nil,
					nil,
					true,
				)
			})

			// Create request
			form := url.Values{}
			form.Set("public_token", tc.PublicToken)
			req, _ := http.NewRequest("POST", "/get-access-token", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			// Recieve response
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

func TestLinkTokenCreate(t *testing.T) {
	// Define a slice of test cases.
	testCases := []struct {
		TestName     string
		Profile      model.Profile
		ExpectedCode int
		ExpectedBody string
		Err          error
	}{
		{
			TestName: "LinkTokenCreated",
			Profile:  model.Profile{Name: "test_user"},
			ExpectedCode: http.StatusOK,
		},
		{
			TestName:   "ProfileFailed",
			Err: errors.New("Failed to get profile"),
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Failed to get profile",
		},
		{
			TestName:     "LinkTokenCreationFailed",
			Profile:      model.Profile{},
			Err:  errors.New("Failed to create token"),
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Failed to create token",
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			// Create mock engine
			gin.SetMode(gin.TestMode)
			r := gin.Default()

			// Handle mock route
			r.POST("/create-link-token", func(c *gin.Context) {
				api.CreateLinkToken(c, api.MockPlaidClient{Err: tc.Err}, apiSql.MockDB{Profile: tc.Profile}, nil, nil)
			})

			// Create request
			form := url.Values{}
			req, _ := http.NewRequest("POST", "/create-link-token", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// Recieve response
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