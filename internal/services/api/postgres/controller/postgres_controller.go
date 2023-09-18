package controller

import (
	"database/sql"
	"math"
	"net/http"
	plaidinterface "services/api/plaid"
	"services/api/postgres"
	"services/api/utils"
	"services/db/postgresql/model"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateProfile handles the creation of a user profile.
func CreateProfile(c *gin.Context, dbs postgresinterface.DBHandler, db *sql.DB, debug bool) {
	// Extract the username and password from the HTTP POST request.
	user := c.PostForm("username")
	pass := c.PostForm("password")

	// Test if username is already taken by attempting to retrieve the profile.
	profile, _ := dbs.RetrieveProfile(db, user, false)
	if profile.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{})
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
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}

	// Create the user profile with the hashed password.
	uid := utils.GenerateRandomString(64)
	err = dbs.CreateProfile(db, strings.ToLower(user), string(hashedPass), uid)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
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
func RetrieveProfile(c *gin.Context, dbs postgresinterface.DBHandler, db *sql.DB, debug bool) {
	// Extract the username and password from the HTTP GET request.
	user := c.PostForm("username")
	pass := c.PostForm("password")
	// Retrieve the user's profile based on the username.
	userProfile, _ := dbs.RetrieveProfile(db, user, false)
	if userProfile.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
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
func RetrieveAccounts(c *gin.Context, dbs postgresinterface.DBHandler, db *sql.DB, debug bool) {
	// Extract the session cookie
	uid, _ := c.Cookie("UID")

	// Retrieve the user's profile based on the username.
	profile, err := dbs.RetrieveProfile(db, uid, true)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}

	// Retrieve the user's accounts associated with the profile.
	accounts, err := dbs.RetrieveAccount(db, profile.ID)
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
	})
}

// RetrieveTransactions retrieves a user's transactions.
func RetrieveTransactions(c *gin.Context, dbs postgresinterface.DBHandler, db *sql.DB, debug bool) {
	// Extract the session cookie
	uid, _ := c.Cookie("UID")
	date := c.Query("date")

	// Retrieve the user's profile based on the username.
	profile, err := dbs.RetrieveProfile(db, uid, true)
	var transactions []model.Transaction
	if err == nil {
		// Retrieve the user's transactions based on the profile ID.
		transactions, err = dbs.RetrieveTransaction(db, model.Transaction{
			ProfileID: profile.ID,
			Query: model.Querys{
				Select: model.QueryParameters{
					GreaterThanEq: model.GreaterThanEq {
						Value: date,
						Column: "transaction_date",
					},
				},
			},
		})
		// update expense every time retrieve transactions
		updateExpenses(transactions, dbs, db, profile.ID)
	}

	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
	// TODO: Handle investment transactions
	// TODO: Handle holdings
}

// Retrieves a user's budgeted expenses.
func RetrieveExpenses(c *gin.Context, dbs postgresinterface.DBHandler, db *sql.DB, debug bool) {
	// Extract the session cookie
	uid, _ := c.Cookie("UID")

	// Retrieve the user's profile based on the username.
	profile, err := dbs.RetrieveProfile(db, uid, true)
	var expenses []model.Expense
	if err == nil {
		// Retrieve the user's expenses based on the profile ID.
		expenses, err = dbs.RetrieveExpense(db, model.Expense{
			ProfileID: profile.ID,
		})
	}
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"expenses": expenses,
	})
}

// Create a user's budgeted expenses.
func CreateExpenses(c *gin.Context, dbs postgresinterface.DBHandler, db *sql.DB, debug bool) {
	// Extract the session cookie
	uid, _ := c.Cookie("UID")
	name := c.PostForm("name")
	limit, _ := strconv.ParseFloat(c.PostForm("limit"), 64)
	spent, _ := strconv.ParseFloat(c.PostForm("spent"), 64)

	// Retrieve the user's profile based on the username.
	profile, err := dbs.RetrieveProfile(db, uid, true)
	if err == nil {
		// Create the user's expenses based on the profile ID.
		err = dbs.CreateExpense(db, model.Expense{
			ProfileID: profile.ID,
			Limit: &limit,
			Name: name,
			Spent: &spent,
		})
	}
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// Update a user's budgeted expenses.
func UpdateExpenses(c *gin.Context, dbs postgresinterface.DBHandler, db *sql.DB, debug bool) {
	// Extract the session cookie
	uid, _ := c.Cookie("UID")
	limit, _ := strconv.ParseFloat(c.PostForm("limit"), 64)
	id, _ := strconv.ParseInt(c.PostForm("id"), 10, 32)

	// Check if user exist
	profile, err := dbs.RetrieveProfile(db, uid, true)
	if err == nil && profile.ID != 0 {
		// Create the user's expenses based on the profile ID.
		err = dbs.UpdateExpense(db, model.Expense {
			Limit: &limit,
		}, model.Expense{
			ID: id,
		})
		
	}
	if err != nil {
		utils.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func updateExpenses(transactions []model.Transaction, dbs postgresinterface.DBHandler, db *sql.DB, profileId int) {
	// Get the category totals for all possible expenses
	categoryTotals := make(map[string]float64)
	for _, transaction := range transactions {
		primary := transaction.PrimaryCategory
		detailed := transaction.DetailCategory
		categoryTotals[primary] += transaction.Amount
		categoryTotals[detailed] += transaction.Amount
	}
	// Update expense total spent
	var expenses []model.Expense
	expenses, _ = dbs.RetrieveExpense(db, model.Expense{
		ProfileID: profileId,
	})
	for _, expense := range expenses {
		var total float64
		categoryList := strings.Split(expense.Category, ",")
		for _, category := range categoryList {
			total += categoryTotals[category]
		}
		total = math.Round(total*100) / 100 // round 2decimals
		dbs.UpdateExpense(db, model.Expense{
			Spent: &total,
		}, model.Expense{
			ID: expense.ID,
		})
	}
}