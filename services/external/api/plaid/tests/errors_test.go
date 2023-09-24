package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"

	"services/external/api/plaid"
	"services/internal/utils/testing"
)
func TestGetExpenses(t *testing.T) {
	// Define a slice of test cases.
	testCases := []struct {
		TestName     string
		ExpectedCode int
		Err          error
	}{
		{
			TestName: "Render500",
			Err: errors.New("Not Plaid Type Error"),
			ExpectedCode: http.StatusInternalServerError,
		},
		{
			TestName: "Render200",
			ExpectedCode: http.StatusOK,
		},
	}

	// Run all test cases
	for _, tc := range testCases {
		t.Run(tc.TestName, func(t *testing.T) {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			err := errors.New("Failed to connect")
			api.RenderError(c, err, api.MockPlaidClient{Err: tc.Err})
			// 200 error
			tests.Equals(t, tc.ExpectedCode, c.Writer.Status())
		})
	}
}