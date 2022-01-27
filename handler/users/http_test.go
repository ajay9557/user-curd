package users

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-crud/models"
	"user-crud/services"
)

var usr = models.User{
	Id:    1,
	Name:  "himanshu",
	Email: "himanshu8083@gmail.com",
	Phone: "8083860404",
	Age:   25,
}

func TestUserHandler_GetById(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := services.NewMockUser(ctrl)
	handler := UserHandler{mockService}

	testCases := []struct {
		desc    string
		id      string
		mock    []*gomock.Call
		expCode int
		expBody []byte
	}{
		{
			desc:    "testcase-1",
			id:      "1",
			expCode: http.StatusOK,
			mock: []*gomock.Call{
				mockService.EXPECT().GetById(1).Return(usr, nil),
			},
			expBody: []byte(`{"id":1,"name":"himanshu","email":"himanshu8083@gmail.com","phone":"8083860404","age":25}`),
		},
		{
			desc:    "testcase-2",
			id:      "1a",
			expCode: http.StatusBadRequest,
			expBody: []byte("invalid id"),
		},
		{
			desc:    "testcase-3",
			id:      "100",
			expCode: http.StatusInternalServerError,
			mock: []*gomock.Call{
				mockService.EXPECT().GetById(100).Return(models.User{}, errors.New("id not found")),
			},
			expBody: []byte("id not found"),
		},
	}

	for _, tcs := range testCases {
		r := httptest.NewRequest("GET", fmt.Sprintf("/user/%s", tcs.id), nil)
		rw := httptest.NewRecorder()
		r = mux.SetURLVars(r, map[string]string{
			"id": tcs.id,
		})
		handler.GetById(rw, r)
		if rw.Body.String() != string(tcs.expBody) {
			t.Errorf("%v, Expected %v got %v", tcs.desc, string(tcs.expBody), rw.Body.String())
		}
	}
}

func TestUserHandler_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := services.NewMockUser(ctrl)
	handler := UserHandler{mockService}

	testCases := []struct {
		desc    string
		mock    []*gomock.Call
		expCode int
		expBody []byte
	}{
		{
			desc:    "testcase-1",
			expCode: http.StatusOK,
			mock: []*gomock.Call{
				mockService.EXPECT().GetAll().Return([]models.User{usr}, nil),
			},
			expBody: []byte(`[{"id":1,"name":"himanshu","email":"himanshu8083@gmail.com","phone":"8083860404","age":25}]`),
		},
		{
			desc:    "testcase-2",
			expCode: http.StatusBadRequest,
			mock: []*gomock.Call{
				mockService.EXPECT().GetAll().Return([]models.User{}, sql.ErrNoRows),
			},
			expBody: []byte("no data found"),
		},
	}

	for _, tcs := range testCases {
		r := httptest.NewRequest("GET", fmt.Sprintf("/user"), nil)
		rw := httptest.NewRecorder()
		handler.GetAll(rw, r)
		if rw.Body.String() != string(tcs.expBody) {
			t.Errorf("%v, Expected %v got %v", tcs.desc, string(tcs.expBody), rw.Body.String())
		}
	}

}

func TestUserHandler_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := services.NewMockUser(ctrl)
	handler := UserHandler{mockService}

	testCases := []struct {
		desc    string
		usr     models.User
		mock    []*gomock.Call
		expCode int
		expBody []byte
	}{
		{
			desc:    "testcase-1",
			usr:     usr,
			expCode: http.StatusOK,
			mock: []*gomock.Call{
				mockService.EXPECT().Insert(usr).Return(usr, nil),
			},
			expBody: []byte(`{"id":1,"name":"himanshu","email":"himanshu8083@gmail.com","phone":"8083860404","age":25}user created`),
		},
	}

	for _, tcs := range testCases {
		jsonUser, _ := json.Marshal(tcs.usr)
		r := httptest.NewRequest("POST", fmt.Sprintf("/user"), bytes.NewBuffer(jsonUser))
		rw := httptest.NewRecorder()
		handler.Insert(rw, r)
		if rw.Body.String() != string(tcs.expBody) {
			t.Errorf("%v, Expected %v got %v", tcs.desc, string(tcs.expBody), rw.Body.String())
		}
	}
}

func TestUserHandler_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := services.NewMockUser(ctrl)
	handler := UserHandler{mockService}

	testCases := []struct {
		desc    string
		mock    []*gomock.Call
		expCode int
		expBody []byte
	}{
		{
			desc:    "testcase-1",
			expCode: http.StatusOK,
			mock: []*gomock.Call{
				mockService.EXPECT().Update(1, "rahul").Return(nil),
			},
			expBody: []byte(`{"id":1,"name":"rahul","email":"himanshu8083@gmail.com","phone":"8083860404","age":25}`),
		},
	}

	for _, tcs := range testCases {
		jsonUser, _ := json.Marshal(models.User{
			Id:    1,
			Name:  "rahul",
			Email: "himanshu8083@gmail.com",
			Phone: "8083860404",
			Age:   25,
		})
		r := httptest.NewRequest("PUT", fmt.Sprintf("/user"), bytes.NewBuffer(jsonUser))
		rw := httptest.NewRecorder()
		handler.Update(rw, r)
		if rw.Body.String() != string(tcs.expBody) {
			t.Errorf("%v, Expected %v got %v", tcs.desc, string(tcs.expBody), rw.Body.String())
		}
	}
}

func TestUserHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := services.NewMockUser(ctrl)
	handler := UserHandler{mockService}

	testCases := []struct {
		desc    string
		id      string
		mock    []*gomock.Call
		expCode int
		expBody []byte
	}{
		{
			desc:    "testcase-1",
			id:      "1",
			expCode: http.StatusOK,
			mock: []*gomock.Call{
				mockService.EXPECT().Delete(1).Return(usr, nil),
			},
			expBody: []byte(`{"id":1,"name":"himanshu","email":"himanshu8083@gmail.com","phone":"8083860404","age":25}`),
		},
		//{
		//	desc:    "testcase-2",
		//	id:      "1a",
		//	expCode: http.StatusBadRequest,
		//	expBody: []byte("invalid id"),
		//},
		{
			desc:    "testcase-3",
			id:      "100",
			expCode: http.StatusInternalServerError,
			mock: []*gomock.Call{
				mockService.EXPECT().Delete(100).Return(models.User{}, errors.New("id not found")),
			},
			expBody: []byte("could not Delete user"),
		},
	}

	for _, tcs := range testCases {
		r := httptest.NewRequest("GET", fmt.Sprintf("/user/%s", tcs.id), nil)
		rw := httptest.NewRecorder()
		r = mux.SetURLVars(r, map[string]string{
			"id": tcs.id,
		})
		handler.Delete(rw, r)
		if rw.Body.String() != string(tcs.expBody) {
			t.Errorf("%v, Expected %v got %v", tcs.desc, string(tcs.expBody), rw.Body.String())
		}
	}
}
