package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewRouter(e *gin.Engine) {

	e.Use(CORSMiddleware())

	// router.POST("/callback", handler.HandleLineMessage)

	fmt.Printf("Run")
}

func CORSMiddleware() gin.HandlerFunc {
	// allowOrigin := service.GetConfig("server.allowOrigin")
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "GET, POST")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Cache-Control, Strict-Transport-Security, token, user_token, account_id")
		c.Writer.Header().Add("Strict-Transport-Security", "max-age=31557600; includeSubDomains")
		c.Writer.Header().Add("Cache-Control", "no-cache, no-store")
		c.Writer.Header().Add("X-Frame-Options", "deny")
		c.Writer.Header().Add("Referrer-Policy", "no-referrer")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
