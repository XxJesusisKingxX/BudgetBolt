package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	plaidinterface "services/api/plaid"
	"services/api/postgres"
	"services/api/utils"
	"services/db/postgresql/model"
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
	fmt.Println(auth)
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

	// Retrieve the user's profile based on the username.
	profile, err := dbs.RetrieveProfile(db, uid, true)
	var transactions []model.Transaction
	if err == nil {
		// Retrieve the user's transactions based on the profile ID.
		transactions, err = dbs.RetrieveTransaction(db, model.Transaction{
			ProfileID: profile.ID,
			Query: model.Querys{
				Select: model.QueryParameters{
					Desc:     true,
					OrderBy: "transaction_date",
				},
			},
		})
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
