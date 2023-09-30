package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup, db *sql.DB) {
    router.POST("profile/create", func(c *gin.Context) {
        CreateProfile(c, DB{}, db, false)
    })
    router.POST("profile/get", func(c *gin.Context) {
        RetrieveProfile(c, DB{}, db, false)
    })
    router.POST("token/create", func(c *gin.Context) {
        CreateToken(c, DB{}, db, false)
    })
    router.GET("token/get", func(c *gin.Context) {
        RetrieveToken(c, DB{}, db, false)
    })
}