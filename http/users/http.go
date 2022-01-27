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
	userService service.UserServiceHandler
}

func New(userService service.UserServiceHandler) *userApi {

	return &userApi{userService: userService}
}

func (usrApi *userApi) GetUserByIdHandler(wr http.ResponseWriter, req *http.Request) {

	// set content-type to json
	wr.Header().Set("content-type", "application/json")

	// get id from url
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid id"}
		wr.WriteHeader(http.StatusInternalServerError)
		res, er := json.Marshal(resErr)
		if er == nil {
			_, _ = wr.Write(res)
		}
		return
	}

	// get data from service layer
	usrData, err := usrApi.userService.GetUserByIdService(id)
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: err.Error()}
		wr.WriteHeader(http.StatusBadRequest)
		res, er := json.Marshal(resErr)
		if er == nil {
			_, _ = wr.Write(res)
		}
		return
	}

	// convert user data to json object and write on the response
	resp, err := json.Marshal(usrData)
	if err == nil {
		_, _ = wr.Write(resp)
	}
}

func (usrApi *userApi) GetAllUserHandler(wr http.ResponseWriter, req *http.Request) {

	// set content-type to json
	wr.Header().Set("content-type", "application/json")

	// get data from service layer
	data, err := usrApi.userService.GetAllUsersService()
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusNotFound, ErrMsg: err.Error()}
		wr.WriteHeader(http.StatusNotFound)
		res, er := json.Marshal(resErr)
		if er == nil {
			_, _ = wr.Write(res)
		}
		return
	}

	// convert user data to json object and write on the response
	resp, err := json.Marshal(data)
	if err == nil {
		_, _ = wr.Write(resp)
	}
}

func (usrApi *userApi) CreateUserHandler(wr http.ResponseWriter, req *http.Request) {

	// set content-type to json
	wr.Header().Set("content-type", "application/json")
	body := req.Body
	var userData entities.User
	_ = json.NewDecoder(body).Decode(&userData)

	err := usrApi.userService.CreateUserService(userData)
	if err != nil {
		respErr := entities.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: "bad request"}
		res, er := json.Marshal(respErr)
		wr.WriteHeader(http.StatusBadRequest)
		if er == nil {
			_, _ = wr.Write(res)
		}
		return
	}
	wr.WriteHeader(http.StatusCreated)
}

func (usrApi *userApi) UpdateUserHandler(wr http.ResponseWriter, req *http.Request) {

	// set content-type to json
	wr.Header().Set("content-type", "application/json")
	body := req.Body
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid id"}
		wr.WriteHeader(http.StatusInternalServerError)
		res, er := json.Marshal(resErr)
		if er == nil {
			_, _ = wr.Write(res)
		}
		return
	}

	var userData struct {
		Name  string
		Email string
		Phone string
		Age   int
	}
	_ = json.NewDecoder(body).Decode(&userData)
	inp := entities.User{
		Id:    id,
		Name:  userData.Name,
		Email: userData.Email,
		Phone: userData.Phone,
		Age:   userData.Age,
	}

	// call to service layer
	err = usrApi.userService.UpdateUserService(inp)
	if err != nil {
		respErr := entities.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: "bad request"}
		res, er := json.Marshal(respErr)
		wr.WriteHeader(http.StatusBadRequest)
		if er == nil {
			_, _ = wr.Write(res)
		}
		return
	}

	// give status OK (200) if everything goes OK
	wr.WriteHeader(http.StatusOK)
}

func (usrApi *userApi) DeleteUserHandler(wr http.ResponseWriter, req *http.Request) {

	// set content-type to json
	wr.Header().Set("content-type", "application/json")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid id"}
		wr.WriteHeader(http.StatusInternalServerError)
		res, er := json.Marshal(resErr)
		if er == nil {
			_, _ = wr.Write(res)
		}
		return
	}

	// call to service layer
	err = usrApi.userService.DeleteUserService(id)
	if err != nil {
		respErr := entities.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: "bad request"}
		res, er := json.Marshal(respErr)
		wr.WriteHeader(http.StatusBadRequest)
		if er == nil {
			_, _ = wr.Write(res)
		}
		return
	}

	// give status OK (200) if everything goes OK
	wr.WriteHeader(http.StatusOK)
}
