package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	userHttp "github.com/tejas/user-crud/http/users"
	userServices "github.com/tejas/user-crud/services/users"
	userStore "github.com/tejas/user-crud/stores/users"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")

	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection Error")
	}

	st := userStore.New(db)
	sr := userServices.New(st)
	handler := userHttp.New(sr)

	router := mux.NewRouter()
	router.Path("/api/users/{id}").Methods("GET").HandlerFunc(handler.FindUserById)
	router.Path("/api/users").Methods("GET").HandlerFunc(handler.FindAllUsers)
	router.Path("/api/users/{id}").Methods("PUT").HandlerFunc(handler.UpdateUserById)
	router.Path("/api/users/{id}").Methods("DELETE").HandlerFunc(handler.DeleteUserById)
	router.Path("/api/users").Methods("POST").HandlerFunc(handler.CreateUsers)

	http.Handle("/", router)

	fmt.Println("Listening to port 3000")
	err = http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println(err)
	}

}
