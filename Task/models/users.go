package models

type User struct {
	
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Age int `json:"age"`
}
type HttpResponse struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
}
type HttpErrs struct {
	ErrCode int `json:"err-code"`
	ErrMsg string `json:"err-msg"`

}
