package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/tejas/user-crud/models"
	"github.com/tejas/user-crud/services"
)

func TestGetUserById(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockUserService := services.NewMockUser(ctrl)
	userApi := New(mockUserService)

	defer ctrl.Finish()

	testCases := []struct {
		desc           string
		input          string
		expectedOut    models.User
		expectedErr    error
		expectedStatus int
	}{

		{
			desc:  "Case 1: Success case",
			input: "1",
			expectedOut: models.User{
				Id:    1,
				Name:  "prasath",
				Email: "naveeth99@hello.com",
				Phone: "8967457789",
				Age:   19,
			},
			expectedErr:    nil,
			expectedStatus: http.StatusOK,
		},

		{
			desc:           "Case 2: Failure case server error",
			input:          "abc",
			expectedStatus: http.StatusInternalServerError,
			expectedErr:    errors.New("invalid id"),
		},

		{
			desc:           "Case 3: Bad request error",
			input:          "-1",
			expectedStatus: http.StatusBadRequest,
			expectedErr:    errors.New("error invalid id type"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+tc.desc, func(t *testing.T) {

			url := "/users/" + tc.input
			req := httptest.NewRequest("GET", url, nil)
			wr := httptest.NewRecorder()

			req = mux.SetURLVars(req, map[string]string{
				"id": tc.input,
			})
			id, err := strconv.Atoi(tc.input)
			if err == nil {
				mockUserService.EXPECT().GetUserById(id).
					Return(&tc.expectedOut, tc.expectedErr)
			}

			userApi.FindUserById(wr, req)
			if wr.Code != tc.expectedStatus {
				t.Errorf("Expected: %v, Got: %v, ", tc.expectedStatus, wr.Code)
			}
		})
	}
}

func TestGetUsers(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockUserService := services.NewMockUser(ctrl)
	userApi := New(mockUserService)

	defer ctrl.Finish()

	testCases := []struct {
		desc           string
		expectedOut    []*models.User
		expectedErr    error
		expectedStatus int
	}{

		{
			desc: "Case 1: Success case",
			expectedOut: []*models.User{
				{
					Id:    1,
					Name:  "sumeet",
					Email: "sumeet@gmail.com",
					Phone: "8951432891",
					Age:   21,
				},
				{
					Id:    2,
					Name:  "madhu",
					Email: "madhuH99@gmail.com",
					Phone: "9669696969",
					Age:   23,
				},
			},
			expectedErr:    nil,
			expectedStatus: http.StatusOK,
		},

		{
			desc:           "Case 2: Failure case",
			expectedErr:    errors.New("error fetching data from databases"),
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+tc.desc, func(t *testing.T) {

			url := "/users"
			req := httptest.NewRequest("GET", url, nil)
			wr := httptest.NewRecorder()
			mockUserService.EXPECT().GetUsers().
				Return(tc.expectedOut, tc.expectedErr)

			userApi.FindUsers(wr, req)
			if wr.Code != tc.expectedStatus {
				t.Errorf("Expected: %v, Got: %v, ", tc.expectedStatus, wr.Code)
			}
		})
	}

}

func TestUserApi_CreateUserHandler(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockUserService := services.NewMockUser(ctrl)
	userApi := New(mockUserService)

	defer ctrl.Finish()

	testCases := []struct {
		desc  string
		input struct {
			Id    interface{}
			Name  interface{}
			Email interface{}
			Phone interface{}
			Age   interface{}
		}
		mockCall       bool
		expectedErr    error
		expectedStatus int
	}{

		{
			desc: "Case 1: Success Case",
			input: struct {
				Id    interface{}
				Name  interface{}
				Email interface{}
				Phone interface{}
				Age   interface{}
			}{
				Id:    3,
				Name:  "prasath",
				Email: "prasath66@gmail.com",
				Phone: "7892567212",
				Age:   21,
			},
			mockCall:       true,
			expectedErr:    nil,
			expectedStatus: http.StatusCreated,
		},

		{
			desc: "Case 2: Invalid Phone",
			input: struct {
				Id    interface{}
				Name  interface{}
				Email interface{}
				Phone interface{}
				Age   interface{}
			}{
				Id:    1,
				Name:  "prasath",
				Email: "prasath2@lmnop.com",
				Phone: "",
				Age:   21,
			},
			mockCall:       true,
			expectedErr:    errors.New("error invalid phone"),
			expectedStatus: http.StatusBadRequest,
		},

		{
			desc: "Case 3: Server error",
			input: struct {
				Id    interface{}
				Name  interface{}
				Email interface{}
				Phone interface{}
				Age   interface{}
			}{
				Id:    1,
				Name:  "prasath",
				Email: "prasath@lmnop.com",
				Phone: "9080706655",
				Age:   "21",
			},
			mockCall:       false,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+tc.desc, func(t *testing.T) {

			url := "/users"
			jsnBdy, _ := json.Marshal(tc.input)
			buff := bytes.NewBuffer(jsnBdy)
			req := httptest.NewRequest("POST", url, buff)
			wr := httptest.NewRecorder()

			if tc.mockCall {
				in := models.User{
					Id:    tc.input.Id.(int),
					Name:  tc.input.Name.(string),
					Email: tc.input.Email.(string),
					Phone: tc.input.Phone.(string),
					Age:   tc.input.Age.(int),
				}
				mockUserService.EXPECT().CreateUser(in).
					Return(tc.expectedErr)
			}
			userApi.CreateUser(wr, req)
			if wr.Code != tc.expectedStatus {
				t.Errorf("Expected: %v, Got: %v, ", tc.expectedStatus, wr.Code)
			}
		})
	}

}

func TestUpdateUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockUserService := services.NewMockUser(ctrl)
	userApi := New(mockUserService)

	defer ctrl.Finish()

	testCases := []struct {
		desc      string
		inputId   string
		inputData struct {
			Name  interface{}
			Email interface{}
			Phone interface{}
			Age   interface{}
		}
		expectedErr    error
		expectedStatus int
		mockCall       bool
	}{

		{
			desc:    "Case 1: Success case",
			inputId: "2",
			inputData: struct {
				Name  interface{}
				Email interface{}
				Phone interface{}
				Age   interface{}
			}{
				Name:  "naveeth",
				Email: "naveeth@gmail.com",
				Phone: "7891232167",
				Age:   32,
			},
			mockCall:       true,
			expectedErr:    nil,
			expectedStatus: http.StatusOK,
		},

		{
			desc:    "Case 2: Server error",
			inputId: "1",
			inputData: struct {
				Name  interface{}
				Email interface{}
				Phone interface{}
				Age   interface{}
			}{
				Name:  "naveeth",
				Email: "naveeth33@gmail.com",
				Phone: "7891232167",
				Age:   "32",
			},
			mockCall:       false,
			expectedErr:    nil,
			expectedStatus: http.StatusInternalServerError,
		},

		{
			desc:    "Case 3: Server error",
			inputId: "abc",
			inputData: struct {
				Name  interface{}
				Email interface{}
				Phone interface{}
				Age   interface{}
			}{
				Name:  "vinod",
				Email: "vinod@masrshal.com",
				Phone: "9087989898",
				Age:   32,
			},
			mockCall:       false,
			expectedErr:    nil,
			expectedStatus: http.StatusInternalServerError,
		},

		{
			desc:    "Case 4: Bad Request",
			inputId: "4",
			inputData: struct {
				Name  interface{}
				Email interface{}
				Phone interface{}
				Age   interface{}
			}{
				Name:  "naveeth",
				Email: "naveeth22@gmail.com",
				Phone: "7891232167",
				Age:   32,
			},
			mockCall:       true,
			expectedErr:    errors.New("error, no id provided, cannot update"),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+tc.desc, func(t *testing.T) {

			url := "/users/" + tc.inputId

			jsnBdy, _ := json.Marshal(tc.inputData)
			buff := bytes.NewBuffer(jsnBdy)
			req := httptest.NewRequest("PUT", url, buff)
			req = mux.SetURLVars(req, map[string]string{
				"id": tc.inputId,
			})
			wr := httptest.NewRecorder()

			id, err := strconv.Atoi(tc.inputId)

			if err == nil && tc.mockCall {
				inp := models.User{
					Id:    id,
					Name:  tc.inputData.Name.(string),
					Email: tc.inputData.Email.(string),
					Phone: tc.inputData.Phone.(string),
					Age:   tc.inputData.Age.(int),
				}
				mockUserService.EXPECT().UpdateUser(inp).
					Return(tc.expectedErr)
			}

			userApi.UpdateById(wr, req)
			if wr.Code != tc.expectedStatus {
				t.Errorf("Expected: %v, Got: %v, ",tc.expectedStatus, wr.Code)
			}
		})
	}
}

