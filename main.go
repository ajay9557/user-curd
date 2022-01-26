package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	userHandler "user-curd/http/users"
	userServices "user-curd/services/users"
	userStores "user-curd/stores/users"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

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
	r.HandleFunc("/insert", userHandle.PostUser).Methods("POST")
	r.HandleFunc("/users", userHandle.GetUsers).Methods("GET")
	r.HandleFunc("/user", userHandle.GetUserById).Methods("GET")
	r.HandleFunc("/update", userHandle.UpdateUser).Methods("PUT")
	r.HandleFunc("/delete", userHandle.DeleteUser).Methods("DELETE")
	fmt.Println("Listening at port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))

}
