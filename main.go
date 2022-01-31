package main

import (
	"database/sql"

	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	user1 "Icrud/Http/Users"
	"Icrud/Middlewares"
	user2 "Icrud/Services/Users"
	user3 "Icrud/Stores/Users"
)

func main() {

	db, err := sql.Open("mysql", "root:satyasusi@tcp(127.0.0.1:3306)/test")

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error in connection establishment!")
		return
	}

	defer db.Close()

	fmt.Println("Server starting....")

	sudb := user3.New(db)

	su := user2.New(sudb)

	ht := user1.Handler{Sev: su}

	r := mux.NewRouter()
	// r.HandleFunc("/api/users/{id}", Middlewares.Logger(ht.UserById)).Methods("GET")
	r.Path("/api/users/{id}").Methods("GET").HandlerFunc(ht.UserById)

	// r.HandleFunc("/api/users", ht.GetUsers).Methods("GET")
	r.Path("/api/users").Methods("GET").HandlerFunc(ht.GetUsers)

	// r.HandleFunc("/api/users", ht.InsertUser).Methods("POST")
	r.Path("/api/users").Methods("POST").HandlerFunc(ht.InsertUser)

	r.HandleFunc("/api/users/{id}", ht.UpdateUserById).Methods("PUT")

	// r.HandleFunc("/api/users/{id}", ht.DeleteUserById).Methods("DELETE")
	r.HandleFunc("/api/users/{id}", Middlewares.Logger(ht.DeleteUserById)).Methods("DELETE")

	http.Handle("/", r)

	log.Println("Listening at :3000")
	http.ListenAndServe(":3000", nil)
}
