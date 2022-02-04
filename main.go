package main

import (
	stores "user-curd/stores/Users"

	"database/sql"
	"fmt"
	"net/http"
	httplayer "user-curd/Http/Users"
	servicelayer "user-curd/services/Users"

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
	store := stores.New(db)
	service := servicelayer.New(store)
	ht := httplayer.Handler{service}

	fmt.Print("Server Starting....")
	m := mux.NewRouter().StrictSlash(true)
	m.HandleFunc("/Users", ht.GetAll).Methods("GET")
	m.HandleFunc("/Users/{id}", ht.DeleteId).Methods("DELETE")
	m.HandleFunc("/Update", ht.UpdateUser).Methods("PUT")
	m.HandleFunc("/Users/{id}", ht.Search).Methods("GET")
	m.HandleFunc("/Users", ht.Create).Methods("POST")
	http.ListenAndServe(":8050", m)
}
