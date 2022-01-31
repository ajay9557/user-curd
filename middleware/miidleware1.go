package middleware

import (
	"log"
	"net/http"
)

func CT(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Content type middleware called.")
		w.Header().Set("Content-Type", "application/json")
		inner.ServeHTTP(w, r)
	})
}
