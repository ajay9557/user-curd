package Users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	service "user-curd/Service"
	"user-curd/model"

	"github.com/gorilla/mux"
)

type Handler struct {
	Sev service.Service
}

func (serv Handler) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	v := params["id"]
	id, _ := strconv.Atoi(v)
	fmt.Println(id)
	userdata, err := serv.Sev.SearchByUserId(id)
	fmt.Println(userdata)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	res, err := json.Marshal(userdata)
	if err == nil {
		w.Write(res)
	}

}

func (serv Handler) Create(w http.ResponseWriter, r *http.Request) {
	var users model.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &users)
	if err != nil {
		_, _ = w.Write([]byte("invalid body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	addr := serv.Sev.IsEmailValid(users.Email)
	if !addr {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Enter Valid Email"))

	}
	usr, err := serv.Sev.InsertUserDetails(users)
	res, _ := json.Marshal(usr)
	if err != nil {
		_, _ = w.Write([]byte("could not create User"))
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		w.Write(res)
	}
}

func (serv Handler) DeleteId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	v := params["id"]
	id, err := strconv.Atoi(v)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	err = serv.Sev.DeleteByUserId(id)
	fmt.Println(err)
	w.Write([]byte("Deleted User Successfully!!"))

}
func (serv Handler) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var user model.User
	resBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(resBody, &user)
	if err != nil {
		fmt.Println(err)
		_, _ = rw.Write([]byte("invalid body"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = serv.Sev.UpdateByUserId(user)
	if err != nil {
		fmt.Println(err)
		_, _ = rw.Write([]byte("Database error"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(user)
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User updated"))
}

func (serv Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	usr, err := serv.Sev.SearchAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Retrieving Failed."))
	}
	res, err := json.Marshal(usr)
	if err != nil {
		_, _ = w.Write([]byte("could not get User"))
		w.WriteHeader(http.StatusInternalServerError)
	} else {

		w.Write(res)
	}
}
