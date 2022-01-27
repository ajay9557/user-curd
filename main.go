package main

import (
	"database/sql"
	"fmt"
	"net/http"
	userHandler "zopsmart/user-curd/handler/user"
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
	u := userStore.New(db)
	p := userService.New(&u)
	ht := userHandler.Handler{p}
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/user/{id}", ht.UserWithId).Methods(http.MethodGet)
	r.HandleFunc("/user", ht.AddUser).Methods(http.MethodPost)
	r.HandleFunc("/user", ht.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/user/delete/{id}", ht.DeleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/user/update/{id}",ht.UpdateUser).Methods(http.MethodPatch)
	http.ListenAndServe(":8080", r)
}
