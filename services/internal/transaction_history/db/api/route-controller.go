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
	if err == nil {
		// Retrieve the user's transactions based on the profile ID.
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

		if err != nil || len(transactions) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error":"TRANSACTIONS NOT FOUND"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
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
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	if debug != true {
		utils.TransactionsToDB(db, id, data)
	}

	c.JSON(http.StatusOK, gin.H{})
}