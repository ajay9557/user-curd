package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	UserHandler "zopsmart/Task/http"
	UserService "zopsmart/Task/services"
	UserStore "zopsmart/Task/stores"

	"github.com/gorilla/mux"
)


func main() {
	db,err := sql.Open("mysql","root:raramuri@localhost(3530)/user")

	if err != nil {
		fmt.Println(err)
		fmt.Println("Connection establishment error")
	}
	defer db.Close()

	userStore := UserStore.New(db)
	userService := UserService.New(userStore)
	userHandler := UserHandler.New(userService)

	r := mux.NewRouter()
	r.HandleFunc("/insert", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/user",userHandler.GetUserById).Methods("GET")
	r.HandleFunc("/delete",userHandler.DeleteUser).Methods("DELETE")
	r.HandlerFunc("/update",userHandler.UpdateUser).Methods("PUT")
	fmt.Println("Listening to 3036")
	log.Fatal(http.ListenAndServe(":3036",r))
		
}
