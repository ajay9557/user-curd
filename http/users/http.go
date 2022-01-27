package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"user-curd/models"
	"user-curd/services"

	"github.com/gorilla/mux"
)

type Handler struct {
	Sev services.Service
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
	var users models.User
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
func (serv Handler) UpdateUserDetails(rw http.ResponseWriter, r *http.Request) {
	var user models.User
	resBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(resBody, &user)
	if err != nil {
		fmt.Println(err)
		_, _ = rw.Write([]byte("invalid body"))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	if user.Id == 0 {
		_, _ = rw.Write([]byte("Id shouldn't be zero"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	ok := serv.Sev.IsEmailValid(user.Email)
	if !ok {
		_, _ = rw.Write([]byte("Email already there,create new email"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		err = serv.Sev.UpdateByUserId(user)
		fmt.Println(err)
		fmt.Println(user)
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("user updated"))
	}
}

// func (serv Handler) UpdateId(w http.ResponseWriter, r *http.Request) {
// 	var userdata models.User
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	v := params["id"]
// 	id, err := strconv.Atoi(v)
// 	if err != nil {
// 		fmt.Errorf("%v", err)
// 	}
// 	userdata, err = serv.Sev.SearchByUserId(id)
// 	//fmt.Print(userdata.Age)
// 	err = serv.Sev.UpdateByUserId(userdata.Age, id)
// 	if err != nil {
// 		fmt.Println(err)
// 		_, _ = w.Write([]byte("Updation Failed"))
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	w.Write([]byte("Updated User Successfully!!"))
// }

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
		//w.Write([]byte("User Successfully retrieved"))
		w.Write(res)
	}
}
