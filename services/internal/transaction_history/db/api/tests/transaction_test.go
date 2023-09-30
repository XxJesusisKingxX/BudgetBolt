package test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"

	"services/internal/transaction_history/db/api"
	"services/internal/transaction_history/db/model"
	"services/internal/utils/testing"
    "services/internal/utils/http"
)

func TestGetTransactions(t *testing.T) {
    testCases := []struct {
        TestName       string
        ProfileErr     error
        TransactionErr error
        Transactions   []model.Transaction
        ExpectedCode   int
        Response       map[string]request.MockResponse
        ShouldContain  string
    }{
        {
            TestName: "TransactionsReceived",
            Transactions: []model.Transaction{
                {
                    ID:          "1",
                    Date:        "2023/01/01",
                    Amount:      12.34,
                    Method:      "",
                    AccountName: "Test Account",
                    Vendor:      "Test",
                    Description: "A test case was received",
                    ProfileID:   1,
                },
            },
            Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
			},
            ExpectedCode:  http.StatusOK,
            ShouldContain: "\"A test case was received\"",
        },
        {
            TestName:   "ProfileIdNotFound",
            ProfileErr: errors.New(""),
            Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusInternalServerError,
				},
			},
            ExpectedCode:  http.StatusInternalServerError,
        },
        {
            TestName: "TransactionsNotFound",
            TransactionErr: errors.New(""),
            Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
			},
            ExpectedCode:  http.StatusNotFound,
            ShouldContain: "\"TRANSACTIONS NOT FOUND\"",
        },
        {
            TestName: "TransactionsEmpty",
            Transactions: []model.Transaction{},
            Response: map[string]request.MockResponse{
				"profile/get": {
					Code: http.StatusOK,
				},
			},
            ExpectedCode:  http.StatusNotFound,
            ShouldContain: "\"TRANSACTIONS NOT FOUND\"",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.TestName, func(t *testing.T) {
            // Create mock engine
            gin.SetMode(gin.TestMode)
            r := gin.Default()
            // Mock requests
            mockClient := request.MockHTTPClient{}
            mockClient.Responses = tc.Response
            // Handle mock route
            r.GET("/get-transactions", func(c *gin.Context) {
                api.RetrieveTransactions(c,
                    api.MockDB{
                        ProfileErr:     tc.ProfileErr,
                        TransactionErr: tc.TransactionErr,
                        Transaction:    tc.Transactions,
                    },
                    nil,
                    mockClient,
                    true,
                )
            })

            // Create request
            form := url.Values{}
            req, _ := http.NewRequest("GET", "/get-transactions", strings.NewReader(form.Encode()))
            req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
            w := httptest.NewRecorder()
            r.ServeHTTP(w, req)

            // Received response
            responseBody, _ := io.ReadAll(w.Result().Body)
            defer w.Result().Body.Close()
            isReceived := strings.Contains(string(responseBody), tc.ShouldContain)

            // Assert
            tests.Equals(t, tc.ExpectedCode, w.Code)
            tests.Equals(t, true, isReceived)
        })
    }
}

func TestStoreTransactions(t *testing.T) {
    testCases := []struct {
        TestName       string
        Form           url.Values
        ExpectedCode   int
    }{
        {
            TestName: "IDInvalid",
            Form: url.Values{
                "id" : {"invalid"},
            },
            ExpectedCode:  http.StatusInternalServerError,
        },
        {
            TestName: "StoreFailed",
            Form: url.Values{
                "id" : {"12"},
                "transactions": {""},
            },
            ExpectedCode:  http.StatusInternalServerError,
        },
        {
            TestName: "StoreSuccess",
            Form: url.Values{
                "id" : {"12"},
                "transactions": {`[{"account_id": "account123"},{"account_id": "account456"}]`},
            },
            ExpectedCode:  http.StatusOK,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.TestName, func(t *testing.T) {
            // Create mock engine
            gin.SetMode(gin.TestMode)
            r := gin.Default()
            // Handle mock route
            r.POST("/get-transactions", func(c *gin.Context) {
                api.StoreTransactions(c,
                    api.MockDB{
                    },
                    nil,
                    true,
                )
            })

            // Create request
            req, _ := http.NewRequest("POST", "/get-transactions", strings.NewReader(tc.Form.Encode()))
            req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
            w := httptest.NewRecorder()
            r.ServeHTTP(w, req)

            // Assert
            tests.Equals(t, tc.ExpectedCode, w.Code)
        })
    }
}
