package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joeytatu/restaurant-management-system/helpers"
)

// Authentication is a middleware function to authenticate requests using JWT
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Missing Authorization Header"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authorizationHeader, " ")[1]

		secretKey := getSecretKey()

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid Token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func getSecretKey() string {
	secretKey, err := helpers.GenerateSecretKey(16)
	if err != nil {
		fmt.Println("Error getting secret key. Error:", err)
	}
	return secretKey
}
