package middleware

import (
	"crypto/rand"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const secretKey = "your_secret_key"

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

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid Token"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Invalid Token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func GenerateSecretKey(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	key := make([]byte, length)

	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	for i := range key {
		key[i] = charset[int(key[i])%len(charset)]
	}

	return string(key), nil
}
