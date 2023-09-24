package api

import (
	budget "services/internal/budgeting/db/model"
	transaction "services/internal/transaction_history/db/model"
	"services/internal/utils"

	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateProfile handles the creation of a user profile.
func CreateProfile(c *gin.Context, dbs DBHandler, db map[string]*sql.DB, debug bool) {
	// Extract the username and password from the HTTP POST request.
	user := c.PostForm("username")
	pass := c.PostForm("password")

	// Test if username is already taken by attempting to retrieve the profile.
	profile, _ := dbs.RetrieveProfile(db["user"], user, false)
	if profile.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error":"PROFILE ALREADY EXISTS"})
		return
	}

	var err error
	var hashedPass []byte
	saltRounds := 17
	if debug {
		saltRounds = 1
	}
	// Generate a bcrypt hash of the user's password.
	hashedPass, err = bcrypt.GenerateFromPassword([]byte(pass), saltRounds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"PASSWORD TOO LONG"})
		return
	}

	// Create the user profile with the hashed password.
	uid := utils.GenerateRandomString(64)
	err = dbs.CreateProfile(db["user"], strings.ToLower(user), string(hashedPass), uid)
	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error":"PROFILE NOT CREATED"})
		return
	}

	// Set a cookie with the SID token
	cookie := &http.Cookie{
		Name:     "UID",
		Value:    uid,
		Expires:  time.Now().Add(24 * time.Hour), // Set an expiration time
		HttpOnly: false,                           // Cookie is not accessible via JavaScript
		Secure:   false,                           // Cookie is transmitted over HTTPS only
		Path: "/",
	}
	http.SetCookie(c.Writer, cookie)
	// test cookie creation
	c.JSON(http.StatusOK, gin.H{
		"uid": uid,
	})
}

// RetrieveProfile retrieves a user's profile and checks the provided password.
func RetrieveProfile(c *gin.Context, dbs DBHandler, db map[string]*sql.DB, debug bool) {
	// Extract the username and password from the HTTP GET request.
	user := c.PostForm("username")
	pass := c.PostForm("password")
	// Retrieve the user's profile based on the username.
	userProfile, _ := dbs.RetrieveProfile(db["user"], user, false)
	if userProfile.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error":"PROFILE NOT FOUND"})
		return
	}

	// Compare the provided password with the stored hashed password.
	auth := bcrypt.CompareHashAndPassword([]byte(userProfile.Password), []byte(pass))
	if auth == nil {
		// Set a cookie with the session token
		cookie := &http.Cookie{
			Name:     "UID",
			Value:    userProfile.RandomUID,
			Expires:  time.Now().Add(24 * time.Hour), // Set an expiration time
			HttpOnly: false,                           // Cookie is not accessible via JavaScript
			Secure:   false,                           // Cookie is transmitted over HTTPS only
			Path: "/",
		}
		http.SetCookie(c.Writer, cookie)
		
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}
}

// RetrieveAccounts retrieves a user's accounts.
func RetrieveAccounts(c *gin.Context, dbs DBHandler, db map[string]*sql.DB, debug bool) {
	// Extract the session cookie
	uid, _ := c.Cookie("UID")

	// Retrieve the user's profile based on the username.
	profile, err := dbs.RetrieveProfile(db["user"], uid, true)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"PROFILE NOT FOUND"})
		return
	}

	// Retrieve the user's accounts associated with the profile.
	accounts, err := dbs.RetrieveAccount(db["transaction"], profile.ID)
	if err != nil || len(accounts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error":"ACCOUNT NOT FOUND"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
	})
}

// RetrieveTransactions retrieves a user's transactions.
func RetrieveTransactions(c *gin.Context, dbs DBHandler, db map[string]*sql.DB, debug bool) {
	// Extract the session cookie
	uid, _ := c.Cookie("UID")
	date := c.Query("date")

	// Retrieve the user's profile based on the username.
	profile, err := dbs.RetrieveProfile(db["user"], uid, true)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"PROFILE NOT FOUND"})
		return
	}

	var transactions []transaction.Transaction
	if err == nil {
		// Retrieve the user's transactions based on the profile ID.
		transactions, err = dbs.RetrieveTransaction(db["transaction"], transaction.Transaction{
			ProfileID: profile.ID,
			Query: transaction.Querys{
				Select: transaction.QueryParameters{
					GreaterThanEq: transaction.GreaterThanEq {
						Value: date,
						Column: "transaction_date",
					},
				},
			},
		})
		if err != nil || len(transactions) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error":"TRANSACTIONS NOT FOUND"})
			return
		}
		// update expense every time retrieve transactions
		UpdateAllExpenses(transactions, dbs, db["budget"], profile.ID)
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}

// Retrieves a user's budgeted expenses.
func RetrieveExpenses(c *gin.Context, dbs DBHandler, db map[string]*sql.DB, debug bool) {
	// Extract the session cookie
	uid, _ := c.Cookie("UID")

	// Retrieve the user's profile based on the username.
	profile, err := dbs.RetrieveProfile(db["user"], uid, true)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"PROFILE NOT FOUND"})
		return
	}

	var expenses []budget.Expense
	if err == nil {
		// Retrieve the user's expenses based on the profile ID.
		expenses, err = dbs.RetrieveExpense(db["budget"], budget.Expense{
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
func CreateExpenses(c *gin.Context, dbs DBHandler, db map[string]*sql.DB, debug bool) {
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
	profile, err := dbs.RetrieveProfile(db["user"], uid, true)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"PROFILE NOT FOUND"})
		return
	}

	if err == nil {
		// Create the user's expenses based on the profile ID.
		err = dbs.CreateExpense(db["budget"], budget.Expense{
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
func UpdateExpenses(c *gin.Context, dbs DBHandler, db map[string]*sql.DB, debug bool) {
	// Extract the session cookie
	uid, _ := c.Cookie("UID")
	limit, limitErr := strconv.ParseFloat(c.PostForm("limit"), 64)
	id, idErr := strconv.ParseInt(c.PostForm("id"), 10, 32)
	if limitErr != nil || idErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"INVALID LIMIT AND/OR ID"})
		return
	}
	// Check if user exist
	_, err := dbs.RetrieveProfile(db["user"], uid, true)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":"PROFILE NOT FOUND"})
		return
	}
	if err == nil {
		// Create the user's expenses based on the profile ID.
		err = dbs.UpdateExpense(db["budget"], budget.Expense {
			Limit: &limit,
		}, budget.Expense{
			ID: id,
		})
		if err != nil {
			c.JSON(http.StatusNotImplemented, gin.H{"error":"EXPENSES NOT UPDATED"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}