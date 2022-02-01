package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tejas/user-crud/models"
	"github.com/tejas/user-crud/services"
)

type Handler struct {
	handler services.User
}

func New(s services.User) Handler {
	return Handler{handler: s}
}

func (h *Handler) FindUserById(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")


	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		newError := models.ErrorResponse{Code: http.StatusInternalServerError, Message: "Invalid User Id"}
		jsonData, _ := json.Marshal(newError)

		_, _ = rw.Write(jsonData)

		return
	}

	user, err := h.handler.GetUserById(id)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{Code: http.StatusBadRequest, Message: err.Error()}
		jsonData, _ := json.Marshal(newError)
		_, _ = rw.Write(jsonData)
		return
	}

	jsonData, _ := json.Marshal(user)
	_, _ = rw.Write(jsonData)

	rw.WriteHeader(http.StatusOK)
}

func (h *Handler) FindUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	users, err := h.handler.GetUsers()

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		newErr := models.ErrorResponse{Code: http.StatusNotFound, Message: err.Error()}
		_, _ = json.Marshal(newErr)

		return
	}

	jsonData, _ := json.Marshal(users)

	rw.Write(jsonData)
}

func (h *Handler) UpdateById(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	body := r.Body

	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		newErr := models.ErrorResponse{Code: http.StatusInternalServerError, Message: "invalid id"}

		jsonData, _ := json.Marshal(newErr)

		_, _ = rw.Write(jsonData)

		return

	}

	var userData struct {
		Name  string
		Email string
		Phone string
		Age   int
	}

	err = json.NewDecoder(body).Decode(&userData)

	if err != nil {

		rw.WriteHeader(http.StatusInternalServerError)
		newErr := models.ErrorResponse{Code: http.StatusInternalServerError, Message: "invalid id"}

		jsonData, _ := json.Marshal(newErr)

		_, _ = rw.Write(jsonData)
		return
	}

	input := models.User{
		Id:    id,
		Name:  userData.Name,
		Email: userData.Email,
		Phone: userData.Phone,
		Age:   userData.Age,
	}

	err = h.handler.UpdateUser(input)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		newErr := models.ErrorResponse{Code: http.StatusBadRequest, Message: "invalid id"}

		jsonData, _ := json.Marshal(newErr)

		_, _ = rw.Write(jsonData)

		return
	}

	rw.WriteHeader(http.StatusOK)

	res := models.Response{
		Data:       models.User{},
		Message:    "user data updated",
		StatusCode: http.StatusOK,
	}

	jsonData, _ := json.Marshal(res)

	rw.Write(jsonData)

}

func (h *Handler) CreateUser(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("content-type", "application/json")
	body := r.Body
	var user models.User
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		respErr := models.ErrorResponse{Code: http.StatusInternalServerError, Message: "cannot insert user data"}
		jsonData, _ := json.Marshal(respErr)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(jsonData)
		return
	}

	err = h.handler.CreateUser(user)
	if err != nil {
		respErr := models.ErrorResponse{Code: http.StatusBadRequest, Message: "cannot insert user data"}
		jsonData, _ := json.Marshal(respErr)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(jsonData)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	response := models.Response{
		Data:       user,
		Message:    "User created",
		StatusCode: http.StatusOK,
	}
	jsonData, _ := json.Marshal(response)
	rw.Write(jsonData)
}

func (h *Handler) DeleteUser(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		resErr := models.ErrorResponse{Code: http.StatusInternalServerError, Message: "invalid id"}
		rw.WriteHeader(http.StatusInternalServerError)
		jsonData, _ := json.Marshal(resErr)
		_, _ = rw.Write(jsonData)
		return
	}

	err = h.handler.DeleteUser(id)
	if err != nil {
		respErr := models.ErrorResponse{Code: http.StatusBadRequest, Message: "error while deleting user data"}
		res, _ := json.Marshal(respErr)
		rw.WriteHeader(http.StatusBadRequest)
		_, _ = rw.Write(res)
		return
	}

	rw.WriteHeader(http.StatusOK)
	response := models.Response{
		Data:       id,
		Message:    "user deleted",
		StatusCode: http.StatusOK,
	}
	jsonData, _ := json.Marshal(response)
	rw.Write(jsonData)
}
