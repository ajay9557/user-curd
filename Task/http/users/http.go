package http

import (
	
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
	"strconv"

	"zopsmart/Task/services"
	"zopsmart/Task/models"
	_ "github.com/go-sql-driver/mysql"
)

type UserHandler struct {
	s services.Services
}

func New(service services.Services) UserHandler {
	return UserHandler{s :service}
}


func (u UserHandler) GetUserById(writer http.ResponseWriter,req *http.Request){
		qp := req.URL.Query()
		ids := qp.Get("id")
		id ,err := strconv.Atoi(ids)
		if err!=nil {
			_,_ = writer.Write([]byte("invalid Id "))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		res, err := u.s.GetUserById(id)
		if err!=nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		} 

		b,_ := json.Marshal(res)
		writer.WriteHeader(http.StatusOK)
		writer.Write(b)

}

func (u UserHandler) DeleteUser(writer http.ResponseWriter,req *http.Request) {
	qp := req.URL.Query()
	ids := qp.Get("id")
	id,err := strconv.Atoi(ids)
	if err!=nil {
		_,_ = writer.Write([]byte("invalid Id "))
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	er := u.s.DeletebyId(id)
	if er != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Error"))
		return
	} 
		
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write([]byte("user deleted"))
	if err != nil {
		fmt.Println(err)
		return
	}
	
}

func (u UserHandler) UpdateUser(writer http.ResponseWriter,req *http.Request) {
	var user *models.User
	resBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		fmt.Println(err)
		_, _ = writer.Write([]byte("invalid body"))
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if user.Id == 0 {
		_, _ = writer.Write([]byte("Id cannot be zero"))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	ok, err := u.s.ValidateEmail(user.Email)
	if err != nil {
		fmt.Println(err)
		_, _ = writer.Write([]byte("Error"))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	} else if !ok {
		_, _ = writer.Write([]byte("email already exist"))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
		
	fmt.Println(user)
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("User updated"))
	
}

func (u UserHandler) CreateUser(writer http.ResponseWriter,req *http.Request) {
	var user *models.User
	resBody,err := ioutil.ReadAll(req.Body)
	if err!=nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		fmt.Println(err)
		_, _ = writer.Write([]byte("invalid body"))
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_,er := u.s.ValidateId(user.Id)
	if er != nil {
		fmt.Println(err)
		_, _ = writer.Write([]byte("Error"))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	ok, err := u.s.ValidateEmail(user.Email)
	if err != nil {
		fmt.Println(err)
		_, _ = writer.Write([]byte("Error"))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	} else if !ok {
		_, _ = writer.Write([]byte("email already exist"))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	} 
	
	fmt.Println(user)
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("User created"))
	
}
