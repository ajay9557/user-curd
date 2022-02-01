package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"user-curd/middlewares"
	"user-curd/stores/users"

	servicelayer "user-curd/services/users"

	httplayer "user-curd/http/users"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "SudheerKumar:Puppala@tcp(127.0.0.1:3306)/Task")
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	store := users.New(db)
	service := servicelayer.New(store)
	ht := httplayer.Handler{service}

	fmt.Print("Server Starting....")
	m := mux.NewRouter().StrictSlash(true)
	m.HandleFunc("/users", ht.GetAll).Methods("GET")
	m.HandleFunc("/users/{id}", middlewares.Logger(ht.DeleteId)).Methods("DELETE")
	m.HandleFunc("/users", middlewares.Logger(ht.UpdateUserDetails)).Methods("PUT")
	m.HandleFunc("/users/{id}", middlewares.Logger(ht.Search)).Methods("GET")
	m.HandleFunc("/users", middlewares.Logger(ht.Create)).Methods("POST")
	http.ListenAndServe(":8040", m)
}
