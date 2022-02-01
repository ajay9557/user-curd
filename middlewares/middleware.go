package middlewares

import (
	"log"
	"net/http"
)

func Authentication(h http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Content type middleware called.")
		w.Header().Set("Content-Type", "application/json")
		userName, password, _ := r.BasicAuth()
		if userName == "himanshu" && password == "abcd" {
			log.Println("validated...")
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid user name or password"))
		}
	})
}
