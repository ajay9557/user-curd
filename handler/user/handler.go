package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"zopsmart/user-curd/model"
	"zopsmart/user-curd/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	Usr service.User
}

func (us Handler) UserWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	v := params["id"]
	id, err := strconv.Atoi(v)

	userdata, err := us.Usr.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Enter Valid Id"))
		return
	}

	res, err := json.Marshal(userdata)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

}

func (us Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userdata, _ := us.Usr.GetUsers()

	res, err := json.MarshalIndent(userdata, "", "")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
	}

}

func (us Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ur model.User
	body := r.Body
	json.NewDecoder(body).Decode(&ur)

	isValid, _ := us.Usr.CheckMail(ur.Email)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Enter Valid email address"))

		return
	}

	u, _ := us.Usr.PostUser(ur.Name, ur.Email, ur.Phone, ur.Age)
	res, err := json.Marshal(u)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}

}

func (us Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ur model.User
	body := r.Body
	json.NewDecoder(body).Decode(&ur)
	params := mux.Vars(r)
	v := params["id"]
	id, _ := strconv.Atoi(v)

	userdata, err := us.Usr.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userdt, _ := us.Usr.Update(userdata, ur)
	users, _ := us.Usr.GetUsers()
	users = append(users, userdt)
	res, _ := json.Marshal(userdt)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func (us Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	v := params["id"]
	id, err := strconv.Atoi(v)

	err = us.Usr.DeleteByID(id)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User Deletion Failed"))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("User Deleted"))

}
