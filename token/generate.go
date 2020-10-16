package token

import (
	//	"encoding/json"
	//	"fmt"
	jwt "github.com/dgrijalva/jwt-go/v4"
	//	"io/ioutil"
	//	"log"
	//	"net/http"
	"time"
)

func GenerateToken() (string, error) {
	// Create the Claims
	mySigningKey := []byte(SECRET)

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(time.Hour * 4)),
		IssuedAt:  jwt.At(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(mySigningKey)
}
