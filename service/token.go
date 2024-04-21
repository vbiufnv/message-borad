package service

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"message-borad/model"
	"time"
)

var Secret = []byte("JwtSecret")

func CreateToken(username string) string {
	claim := model.Claim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "http://example.com", //签发人
		},
	}
	//Header&Payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	//+Signature
	tokenString, err := token.SignedString(Secret)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}