func TestUserApi_DeleteUserHandler(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockUserService := services.NewMockUser(ctrl)
	userApi := New(mockUserService)

	defer ctrl.Finish()

	testCases := []struct {
		desc           string
		inputId        string
		expectedStatus int
		expectedErr    error
	}{

		{
			desc:           "Case 1: Success case",
			inputId:        "2",
			expectedStatus: http.StatusOK,
			expectedErr:    nil,
		},

		{
			desc:           "Case 2: Server Error",
			inputId:        "abc",
			expectedStatus: http.StatusInternalServerError,
			expectedErr:    nil,
		},

		{
			desc:           "Case 3: Bad Request",
			inputId:        "-1",
			expectedStatus: http.StatusBadRequest,
			expectedErr:    errors.New("error invalid id"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+tc.desc, func(t *testing.T) {

			url := "/users/" + tc.inputId
			req := httptest.NewRequest("DELETE", url, nil)

			req = mux.SetURLVars(req, map[string]string{
				"id": tc.inputId,
			})
			wr := httptest.NewRecorder()

			id, err := strconv.Atoi(tc.inputId)
			if err == nil {
				mockUserService.EXPECT().DeleteUser(id).
					Return(tc.expectedErr)
			}
			userApi.DeleteUser(wr, req)
			if wr.Code != tc.expectedStatus {
				t.Errorf("Expected: %v, Got: %v, ", tc.expectedStatus, wr.Code)
			}
		})
	}
}
