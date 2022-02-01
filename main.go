package main

import (
	"database/sql"
	"fmt"
	"net/http"
	userHandler "zopsmart/user-curd/handler/user"
	"zopsmart/user-curd/middlewares"
	userService "zopsmart/user-curd/service/user"
	userStore "zopsmart/user-curd/store/user"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "UserInfo",
	}
	var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error in connection establishment!")
		return
	}
	store := userStore.New(db)
	service := userService.New(&store)
	h := userHandler.Handler{service}
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/user/{id}", middlewares.Auth(h.GetUserWithId)).Methods(http.MethodGet)
	r.HandleFunc("/users", h.AddUser).Methods(http.MethodPost)
	r.HandleFunc("/users", h.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", h.DeleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/users/{id}", h.UpdateUser).Methods(http.MethodPut)
	http.ListenAndServe(":8000", r)
}
