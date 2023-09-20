package main

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"services/api/postgres"
	postgresc "services/api/postgres/controller"
	"services/db/postgresql/model"
	"services/utils/testing"
	"strings"
	"testing"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateProfile_ProfileNameTaken(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/api/profile/create", func(c *gin.Context) {
		postgresc.CreateProfile(c,
			postgresinterface.MockDB{
				Profile: model.Profile{ID: 1},
			},
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("POST", "/api/profile/create", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Assert
	tests.Equals(t, http.StatusConflict, w.Code)
}
func TestCreateProfile_PasswordTooLong(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/api/profile/create", func(c *gin.Context) {
		postgresc.CreateProfile(c,
			postgresinterface.MockDB{
				Profile: model.Profile{ID: 0},
			},
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	form.Set("username", "test_user")
	form.Set("password", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz")
	req, _ := http.NewRequest("POST", "/api/profile/create", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Receive response
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isLong := strings.Contains(string(responseBody), "password length exceeds 72 bytes")
	// Assert
	tests.Equals(t, http.StatusInternalServerError, w.Code)
	tests.Equals(t, true, isLong)
}
func TestCreateProfile_ProfileNotCreated(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/api/profile/create", func(c *gin.Context) {
		postgresc.CreateProfile(c,
			postgresinterface.MockDB{
				Profile:    model.Profile{ID: 0},
				ProfileErr: errors.New("Failed to create profile"),
			},
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	form.Set("username", "test_user")
	form.Set("password", "abcdefghijklmnopqrstuvwxyz")
	req, _ := http.NewRequest("POST", "/api/profile/create", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Receive response
	responseBody, _ := io.ReadAll(w.Result().Body)
	defer w.Result().Body.Close()
	isCreated := !strings.Contains(string(responseBody), "\"Failed to create profile\"")
	// Assert
	tests.Equals(t, http.StatusInternalServerError, w.Code)
	tests.Equals(t, false, isCreated)
}
func TestCreateProfile_ProfileCreated(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.POST("/api/profile/create", func(c *gin.Context) {
		postgresc.CreateProfile(c,
			postgresinterface.MockDB{
				Profile: model.Profile{ID: 0},
			},
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	form.Set("username", "test_user")
	form.Set("password", "abcdefghijklmnopqrstuvwxyz")
	req, _ := http.NewRequest("POST", "/api/profile/create", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Assert
	tests.Equals(t, http.StatusOK, w.Code)
}





func TestRetrieveProfile_NotExist(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/api/profile/get", func(c *gin.Context) {
		postgresc.RetrieveProfile(c,
			postgresinterface.MockDB{
				Profile: model.Profile{
					ID: 0,
				},
			},
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	req, _ := http.NewRequest("GET", "/api/profile/get", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Assert
	tests.Equals(t, http.StatusNotFound, w.Code)
}
func TestRetrieveProfile_AuthFailed(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	r.GET("/api/profile/get", func(c *gin.Context) {
		postgresc.RetrieveProfile(c,
			postgresinterface.MockDB{
				Profile: model.Profile{
					ID: 1,
					Password: "abcdefghijklmnopqrstuvwxy",
				},
			},
			nil,
			true,
		)
	})
	// Create request
	form := url.Values{}
	form.Set("username", "test_user")
	form.Set("password", "abcdefghijklmnopqrstuvwxyz")
	req, _ := http.NewRequest("GET", "/api/profile/get", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Assert
	tests.Equals(t, http.StatusUnauthorized, w.Code)
}
func TestRetrieveProfile_AuthSucceed(t *testing.T) {
	// Create mock engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Handle mock route
	hashPass, _ := bcrypt.GenerateFromPassword([]byte("abc"), bcrypt.DefaultCost)
	r.POST("/api/profile/get", func(c *gin.Context) {
		postgresc.RetrieveProfile(c,
			postgresinterface.MockDB{
				Profile: model.Profile{
					Name:     "test",
					ID:       1,
					Password: string(hashPass),
				},
			},
			nil,
			true,
		)
	})
	// Create reque
	form := url.Values{}
	form.Set("username", "test")
	form.Set("password", "abc")
	req, _ := http.NewRequest("POST", "/api/profile/get", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Assert
	tests.Equals(t, http.StatusOK, w.Code)
}