package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"user-curd/entities"
	"user-curd/service"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestUserApi_GetUserByIdHandler(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockUserService := service.NewMockUserServiceHandler(ctrl)
	userApi := New(mockUserService)

	defer ctrl.Finish()

	testCases := []struct {
		caseId int
		//mockCall       *gomock.Call
		input          string
		expectedOut    entities.User
		expectedErr    error
		expectedStatus int
	}{
		// Success case
		{
			caseId: 1,
			input:  "1",
			//expectedOut: entities.User{
			//	Id:    1,
			//	Name:  "jojo",
			//	Email: "jonathan99@bizzar.com",
			//	Phone: "8967457789",
			//	Age:   19,
			//},
			expectedErr:    nil,
			expectedStatus: http.StatusOK,
		},
		// Error Internal Server error
		{
			caseId:         2,
			input:          "abc",
			expectedStatus: http.StatusInternalServerError,
			expectedErr:    errors.New("invalid id"),
		},
		// Error Bad request
		{
			caseId:         3,
			input:          "-1",
			expectedStatus: http.StatusBadRequest,
			expectedErr:    errors.New("error invalid id type"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {

			// define url and test request and response
			url := "/users/" + tc.input
			req := httptest.NewRequest("GET", url, nil)
			wr := httptest.NewRecorder()

			req = mux.SetURLVars(req, map[string]string{
				"id": tc.input,
			})
			id, err := strconv.Atoi(tc.input)
			if err == nil {
				mockUserService.EXPECT().GetUserByIdService(id).
					Return(&tc.expectedOut, tc.expectedErr)
			}

			userApi.GetUserByIdHandler(wr, req)
			if wr.Code != tc.expectedStatus {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedStatus, wr.Code)
			}
		})
	}
}

func TestUserApi_GetAllUserHandler(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockUserService := service.NewMockUserServiceHandler(ctrl)
	userApi := New(mockUserService)

	defer ctrl.Finish()

	testCases := []struct {
		caseId         int
		expectedOut    []*entities.User
		expectedErr    error
		expectedStatus int
	}{
		// Success case
		{
			caseId: 1,
			expectedOut: []*entities.User{
				{
					Id:    1,
					Name:  "john",
					Email: "john19@example.com",
					Phone: "9823218772",
					Age:   21,
				},
				{
					Id:    2,
					Name:  "jess",
					Email: "jessH99@example.com",
					Phone: "9821328772",
					Age:   19,
				},
			},
			expectedErr:    nil,
			expectedStatus: http.StatusOK,
		},
		// No data in database
		{
			caseId:         2,
			expectedErr:    errors.New("error fetching data from databases"),
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {

			// define url and test request and response
			url := "/users"
			req := httptest.NewRequest("GET", url, nil)
			wr := httptest.NewRecorder()
			mockUserService.EXPECT().GetAllUsersService().
				Return(tc.expectedOut, tc.expectedErr)

			userApi.GetAllUserHandler(wr, req)
			if wr.Code != tc.expectedStatus {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedStatus, wr.Code)
			}
		})
	}

}

func TestUserApi_CreateUserHandler(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockUserService := service.NewMockUserServiceHandler(ctrl)
	userApi := New(mockUserService)

	defer ctrl.Finish()

	testCases := []struct {
		caseId         int
		input          entities.User
		expectedErr    error
		expectedStatus int
	}{
		// Success Case
		{
			caseId: 1,
			input: entities.User{
				Id:    3,
				Name:  "jotaro",
				Email: "jojo66@example.com",
				Phone: "7892567212",
				Age:   21,
			},
			expectedErr:    nil,
			expectedStatus: http.StatusCreated,
		},
		// Internal server err with invalid phone
		{
			caseId: 2,
			input: entities.User{
				Id:    1,
				Name:  "jotaro",
				Phone: "",
			},
			expectedErr:    errors.New("error invalid phone"),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {

			// define url and test request and response
			url := "/users"
			jsnBdy, _ := json.Marshal(tc.input)
			buff := bytes.NewBuffer(jsnBdy)
			req := httptest.NewRequest("POST", url, buff)
			wr := httptest.NewRecorder()

			if tc.expectedStatus != http.StatusInternalServerError {
				mockUserService.EXPECT().CreateUserService(tc.input).
					Return(tc.expectedErr)
			}
			userApi.CreateUserHandler(wr, req)
			if wr.Code != tc.expectedStatus {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedStatus, wr.Code)
			}
		})
	}

}

