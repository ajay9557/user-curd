package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	UserHandler "zopsmart/Task/http/users"
	UserService "zopsmart/Task/services/users"
	UserStore "zopsmart/Task/stores/users"

	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "root:raramuri@localhost(8000)/users")

	if err != nil {
		fmt.Println(err)
		fmt.Println("Connection establishment error")
	}
	defer db.Close()

	userStore := UserStore.New(db)
	userService := UserService.New(userStore)
	userHandler := UserHandler.New(userService)

	r := mux.NewRouter()
	r.HandleFunc("/user", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", userHandler.GetUserById).Methods("GET")
	r.HandleFunc("/user/{id}", userHandler.DeleteUser).Methods("DELETE")
	r.HandleFunc("/user/{id}", userHandler.UpdateUser).Methods("PUT")
	fmt.Println("Listening to 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
