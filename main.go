package main

import (
	"database/sql"
	"log"
	"net/http"
	storeUser "user-curd/datastore/users"
	"user-curd/driver"
	httpUser "user-curd/http/users"
	srvcUser "user-curd/service/users"

	"github.com/gorilla/mux"
)

func main() {

	sqlConf := driver.MySQLConfig{
		Host:     "localhost",
		User:     "vips",
		Password: "1234",
		Port:     "3306",
		Db:       "users",
	}

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

	str := storeUser.New(db)
	serv := srvcUser.New(str)
	usrHandler := httpUser.New(serv)

	r := mux.NewRouter()
	r.HandleFunc("/user", usrHandler.GetAllUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", usrHandler.GetUserByIdHandler).Methods(http.MethodGet)
	r.HandleFunc("/user", usrHandler.CreateUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", usrHandler.UpdateUserHandler).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", usrHandler.DeleteUserHandler).Methods(http.MethodDelete)

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		return
	}
}
