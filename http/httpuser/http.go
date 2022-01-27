package httpuser

import (
	"encoding/json"
	"go_lang/Assignment/user-curd/models"
	"go_lang/Assignment/user-curd/services/user"
	"net/http"
	"strconv"
)

type HttpService struct {
	HttpServ user.Services
}

func (hs HttpService) Handler(response http.ResponseWriter, request *http.Request) {

	response.Header().Add("content-type", "text/html")

	response.Write([]byte("<h2>Server Online</h2>"))

	if request.Method == http.MethodGet {

		userId := request.URL.Query().Get("id")

		if userId == "" {
			userList, err := hs.HttpServ.GetAllUser()

			if err != nil {
				errorResponse := models.ErrorResponse{http.StatusBadRequest, "ERROR IN FETCHING DATA"}
				errorResp, _ := json.Marshal(errorResponse)
				response.Write([]byte(errorResp))
			} else {
				userListFormat, _ := json.Marshal(userList)
				response.Write([]byte(userListFormat))
			}

		} else {

			id, err := strconv.Atoi(userId)

			if err != nil {
				errorResponse := models.ErrorResponse{http.StatusBadRequest, "INVALID ID"}
				errorResp, _ := json.Marshal(errorResponse)
				response.Write([]byte(errorResp))
			}

			user, err := hs.HttpServ.GetUser(id)

			if err != nil {
				errorResponse := models.ErrorResponse{http.StatusBadRequest, "USER NOT AVAILABLE"}
				errorResp, _ := json.Marshal(errorResponse)
				response.Write([]byte(errorResp))
			}

			jsonUserObject, _ := json.Marshal(user)
			response.Write([]byte(jsonUserObject))
		}

	} else if request.Method == http.MethodPost {

		decoder := json.NewDecoder(request.Body)
		var usr models.User
		err := decoder.Decode(&usr)

		if err != nil {
			errorResponse := models.ErrorResponse{http.StatusNoContent, "INVALID DATA FORMAT"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		}

		err = hs.HttpServ.AddUser(usr)

		if err != nil {
			errorResponse := models.ErrorResponse{http.StatusNoContent, "ERROR IN ADDING USER"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		} else {
			errorResponse := models.ErrorResponse{http.StatusOK, "USER ADDED SUCCESSFULLY"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		}
	} else if request.Method == http.MethodPut {

		decoder := json.NewDecoder(request.Body)
		var usr models.User
		err := decoder.Decode(&usr)

		if err != nil {
			errorResponse := models.ErrorResponse{http.StatusNoContent, "INVALID DATA FORMAT"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		}

		err = hs.HttpServ.UpdateUser(usr)

		if err != nil {
			errorResponse := models.ErrorResponse{http.StatusBadRequest, "ERROR IN UPDATING USER"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		} else {
			errorResponse := models.ErrorResponse{http.StatusOK, "UPDATED SUCCESSFULLY"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		}
	} else if request.Method == http.MethodDelete {

		var uidd struct {
			Id int `json: "id"`
		}
		_ = json.NewDecoder(request.Body).Decode(&uidd)

		err := hs.HttpServ.DeleteUser(uidd.Id)

		if err != nil {
			errorResponse := models.ErrorResponse{http.StatusBadRequest, "ERROR IN DELETING USER"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		} else {
			errorResponse := models.ErrorResponse{http.StatusOK, "DELETED SUCCESSFULLY"}
			errorResp, _ := json.Marshal(errorResponse)
			response.Write([]byte(errorResp))
		}
	}
}
