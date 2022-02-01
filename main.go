package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
	userHandler "github.com/tejas/user-crud/http/users"
	"github.com/tejas/user-crud/middleware"
	userService "github.com/tejas/user-crud/services/users"
	userStore "github.com/tejas/user-crud/store/users"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")

	if err != nil {
		fmt.Println(err)
	}

	userStore := userStore.New(db)
	userService := userService.New(userStore)
	handler := userHandler.New(userService)

	router := mux.NewRouter()

	router.Handle("/api/user/{id}", middleware.Authentication(http.HandlerFunc(handler.FindUserById))).Methods(http.MethodGet)
	router.Handle("/api/users", middleware.Authentication(http.HandlerFunc(handler.FindUsers))).Methods(http.MethodGet)
	router.Handle("/api/user/{id}", middleware.Authentication(http.HandlerFunc(handler.UpdateById))).Methods(http.MethodPut)
	router.Handle("/api/user", middleware.Authentication(http.HandlerFunc(handler.CreateUser))).Methods(http.MethodPost)
	router.Handle("/api/user/{id}", middleware.Authentication(http.HandlerFunc(handler.DeleteUser))).Methods(http.MethodDelete)

	http.Handle("/", router)

	fmt.Println("Listening to port 8080...")
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}
