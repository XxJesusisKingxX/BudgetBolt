package api

import (
	"services/internal/transaction_history/db/model"
	"services/internal/transaction_history/db/utils"
	user "services/internal/user_management/db/model"
	"services/internal/utils/http"
	"strings"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"math"

	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v12/plaid"
)

func StoreAccounts(c *gin.Context, dbs DBHandler, db *sql.DB, debug bool) {
	// Extract the session cookie
	id, idErr := strconv.ParseInt(c.PostForm("id"), 10, 32)
	if idErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var data []plaid.AccountBase
	accounts := c.PostForm("accounts")
	err := json.Unmarshal([]byte(accounts), &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	if debug != true{
		utils.AccountsToDB(db, id, data)
	}

	c.JSON(http.StatusOK, gin.H{})
}

// RetrieveTransactions retrieves a user's transactions.
func RetrieveTransactions(c *gin.Context, dbs DBHandler, db *sql.DB,httpClient request.HTTP, debug bool) {
	// Extract the session cookie
	uid, _ := c.Cookie("UID")
	uidP := c.Query("uid")
	date := c.Query("date")
	category := c.Query("category")
	recurring := c.Query("recurring")

	// Retrieve the user's profile based on the uid.
	var profile user.Profile
	var body string
	if len(uidP) != 0 {
		body = fmt.Sprintf("uid=%v", uidP)
	} else {
		body = fmt.Sprintf("uid=%v", uid)
	}
	status, resp, err := httpClient.POST("profile/get", body)
	request.ParseResponse(resp, &profile)
	if status != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var transactions []model.Transaction
	var bills = make(map[string]*model.Bill)

	// var recurringTransactions []model.RecurringTransaction

	// if recurring == "enable" {
	// 	recurringTransactions, err = dbs.RetrieveRecurringTransaction(db, model.RecurringTransaction{
	// 		ProfileID: int64(profile.ID),
	// 		Query: model.Querys{
	// 			Select: model.QueryParameters{
	// 				OrderBy: model.OrderBy {
	// 					Desc: true,
	// 					Column: "last_date",
	// 				},
	// 				GreaterThanEq: model.GreaterThanEq{
	// 					Value: date,
	// 					Column: "last_date",
	// 				},
	// 			},
	// 		},
	// 	})

	// 	if err != nil || len(recurringTransactions) == 0 {
	// 		c.JSON(http.StatusNotFound, gin.H{"error":"TRANSACTIONS NOT FOUND"})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"transactions": recurringTransactions,
	// 	})

	// } 

	// Calculate recurring
	if recurring != "" {
		transactions, err = dbs.RetrieveTransaction(db, model.Transaction{
			ProfileID: int64(profile.ID),
			Query: model.Querys{
				Select: model.QueryParameters{
					OrderBy: model.OrderBy {
						Asc: true,
						Column: "transaction_date",
					},
					GreaterThanEq: model.GreaterThanEq{
						Value: utils.LookBackView(date, recurring),
						Column: "transaction_date",
					},
				},
			},
		})

		for _,transaction := range transactions {
			// TODO ADD QUERY ABILTITY TO SPEED UP
			if transaction.PrimaryCategory == "INCOME" || transaction.PrimaryCategory == "TRANSFER_IN" || transaction.PrimaryCategory == "TRANSFER_OUT" || transaction.PrimaryCategory == "BANK_FEES" {
				continue
			}

			var key string
			if transaction.Vendor != "" {
				key = transaction.Vendor
			} else {
				key = transaction.Description
			}

			if bill, ok := bills[key]; !ok {
				// Initialize first time seeing
				bills[key] = &model.Bill{
					Name:              key,
					Total:             transaction.Amount,
					MaxAmount:         transaction.Amount,
					AverageAmount:     transaction.Amount,
					Status:            "UNKNOWN",
					Frequency:         1,
					EarliestDate:      transaction.Date,
					PreviousDateCycle: transaction.Date,
					LastDateCycle:     transaction.Date,
					Category:          transaction.PrimaryCategory,
				}
				continue
			} else {
				// Add total
				bill.Total += transaction.Amount
				// Add amount of times seen (reset if trend breaks)
				bill.Frequency += 1
				// Calculate average amount over all recurring transactions
				bill.AverageAmount = bill.Total / float64(bill.Frequency)
				// Calculate max amount spent over all recurring transactions
				bill.MaxAmount = math.Max(bill.MaxAmount, transaction.Amount)
				// Check if the trend is healthy; if not, set "true" to "DEGRADED"
				if !utils.IsTrendHealthy(bill.PreviousDateCycle, bill.LastDateCycle, recurring) {
					bill.Degraded += 1
					// Reset frequency
					bill.Frequency = 1
					// Reset prev and current dater pointer to same
					bill.PreviousDateCycle = transaction.Date
					bill.LastDateCycle = transaction.Date
					continue
				}
				// Swap current pointer to prev before updating to continue streak
				bill.PreviousDateCycle = bill.LastDateCycle
				bill.LastDateCycle = transaction.Date
				// Get due date next
				bill.DueDate = utils.PredictNextDueDate(bill.PreviousDateCycle, bill.LastDateCycle)
				// Set "status" to "MATURE" or "EARLY" based on "freq" if not already set to "DEGRADED"
				if bill.Frequency >= 3 {
					bill.Status = "MATURE"
				} else if bill.Frequency == 2 {
					bill.Status = "EARLY"
				}
			}
		}


	} else {
		transactions, err = dbs.RetrieveTransaction(db, model.Transaction{
			ProfileID: int64(profile.ID),
			Query: model.Querys{
				Select: model.QueryParameters{
					OrderBy: model.OrderBy {
						Desc: true,
						Column: "transaction_date",
					},
					GreaterThanEq: model.GreaterThanEq{
						Value: date,
						Column: "transaction_date",
					},
					Equal: model.Equal{
						Value: strings.ToUpper(category),
						Column: "primary_category",
					},
				},
			},
		})
	}

	if err != nil || len(transactions) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error":"TRANSACTIONS NOT FOUND"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
		"bills": bills,
	})
}



