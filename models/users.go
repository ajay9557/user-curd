package models

type User struct {
	Id    int
	Name  string
	Email string
	Phone string
	Age   int
}

type Response struct {
	Data          interface{} `json:"data"`
	Message       string      `json:"message"`
	StatusCode int         `json:"statusCode"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
