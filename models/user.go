package models

type User struct {
	Id    int    `json: "id"`
	Name  string `json: "name"`
	Email string `json: "email"`
	Phone string `json: "phone"`
	Age   int    `json: "age"`
}

type ErrorResponse struct {
	StatusCode   int    `json: "statuscode"`
	ErrorMessage string `json: "error"`
}
