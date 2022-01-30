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
				errorResponse := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "ERROR IN FETCHING DATA"}
				errorResp, _ := json.Marshal(errorResponse)
				response.Write([]byte(errorResp))
				return
			} else {
				userListFormat, _ := json.Marshal(userList)
				response.Write([]byte(userListFormat))
			}

		} else {

			id, err := strconv.Atoi(userId)

			if err != nil {
				response.WriteHeader(http.StatusBadRequest)
				errorResponse := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "INVALID ID"}
				errorResp, _ := json.Marshal(errorResponse)
				response.Write([]byte(errorResp))
				return
			}

			user, err := hs.HttpServ.GetUser(id)

			if err != nil {
				response.WriteHeader(http.StatusBadRequest)
				errorResponse := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "USER NOT AVAILABLE"}
				errorResp, _ := json.Marshal(errorResponse)
				response.Write([]byte(errorResp))
			}

			jsonUserObject, _ := json.Marshal(user)
			response.Write([]byte(jsonUserObject))
		}
	case http.MethodPost:
		decoder := json.NewDecoder(request.Body)
		var usr models.User
		err := decoder.Decode(&usr)

		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			errorResponse := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "EMPTY USER"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		}

		err = hs.HttpServ.AddUser(usr)

		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			errorResponse := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "ERROR IN ADDING USER"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		} else {
			response.WriteHeader(http.StatusOK)
			errorResponse := models.ErrorResponse{StatusCode: http.StatusOK, ErrorMessage: "USER ADDED SUCCESSFULLY"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		}
	case http.MethodPut:
		decoder := json.NewDecoder(request.Body)
		var usr models.User
		err := decoder.Decode(&usr)

		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			errorResponse := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "EMPTY USER"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		}

		err = hs.HttpServ.UpdateUser(usr)

		if err != nil {
			response.WriteHeader(http.StatusBadRequest)
			errorResponse := models.ErrorResponse{StatusCode: http.StatusBadRequest, ErrorMessage: "ERROR IN UPDATING USER"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		} else {
			response.WriteHeader(http.StatusOK)
			errorResponse := models.ErrorResponse{StatusCode: http.StatusOK, ErrorMessage: "USER UPDATED SUCCESSFULLY"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		}
	case http.MethodDelete:

		userId := request.URL.Query().Get("id")
		id, _ := strconv.Atoi(userId)

		err := hs.HttpServ.DeleteUser(id)

		if err != nil || id == 0 {
			response.WriteHeader(http.StatusBadRequest)
			errorResponse := models.ErrorResponse{StatusCode: http.StatusOK, ErrorMessage: "ERROR IN DELETING"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		} else {
			response.WriteHeader(http.StatusOK)
			errorResponse := models.ErrorResponse{StatusCode: http.StatusOK, ErrorMessage: "DELETED SUCCESSFULLY"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		}
	}
}
