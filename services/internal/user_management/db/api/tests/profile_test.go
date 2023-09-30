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
	"golang.org/x/crypto/bcrypt"

	"services/internal/user_management/db/api"
	user "services/internal/user_management/db/model"
	"services/internal/utils/testing"
)
func TestCreateProfile(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)

	// Define a slice of test cases.
	testCases := []struct {
		TestName     string
		Form         url.Values
		ProfileId    int64
		ExpectedCode int
		ExpectedBody string
		ProfileErr   error
	}{
		{
			TestName: "NameTaken",
			ProfileId: 1,
			ExpectedCode: http.StatusConflict,
			ExpectedBody: "PROFILE ALREADY EXISTS",
		},
		{
			TestName: "PasswordTooLong",
			Form: url.Values {
				"password":{"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"},
			},
			ProfileId: 0,
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "PASSWORD TOO LONG",
		},
		{
			TestName: "ProfileNotCreated",
			ProfileId: 0,
			ProfileErr:   errors.New(""),
			ExpectedCode: http.StatusNotImplemented,
			ExpectedBody: "PROFILE NOT CREATED",
		},
		{
			TestName: "ProfileCreated",
			ProfileId: 0,
			ExpectedCode: http.StatusOK,
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			r := gin.Default()
			// Handle mock route
			r.POST("/create-profile", func(c *gin.Context) {
				api.CreateProfile(c,
					api.MockDB{
						Profile:    user.Profile{ID: tc.ProfileId},
						ProfileErr: tc.ProfileErr,
					},
					nil,
					true,
				)
			})
			// Create request
			req, _ := http.NewRequest("POST", "/create-profile", strings.NewReader(tc.Form.Encode()))
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

func TestRetrieveProfile(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)

	// Define a slice of test cases.
	testCases := []struct {
		TestName     string
		Form         url.Values
		Password     string
		ProfileId    int64
		ExpectedCode int
		ExpectedBody string
		ProfileErr   error
	}{
		{
			TestName: "AuthFailed",
			Form: url.Values{
				"username": {"test"},
				"password": {"abc"},
			},
			Password: "abcd",
			ProfileId: 1,
			ExpectedCode: http.StatusUnauthorized,
		},
		{
			TestName: "AuthSucceed",
			ProfileId: 1,
			Form: url.Values{
				"username": {"test"},
				"password": {"a"},
			},
			Password: "a",
			ExpectedCode: http.StatusOK,
		},
		{
			TestName: "ProfileNotFound",
			ProfileId: 0,
			ExpectedCode: http.StatusNotFound,
			ExpectedBody: "PROFILE NOT FOUND",
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			r := gin.Default()
			// Handle mock route
			hashPass, _ := bcrypt.GenerateFromPassword([]byte(tc.Password), bcrypt.DefaultCost)
			r.POST("/get-profile", func(c *gin.Context) {
				api.RetrieveProfile(c,
					api.MockDB{
						Profile: user.Profile{
							Name:     "test",
							ID:       tc.ProfileId,
							Password: string(hashPass),
						},
					},
					nil,
					true,
				)
			})
			// Create request
			req, _ := http.NewRequest("POST", "/get-profile", strings.NewReader(tc.Form.Encode()))
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
