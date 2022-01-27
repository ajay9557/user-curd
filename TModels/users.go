package TModels

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Age   int    `json:"age"`
}

type RetErr struct {
	StCode     int    `json:"st_code"`
	Errmessage string `json:"errmessage"`
}
