package users

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-curd/entities"
	"user-curd/service"

	"github.com/gorilla/mux"
)

type userApi struct {
	userService service.UserService
}

func New(userService service.UserService) *userApi {
	return &userApi{userService: userService}
}

func (ua *userApi) GetUserByIdHandler(wr http.ResponseWriter, req *http.Request) {

	// set content-type to json
	wr.Header().Set("content-type", "application/json")

	// get id from url
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid id"}
		wr.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(resErr)
		_, _ = wr.Write(res)
		return
	}

	// get data from service layer
	usrData, err := ua.userService.GetUserByIdService(id)
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: err.Error()}
		wr.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(resErr)
		_, _ = wr.Write(res)
		return
	}

	// convert user data to json object and write on the response
	resp, _ := json.Marshal(usrData)
	_, _ = wr.Write(resp)
	wr.WriteHeader(http.StatusOK)
}

func (ua *userApi) GetAllUserHandler(wr http.ResponseWriter, req *http.Request) {

	// set content-type to json
	wr.Header().Set("content-type", "application/json")

	// get data from service layer
	data, err := ua.userService.GetAllUsersService()
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusNotFound, ErrMsg: err.Error()}
		wr.WriteHeader(http.StatusNotFound)
		res, _ := json.Marshal(resErr)
		_, _ = wr.Write(res)
		return
	}

	// convert user data to json object and write on the response
	resp, _ := json.Marshal(data)
	_, _ = wr.Write(resp)
}

func (ua *userApi) CreateUserHandler(wr http.ResponseWriter, req *http.Request) {

	// set content-type to json
	wr.Header().Set("content-type", "application/json")
	body := req.Body
	var userData entities.User
	err := json.NewDecoder(body).Decode(&userData)
	if err != nil {
		respErr := entities.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "bad request"}
		res, _ := json.Marshal(respErr)
		wr.WriteHeader(http.StatusInternalServerError)
		wr.Write(res)
		return
	}

	err = ua.userService.CreateUserService(userData)
	if err != nil {
		respErr := entities.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: "bad request"}
		res, _ := json.Marshal(respErr)
		wr.WriteHeader(http.StatusBadRequest)
		wr.Write(res)
		return
	}

	wr.WriteHeader(http.StatusCreated)
	response := entities.HttpResponse{
		Data:       userData,
		Message:    "User created",
		StatusCode: http.StatusOK,
	}
	resp, _ := json.Marshal(response)
	wr.Write(resp)
}

func (ua *userApi) UpdateUserHandler(wr http.ResponseWriter, req *http.Request) {

	// set content-type to json
	wr.Header().Set("content-type", "application/json")
	body := req.Body
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid id"}
		wr.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(resErr)
		_, _ = wr.Write(res)
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
		resErr := entities.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid id"}
		wr.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(resErr)
		_, _ = wr.Write(res)
		return
	}

	inp := entities.User{
		Id:    id,
		Name:  userData.Name,
		Email: userData.Email,
		Phone: userData.Phone,
		Age:   userData.Age,
	}
	// call to service layer
	err = ua.userService.UpdateUserService(inp)
	if err != nil {
		respErr := entities.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: "bad request"}
		res, _ := json.Marshal(respErr)
		wr.WriteHeader(http.StatusBadRequest)
		_, _ = wr.Write(res)
		return
	}

	// give status OK (200) if everything goes OK
	wr.WriteHeader(http.StatusOK)
	response := entities.HttpResponse{
		Data:       inp,
		Message:    "Data updated",
		StatusCode: http.StatusOK,
	}
	resp, _ := json.Marshal(response)
	wr.Write(resp)
}

func (ua *userApi) DeleteUserHandler(wr http.ResponseWriter, req *http.Request) {

	// set content-type to json
	wr.Header().Set("content-type", "application/json")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid id"}
		wr.WriteHeader(http.StatusInternalServerError)
		res, _ := json.Marshal(resErr)
		_, _ = wr.Write(res)
		return
	}

	// call to service layer
	err = ua.userService.DeleteUserService(id)
	if err != nil {
		respErr := entities.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: "bad request"}
		res, _ := json.Marshal(respErr)
		wr.WriteHeader(http.StatusBadRequest)
		_, _ = wr.Write(res)
		return
	}

	// give status OK (200) if everything goes OK
	wr.WriteHeader(http.StatusOK)
	response := entities.HttpResponse{
		Data:       id,
		Message:    "User deleted",
		StatusCode: http.StatusOK,
	}
	resp, _ := json.Marshal(response)
	wr.Write(resp)
}
