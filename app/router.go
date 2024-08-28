package app

import (
	"example/controller"
	"example/service"
	"net/http"
	"strings"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		fmt.Println(c)
		// Check if the Authorization header is present
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Optionally, you can validate the token here
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization token"})
			c.Abort()
			return
		}
		token := authHeader[7:]
		_, err := service.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}
		// If everything is fine, proceed to the next handler
		c.Next()
	}
}

func Router() {
	// make a group of authorization...
	auth := r.Group("/api/v1")
	auth.Use(AuthMiddleware())
	//BEGIN user endpoints
	r.POST("/api/v1/register", controller.Register)
	r.POST("/api/v1/login", controller.Login)
	//[] user

	//BEGIN admin endpoints
	auth.GET("/users", controller.GetUsers)
	auth.DELETE("/users", controller.DelUserUID)
	auth.POST("/users", controller.UsersUID)
	//[] admin
}
