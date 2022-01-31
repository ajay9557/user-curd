package users

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/usercurd/models"
	"github.com/usercurd/services"
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
	users, err := uh.Usev.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("no data found"))
		return
	}
	b, _ := json.Marshal(users)
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func (uh *UserHandler) Insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var usr models.User
	err := json.Unmarshal(body, &usr)

	if err != nil {
		fmt.Println(err)
		_, _ = w.Write([]byte("invalid body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	usr, err = uh.Usev.Insert(usr)

	if err != nil {
		_, _ = w.Write([]byte("could not create user"))
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	b, _ := json.Marshal(usr)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	w.Write([]byte("user created"))
}

func (uh *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var usr models.User
	err := json.Unmarshal(body, &usr)

	if err != nil {
		fmt.Println(err)
		_, _ = w.Write([]byte("invalid body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = uh.Usev.Update(usr.Id, usr.Name)

	if err != nil {
		_, _ = w.Write([]byte("could not update user"))
		w.WriteHeader(http.StatusInternalServerError)

		return

	}
	b, _ := json.Marshal(usr)
	w.Write(b)
}

func (uh *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	v := params["id"]

	id, err := strconv.Atoi(v)

	if err != nil {
		_, _ = w.Write([]byte("invalid user id"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	usr, err := uh.Usev.Delete(id)

	if err != nil {
		_, _ = w.Write([]byte("could not Delete user"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, _ := json.Marshal(usr)
	w.Write(b)
}
