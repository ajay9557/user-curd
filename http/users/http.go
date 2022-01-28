package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"user-curd/models"
	"user-curd/services"
)

type UserHandler struct {
	serv services.Services
}

func New(service services.Services) UserHandler {
	return UserHandler{serv: service}
}

func (u UserHandler) PostUser(rw http.ResponseWriter, r *http.Request) {
	var user models.User
	rw.Header().Set("Content-Type", "application/json")
	resBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(resBody, &user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte("invalid body"))
		return
	}
	err = u.serv.InsertUserDetails(user)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write([]byte("internal error"))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User created"))
}

func (u UserHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	allUsers, err := u.serv.FetchAllUserDetails()
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("error generated"))
	} else {
		b, _ := json.Marshal(allUsers)
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write(b)
	}
}

func (u UserHandler) GetUserById(rw http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	id := q.Get("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte("invalid parameter id"))
		return
	}
	user, err := u.serv.FetchUserDetailsById(int_id)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte("internal error"))
		return
	} else {
		b, _ := json.Marshal(user)
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write(b)
	}

}

func (u UserHandler) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var user models.User
	rw.Header().Set("Content-Type", "application/json")
	resBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(resBody, &user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write([]byte("invalid body"))
		return
	}
	err = u.serv.UpdateUserDetails(user)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write([]byte("internal error"))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User updated"))
}

func (u UserHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	id := q.Get("id")
	int_id, _ := strconv.Atoi(id)
	if int_id == 0 {
		_, _ = rw.Write([]byte("Id shouldn't be zero"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err := u.serv.DeleteUserDetailsById(int_id)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("error generated"))
		return
	} else {
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte("User deleted successfully"))
	}

}
