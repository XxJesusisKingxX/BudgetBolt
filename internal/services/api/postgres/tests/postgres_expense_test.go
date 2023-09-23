package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	postgresinterface "services/api/postgres"
	postgresc "services/api/postgres/controller"
	"services/db/postgresql/model"
	tests "services/utils/testing"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetExpenses(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)

	// Define a slice of test cases.
	limit := 1.0
	spent := 1.0
	testCases := []struct {
		Expense      []model.Expense
		ExpectedCode int
		ExpectedBody string
		ProfileErr   error
		ExpenseErr   error
	}{
		{
			// Case: expenses found
			Expense: []model.Expense{
				{
					Name:  "Test",
					Limit: &limit,
					Spent: &spent,
				},
			},
			ExpectedCode: http.StatusOK,
		},
		{
			// Case: profile not found
			ProfileErr:   errors.New("Failed to get profile id"),
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Failed to get profile id",
		},
		{
			// Case: expenses not found
			ExpenseErr:   errors.New("Failed to get expenses"),
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Failed to get expenses",
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		r := gin.Default()
		// Handle mock route
		r.GET("/get-expenses", func(c *gin.Context) {
			postgresc.RetrieveExpenses(c,
				postgresinterface.MockDB{
					Profile:    model.Profile{ID: 1},
					Expense:    tc.Expense,
					ProfileErr: tc.ProfileErr,
					ExpenseErr: tc.ExpenseErr,
				},
				nil,
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
	}
}

func TestCreateExpenses(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)

	// Define a slice of test cases.
	testCases := []struct {
		Form         url.Values
		ExpectedCode int
		ExpectedBody string
		ProfileErr   error
		ExpenseErr   error
	}{
		{
			// Case: expense not created
			Form: url.Values{
				"name":  {"TestExpense"},
				"limit": {"100.00"},
				"spent": {"50.00"},
			},
			ExpenseErr:   errors.New("Failed to create expense"),
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Failed to create expense",
		},
		{
			// Case: profile not found
			Form: url.Values{
				"name":  {"TestExpense"},
				"limit": {"100.00"},
				"spent": {"50.00"},
			},
			ProfileErr:   errors.New("Failed to get profile id"),
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Failed to get profile id",
		},
		{
			// Case: REQUEST all valid items
			Form: url.Values{
				"name":  {"TestExpense"},
				"limit": {"100.00"},
				"spent": {"50.00"},
			},
			ExpectedCode: http.StatusOK,
		},
		{
			// Case: REQUEST limit is invalid
			Form: url.Values{
				"name":  {"TestExpense"},
				"limit": {"invalid-limit"},
				"spent": {"50.00"},
			},
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Invalid limit and/or spent amount",
		},
		{
			// Case: REQUEST spent is invalid
			Form: url.Values{
				"name":  {"TestExpense"},
				"limit": {"50.00"},
				"spent": {"invalid-spent"},
			},
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Invalid limit and/or spent amount",
		},
		{
			// Case: REQUEST Both limit and spent are invalid
			Form: url.Values{
				"name":  {"TestExpense"},
				"limit": {"invalid-limit"},
				"spent": {"invalid-spent"},
			},
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Invalid limit and/or spent amount",
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		r := gin.Default()
		// Handle mock route
		r.POST("/create-expenses", func(c *gin.Context) {
			postgresc.CreateExpenses(c,
				postgresinterface.MockDB{
					ProfileErr: tc.ProfileErr,
					ExpenseErr: tc.ExpenseErr,
				},
				nil,
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
	}
}

func TestUpdateExpenses(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)

	// Define a slice of test cases.
	testCases := []struct {
		Form         url.Values
		ExpectedCode int
		ExpectedBody string
		ProfileErr   error
		ExpenseErr   error
	}{
		{
			// Case: expense updated
			Form: url.Values{
				"limit": {"100.00"},
				"id":    {"1"},
			},
			ExpectedCode: http.StatusOK,
		},
		{
			// Case: expense not updated
			Form: url.Values{
				"limit": {"100.00"},
				"id":    {"1"},
			},
			ExpenseErr:   errors.New("Failed to update expense"),
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Failed to update expense",
		},
		{
			// Case: profile not found
			Form: url.Values{
				"limit": {"100.00"},
				"id":    {"1"},
			},
			ProfileErr:   errors.New("Failed to get profile id"),
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Failed to get profile id",
		},
		{
			// Case: REQUEST limit is invalid
			Form: url.Values{
				"limit": {"invalid-limit"},
				"spent": {"1"},
			},
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Invalid limit amount and/or id",
		},
		{
			// Case: REQUEST id is invalid
			Form: url.Values{
				"limit": {"100.00"},
				"id":    {"invalid-id"},
			},
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Invalid limit amount and/or id",
		},
		{
			// Case: REQUEST Both limit and id are invalid
			Form: url.Values{
				"name":  {"TestExpense"},
				"limit": {"invalid-limit"},
				"id":    {"invalid-id"},
			},
			ExpectedCode: http.StatusInternalServerError,
			ExpectedBody: "Invalid limit amount and/or id",
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		r := gin.Default()
		// Handle mock route
		r.POST("/update-expenses", func(c *gin.Context) {
			postgresc.UpdateExpenses(c,
				postgresinterface.MockDB{
					ProfileErr: tc.ProfileErr,
					ExpenseErr: tc.ExpenseErr,
				},
				nil,
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
	}
}
