package httpuser

import (
	"bytes"
	"encoding/json"
	"errors"
	"go_lang/Assignment/user-curd/models"
	"go_lang/Assignment/user-curd/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_GetOneUserHandler(t *testing.T) {

	controller := gomock.NewController(t)
	mockUserService := services.NewMockServices(controller)
	handler := HttpService{mockUserService}

	testCaseGetById := []struct {
		desc     string
		input    string
		param    string
		expected int
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case1",
			input:    "GET",
			param:    "1",
			expected: http.StatusOK,
			mock:     []*gomock.Call{mockUserService.EXPECT().GetUser(1).Return(models.User{Id: 11, Name: "Ravi", Email: "ravi34@gmail.com", Phone: "9876789098", Age: 45}, nil)},
		},
		{
			desc:     "Test Case2",
			input:    "GET",
			param:    "1000",
			expected: http.StatusBadRequest,
			mock:     []*gomock.Call{mockUserService.EXPECT().GetUser(1000).Return(models.User{}, errors.New("INVLID ID"))},
		},
		{
			desc:     "Test Case3",
			input:    "GET",
			param:    "abc",
			expected: http.StatusBadRequest,
			mock:     nil,
		},
	}
	for _, tcs := range testCaseGetById {
		t.Run(tcs.desc, func(t *testing.T) {
			req := httptest.NewRequest(tcs.input, "/?id="+tcs.param, nil)
			res := httptest.NewRecorder()

			handler.GetOneUserHandler(res, req)
			if res.Code != tcs.expected {
				t.Errorf("Output: %v, Expected: %v", res.Code, tcs.expected)
			}
		})
	}
}

func Test_GetAllUserHandler(t *testing.T) {
	controller := gomock.NewController(t)
	mockUserService := services.NewMockServices(controller)
	handler := HttpService{mockUserService}
	testCaseGet := []struct {
		desc     string
		input    string
		expected int
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case1",
			input:    "GET",
			expected: http.StatusOK,
			mock:     []*gomock.Call{mockUserService.EXPECT().GetAllUser().Return([]models.User{{Id: 11, Name: "Ravi", Email: "ravi34@gmail.com", Phone: "9876789098", Age: 45}}, nil)},
		},
		{
			desc:     "Test Case2",
			input:    "GET",
			expected: http.StatusBadRequest,
			mock:     []*gomock.Call{mockUserService.EXPECT().GetAllUser().Return([]models.User{}, errors.New("ERROR IN FETCHING ROWS"))},
		},
	}
	for _, tcs := range testCaseGet {
		t.Run(tcs.desc, func(t *testing.T) {
			req := httptest.NewRequest(tcs.input, "/", nil)
			res := httptest.NewRecorder()
			handler.GetAllUserHandler(res, req)
			if res.Code != tcs.expected {
				t.Errorf("Output: %v, Expected: %v", req, res)
			}
		})
	}
}

func Test_AddUserHandler(t *testing.T) {
	controller := gomock.NewController(t)
	mockUserService := services.NewMockServices(controller)
	handler := HttpService{mockUserService}

	testCasePost := []struct {
		desc     string
		input    string
		param    interface{}
		expected int
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case1",
			input:    "POST",
			param:    models.User{Id: 11, Name: "Ravi", Email: "ravi34@gmail.com", Phone: "9876789098", Age: 45},
			expected: http.StatusOK,
			mock:     []*gomock.Call{mockUserService.EXPECT().AddUser(models.User{Id: 11, Name: "Ravi", Email: "ravi34@gmail.com", Phone: "9876789098", Age: 45}).Return(nil)},
		},
		{
			desc:     "Test Case2",
			input:    "POST",
			param:    "",
			expected: http.StatusBadRequest,
			mock:     []*gomock.Call{mockUserService.EXPECT().AddUser(models.User{}).Return(errors.New("EMPTY USER"))},
		},
		{
			desc:     "Test Case3",
			input:    "POST",
			param:    models.User{},
			expected: http.StatusBadRequest,
			mock:     []*gomock.Call{mockUserService.EXPECT().AddUser(models.User{}).Return(errors.New("EMPTY USER"))},
		},
	}
	for _, tcs := range testCasePost {
		t.Run(tcs.desc, func(t *testing.T) {
			data, _ := json.Marshal(tcs.param)
			req := httptest.NewRequest(tcs.input, "/", bytes.NewBuffer(data))
			res := httptest.NewRecorder()

			handler.AddUserHandler(res, req)
			if res.Code != tcs.expected {
				t.Errorf("Output: %v, Expected: %v", res.Code, tcs.expected)
			}
		})
	}
}

func Test_UpdateUserHandler(t *testing.T) {
	controller := gomock.NewController(t)
	mockUserService := services.NewMockServices(controller)
	handler := HttpService{mockUserService}

	testCasePut := []struct {
		desc     string
		input    string
		param    interface{}
		expected int
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case1",
			input:    "PUT",
			param:    models.User{Id: 11, Name: "Ravi", Email: "ravi34@gmail.com", Phone: "9876789098", Age: 45},
			expected: http.StatusOK,
			mock:     []*gomock.Call{mockUserService.EXPECT().UpdateUser(models.User{Id: 11, Name: "Ravi", Email: "ravi34@gmail.com", Phone: "9876789098", Age: 45}).Return(nil)},
		},
		{
			desc:     "Test Case2",
			input:    "PUT",
			param:    "",
			expected: http.StatusBadRequest,
			mock:     []*gomock.Call{mockUserService.EXPECT().UpdateUser(models.User{}).Return(nil)},
		},
		{
			desc:     "Test Case3",
			input:    "PUT",
			param:    models.User{},
			expected: http.StatusBadRequest,
			mock:     []*gomock.Call{mockUserService.EXPECT().UpdateUser(models.User{}).Return(errors.New("ERROR IN UPDATING USER"))},
		},
	}
	for _, tcs := range testCasePut {
		t.Run(tcs.desc, func(t *testing.T) {
			data, _ := json.Marshal(tcs.param)
			req := httptest.NewRequest(tcs.input, "/", bytes.NewBuffer(data))
			res := httptest.NewRecorder()

			handler.UpdateUserHandler(res, req)
			if res.Code != tcs.expected {
				t.Errorf("Output: %v, Expected: %v", res.Code, tcs.expected)
			}
		})
	}
}

func Test_DeleteUserHandler(t *testing.T) {
	controller := gomock.NewController(t)
	mockUserService := services.NewMockServices(controller)
	handler := HttpService{mockUserService}

	testCaseDelete := []struct {
		desc     string
		input    string
		param    string
		expected int
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case1",
			input:    "DELETE",
			param:    "10",
			expected: http.StatusOK,
			mock:     []*gomock.Call{mockUserService.EXPECT().DeleteUser(10).Return(nil)},
		},
		{
			desc:     "Test Case2",
			input:    "DELETE",
			param:    "1111111111111111",
			expected: http.StatusBadRequest,
			mock:     []*gomock.Call{mockUserService.EXPECT().DeleteUser(1111111111111111).Return(errors.New("ERROR IN DELETING"))},
		},
	}
	for _, tcs := range testCaseDelete {
		t.Run(tcs.desc, func(t *testing.T) {
			req := httptest.NewRequest(tcs.input, "/?id="+tcs.param, nil)
			res := httptest.NewRecorder()

			handler.DeleteUserHandler(res, req)
			if res.Code != tcs.expected {
				t.Errorf("Output: %v, Expected: %v", res.Code, tcs.expected)
			}
		})
	}
}
