package api

import (
	"fmt"
	"math"
	"services/internal/budgeting/db/model"
	transaction "services/internal/transaction_history/db/model"
	user "services/internal/user_management/db/model"
	"services/internal/utils/http"

	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Retrieves a user's budgeted expenses.
func RetrieveExpenses(c *gin.Context, dbs DBHandler, db *sql.DB, httpClient request.HTTP, debug bool) {
	// Extract the session cookie or form
	uid, _ := c.Cookie("UID")
	uidP := c.Query("uid")

	// Retrieve the user's profile based on the cookie or form.
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

	var expenses []model.Expense
	if err == nil {
		// Retrieve the user's expenses based on the profile ID.
		expenses, err = dbs.RetrieveExpense(db, model.Expense{
			ProfileID: profile.ID,
		})
	}
	if err != nil || len(expenses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error":"EXPENSES NOT FOUND"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"expenses": expenses,
	})
}

// Create a user's budgeted expenses.
func CreateExpenses(c *gin.Context, dbs DBHandler, db *sql.DB, httpClient request.HTTP, debug bool) {
	// Extract the session cookie
	uid, _ := c.Cookie("UID")
	name := c.PostForm("name")
	limit, limitErr := strconv.ParseFloat(c.PostForm("limit"), 64)
	spent, spentErr := strconv.ParseFloat(c.PostForm("spent"), 64)
	if limitErr != nil || spentErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"INVALID LIMIT AND/OR SPENT"})
		return
	}
	// Retrieve the user's profile based on the username.
	var profile user.Profile
	body := fmt.Sprintf("uid=%v", uid)
	status, resp, err := httpClient.POST("profile/get", body)
	request.ParseResponse(resp, &profile)

	if status != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if err == nil {
		// Create the user's expenses based on the profile ID.
		err = dbs.CreateExpense(db, model.Expense{
			ProfileID: profile.ID,
			Limit: &limit,
			Name: name,
			Spent: &spent,
		})

		if err != nil {
			c.JSON(http.StatusNotImplemented, gin.H{"error":"EXPENSES NOT CREATED"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}

// Update a user's budgeted expenses.
func UpdateExpenses(c *gin.Context, dbs DBHandler, db *sql.DB,httpClient request.HTTP, debug bool) {
	// Extract the session cookie
	uid, _ := c.Cookie("UID")
	uidP := c.PostForm("uid")
	limit, limitErr := strconv.ParseFloat(c.PostForm("limit"), 64)
	id, idErr := strconv.ParseInt(c.PostForm("id"), 10, 32)
	if limitErr != nil || idErr != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"INVALID LIMIT AND/OR ID"})
		return
	}
	// Check if user exist
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
	if err == nil {
		// Create the user's expenses based on the profile ID.
		err = dbs.UpdateExpense(db, model.Expense {
			Limit: &limit,
		}, model.Expense{
			ID: id,
		})
		if err != nil {
			c.JSON(http.StatusNotImplemented, gin.H{"error":"EXPENSES NOT UPDATED"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}

func UpdateAllExpenses(c *gin.Context, dbs DBHandler, db *sql.DB, httpClient request.HTTP, debug bool) {
	uid, _ := c.Cookie("UID")
	date := c.PostForm("date")

	// Get transactions
	var transaction transaction.Transactions
	url := fmt.Sprintf("transactions/get?uid=%v&date=%v", uid, date)
	status, resp, _ := httpClient.GET(url)
	request.ParseResponse(resp, &transaction)

	if status != http.StatusOK && status != http.StatusNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	// Get the category totals for all possible expenses
	categoryTotals := make(map[string]float64)
	for _, transaction := range transaction.Transactions {
		if transaction.Amount > 0 {
			primary := transaction.PrimaryCategory
			detailed := transaction.DetailCategory
			categoryTotals[primary] += transaction.Amount
			categoryTotals[detailed] += transaction.Amount
		}
	}

	// Get profile
	var profile user.Profile
	body := fmt.Sprintf("uid=%v", uid)
	status, resp, _ = httpClient.POST("profile/get", body)
	request.ParseResponse(resp, &profile)
	if status != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	expenses, _ := dbs.RetrieveExpense(db, model.Expense{
		ProfileID: profile.ID,
	})
	// Update expense total spent
	for _, expense := range expenses {
		var total float64
		categories := expense.Category
		for _, category := range categories {
			total += categoryTotals[category]
		}
		total = math.Round(total*100) / 100 // round 2decimals
		dbs.UpdateExpense(db, model.Expense{
			Spent: &total,
		}, model.Expense{
			ID: expense.ID,
		})
	}

	c.JSON(http.StatusOK, gin.H{})
}

// Create/Update a user's budgeted expenses.
func UpsertIncome(c *gin.Context, dbs DBHandler, db *sql.DB, httpClient request.HTTP, debug bool) {
	// Extract the session cookie
	uid, _ := c.Cookie("UID")
	date := c.PostForm("date")

	// Get transactions
	var transaction transaction.Transactions
	url := fmt.Sprintf("transactions/get?uid=%v&date=%v&category=income", uid, date)
	status, resp, _ := httpClient.GET(url)
	request.ParseResponse(resp, &transaction)

	if status != http.StatusOK && status != http.StatusNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	// Retrieve the user's profile based on the username.
	var profile user.Profile
	body := fmt.Sprintf("uid=%v", uid)
	status, resp, err := httpClient.POST("profile/get", body)
	request.ParseResponse(resp, &profile)

	if status != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	// Check if incomes already exist
	incomeTotals := make(map[string]float64)
	if len(transaction.Transactions) != 0 {
		for _, transaction := range transaction.Transactions {
			if transaction.Vendor != "" {
				incomeTotals[transaction.Vendor] += transaction.Amount * -1 // remove negative
			} else {
				incomeTotals[transaction.Description] += transaction.Amount * -1
			}
		}

		for name, amount := range incomeTotals {
			incomes, _ := dbs.RetrieveIncome(db, model.Income{
				Name: name,
			})
			if incomes != nil {
				// Income already exists, update it
				err = dbs.UpdateIncome(db, model.Income{
					Amount: &amount,
				},model.Income{
					Name: name,
				})
			} else {
				// Income doesn't exist, create a new one
				err = dbs.CreateIncome(db, model.Income{
					Name: name,
					Amount: &amount,
					ProfileID: profile.ID,
				})
			}
		}
	} else {
		incomes, _ := dbs.RetrieveIncome(db, model.Income{
			ProfileID: profile.ID,
		})
		for _, income := range incomes {
			amount := 0.00
			err = dbs.UpdateIncome(db, model.Income{
				Amount: &amount,
			},model.Income{
				Name: income.Name,
			})
		}
	}

	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error":"INCOMES NOT CREATED/UPDATED"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func RetrieveIncome(c *gin.Context, dbs DBHandler, db *sql.DB, httpClient request.HTTP, debug bool) {
	// Extract the session cookie or form
	uid, _ := c.Cookie("UID")
	uidP := c.Query("uid")

	// Retrieve the user's profile based on the cookie or form.
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

	var incomes []model.Income
	if err == nil {
		// Retrieve the user's expenses based on the profile ID.
		incomes, err = dbs.RetrieveIncome(db, model.Income{
			ProfileID: profile.ID,
		})
	}
	if err != nil || len(incomes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error":"INCOMES NOT FOUND"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"incomes": incomes,
	})
}