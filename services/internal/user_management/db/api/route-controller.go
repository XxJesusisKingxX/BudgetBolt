package api

import (
	"services/internal/user_management/db/model"
	"services/internal/utils"
	"strconv"

	"database/sql"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// CreateProfile handles the creation of a user profile.
func CreateProfile(c *gin.Context, dbs DBHandler, db *sql.DB, debug bool) {
	// Extract the username and password from the HTTP POST request.
	user := c.PostForm("username")
	pass := c.PostForm("password")

	// Test if username is already taken by attempting to retrieve the profile.
	profile, _ := dbs.RetrieveProfile(db, user, false)
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
	err = dbs.CreateProfile(db, strings.ToLower(user), string(hashedPass), uid)
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
func RetrieveProfile(c *gin.Context, dbs DBHandler, db *sql.DB, debug bool) {
	// Extract the username and password or UID from the HTTP GET request.
	user := c.PostForm("username")
	pass := c.PostForm("password")
	uid := c.PostForm("uid")

	// Retrieve the user's profile based on the username or UID.
	var userProfile model.Profile
	if len(uid) == 0 {
		userProfile, _ = dbs.RetrieveProfile(db, user, false)
	} else {
		userProfile, _ = dbs.RetrieveProfile(db, uid, true)
	}
	if userProfile.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error":"PROFILE NOT FOUND"})
		return
	}

	// Compare the provided password with the stored hashed password.
	if len(uid) == 0 {
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
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"profile_id": userProfile.ID,
	})
}

func CreateToken(c *gin.Context, dbs DBHandler, db *sql.DB, debug bool) {
	itemId := c.PostForm("itemId")
	token := c.PostForm("token")
	id, idErr := strconv.ParseInt(c.PostForm("id"), 10, 32)
	if idErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	err := dbs.CreateToken(db, model.Token{
		ProfileID: id,
		Item: itemId,
		Token: token,
	})

	if err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error":"TOKEN NOT CREATED"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func RetrieveToken(c *gin.Context, dbs DBHandler, db *sql.DB, debug bool) {
	uid := c.Query("uid")
	profile, err := dbs.RetrieveProfile(db, uid, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	token, err := dbs.RetrieveToken(db, profile.ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokens": token,
	})
}