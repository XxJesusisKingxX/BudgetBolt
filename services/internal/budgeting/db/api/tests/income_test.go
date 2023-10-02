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

	"services/internal/budgeting/db/api"
	"services/internal/budgeting/db/model"
	user "services/internal/user_management/db/model"
	"services/internal/utils/http"
	"services/internal/utils/testing"
)

func TestGetIncomes(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	// Define a slice of test cases.
	amt := 1.0
	testCases := []struct {
		TestName     string
		Income      []model.Income
		ExpectedCode int
		Response     map[string]request.MockResponse
		ExpectedBody string
		ProfileErr   error
		IncomeErr   error
	}{
		{
			TestName: "IncomesFound",
			Income: []model.Income{
				{
					Name:  "Test",
					Amount: &amt,
				},
			},
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
			},
			ExpectedCode: http.StatusOK,
		},
		{
			TestName: "ProfileNotFound",
			ProfileErr:   errors.New(""),
			ExpectedCode: http.StatusInternalServerError,
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusInternalServerError,
				},
			},
		},
		{
			TestName: "IncomesNotFound",
			IncomeErr:   errors.New(""),
			ExpectedCode: http.StatusNotFound,
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
			},
			ExpectedBody: "INCOMES NOT FOUND",
		},
		{
			TestName: "IncomesEmpty",
			Income: []model.Income{},
			ExpectedCode: http.StatusNotFound,
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
			},
			ExpectedBody: "INCOMES NOT FOUND",
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			r := gin.Default()
			// Mock requests
			mockClient := request.MockHTTPClient{}
			mockClient.Responses = tc.Response
			// Handle mock route
			r.GET("/get-income", func(c *gin.Context) {
				api.RetrieveIncome(c,
					api.MockDB{
						Profile:    user.Profile{ID: 1},
						Income:    tc.Income,
						ProfileErr: tc.ProfileErr,
						IncomeErr: tc.IncomeErr,
					},
					nil,
					mockClient,
					true,
				)
			})
			// Create request
			req, _ := http.NewRequest("GET", "/get-income", nil)
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
			if value, exists := body["error"]; exists {
				tests.Equals(t, tc.ExpectedBody, value)
			} else {
				tests.Equals(t, tc.ExpectedBody, body["incomes"])
			}
		})
	}
}

func TestUpsertIncomes(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	// Define a slice of test cases.
	amt := 1.0
	testCases := []struct {
		TestName     string
		Income      []model.Income
		Form         url.Values
		ExpectedCode int
		Response     map[string]request.MockResponse
		ExpectedBody string
		ProfileErr   error
		IncomeErr    error
	}{
		{
			TestName: "TransactionsNotFound",
			Income: []model.Income{
				{
					Name:  "Test",
					Amount: &amt,
				},
			},
			Form: url.Values{
				"date":  {"123"},
			},
			Response: map[string]request.MockResponse{
				"transactions/get?uid=test&date=123&category=income": {
					Code: http.StatusInternalServerError,
				},
			},
			ExpectedCode: http.StatusInternalServerError,
		},
		{
			TestName: "ProfileNotFound",
			Form: url.Values{
				"date":  {"123"},
			},
			ExpectedCode: http.StatusInternalServerError,
			Response: map[string]request.MockResponse{
				"transactions/get?uid=test&date=123&category=income": {
					Code: http.StatusOK,
				},
				"profile/get": {
					Code: http.StatusInternalServerError,
				},
			},
		},
		{
			TestName: "IncomesUpdated&Created",
			Form: url.Values{
				"date":  {"123"},
			},
			ExpectedCode: http.StatusOK,
			Response: map[string]request.MockResponse{
				"transactions/get?uid=test&date=123&category=income": {
					Code: http.StatusOK,
				},
				"profile/get": {
					Code: http.StatusOK,
				},
			},
		},
		{
			TestName: "IncomesFailedUpdate/Create/Retrieve",
			Form: url.Values{
				"date":  {"123"},
			},
			IncomeErr: errors.New(""),
			Income: []model.Income{
				{
					Name: "Test1",
					Amount: &amt,
				},
				{
					Name: "Test1",
					Amount: &amt,
				},
			},
			ExpectedCode: http.StatusNotImplemented,
			Response: map[string]request.MockResponse{
				"transactions/get?uid=test&date=123&category=income": {
					Code: http.StatusOK,
				},
				"profile/get": {
					Code: http.StatusOK,
				},
			},
			ExpectedBody: "INCOMES NOT CREATED/UPDATED",
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			r := gin.Default()
			// Mock requests
			mockClient := request.MockHTTPClient{}
			mockClient.Responses = tc.Response
			// Handle mock route
			r.POST("/upsert-income", func(c *gin.Context) {
				api.UpsertIncome(c,
					api.MockDB{
						Profile:    user.Profile{ID: 1},
						Income:    tc.Income,
						ProfileErr: tc.ProfileErr,
						IncomeErr: tc.IncomeErr,
					},
					nil,
					mockClient,
					true,
				)
			})
			// Create request
			req, _ := http.NewRequest("POST", "/upsert-income", strings.NewReader(tc.Form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			// Create a new cookie.
			cookie := &http.Cookie{
				Name:  "UID",
				Value: "test",
			}
			req.AddCookie(cookie)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			// Receive response
			var body map[string]string
			responseBody, _ := io.ReadAll(w.Result().Body)
			json.Unmarshal(responseBody, &body)
			defer w.Result().Body.Close()
			// Assert
			tests.Equals(t, tc.ExpectedCode, w.Code)
			if value, exists := body["error"]; exists {
				tests.Equals(t, tc.ExpectedBody, value)
			} else {
				tests.Equals(t, tc.ExpectedBody, body["incomes"])
			}
		})
	}
}