func TestUserApi_UpdateUserHandler(t *testing.T) {

	// mock for service layer
	ctrl := gomock.NewController(t)
	mockUserService := service.NewMockUserServiceHandler(ctrl)
	userApi := New(mockUserService)

	defer ctrl.Finish()

	testCases := []struct {
		caseId    int
		inputId   string
		inputData struct {
			Name  string
			Email string
			Phone string
			Age   int
		}
		expectedErr    error
		expectedStatus int
	}{
		// Success case
		{
			caseId:  1,
			inputId: "2",
			inputData: struct {
				Name  string
				Email string
				Phone string
				Age   int
			}{
				Name:  "jonathan",
				Email: "jon34@example.com",
				Phone: "7891232167",
				Age:   32,
			},
			expectedErr:    nil,
			expectedStatus: http.StatusOK,
		},
		// Error Bad request
		{
			caseId:  2,
			inputId: "anc",
			inputData: struct {
				Name  string
				Email string
				Phone string
				Age   int
			}{
				Name:  "jonathan",
				Email: "jon34@example.com",
				Phone: "7891232167",
				Age:   32,
			},
			expectedErr:    nil,
			expectedStatus: http.StatusInternalServerError,
		},
		// Error bad request
		{
			caseId:  1,
			inputId: "4",
			inputData: struct {
				Name  string
				Email string
				Phone string
				Age   int
			}{
				Name:  "jonathan",
				Email: "jon34@example.com",
				Phone: "7891232167",
				Age:   32,
			},
			expectedErr:    errors.New("error, no id provided, cannot update"),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {

			// define url and test request and response
			url := "/users/" + tc.inputId

			// json input
			jsnBdy, _ := json.Marshal(tc.inputData)
			buff := bytes.NewBuffer(jsnBdy)
			req := httptest.NewRequest("PUT", url, buff)
			req = mux.SetURLVars(req, map[string]string{
				"id": tc.inputId,
			})
			wr := httptest.NewRecorder()

			id, err := strconv.Atoi(tc.inputId)
			if err == nil {
				inp := entities.User{
					Id:    id,
					Name:  tc.inputData.Name,
					Email: tc.inputData.Email,
					Phone: tc.inputData.Phone,
					Age:   tc.inputData.Age,
				}
				mockUserService.EXPECT().UpdateUserService(inp).
					Return(tc.expectedErr)
			}

			userApi.UpdateUserHandler(wr, req)
			if wr.Code != tc.expectedStatus {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedStatus, wr.Code)
			}
		})
	}
}

func TestUserApi_DeleteUserHandler(t *testing.T) {

	// mock for service layer
	ctrl := gomock.NewController(t)
	mockUserService := service.NewMockUserServiceHandler(ctrl)
	userApi := New(mockUserService)

	defer ctrl.Finish()

	testCases := []struct {
		caseId         int
		inputId        string
		expectedStatus int
		expectedErr    error
	}{
		// Success case
		{
			caseId:         1,
			inputId:        "2",
			expectedStatus: http.StatusOK,
			expectedErr:    nil,
		},
		// Error case : status internal server error
		{
			caseId:         2,
			inputId:        "abc",
			expectedStatus: http.StatusInternalServerError,
			expectedErr:    nil,
		},
		// Error case : Status bad request
		{
			caseId:         2,
			inputId:        "-1",
			expectedStatus: http.StatusBadRequest,
			expectedErr:    errors.New("error invalid id"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {

			// define url and test request and response
			url := "/users/" + tc.inputId
			req := httptest.NewRequest("DELETE", url, nil)

			// set url variables
			req = mux.SetURLVars(req, map[string]string{
				"id": tc.inputId,
			})
			wr := httptest.NewRecorder()

			id, err := strconv.Atoi(tc.inputId)
			if err == nil {
				mockUserService.EXPECT().DeleteUserService(id).
					Return(tc.expectedErr)
			}
			userApi.DeleteUserHandler(wr, req)
			if wr.Code != tc.expectedStatus {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedStatus, wr.Code)
			}
		})
	}
}
