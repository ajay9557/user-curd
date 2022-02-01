package main

import (
	"database/sql"
	"log"
	"net/http"
	storeUser "user-curd/datastore/users"
	"user-curd/driver"
	httpUser "user-curd/http/users"
	"user-curd/middleware"
	serviceUser "user-curd/service/users"

	"github.com/gorilla/mux"
)

func main() {
	// define the mysql configuration
	sqlConf := driver.MySQLConfig{
		Driver:   "mysql",
		Host:     "localhost",
		User:     "vips",
		Password: "1234",
		Port:     "3306",
		Db:       "users",
	}

	// handle opening sql connection
	db, err := driver.ConnectToMySQL(sqlConf)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("error closing connection to sql %v", err)
		}
	}(db)
	if err != nil {
		log.Printf("error connecting to sql server %v", err)
	}

	// define each layer handlers
	s := storeUser.New(db)
	sv := serviceUser.New(s)
	usrHandler := httpUser.New(sv)

	// define mux and routes with their handlers
	r := mux.NewRouter()
	r.Handle("/users", middleware.Authentication(http.HandlerFunc(usrHandler.GetAllUserHandler))).Methods(http.MethodGet)
	r.Handle("/users/{id}", middleware.Authentication(http.HandlerFunc(usrHandler.GetUserByIdHandler))).Methods(http.MethodGet)
	r.Handle("/users", middleware.Authentication(http.HandlerFunc(usrHandler.CreateUserHandler))).Methods(http.MethodPost)
	r.Handle("/users/{id}", middleware.Authentication(http.HandlerFunc(usrHandler.UpdateUserHandler))).Methods(http.MethodPut)
	r.Handle("/users/{id}", middleware.Authentication(http.HandlerFunc(usrHandler.DeleteUserHandler))).Methods(http.MethodDelete)

	// Run the server
	log.Printf("Listening on port 8000...")
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Printf("error creating server: %v", err)
	}
}
