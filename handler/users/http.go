package users

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"user-crud/models"
	"user-crud/services"
)

type UserHandler struct {
	Usev services.User
}

func (uh *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	v := params["id"]

	id, err := strconv.Atoi(v)

	if err != nil {
		w.Write([]byte("invalid id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	usr, err := uh.Usev.GetById(id)

	if err != nil {
		w.Write([]byte("id not found"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(usr)
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	response.Data = nil
	users, err := uh.Usev.GetAll()
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Message = "no data found"
		b, _ := json.Marshal(response)
		w.WriteHeader(response.StatusCode)
		w.Write(b)
		return
	}
	response.Data = users
	response.Message = "all users fetched"
	response.StatusCode = http.StatusOK
	b, _ := json.Marshal(response)
	w.Write(b)
	w.WriteHeader(response.StatusCode)
}

func (uh *UserHandler) Insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var usr models.User
	var response models.Response
	err := json.Unmarshal(body, &usr)

	if err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "invalid body"
		response.Data = nil
		b, _ := json.Marshal(response)
		_, _ = w.Write(b)
		w.WriteHeader(response.StatusCode)
		return
	}
	user, err := uh.Usev.Insert(&usr)
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Message = err.Error()
		response.Data = nil
		b, _ := json.Marshal(response)
		_, _ = w.Write(b)
		w.WriteHeader(response.StatusCode)
		return
	}
	response.StatusCode = http.StatusOK
	response.Message = "user created"
	response.Data = user
	b, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (uh *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	body, _ := ioutil.ReadAll(r.Body)
	var usr models.User
	err := json.Unmarshal(body, &usr)

	if err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "invalid body"
		response.Data = nil
		b, _ := json.Marshal(response)
		_, _ = w.Write(b)
		w.WriteHeader(response.StatusCode)
		return
	}
	user, err := uh.Usev.Update(&usr)

	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Message = "could not update user"
		response.Data = nil
		b, _ := json.Marshal(response)
		_, _ = w.Write(b)
		w.WriteHeader(response.StatusCode)
		return
	}
	response.StatusCode = http.StatusOK
	response.Message = "user updated"
	response.Data = user
	b, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (uh *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	params := mux.Vars(r)

	v := params["id"]

	id, err := strconv.Atoi(v)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.StatusCode = http.StatusBadRequest
		response.Message = "invalid id"
		response.Data = nil
		b, _ := json.Marshal(response)
		_, _ = w.Write(b)
		return
	}
	err = uh.Usev.Delete(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.StatusCode = http.StatusInternalServerError
		response.Message = err.Error()
		response.Data = nil
		b, _ := json.Marshal(response)
		_, _ = w.Write(b)
		return
	}
	response.StatusCode = http.StatusOK
	response.Message = "user deleted"
	response.Data = nil
	b, _ := json.Marshal(response)
	_, _ = w.Write(b)
	w.WriteHeader(response.StatusCode)
}
