package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
	// condition check 1 :  if the Authorization header is present in the HTTP request	
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
			c.Abort()
			return
		}

		// Token usually comes in the format "Bearer <token>"
		//condition check 2 it expects the token to be in the format "Bearer <token>
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			//tokenParts[0] == "Bearer
			//tokenParts[1] represents the JWT token string that carries the actual authentication information.
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			c.Abort()
			return
		}
               //condition check 3 //  It reads the token and checks if it's valid and hasn't been tampered with
		//Tokens are like secret messages. You need to decode (parse) them to read what's inside and verify (validate) that they're real and not fake.
		
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenParts[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// Set user ID in context
		//comdition check 4 // It takes out the user_id from the token's secret message.
		//This part is like finding out who is knocking on your door before letting them in.
		if userID, ok := claims["user_id"].(float64); ok {
			c.Set("userID", int(userID))
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Next()  //If everything is okay (token is there, format is right, it's valid, and you found out who it is), 
		//you let the request continue to the next step, like showing the visitor into your house.
	}
}
