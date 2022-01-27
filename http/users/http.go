package users

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"user-curd/entities"
	"user-curd/service"
)

type userApi struct {
	userService service.UserServiceHandler
}

func New(userService service.UserServiceHandler) *userApi {
	return &userApi{userService: userService}
}

func (usrApi *userApi) GetUserByIdHandler(wr http.ResponseWriter, req *http.Request) {

	wr.Header().Set("content-type", "application/json")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid id"}
		wr.WriteHeader(http.StatusInternalServerError)
		res, er := json.Marshal(resErr)
		if er == nil {
			wr.Write(res)
		}
		return
	}

	usrData, err := usrApi.userService.GetUserByIdService(id)
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: err.Error()}
		wr.WriteHeader(http.StatusBadRequest)
		res, er := json.Marshal(resErr)
		if er == nil {
			wr.Write(res)
		}
		return
	}

	resp, err := json.Marshal(usrData)
	if err == nil {
		wr.Write(resp)
	}
}

func (usrApi *userApi) GetAllUserHandler(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("content-type", "application/json")
	data, err := usrApi.userService.GetAllUsersService()
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusNotFound, ErrMsg: err.Error()}
		wr.WriteHeader(http.StatusNotFound)
		res, er := json.Marshal(resErr)
		if er == nil {
			wr.Write(res)
		}
		return
	}

	resp, err := json.Marshal(data)
	if err == nil {
		wr.Write(resp)
	}
}

func (usrApi *userApi) CreateUserHandler(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("content-type", "application/json")
	body := req.Body
	var userData entities.User
	err := json.NewDecoder(body).Decode(&userData)
	//if err != nil {
	//	respErr := entities.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid input"}
	//	res, er := json.Marshal(respErr)
	//	wr.WriteHeader(http.StatusInternalServerError)
	//	if er == nil {
	//		wr.Write(res)
	//	}
	//	return
	//}

	err = usrApi.userService.CreateUserService(userData)
	if err != nil {
		respErr := entities.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: "bad request"}
		res, er := json.Marshal(respErr)
		wr.WriteHeader(http.StatusBadRequest)
		if er == nil {
			wr.Write(res)
		}
		return
	}
	wr.WriteHeader(http.StatusCreated)
}

func (usrApi *userApi) UpdateUserHandler(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("content-type", "application/json")
	body := req.Body
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid id"}
		wr.WriteHeader(http.StatusInternalServerError)
		res, er := json.Marshal(resErr)
		if er == nil {
			wr.Write(res)
		}
		return
	}

	var userData struct {
		Name  string
		Email string
		Phone string
		Age   int
	}
	err = json.NewDecoder(body).Decode(&userData)
	inp := entities.User{
		Id:    id,
		Name:  userData.Name,
		Email: userData.Email,
		Phone: userData.Phone,
		Age:   userData.Age,
	}
	err = usrApi.userService.UpdateUserService(inp)
	if err != nil {
		respErr := entities.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: "bad request"}
		res, er := json.Marshal(respErr)
		wr.WriteHeader(http.StatusBadRequest)
		if er == nil {
			wr.Write(res)
		}
		return
	}
	wr.WriteHeader(http.StatusOK)
}

func (usrApi *userApi) DeleteUserHandler(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("content-type", "application/json")
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		resErr := entities.HttpErrs{ErrCode: http.StatusInternalServerError, ErrMsg: "invalid id"}
		wr.WriteHeader(http.StatusInternalServerError)
		res, er := json.Marshal(resErr)
		if er == nil {
			wr.Write(res)
		}
		return
	}

	err = usrApi.userService.DeleteUserService(id)
	if err != nil {
		respErr := entities.HttpErrs{ErrCode: http.StatusBadRequest, ErrMsg: "bad request"}
		res, er := json.Marshal(respErr)
		wr.WriteHeader(http.StatusBadRequest)
		if er == nil {
			wr.Write(res)
		}
		return
	}
	wr.WriteHeader(http.StatusOK)
}
