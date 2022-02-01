package middleware

import (
	"encoding/json"
	"go_lang/Assignment/user-curd/models"
	"go_lang/Assignment/user-curd/services/user"
	"net/http"

	"github.com/gorilla/mux"
)

func Filter(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ide := mux.Vars(r)["id"]
		isValid := user.IsValidId(ide)
		if ide == "" {
			handler.ServeHTTP(rw, r)
		}
		if !isValid {
			responseMsg := models.ErrorResponse{Code: http.StatusBadRequest, Message: "INVALID ID"}
			jsonResp, _ := json.Marshal(responseMsg)
			rw.Write(jsonResp)
			return
		}
		handler.ServeHTTP(rw, r)
	})
}
