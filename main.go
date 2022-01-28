package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	userHandler "user-curd/http/users"
	"user-curd/middlewares"
	userServices "user-curd/services/users"
	userStores "user-curd/stores/users"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func addMiddlewares(h http.HandlerFunc, middlewares ...func(handlerFun http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func main() {
	db, err := sql.Open("mysql", "gopi:gopi@123@tcp(0.0.0.0:3306)/usersDB")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	userStore := userStores.New(db) //store layer
	userService := userServices.New(userStore)
	userHandle := userHandler.New(userService)
	r := mux.NewRouter()
	r.Handle("/api/log", addMiddlewares(userHandle.GetUsers, middlewares.BasicAuth))
	r.HandleFunc("/insert", userHandle.PostUser).Methods("POST")
	r.HandleFunc("/users", userHandle.GetUsers).Methods("GET")
	r.HandleFunc("/user", userHandle.GetUserById).Methods("GET")
	r.HandleFunc("/update", userHandle.UpdateUser).Methods("PUT")
	r.HandleFunc("/delete", userHandle.DeleteUser).Methods("DELETE")
	fmt.Println("Listening at port 8086")
	log.Fatal(http.ListenAndServe(":8086", r))

}
