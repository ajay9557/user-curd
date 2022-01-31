package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
	"zopsmart/Task/models"

	"github.com/dgrijalva/jwt-go"
)

func Authentication(h http.Handler) http.Handler {
	// load project env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("error loading env file %v", err)
	}

	authJ := http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		authHeader := strings.Split(req.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			wr.Header().Set("content-type", "application/json")
			wr.WriteHeader(http.StatusUnauthorized)
			respErr, _ := json.Marshal(models.HttpErrs{ErrMsg: "Unauthorized", ErrCode: http.StatusUnauthorized})
			wr.Write(respErr)
		} else {
			// get the jwt token if it exists
			jwtToken := authHeader[1]

			// check if the token is signed with correct algo and signed with the given secret key
			token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("SECRETKEY")), nil
			})

			// check if the token is valid
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				v := true
				if !(claims["username"] != "" && strings.Split(claims["username"].(string), "@")[1] == "zopsmart.com") {
					v = false
				}
				// can get from the database
				if !(claims["password"] != "" && claims["password"].(string) == "1234") {
					v = false
				}

				if v {
					h.ServeHTTP(wr, req)
				} else {
					wr.Header().Set("content-type", "application/json")
					wr.WriteHeader(http.StatusUnauthorized)
					respErr, _ := json.Marshal(models.HttpErrs{ErrMsg: "Unauthorized", ErrCode: http.StatusUnauthorized})
					wr.Write(respErr)
				}
			} else {
				wr.Header().Set("content-type", "application/json")
				wr.WriteHeader(http.StatusUnauthorized)
				respErr, _ := json.Marshal(models.HttpErrs{ErrMsg: "Unauthorized", ErrCode: http.StatusUnauthorized})
				wr.Write(respErr)
			}
		}
	})

	return authJ
}

func BasicAuthentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		email, pass, ok := req.BasicAuth()
		if ok {
			emailHash := sha256.Sum256([]byte(email))
			passHash := sha256.Sum256([]byte(pass))
			expectedEmailHash := sha256.Sum256([]byte("prasath.k@zopsmart.com"))
			expectedPassHash := sha256.Sum256([]byte("1234"))

			emailMatch := (subtle.ConstantTimeCompare(emailHash[:], expectedEmailHash[:])) == 1
			passMatch := (subtle.ConstantTimeCompare(passHash[:], expectedPassHash[:])) == 1

			if emailMatch && passMatch {
				h.ServeHTTP(wr, req)
				return
			}
		}
		wr.Header().Set("content-type", "application/json")
		wr.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		respErr, _ := json.Marshal(models.HttpErrs{ErrMsg: "Unauthorized access", ErrCode: http.StatusUnauthorized})
		wr.WriteHeader(http.StatusUnauthorized)
		wr.Write(respErr)
	})
}