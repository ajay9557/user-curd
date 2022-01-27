package Users

import (
	"Icrud/Services"
	"Icrud/TModels"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService := Services.NewMockISUser(ctrl)
	h := Handler{mockUserService}

	testUser := TModels.User{Id: 1, Name: "Naruto", Email: "naruto@japan.com", Phone: "9999999999", Age: 180}

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
			mockCall:           mockUserService.EXPECT().UserById(1).Return(testUser, nil),
		},
		{
			desc:               "Case2",
			id:                 "2",
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           mockUserService.EXPECT().UserById(2).Return(TModels.User{}, errors.New("Invalid Id")),
		},
		{
			desc:               "Case3",
			id:                 "id",
			expectedStatusCode: http.StatusInternalServerError,
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

		h.UserById(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}

func TestGetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService := Services.NewMockISUser(ctrl)
	h := Handler{mockUserService}

	data1 := []TModels.User{
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
			mockCall:           mockUserService.EXPECT().GetUsers().Return([]TModels.User{}, errors.New("Invalid Id")),
		},
	}

	for _, test := range tests {
		// Creating test request and response object
		req := httptest.NewRequest("GET", "/api/users/", nil)
		res := httptest.NewRecorder()

		h.GetUsers(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}

func TestDeleteUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService := Services.NewMockISUser(ctrl)
	h := Handler{mockUserService}

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
			mockCall:           mockUserService.EXPECT().DeleteUserById(1).Return(1, nil),
		},
		{
			desc:               "Case2",
			id:                 "2",
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           mockUserService.EXPECT().DeleteUserById(2).Return(0, errors.New("Invalid Id")),
		},
		{
			desc:               "Case3",
			id:                 "abcd",
			expectedStatusCode: http.StatusInternalServerError,
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

		h.DeleteUserById(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}

func TestInsertUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService := Services.NewMockISUser(ctrl)
	h := Handler{mockUserService}

	testUser := TModels.User{Name: "Ridhdhish", Email: "ridhdhish@gmail.com", Phone: "8320578360", Age: 21}

	tests := []struct {
		desc               string
		user               TModels.User
		expectedStatusCode int
		mockCall           *gomock.Call
	}{
		{
			desc:               "Case1",
			user:               testUser,
			expectedStatusCode: http.StatusOK,
			mockCall:           mockUserService.EXPECT().InsertUser(testUser).Return(1, nil),
		},
		{
			desc:               "Case2",
			user:               TModels.User{},
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
		{
			desc:               "Case3",
			user:               testUser,
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           mockUserService.EXPECT().InsertUser(testUser).Return(0, errors.New("Could not create new user")),
		},
	}
	for _, test := range tests {
		body, _ := json.Marshal(test.user)
		// Creating test request and response object
		req := httptest.NewRequest("POST", "/api/users", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		h.InsertUser(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}

func TestUpdateUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserService := Services.NewMockISUser(ctrl)
	h := Handler{mockUserService}

	testUser := TModels.User{Name: "Ridhdhish", Email: "ridhdhish@gmail.com", Phone: "8320578360", Age: 21}

	tests := []struct {
		desc               string
		id                 string
		expectedStatusCode int
		body               TModels.User
		mockCall           *gomock.Call
	}{
		{
			desc:               "Case1",
			id:                 "1",
			body:               testUser,
			expectedStatusCode: http.StatusOK,
			mockCall:           mockUserService.EXPECT().UpdateUserById(testUser, 1).Return(1, nil),
		},
		{
			desc:               "Case2",
			id:                 "2",
			body:               testUser,
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           mockUserService.EXPECT().UpdateUserById(testUser, 2).Return(0, errors.New("Invalid Id")),
		},
		{
			desc:               "Case3",
			id:                 "1",
			body:               TModels.User{},
			expectedStatusCode: http.StatusBadRequest,
			mockCall:           nil,
		},
		{
			desc:               "Case4",
			id:                 "abcd",
			body:               testUser,
			expectedStatusCode: http.StatusInternalServerError,
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

		h.UpdateUserById(res, req)

		if res.Code != test.expectedStatusCode {
			t.Errorf("Expected Status Code: %v, Got: %v", test.expectedStatusCode, res.Code)
		}
	}
}
