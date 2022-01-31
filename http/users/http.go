package users

import (
	"encoding/json"
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
	var res models.Response
	var errRes models.ErrorResponse

	resBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(resBody, &user)
	if err != nil {
		errRes.Code = 400
		errRes.Message = "invalid body"
		b, _ := json.Marshal(errRes)
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write(b)
		return
	}
	err = u.serv.InsertUserDetails(user)
	if err != nil {
		errRes.Code = 500
		errRes.Message = "internal error"
		b, _ := json.Marshal(errRes)
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write([]byte(b))
		return
	}
	res.Data = user
	res.Message = "user created"
	res.StatusCode = 200
	rw.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(res)
	rw.Write([]byte(b))
}

func (u UserHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	var res models.Response
	var errRes models.ErrorResponse
	allUsers, err := u.serv.FetchAllUserDetails()
	if err != nil {
		errRes.Code = 400
		errRes.Message = "internal error"
		b, _ := json.Marshal(errRes)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(b))
	} else {
		res.Data = allUsers
		res.Message = "all users obtained successfully"
		res.StatusCode = 200
		b, _ := json.Marshal(res)
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write(b)
	}
}

func (u UserHandler) GetUserById(rw http.ResponseWriter, r *http.Request) {
	var res models.Response
	var errRes models.ErrorResponse
	q := r.URL.Query()
	id := q.Get("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		errRes.Code = 400
		errRes.Message = "invalid parameter id"
		b, _ := json.Marshal(errRes)
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write(b)
		return
	}
	user, err := u.serv.FetchUserDetailsById(userId)
	if err != nil {
		errRes.Code = 500
		errRes.Message = "internal error"
		b, _ := json.Marshal(errRes)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(b)
		return
	} else {
		res.Data = user
		res.Message = "user obtained successfully"
		res.StatusCode = 200
		b, _ := json.Marshal(res)
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write(b)
	}

}

func (u UserHandler) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var user models.User
	var res models.Response
	var errRes models.ErrorResponse
	rw.Header().Set("Content-Type", "application/json")
	resBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(resBody, &user)
	if err != nil {
		errRes.Code = 400
		errRes.Message = "invalid body"
		b, _ := json.Marshal(errRes)
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write(b)
		return
	}
	err = u.serv.UpdateUserDetails(user)
	if err != nil {
		errRes.Code = 500
		errRes.Message = "internal error"
		b, _ := json.Marshal(errRes)
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write(b)
		return
	}
	res.Data = user
	res.Message = "user updated successfully"
	res.StatusCode = 200
	b, _ := json.Marshal(res)
	rw.WriteHeader(http.StatusOK)
	rw.Write(b)
}

func (u UserHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	var res models.Response
	var errRes models.ErrorResponse
	q := r.URL.Query()
	id := q.Get("id")
	userId, _ := strconv.Atoi(id)
	err := u.serv.DeleteUserDetailsById(userId)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		errRes.Code = 500
		errRes.Message = "internal error"
		b, _ := json.Marshal(errRes)
		rw.Write(b)
		return
	} else {
		res.Message = "deleted successfully"
		res.StatusCode = 200
		b, _ := json.Marshal(res)
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write(b)
	}

}
