package main

import (
	"database/sql"
	"errors"
	"fmt"
	"go_lang/Assignment/user-curd/http/httpuser"
	service "go_lang/Assignment/user-curd/services/user"
	"go_lang/Assignment/user-curd/stores/user"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	connection, error := sql.Open("mysql", "vicky:root@/user")
	if error != nil {
		fmt.Println(errors.New("ERROR IN CONNECTING TO DATABASE"))
		fmt.Println(error)
	}

	return connection
}

func main() {

	connection := Connect()

	store := user.New(connection)

	services := service.New(store)

	httpService := httpuser.HttpService{HttpServ: services}

	http.HandleFunc("/id", httpService.Handler)
	http.HandleFunc("/", httpService.Handler)
	fmt.Println("Listning on port: 5454")
	http.ListenAndServe(":5454", nil)
}
