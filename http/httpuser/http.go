package httpuser

import (
	"encoding/json"
	"go_lang/Assignment/user-curd/models"
	"go_lang/Assignment/user-curd/services"
	"net/http"
	"strconv"
)

type HttpService struct {
	HttpServ services.Services
}

func (hs HttpService) Handler(response http.ResponseWriter, request *http.Request) {

	response.Header().Add("content-type", "text/html")

	switch request.Method {
	case http.MethodGet:

		userId := request.URL.Query().Get("id")

		if userId == "" {
			userList, err := hs.HttpServ.GetAllUser()

			if err != nil {
				response.WriteHeader(http.StatusBadRequest)
				errorResponse := models.ErrorResponse{Code: http.StatusBadRequest, Message: "ERROR IN FETCHING DATA"}
				errorResp, _ := json.Marshal(errorResponse)
				response.Write([]byte(errorResp))
				return
			} else {
				responseObj := models.Response{Data: userList, Message: "DATA FETCHED", StatusCode: http.StatusOK}
				responseObjJson, _ := json.Marshal(responseObj)
				response.Write([]byte(responseObjJson))
			}

		} else {

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

			responseObj := models.Response{Data: user, Message: "DATA FETCHED", StatusCode: http.StatusOK}
			responseObjJson, _ := json.Marshal(responseObj)
			response.Write([]byte(responseObjJson))
		}
	case http.MethodPost:
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
			responseObj := models.Response{Data: "", Message: "USER ADDED SUCCESSFULLY", StatusCode: http.StatusOK}
			responseObjJson, _ := json.Marshal(responseObj)
			response.Write([]byte(responseObjJson))
		}
	case http.MethodPut:
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
			response.WriteHeader(http.StatusOK)
			errorResponse := models.ErrorResponse{Code: http.StatusBadRequest, Message: "USER UPDATED SUCCESSFULLY"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		}
	case http.MethodDelete:

		userId := request.URL.Query().Get("id")
		id, _ := strconv.Atoi(userId)

		err := hs.HttpServ.DeleteUser(id)

		if err != nil || id == 0 {
			response.WriteHeader(http.StatusBadRequest)
			errorResponse := models.ErrorResponse{Code: http.StatusBadRequest, Message: "ERROR IN DELETING"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		} else {
			response.WriteHeader(http.StatusOK)
			errorResponse := models.ErrorResponse{Code: http.StatusOK, Message: "USER DELETED SUCCESSFULLY"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		}
	}
}
