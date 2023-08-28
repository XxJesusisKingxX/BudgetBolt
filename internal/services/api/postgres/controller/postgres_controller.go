package controller

import (
	"database/sql"
	"net/http"
	plaidinterface "services/api/plaid"
	"services/api/postgres"
	"services/api/utils"
	"services/db/postgresql/model"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateProfile(c *gin.Context, dbs postgresinterface.DBHandler, db *sql.DB, debug bool) {
	user := c.PostForm("username")
	pass := c.PostForm("password")
	// Test if username is already taken
	profile, _ := dbs.RetrieveProfile(db, user)
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
	hashedPass, err = bcrypt.GenerateFromPassword([]byte(pass), saltRounds)
	if err != nil {
		api.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	err = dbs.CreateProfile(db, strings.ToLower(user), string(hashedPass))
	if err != nil {
		api.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func RetrieveProfile(c *gin.Context, dbs postgresinterface.DBHandler, db *sql.DB, debug bool) {
	user := c.Query("username")
	pass := c.Query("password")

	userProfile, _ := dbs.RetrieveProfile(db, strings.ToLower(user))
	if userProfile.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}
	auth := bcrypt.CompareHashAndPassword([]byte(userProfile.Password), []byte(pass))
	if auth == nil {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}
}

func RetrieveAccounts(c *gin.Context, dbs postgresinterface.DBHandler, db *sql.DB, debug bool) {
	user := c.Query("username")
	profile, err := dbs.RetrieveProfile(db, user)
	if err != nil {
		api.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	accounts, err := dbs.RetrieveAccount(db, profile.ID)
	if err != nil {
		api.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
	})
}

func RetrieveTransactions(c *gin.Context, dbs postgresinterface.DBHandler, db *sql.DB, debug bool) {
	user := c.Query("username")
	profile, err := dbs.RetrieveProfile(db, user)
	var transactions []model.Transaction
	if err == nil {
		transactions, err = dbs.RetrieveTransaction(db, model.Transaction{
			ProfileID: profile.ID,
			Query: model.Querys{
				Select: model.QueryParameters{
					Desc: true,
					OrderBy: "transaction_date",
			}}})
	}
	if err != nil {
		api.RenderError(c, err, plaidinterface.PlaidClient{})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
	//TODO investment transactions
    //TODO holdings
}
