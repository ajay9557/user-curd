package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test(t *testing.T) {

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value("Username")
		if val == nil {
			t.Error("username not present")
		}
		Password, ok := val.(string)
		if !ok {
			t.Error("not string")
		}
		if Password != "1234" {
			t.Error("wrong password")
		}
	})

	handlerToTest := Logger(nextHandler)

	req := httptest.NewRequest("GET", "http://testing", nil)

	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
}
