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
	Svc service.User
}

func (us Handler) GetUserWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	v := params["id"]
	id, err := strconv.Atoi(v)

	userdata, err := us.Svc.GetByID(id)
	if err != nil {
		errobj := model.ErrorResponse{http.StatusBadRequest, "Id not found"}
		w.WriteHeader(http.StatusBadRequest)
		res, er := json.Marshal(errobj)
		if er == nil {
			w.Write([]byte(res))
		}
		return
	}
	if userdata.Id == 0 {
		errobj := model.ErrorResponse{http.StatusBadRequest, "Id not found"}
		w.WriteHeader(http.StatusBadRequest)
		res, er := json.Marshal(errobj)
		if er == nil {
			w.Write([]byte(res))
		}
		return
	}
	response := model.Response{Data: userdata, Message: "user successfully retrieved", StatusCode: http.StatusOK}
	res, err := json.Marshal(response)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

}

func (us Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userdata, err := us.Svc.GetUsers()
	if err != nil {
		errobj := model.ErrorResponse{http.StatusInternalServerError, "Failed to retrieve users"}
		w.WriteHeader(http.StatusInternalServerError)
		res, er := json.Marshal(errobj)
		if er == nil {
			w.Write([]byte(res))
		}
		return
	}
	response := model.Response{Data: userdata, Message: "users successfully retrieved", StatusCode: http.StatusOK}

	res, err := json.MarshalIndent(response, "", "")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

}

func (us Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ur model.User
	body := r.Body
	json.NewDecoder(body).Decode(&ur)

	u, err := us.Svc.PostUser(ur)
	if err != nil {
		errobj := model.ErrorResponse{http.StatusInternalServerError, "Failed to add user"}
		w.WriteHeader(http.StatusInternalServerError)
		res, er := json.Marshal(errobj)
		if er == nil {
			w.Write([]byte(res))
		}
		return

	} else {
		response := model.Response{Data: u, Message: "users successfully added", StatusCode: http.StatusOK}
		res, err := json.Marshal(response)
		if err == nil {
			w.WriteHeader(http.StatusCreated)
			w.Write(res)
		}
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
	updatedUser, err := us.Svc.Update(id, ur)
	if err != nil {
		errobj := model.ErrorResponse{http.StatusInternalServerError, "Failed to update user"}
		w.WriteHeader(http.StatusInternalServerError)
		res, er := json.Marshal(errobj)
		if er == nil {
			w.Write([]byte(res))
		}
		return
	}
	response := model.Response{Data: updatedUser, Message: "users successfully updated", StatusCode: http.StatusOK}
	res, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func (us Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	v := params["id"]
	id, err := strconv.Atoi(v)

	err = us.Svc.DeleteByID(id)
	if err != nil {
		errobj := model.ErrorResponse{http.StatusInternalServerError, "Failed to delete user"}
		w.WriteHeader(http.StatusInternalServerError)
		res, er := json.Marshal(errobj)
		if er == nil {
			w.Write([]byte(res))
		}
		return
	}
	response := model.Response{Data: "", Message: "users successfully added", StatusCode: http.StatusOK}
	res, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
