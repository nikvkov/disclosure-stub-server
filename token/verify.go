package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"net/http"
	"strings"
)

const SECRET = "SRD2-STUB"

func VerifyToken(r *http.Request) (bool, bool) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	if len(splitToken) != 2 {
		return false, false
	}

	reqToken = strings.TrimSpace(splitToken[1])
	token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err == nil && token.Valid {
		fmt.Println("valid token")
		return true, true
	} else {
		fmt.Println("invalid token")
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET), nil
		})
		if err != nil {
			if strings.Contains(err.Error(), "token is malformed: failed to decode token claims") {
				return true, false
			}
		}
		return false, false
	}
	return false, false
}
