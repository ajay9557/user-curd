package users

import (
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
	h := Handler{mockUserService}

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
	h := Handler{mockUserService}

	data1 := []models.User{
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
			mockCall:           mockUserService.EXPECT().GetUsers().Return([]models.User{}, errors.New("Invalid Id")),
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