func StoreTransactions(c *gin.Context, dbs DBHandler, db *sql.DB, debug bool) {
	id, idErr := strconv.ParseInt(c.PostForm("id"), 10, 64)
	if idErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var data []plaid.Transaction
	transactions := c.PostForm("transactions")
	err := json.Unmarshal([]byte(transactions), &data)
	// var dataStream []plaid.TransactionStream
	// recurringTransactions := c.PostForm("recurrings")
	// err = json.Unmarshal([]byte(recurringTransactions), &dataStream)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	if debug != true {
		utils.TransactionsToDB(db, id, data)
		// utils.RecurringTransactionsToDB(db, id, dataStream)
	}

	c.JSON(http.StatusOK, gin.H{})
}

func DeletePendingTransactions(c *gin.Context, dbs DBHandler, db *sql.DB,httpClient request.HTTP, debug bool) {
	uid, _ := c.Cookie("UID")

	var profile user.Profile
	body := fmt.Sprintf("uid=%v", uid)
	status, resp, err := httpClient.POST("profile/get", body)
	request.ParseResponse(resp, &profile)

	if status != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	transactions, err := dbs.RetrieveTransaction(db, model.Transaction{
		ProfileID: int64(profile.ID),
		Query: model.Querys{
			Select: model.QueryParameters{
				Equal: model.Equal{
					Value: true,
					Column: "pending",
				},
			},
		},
	})

	for _, transaction := range transactions {
		const FIVEDAYS = 5 * 24 * time.Hour // 5 days in duration
		postedDate, _ := time.Parse("2006-01-02", transaction.Date)
		timePassed := time.Since(postedDate)
		if timePassed >= FIVEDAYS {
			err = dbs.DeleteTransaction(db, model.Transaction{
				ID: transaction.ID,
			})
		}
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}


	c.JSON(http.StatusOK, gin.H{})
}

// test delete pending