package http

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"zopsmart/Task/models"
	"zopsmart/Task/services"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestApiGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockServices(ctrl)
	mock := New(mockService)

	testCases := []struct {
		id        string
		mock      *gomock.Call
		expectErr  error
		expectOut models.User
		expectstat int
	}{
		{
			id: "1",
			expectErr: nil,
			mock: 
				mockService.EXPECT().GetUserById(1).Return(models.User{
					Id:    1,
					Name:  "prasath",
					Email: "prasath@gmail.com",
					Phone: "12345",
					Age:   20,
				}, nil),
			expectOut: models.User{
				Id:1,
				Name:"prasath",
				Email:"prasath@gmail.com",
				Phone:"12345",
				Age:20,
			},
			expectstat: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run("sucess test case",func(t *testing.T) {

			url := "/users/" + tc.id
			req := httptest.NewRequest("GET", url, nil)
			wr := httptest.NewRecorder()

			req = mux.SetURLVars(req, map[string]string{
				"id": tc.id,
			})

			mock.GetUserById(wr, req)
			if wr.Code != tc.expectstat {
				t.Errorf("Expected: \t%v\nGot: \t%v\n", tc.expectstat, wr.Code)
			}
		})
	}
	
}

func TestUserApi_GetAllUserHandler(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockService := service.NewMockServices(ctrl)
	userApi := New(mockService)

	defer ctrl.Finish()

	testCases := []struct {
		expectedOut    []models.User
		expectedErr    error
		expectedStatus int
	}{
		{
			expectedOut: []models.User{
				{
					Id:    1,
					Name:  "prasath",
					Email: "prasath@gmail.com",
					Phone: "12345",
					Age:   20,
				},
				{
					Id:    2,
					Name:  "rishi kumar",
					Email: "rishi.kumar@example.com",
					Phone: "9821328772",
					Age:   25,
				},
			},
			expectedErr:    nil,
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run("Success case testing ", func(t *testing.T) {

			// define url and test request and response
			url := "/users"
			req := httptest.NewRequest("GET", url, nil)
			wr := httptest.NewRecorder()
			mockService.EXPECT().GetAllUsersService().
				Return(tc.expectedOut, tc.expectedErr)

			userApi.AllUserDetails(wr, req)
			if wr.Code != tc.expectedStatus {
				t.Errorf("Expected: \t%v\nGot: \t%v\n", tc.expectedStatus, wr.Code)
			}
		})
	}

}

func TestApiDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockServices(ctrl)
	mock := New(mockService)

	testCases := []struct {
		id  int
		expectErr  error
		expectstat  int
	}{
		{
			id:       1,
			expectErr: nil,
			expectstat : http.StatusOK,
		},
	}

	for _,tc := range testCases {
		t.Run("Success test case",func (t *testing.T)  {
			url := "/user/"+strconv.Itoa(tc.id)
			req:= httptest.NewRequest("DELETE", url, nil)
			writer := httptest.NewRecorder()
			mock.DeleteUser(writer, req)
			if writer.Code != tc.expectstat {
				t.Errorf("Expected: \t%v\nGot: \t%v\n", tc.expectstat, writer.Code)
			}
			
			
		})
	}
}

func TestApiUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockServices(ctrl)
	mock := New(mockService)

	testCases := []struct {
		id  int
		expectErr  error
		expectstat  int
	}{
		{
			id:       1,
			expectErr: nil,
			expectstat : http.StatusOK,
		},
	}

	for _,tc := range testCases {
		t.Run("Success test case",func (t *testing.T)  {
			url := "/user/"+strconv.Itoa(tc.id)
			req:= httptest.NewRequest("UPDATE", url, nil)
			writer := httptest.NewRecorder()
			mock.UpdateUser(writer, req)
			if writer.Code != tc.expectstat {
				t.Errorf("Expected: \t%v\nGot: \t%v\n", tc.expectstat, writer.Code)
			}
			
			
		})
	}
}


func TestApiCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	 mockService := service.NewMockServices(ctrl)
	 mock := New(mockService)

	testCases := []struct {
		id int
		expectErr error
		expectstat int

	}{
		{
			id:       1,
			expectErr: nil,
			expectstat: http.StatusOK,
		},
	}

	for _,tc := range testCases {
		t.Run("success test case",func(t *testing.T) {
			url := "/user"
			req:= httptest.NewRequest("INSERT", url, nil)
			writer := httptest.NewRecorder()
			mock.CreateUser(writer, req)
			if writer.Code != tc.expectstat {
				t.Errorf("Expected: \t%v\nGot: \t%v\n", tc.expectstat, writer.Code)
			}
		})
	}
}
