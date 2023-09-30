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

func TestGetExpenses(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	// Define a slice of test cases.
	limit := 1.0
	spent := 1.0
	testCases := []struct {
		TestName     string
		Expense      []model.Expense
		ExpectedCode int
		Response     map[string]request.MockResponse
		ExpectedBody string
		ProfileErr   error
		ExpenseErr   error
	}{
		{
			TestName: "ExpensesFound",
			Expense: []model.Expense{
				{
					Name:  "Test",
					Limit: &limit,
					Spent: &spent,
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
			TestName: "ExpensesNotFound",
			ExpenseErr:   errors.New(""),
			ExpectedCode: http.StatusNotFound,
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
			},
			ExpectedBody: "EXPENSES NOT FOUND",
		},
		{
			TestName: "ExpensesEmpty",
			Expense: []model.Expense{},
			ExpenseErr:   errors.New(""),
			ExpectedCode: http.StatusNotFound,
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
			},
			ExpectedBody: "EXPENSES NOT FOUND",
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
			r.GET("/get-expenses", func(c *gin.Context) {
				api.RetrieveExpenses(c,
					api.MockDB{
						Profile:    user.Profile{ID: 1},
						Expense:    tc.Expense,
						ProfileErr: tc.ProfileErr,
						ExpenseErr: tc.ExpenseErr,
					},
					nil,
					mockClient,
					true,
				)
			})
			// Create request
			req, _ := http.NewRequest("GET", "/get-expenses", nil)
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
				tests.Equals(t, tc.ExpectedBody, body["expenses"])
			}
		})
	}
}

