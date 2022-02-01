package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"layer/user/models"
	"layer/user/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

/*
Test for Fetch User by Id
/api/users/{id}
*/
func TestUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService := services.NewMockUser(ctrl)
	h := New(mockUserService)

	testUser := models.User{Id: 1, Name: "Naruto", Email: "naruto@japan.com", Phone: "9999999999", Age: 180}

	tests := []struct {
		desc               string
		id                 string
		expectedStatusCode int
		mockCall           *gomock.Call
	}{
		{
			desc:               "Case1",
			id:                 "1",
			expectedStatusCode: http.StatusOK,
			mockCall:           mockUserService.EXPECT().GetUserById(1).Return(testUser, nil),
		},
		{
			desc:               "Case2",
			id:                 "2",
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           mockUserService.EXPECT().GetUserById(2).Return(models.User{}, errors.New("Invalid Id")),
		},
		{
			desc:               "Case3",
			id:                 "id",
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
		{
			desc:               "Case4",
			id:                 "-1",
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
	}

	for _, test := range tests {
		// Creating test request and response object
		req := httptest.NewRequest("GET", "/api/users/"+test.id, nil)
		res := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": test.id,
		})

		h.GetUserByIdHandler(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}

func TestGetUsersHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService := services.NewMockUser(ctrl)
	h := New(mockUserService)

	data1 := []*models.User{
		{Id: 1, Name: "Naruto", Email: "naruto@gmail.com", Phone: "9999999999", Age: 18},
		{Id: 2, Name: "Itachi", Email: "itachi@gmail.com", Phone: "8320578360", Age: 24},
	}

	tests := []struct {
		desc               string
		expectedStatusCode int
		mockCall           *gomock.Call
	}{
		{
			desc:               "Case1",
			expectedStatusCode: http.StatusOK,
			mockCall:           mockUserService.EXPECT().GetUsers().Return(data1, nil),
		},
		{
			desc:               "Case2",
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           mockUserService.EXPECT().GetUsers().Return(nil, errors.New("Invalid Id")),
		},
	}

	for _, test := range tests {
		// Creating test request and response object
		req := httptest.NewRequest("GET", "/api/users/", nil)
		res := httptest.NewRecorder()

		h.GetUsersHandler(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService := services.NewMockUser(ctrl)
	h := New(mockUserService)

	testUser := models.User{Name: "Ridhdhish", Email: "ridhdhish@gmail.com", Phone: "8320578360", Age: 21}
	updatedUser := models.User{Name: "Ridhdhish", Email: "ridhdhish@gmail.com", Phone: "8320578360", Age: 21}

	tests := []struct {
		desc               string
		id                 string
		expectedStatusCode int
		body               models.User
		mockCall           *gomock.Call
	}{
		{
			desc:               "Case1",
			id:                 "1",
			body:               testUser,
			expectedStatusCode: http.StatusOK,
			mockCall:           mockUserService.EXPECT().UpdateUser(1, testUser).Return(&updatedUser, nil),
		},
		{
			desc:               "Case2",
			id:                 "2",
			body:               testUser,
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           mockUserService.EXPECT().UpdateUser(2, testUser).Return(&models.User{}, errors.New("Invalid Id")),
		},
		{
			desc:               "Case3",
			id:                 "1",
			body:               models.User{},
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
		{
			desc:               "Case4",
			id:                 "abcd",
			body:               testUser,
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
		{
			desc:               "Case5",
			id:                 "-1",
			body:               testUser,
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
	}

	for _, test := range tests {
		// Setting up body of request
		body, _ := json.Marshal(test.body)

		// Creating test request and response object
		req := httptest.NewRequest("PUT", "/api/users/"+test.id, bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": test.id,
		})

		h.UpdateUserHandler(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService := services.NewMockUser(ctrl)
	h := New(mockUserService)

	tests := []struct {
		desc               string
		id                 string
		expectedStatusCode int
		mockCall           *gomock.Call
	}{
		{
			desc:               "Case1",
			id:                 "1",
			expectedStatusCode: http.StatusOK,
			mockCall:           mockUserService.EXPECT().DeleteUser(1).Return(1, nil),
		},
		{
			desc:               "Case2",
			id:                 "2",
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           mockUserService.EXPECT().DeleteUser(2).Return(0, errors.New("Invalid Id")),
		},
		{
			desc:               "Case3",
			id:                 "abcd",
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
		{
			desc:               "Case4",
			id:                 "-1",
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
	}
	for _, test := range tests {
		// Creating test request and response object
		req := httptest.NewRequest("DELETE", "/api/users/"+test.id, nil)
		res := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": test.id,
		})

		h.DeleteUserHandler(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService := services.NewMockUser(ctrl)
	h := New(mockUserService)

	testUser := models.User{Name: "Ridhdhish", Email: "ridhdhish@gmail.com", Phone: "8320578360", Age: 21}

	tests := []struct {
		desc               string
		user               models.User
		expectedStatusCode int
		mockCall           *gomock.Call
	}{
		{
			desc:               "Case1",
			user:               testUser,
			expectedStatusCode: http.StatusOK,
			mockCall:           mockUserService.EXPECT().CreateUser(testUser).Return(&testUser, nil),
		},
		{
			desc:               "Case2",
			user:               models.User{},
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
		{
			desc:               "Case3",
			user:               testUser,
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           mockUserService.EXPECT().CreateUser(testUser).Return(&models.User{}, errors.New("Could not create new user")),
		},
	}
	for _, test := range tests {
		body, _ := json.Marshal(test.user)
		// Creating test request and response object
		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		h.CreateUserHandler(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}
