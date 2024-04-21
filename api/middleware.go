package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"message-borad/model"
	"net/http"
	"strings"
)

var Secret = []byte("JwtSecret")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := GetToken(c)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, "authorization token err")
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &model.Claim{}, func(token *jwt.Token) (interface{}, error) {
			return Secret, nil
		})
		if err != nil {
			log.Fatal(err)
		}

		username := token.Claims.(*model.Claim).Username

		// 传递信息
		c.Set("username", username)

	}
}

func GetToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")
	if bearerToken == "" {
		return ""
	}
	part := strings.SplitN(bearerToken, " ", 2)
	if len(part) != 2 || part[0] != "Bearer" {
		return ""
	}
	return part[1]
}