func TestCreateExpenses(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)

	// Define a slice of test cases.
	testCases := []struct {
		TestName     string
		Form         url.Values
		ExpectedCode int
		Response     map[string]request.MockResponse
		ExpectedBody string
		ProfileErr   error
		ExpenseErr   error
	}{
		{
			TestName: "ExpensesNotCreated",
			Form: url.Values{
				"name":  {"TestExpense"},
				"limit": {"100.00"},
				"spent": {"50.00"},
			},
			ExpenseErr:   errors.New(""),
			ExpectedCode: http.StatusNotImplemented,
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
			},
			ExpectedBody: "EXPENSES NOT CREATED",
		},
		{
			TestName: "ProfileNotFound",
			Form: url.Values{
				"name":  {"TestExpense"},
				"limit": {"100.00"},
				"spent": {"50.00"},
			},
			ProfileErr: errors.New(""),
			ExpectedCode: http.StatusInternalServerError,
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusInternalServerError,
				},
			},
		},
		{
			TestName: "ExpensesCreated",
			Form: url.Values{
				"name":  {"TestExpense"},
				"limit": {"100.00"},
				"spent": {"50.00"},
			},
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
			},
			ExpectedCode: http.StatusOK,
		},
		{
			TestName: "LimitFieldInvalid",
			Form: url.Values{
				"name":  {"TestExpense"},
				"limit": {"invalid-limit"},
				"spent": {"50.00"},
			},
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "INVALID LIMIT AND/OR SPENT",
		},
		{
			TestName: "SpentFieldInvalid",
			Form: url.Values{
				"name":  {"TestExpense"},
				"limit": {"50.00"},
				"spent": {"invalid-spent"},
			},
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "INVALID LIMIT AND/OR SPENT",
		},
		{
			TestName: "Spent&LimitFieldInvalid",
			Form: url.Values{
				"name":  {"TestExpense"},
				"limit": {"invalid-limit"},
				"spent": {"invalid-spent"},
			},
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "INVALID LIMIT AND/OR SPENT",
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
			r.POST("/create-expenses", func(c *gin.Context) {
				api.CreateExpenses(c,
					api.MockDB{
						ProfileErr: tc.ProfileErr,
						ExpenseErr: tc.ExpenseErr,
					},
					nil,
					mockClient,
					true,
				)
			})
			// Create request
			req, _ := http.NewRequest("POST", "/create-expenses", strings.NewReader(tc.Form.Encode()))
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

func TestUpdateExpenses(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)

	// Define a slice of test cases.
	testCases := []struct {
		TestName    string
		Form         url.Values
		ExpectedCode int
		Response     map[string]request.MockResponse
		ExpectedBody string
		ProfileErr   error
		ExpenseErr   error
	}{
		{
			TestName: "ExpensesUpdated",
			Form: url.Values{
				"limit": {"100.00"},
				"id":    {"1"},
			},
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
			},
			ExpectedCode: http.StatusOK,
		},
		{
			TestName: "ExpensesNotUpdated",
			Form: url.Values{
				"limit": {"100.00"},
				"id":    {"1"},
			},
			ExpenseErr:   errors.New(""),
			ExpectedCode: http.StatusNotImplemented,
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
			},
			ExpectedBody: "EXPENSES NOT UPDATED",
		},
		{
			TestName: "ProfileNotFound",
			Form: url.Values{
				"limit": {"100.00"},
				"id":    {"1"},
			},
			ProfileErr:   errors.New(""),
			ExpectedCode: http.StatusInternalServerError,
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusInternalServerError,
				},
			},
		},
		{
			TestName: "LimitFieldInvalid",
			Form: url.Values{
				"limit": {"invalid-limit"},
				"spent": {"1"},
			},
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "INVALID LIMIT AND/OR ID",
		},
		{
			TestName: "IdFieldInvalid",
			Form: url.Values{
				"limit": {"100.00"},
				"id":    {"invalid-id"},
			},
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "INVALID LIMIT AND/OR ID",
		},
		{
			TestName: "Id&LimitFieldInvalid",
			Form: url.Values{
				"name":  {"TestExpense"},
				"limit": {"invalid-limit"},
				"id":    {"invalid-id"},
			},
			ExpectedCode: http.StatusBadRequest,
			ExpectedBody: "INVALID LIMIT AND/OR ID",
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
			r.POST("/update-expenses", func(c *gin.Context) {
				api.UpdateExpenses(c,
					api.MockDB{
						ProfileErr: tc.ProfileErr,
						ExpenseErr: tc.ExpenseErr,
					},
					nil,
					mockClient,
					true,
				)
			})
			// Create request
			req, _ := http.NewRequest("POST", "/update-expenses", strings.NewReader(tc.Form.Encode()))
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

func TestUpdateAllExpenses(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)

	// Define a slice of test cases.
	testCases := []struct {
		TestName    string
		ExpectedCode int
		Response     map[string]request.MockResponse
		ExpectedBody string
		ProfileErr   error
		ExpenseErr   error
	}{
		{
			TestName: "ExpensesUpdated",
			Response: map[string]request.MockResponse{
				"transactions/get?uid=test&date=123": {
					Code: http.StatusOK,
				},
				"profile/get": {
					Code: http.StatusOK,
				},
			},
			ExpectedCode: http.StatusOK,
		},
		{
			TestName: "TransactionsNotFound",
			ExpectedCode: http.StatusInternalServerError,
			Response: map[string]request.MockResponse{
				"transactions/get?uid=test&date=123": {
					Code: http.StatusInternalServerError,
				},
			},
		},
		{
			TestName: "ProfileNotFound",
			Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusInternalServerError,
				},
			},
			ExpectedCode: http.StatusInternalServerError,
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
			r.POST("/update-expenses/all", func(c *gin.Context) {
				api.UpdateAllExpenses(c,
					api.MockDB{
					},
					nil,
					mockClient,
					true,
				)
			})
			// Create request
			form := url.Values{
				"date": {"123"},
			}
			req, _ := http.NewRequest("POST", "/update-expenses/all", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
			tests.Equals(t, tc.ExpectedBody, body["error"])
		})
	}
}
