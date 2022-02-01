package auth

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// Used to validate zopsmart email id. Only email id ending with zopsmart.com can access api.
func ValidateEmail(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-type", "application.json")

		token := strings.Split(req.Header.Get("Authorization"), " ")

		if len(token) != 2 {
			_, _ = res.Write([]byte(`{"error": "Could not parse authentication token"}`))
			return
		}

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(token[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("mysecret"), nil
		})

		if err != nil {
			_, _ = res.Write([]byte(`{"error": "Invalid Token. Could not parse"}`))
			return
		}

		email := fmt.Sprint(claims["email"])

		re := regexp.MustCompile(`.*zopsmart.com`)
		isValid := re.MatchString(email)

		if !isValid {
			_, _ = res.Write([]byte(`{"error": "Only zopsmart member can access this functionality"}`))
			return
		}

		h.ServeHTTP(res, req)
	})
}
