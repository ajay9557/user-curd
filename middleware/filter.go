package middleware

import (
	"encoding/json"
	"go_lang/Assignment/user-curd/models"
	"go_lang/Assignment/user-curd/services/user"
	"net/http"
)

func Filter(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ide := r.URL.Query().Get("id")
		isValid := user.IsValidId(ide)
		if !isValid {
			responseMsg := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "INVALID ID"}
			jsonResp, _ := json.Marshal(responseMsg)
			rw.Write(jsonResp)
			return
		}
		handler.ServeHTTP(rw, r)
	})
}
