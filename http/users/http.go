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
	Svc services.Service
}

func (h Handler) Search(w http.ResponseWriter, r *http.Request) {
	var res models.Response
	var errorRes models.ErrorResponse
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	v := params["id"]
	id, _ := strconv.Atoi(v)
	userdata, err := h.Svc.GetByUserId(id)
	if err != nil {
		errorRes = models.ErrorResponse{
			Code:    500,
			Message: "Internal Server Error",
		}
		b, _ := json.Marshal(errorRes)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write((b))
		return
	}
	res = models.Response{
		Data:       userdata,
		Message:    "Retrieved Successfully",
		StatusCode: 200,
	}
	b, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var users models.User
	var res models.Response
	var errorRes models.ErrorResponse
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &users)
	if err != nil {
		_, _ = w.Write([]byte("Invalid body"))
		errorRes = models.ErrorResponse{
			Code:    400,
			Message: "Status Bad Request",
		}
		b, _ := json.Marshal(errorRes)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write((b))
		return
	}
	usr, err := h.Svc.InsertUserDetails(users)
	if err != nil {
		errorRes = models.ErrorResponse{
			Code:    500,
			Message: "Internal Server Error",
		}
		b, _ := json.Marshal(errorRes)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write((b))
		return
	}
	res = models.Response{
		Data:       usr,
		Message:    "Inserted Successfully",
		StatusCode: 200,
	}
	b, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

func (h Handler) DeleteId(w http.ResponseWriter, r *http.Request) {
	var res models.Response
	var errorRes models.ErrorResponse
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	v := params["id"]
	id, err := strconv.Atoi(v)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	err = h.Svc.DeleteByUserId(id)
	if err != nil {
		errorRes = models.ErrorResponse{
			Code:    500,
			Message: "Internal Server Error",
		}
		b, _ := json.Marshal(errorRes)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write((b))
		return
	}
	res = models.Response{
		Data:       nil,
		Message:    "Deleted successfully",
		StatusCode: 200,
	}
	b, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}
func (h Handler) UpdateUserDetails(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var res models.Response
	var errorRes models.ErrorResponse
	resBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(resBody, &user)
	if err != nil {
		_, _ = w.Write([]byte("invalid body"))
		errorRes = models.ErrorResponse{
			Code:    400,
			Message: "Status Bad Request",
		}
		b, _ := json.Marshal(errorRes)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write((b))
		return
	}
	if user.Id == 0 {
		_, _ = w.Write([]byte("Id shouldn't be zero"))

	}
	err = h.Svc.UpdateByUserId(user)
	if err != nil {
		errorRes = models.ErrorResponse{
			Code:    500,
			Message: "Internal Server Error",
		}
		b, _ := json.Marshal(errorRes)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write((b))
		return
	}
	res = models.Response{
		Data:       resBody,
		Message:    "Updated User Successfully",
		StatusCode: 200,
	}
	b, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
	w.Write([]byte("Updated User Successfully"))
}
func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	var res models.Response
	var errorRes models.ErrorResponse
	usr, err := h.Svc.GetAll()
	if err != nil {
		errorRes = models.ErrorResponse{
			Code:    500,
			Message: "Internal Server Error",
		}
		b, _ := json.Marshal(errorRes)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write((b))
		return
	}
	res = models.Response{
		Data:       usr,
		Message:    "Retrieved Successfully",
		StatusCode: 200,
	}
	b, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
