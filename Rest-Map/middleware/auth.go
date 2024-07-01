package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		//condition check 1
		tokenstring := c.GetHeader("Authorization")

		if tokenstring == "" {

			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		tokenparts := strings.Split(tokenstring, " ")

		//condition check 2
		//tokenParts[0] == "Bearer
		////tokenParts[1] represents the JWT token string that carries the actual authentication information.

		if len(tokenparts) != 2 || tokenparts[0] != "Bearer" {

			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			c.Abort()
			return
		}

		//condition check 3:it reads the token and checks if it's valid and hasn't been tampered with
		//Tokens are like secret messages.

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenparts[1], claims, func(token *jwt.Token) (interface{}, error) {

			return []byte("your_secret_key"), nil
		})

		if err != nil || !token.Valid {

			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		//conndition check 4 // It takes out the user_id from the token's secret message.

		if userid, ok := claims["user_id"].(float64); ok {

			c.Set("userid", uint(userid))
		} else {

			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

//condition check 1 :  if the Authorization header is present in the HTTP request
//condition check 2 : it expects the token to be in the format "Bearer <token>

//condition check 3:it reads the token and checks if it's valid and hasn't been tampered with
//Tokens are like secret messages. You need to decode (parse) them to read what's inside and verify (validate)

//comdition check 4 // It takes out the user_id from the token's secret message.
//This part is like finding out who is knocking on your door before letting them in.
