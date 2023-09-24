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

	"services/internal/api/sql"
	transaction "services/internal/transaction_history/db/model"
	"services/internal/utils/testing"
)

func TestGetTransactions(t *testing.T) {
    testCases := []struct {
        testName       string
        profileErr     error
        transactionErr error
        transactions   []transaction.Transaction
        expectedCode   int
        shouldContain  string
    }{
        {
            testName: "TransactionsReceived",
            transactions: []transaction.Transaction{
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
            expectedCode:  http.StatusOK,
            shouldContain: "\"A test case was received\"",
        },
        {
            testName:   "ProfileIdNotFound",
            profileErr: errors.New(""),
            expectedCode:  http.StatusNotFound,
            shouldContain: "\"PROFILE NOT FOUND\"",
        },
        {
            testName: "TransactionsNotFound",
            transactionErr: errors.New(""),
            expectedCode:  http.StatusNotFound,
            shouldContain: "\"TRANSACTIONS NOT FOUND\"",
        },
        {
            testName: "TransactionsEmpty",
            transactions: []transaction.Transaction{},
            expectedCode:  http.StatusNotFound,
            shouldContain: "\"TRANSACTIONS NOT FOUND\"",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.testName, func(t *testing.T) {
            // Create mock engine
            gin.SetMode(gin.TestMode)
            r := gin.Default()

            // Handle mock route
            r.GET("/get-transactions", func(c *gin.Context) {
                api.RetrieveTransactions(c,
                    api.MockDB{
                        ProfileErr:     tc.profileErr,
                        TransactionErr: tc.transactionErr,
                        Transaction:    tc.transactions,
                    },
                    nil,
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
            isReceived := strings.Contains(string(responseBody), tc.shouldContain)

            // Assert
            tests.Equals(t, tc.expectedCode, w.Code)
            tests.Equals(t, true, isReceived)
        })
    }
}
