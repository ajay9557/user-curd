package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	handler "user-crud/handler/users"
	"user-crud/middlewares"
	service "user-crud/services/users"
	store "user-crud/stores/users"
)

func main() {
	db, err := sql.Open("mysql", "test:1234@/test")
	if err != nil {
		log.Print("could not open mysql")
	}
	st := store.New(db)
	sr := service.New(st)

	handler := handler.UserHandler{sr}

	r := mux.NewRouter()

	r.HandleFunc("/user", middlewares.Authentication(http.HandlerFunc(handler.Insert))).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", handler.Update).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", handler.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/user/{id}", handler.GetById).Methods(http.MethodGet)
	r.HandleFunc("/user", handler.GetAll).Methods(http.MethodGet)
	err = http.ListenAndServe(":3000", r)
	println(err.Error())
}
