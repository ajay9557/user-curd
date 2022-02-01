package main

import (
	"fmt"
	"layer/user/driver"
	userHttp "layer/user/http/users"
	"layer/user/middlewares/auth"
	userServices "layer/user/services/users"
	userstore "layer/user/stores/users"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := driver.ConnectToMySql("mysql")

	if err != nil {
		// TODO: Handle error
		return
	}

	st := userstore.New(db)
	sr := userServices.New(st)
	// handler := userHttp.Handler{S: sr}
	handler := userHttp.New(sr)

	router := mux.NewRouter()
	router.Path("/api/users/{id}").Methods("GET").Handler(func() http.Handler {
		return auth.ValidateEmail(http.HandlerFunc(handler.GetUserByIdHandler))
	}())

	router.Path("/api/users").Methods("GET").Handler(func() http.Handler {
		return auth.ValidateEmail(http.HandlerFunc(handler.GetUsersHandler))
	}())
	router.Path("/api/users/{id}").Methods("PUT").Handler(func() http.Handler {
		return auth.ValidateEmail(http.HandlerFunc(handler.UpdateUserHandler))
	}())
	router.Path("/api/users/{id}").Methods("DELETE").Handler(func() http.Handler {
		return auth.ValidateEmail(http.HandlerFunc(handler.DeleteUserHandler))
	}())
	router.Path("/api/users").Methods("POST").HandlerFunc(handler.CreateUserHandler)

	http.Handle("/", router)

	fmt.Println("Listening to port 3000")
	err = http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)
	}
}
