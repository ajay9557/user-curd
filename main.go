package main

import (
	"database/sql"
	"errors"
	"fmt"
	"go_lang/Assignment/user-curd/http/httpuser"
	"go_lang/Assignment/user-curd/middleware"
	service "go_lang/Assignment/user-curd/services/user"
	"go_lang/Assignment/user-curd/stores/user"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Connect() *sql.DB {
	connection, error := sql.Open("mysql", "vicky:root@/user")
	if error != nil {
		fmt.Println(errors.New("ERROR IN CONNECTING TO DATABASE"))
		fmt.Println(error)
	}

	return connection
}

func AddMiddle(handler http.Handler, midware func(handler http.Handler) http.Handler) http.Handler {
	return midware(handler)
}

func main() {

	connection := Connect()

	store := user.New(connection)

	services := service.New(store)

	httpService := httpuser.HttpService{HttpServ: services}

	router := mux.NewRouter()
	router.HandleFunc("/id", httpService.Handler)
	http.Handle("/id", AddMiddle(http.HandlerFunc(httpService.Handler), middleware.Filter))

	router.HandleFunc("/", httpService.Handler)
	http.Handle("/", AddMiddle(http.HandlerFunc(httpService.Handler), middleware.Filter))

	fmt.Println("Listning on port: 5454")
	http.ListenAndServe(":5454", nil)
}
