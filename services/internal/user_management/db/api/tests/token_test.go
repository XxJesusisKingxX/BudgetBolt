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

	"services/internal/user_management/db/api"
	"services/internal/utils/testing"
)
func TestCreateToken(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)

	// Define a slice of test cases.
	testCases := []struct {
		TestName     string
		Form         url.Values
		ExpectedCode int
		ExpectedBody string
		TokenErr   error
	}{
		{
			TestName: "InvalidId",
			Form: url.Values{
				"id":{"invalid"},
			},
			ExpectedCode: http.StatusInternalServerError,
		},
		{
			TestName: "TokenCreated",
			Form: url.Values {
				"itemId":{"123"},
				"token":{"access-sandbox-132"},
				"id":{"12"},
			},
			ExpectedCode: http.StatusOK,
		},
		{
			TestName: "TokenNotCreated",
			Form: url.Values {
				"itemId":{"123"},
				"token":{"access-sandbox-132"},
				"id":{"12"},
			},
			TokenErr: errors.New(""),
			ExpectedCode: http.StatusNotImplemented,
			ExpectedBody: "TOKEN NOT CREATED",
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			r := gin.Default()
			// Handle mock route
			r.POST("/create-token", func(c *gin.Context) {
				api.CreateToken(c,
					api.MockDB{
						TokenErr: tc.TokenErr,
					},
					nil,
					true,
				)
			})
			// Create request
			req, _ := http.NewRequest("POST", "/create-token", strings.NewReader(tc.Form.Encode()))
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

func TestGetToken(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)

	// Define a slice of test cases.
	testCases := []struct {
		TestName     string
		ExpectedCode int
		ExpectedBody string
		TokenErr     error
		ProfileErr   error
	}{
		{
			TestName: "TokenRetrieved",
			ExpectedCode: http.StatusOK,
		},
		{
			TestName: "ProfileFailed",
			ProfileErr: errors.New(""),
			ExpectedCode: http.StatusInternalServerError,
		},
		{
			TestName: "TokenFailed",
			TokenErr: errors.New(""),
			ExpectedCode: http.StatusNotFound,
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			r := gin.Default()
			// Handle mock route
			r.GET("/get-token", func(c *gin.Context) {
				api.RetrieveToken(c,
					api.MockDB{
						ProfileErr: tc.ProfileErr,
						TokenErr: tc.TokenErr,
					},
					nil,
					true,
				)
			})
			// Create request
			req, _ := http.NewRequest("GET", "/get-token", nil)
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
