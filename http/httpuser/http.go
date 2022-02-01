package httpuser

import (
	"encoding/json"
	"go_lang/Assignment/user-curd/models"
	"go_lang/Assignment/user-curd/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HttpService struct {
	HttpServ services.Services
}

func (hs HttpService) GetOneUserHandler(response http.ResponseWriter, request *http.Request) {

	userId := mux.Vars(request)["id"]
	id, err := strconv.Atoi(userId)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		errorResponse := models.ErrorResponse{Code: http.StatusBadRequest, Message: "INVALID ID"}
		errorResp, _ := json.Marshal(errorResponse)
		response.Write([]byte(errorResp))
		return
	}

	user, err := hs.HttpServ.GetUser(id)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		errorResponse := models.ErrorResponse{Code: http.StatusBadRequest, Message: "USER NOT AVAILABLE"}
		errorResp, _ := json.Marshal(errorResponse)
		response.Write([]byte(errorResp))
	}
	response.Header().Add("content-type", "application/json")
	responseObj := models.Response{Data: user, Message: "DATA FETCHED", StatusCode: http.StatusOK}
	responseObjJson, _ := json.Marshal(responseObj)
	response.Write([]byte(responseObjJson))
}

func (hs HttpService) GetAllUserHandler(response http.ResponseWriter, request *http.Request) {

	userList, err := hs.HttpServ.GetAllUser()

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		errorResponse := models.ErrorResponse{Code: http.StatusBadRequest, Message: "ERROR IN FETCHING DATA"}
		errorResp, _ := json.Marshal(errorResponse)
		response.Write([]byte(errorResp))
		return
	} else {
		response.Header().Add("content-type", "application/json")
		responseObj := models.Response{Data: userList, Message: "DATA FETCHED", StatusCode: http.StatusOK}
		responseObjJson, _ := json.Marshal(responseObj)
		response.Write([]byte(responseObjJson))
	}

}

func (hs HttpService) AddUserHandler(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var usr models.User
	err := decoder.Decode(&usr)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		errorResponse := models.ErrorResponse{Code: http.StatusBadRequest, Message: "EMPTY USER"}
		errorResp, _ := json.Marshal(errorResponse)
		response.Write([]byte(errorResp))
	}

	err = hs.HttpServ.AddUser(usr)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		errorResponse := models.ErrorResponse{Code: http.StatusBadRequest, Message: "ERROR IN ADDING USER"}
		errorResp, _ := json.Marshal(errorResponse)
		response.Write([]byte(errorResp))
	} else {
		response.Header().Add("content-type", "application/json")
		responseObj := models.Response{Data: "", Message: "USER ADDED SUCCESSFULLY", StatusCode: http.StatusOK}
		responseObjJson, _ := json.Marshal(responseObj)
		response.Write([]byte(responseObjJson))
	}
}

func (hs HttpService) UpdateUserHandler(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var usr models.User
	err := decoder.Decode(&usr)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		errorResponse := models.ErrorResponse{Code: http.StatusBadRequest, Message: "EMPTY USER"}
		errorResp, _ := json.Marshal(errorResponse)
		response.Write([]byte(errorResp))
	}

	err = hs.HttpServ.UpdateUser(usr)

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		errorResponse := models.ErrorResponse{Code: http.StatusBadRequest, Message: "ERROR IN UPDATING USER"}
		errorResp, _ := json.Marshal(errorResponse)
		response.Write([]byte(errorResp))
	} else {
		response.Header().Add("content-type", "application/json")
		response.WriteHeader(http.StatusOK)
		errorResponse := models.ErrorResponse{Code: http.StatusBadRequest, Message: "USER UPDATED SUCCESSFULLY"}
		errorResp, _ := json.Marshal(errorResponse)
		response.Write([]byte(errorResp))
	}
}

func (hs HttpService) DeleteUserHandler(response http.ResponseWriter, request *http.Request) {
	userId := request.URL.Query().Get("id")
	id, _ := strconv.Atoi(userId)

	err := hs.HttpServ.DeleteUser(id)

	if err != nil || id == 0 {
		response.WriteHeader(http.StatusBadRequest)
		errorResponse := models.ErrorResponse{Code: http.StatusBadRequest, Message: "ERROR IN DELETING"}
		errorResp, _ := json.Marshal(errorResponse)
		response.Write([]byte(errorResp))
	} else {
		response.Header().Add("content-type", "application/json")
		response.WriteHeader(http.StatusOK)
		errorResponse := models.ErrorResponse{Code: http.StatusOK, Message: "USER DELETED SUCCESSFULLY"}
		errorResp, _ := json.Marshal(errorResponse)
		response.Write([]byte(errorResp))
	}
}
