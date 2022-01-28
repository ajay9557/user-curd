package main

import (
	"user-curd/stores/Users"

	"database/sql"
	"fmt"
	"net/http"
	hl "user-curd/Http/Users"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "nayani:1234(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error in connection establishment!")
		return
	}
	defer db.Close()
	st := Users.New(db)
	s := sl.new(st)
	ht := hl.Handlers{s}

	fmt.Print("Server Starting....")
	m := mux.NewRouter().StrictSlash(true)
	m.HandleFunc("/Users", ht.GetAll).Methods("GET")
	m.HandleFunc("/Users/{id}", ht.DeleteId).Methods("DELETE")
	m.HandleFunc("/Update", ht.UpdateUserDetails).Methods("PUT")
	m.HandleFunc("/Users/{id}", ht.Search).Methods("GET")
	m.HandleFunc("/Users", ht.Create).Methods("POST")
	http.ListenAndServe(":8050", m)
}
