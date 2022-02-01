package users

import (
	"encoding/json"
	"layer/user/models"
	"layer/user/services"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	hndlr services.User
}

func New(s services.User) Handler {
	return Handler{s}
}

func (srv Handler) GetUserByIdHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)

	userId := params["id"]

	id, err := strconv.Atoi(userId)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Invalid User id"}
		err, _ := json.Marshal(newError)
		_, _ = res.Write(err)
		return
	}

	if id <= 0 {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Id should be greater than 0"}
		err, _ := json.Marshal(newError)
		_, _ = res.Write(err)
		return
	}

	user, err := srv.hndlr.GetUserById(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: err.Error()}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)
		return
	}

	responseData := models.Response{
		Data: struct {
			User models.User `json:"user"`
		}{
			User: user,
		},
		Message:    "Successful operation",
		StatusCode: 200,
	}

	jsonData, _ := json.Marshal(responseData)
	_, _ = res.Write([]byte(jsonData))

}

func (srv Handler) GetUsersHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	users, err := srv.hndlr.GetUsers()
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: err.Error()}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	responseData := models.Response{
		Data: struct {
			Users []*models.User `json:"users"`
		}{
			Users: users,
		},
		StatusCode: 200,
		Message:    "Successful operation",
	}

	jsonData, _ := json.Marshal(responseData)
	_, _ = res.Write([]byte(string(jsonData)))

}

func (srv Handler) UpdateUserHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil || reflect.DeepEqual(user, models.User{}) {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Cannot parse request data"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	params := mux.Vars(req)
	id := params["id"]

	convId, err := strconv.Atoi(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Invalid id"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	if convId <= 0 {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Id should be greater than 0"}
		err, _ := json.Marshal(newError)
		_, _ = res.Write(err)
		return
	}

	updatedUser, err := srv.hndlr.UpdateUser(convId, user)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: err.Error()}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	responseData := models.Response{
		Data: struct {
			User *models.User `json:"user"`
		}{
			updatedUser,
		},
		StatusCode: 201,
		Message:    "Successfull operation",
	}

	jsonData, _ := json.Marshal(responseData)

	_, _ = res.Write(jsonData)
}

func (srv Handler) DeleteUserHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)
	id := params["id"]

	convId, err := strconv.Atoi(id)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Invalid id"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	if convId <= 0 {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Id should be greater than 0"}
		err, _ := json.Marshal(newError)
		_, _ = res.Write(err)
		return
	}

	err = srv.hndlr.DeleteUser(convId)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: err.Error()}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	responseData := models.Response{
		Message:    "Successfull operation",
		StatusCode: 200,
	}

	jsonData, _ := json.Marshal(responseData)

	_, _ = res.Write([]byte(jsonData))

}

func (srv Handler) CreateUserHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil || reflect.DeepEqual(user, models.User{}) {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "Bad Request. Cannot parse request data"}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	createdUser, err := srv.hndlr.CreateUser(user)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		newError := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: err.Error()}
		jsonData, _ := json.Marshal(newError)
		_, _ = res.Write(jsonData)

		return
	}

	responseData := models.Response{
		Data: struct {
			User *models.User `json:"user"`
		}{
			createdUser,
		},
		Message:    "Successfull operation",
		StatusCode: 201,
	}

	jsonData, _ := json.Marshal(responseData)

	_, _ = res.Write([]byte(string(jsonData)))

}
