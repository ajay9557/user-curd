package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"
	"strconv"

	"zopsmart/Task/models"
	service "zopsmart/Task/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	s service.Services
}

func New(service service.Services) UserHandler {
	return UserHandler{s: service}
}

func (u UserHandler) AllUserDetails(writer http.ResponseWriter,req *http.Request) {

	writer.Header().Set("content-type","application/json")

	data, err := u.s.GetAllUsersService()
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	// convert user data to json object and write on the response
	resp, _ := json.Marshal(data)
	_, _ = writer.Write(resp)
}

func (u UserHandler) GetUserById(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("content-type", "application/json")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		resErr, _ := writer.Write([]byte("invalid Id"))
		writer.WriteHeader(http.StatusInternalServerError)
		b,_ := json.Marshal(resErr)
		writer.Write(b)
		return
	}

	res, err := u.s.GetUserById(id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		b,_ := json.Marshal(res)
		writer.Write(b)
		return
	}

	b, _ := json.Marshal(res)
	writer.WriteHeader(http.StatusOK)
	writer.Write(b)

}

func (u UserHandler) DeleteUser(writer http.ResponseWriter, req *http.Request) {
	
	writer.Header().Set("content-type", "application/json")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		resErr, _ := writer.Write([]byte("invalid Id"))
		writer.WriteHeader(http.StatusInternalServerError)
		b,_ := json.Marshal(resErr)
		writer.Write(b)
		return
	}
	er := u.s.DeletebyId(id)
	if er != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
		
	writer.WriteHeader(http.StatusOK)

}

func (u UserHandler) UpdateUser(writer http.ResponseWriter, req *http.Request) {
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

	writer.WriteHeader(http.StatusOK)
	

}

func (u UserHandler) CreateUser(writer http.ResponseWriter, req *http.Request) {
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
	_, er := u.s.ValidateId(user.Id)
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
	writer.WriteHeader(http.StatusOK)
	
}
