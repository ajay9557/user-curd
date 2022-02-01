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

func (u UserHandler) AllUserDetails(w http.ResponseWriter,r *http.Request) {

	w.Header().Set("content-type","application/json")

	data, err := u.s.GetAllUsersService()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (u UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		resErr, _ := w.Write([]byte("invalid Id"))
		w.WriteHeader(http.StatusBadRequest)
		b,_ := json.Marshal(resErr)
		w.Write(b)
		return
	}

	res, err := u.s.GetUserByIdService(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		b,_ := json.Marshal(res)
		w.Write(b)
		return
	} 

	b, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

func (u UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	
    w.Header().Set("content-type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		resErr, _ := w.Write([]byte("invalid Id"))
		w.WriteHeader(http.StatusBadRequest)
		b,_ := json.Marshal(resErr)
		w.Write(b)
		return
	}
	er := u.s.DeletebyIdService(id)
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error"))
		return
	}
	w.Write([]byte("Successfully deleted"))
	w.WriteHeader(http.StatusOK)
}

func (u UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		fmt.Println(err)
		resErr, _:=w.Write([]byte("invalid body"))
		b,_ := json.Marshal(resErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(b)
		return
	}
	if user.Id == 0 {
		_, _ = w.Write([]byte("Id cannot be zero"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	er := u.s.UpdatebyIdService(user)

	if er != nil {
		fmt.Println(er)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error"))
		return
	}

	w.Write([]byte("successfully updated"))
	w.WriteHeader(http.StatusOK)
	
}

func (u UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("invalid body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	e := u.s.CreateUserService(user)

	if e != nil {
		fmt.Println(e)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully created"))
	
}
