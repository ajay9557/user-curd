package main

import (
	"database/sql"
	"fmt"
	"net/http"
	UserHandler "zopsmart/Task/http/users"
	"zopsmart/Task/middleware"
	UserService "zopsmart/Task/services/users"
	UserStore "zopsmart/Task/stores/users"

	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "root:yes@tcp(localhost:3306)/user")

	if err != nil {
		fmt.Println(err)
		fmt.Println("Connection establishment error")
	}

	userStore := UserStore.New(db)
	userService := UserService.New(userStore)
	userHandler := UserHandler.New(userService)



r := mux.NewRouter()

r.Handle("/user", middleware.Authentication(http.HandlerFunc(userHandler.CreateUser))).Methods("POST")
r.Handle("/user", middleware.Authentication(http.HandlerFunc(userHandler.GetUserById))).Methods("GET")
r.Handle("/user", middleware.Authentication(http.HandlerFunc(userHandler.DeleteUser))).Methods("DELETE")
r.Handle("/user", middleware.Authentication(http.HandlerFunc(userHandler.UpdateUser))).Methods("PUT")
r.Handle("/user", middleware.Authentication(http.HandlerFunc(userHandler.AllUserDetails))).Methods("GET")

err = http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)
	}
}