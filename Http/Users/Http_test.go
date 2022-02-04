package Users

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-curd/model"
	"user-curd/services"

	gomock "github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var user = model.User{
	Id:    1,
	Name:  "sai",
	Email: "Nayani@gmail.com",
	Phone: "6303880131",
	Age:   "25",
}

func TestHandler_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := services.NewMockUser(ctrl)
	handler := Handler{mockService}

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
				mockService.EXPECT().SearchByUserId(1).Return(&user, nil),
			},
			expBody: []byte(`{"id":1,"name":"sai","email":"Nayani@gmail.com","phone":"6303880131","age":"25"}`),
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
				mockService.EXPECT().SearchByUserId(100).Return(&model.User{}, errors.New("id not found")),
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
		handler.Search(rw, r)
		if rw.Body.String() != string(tcs.expBody) {
			t.Errorf("%v, Expected %v got %v", tcs.desc, string(tcs.expBody), rw.Body.String())
		}
	}
}

func TestHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := services.NewMockUser(ctrl)
	handler := Handler{mockService}

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
				mockService.EXPECT().DeleteByUserId(1).Return(nil),
			},
			expBody: []byte(`{"id":1,"name":"ai","email":"Nayani@gmail.com","phone":"6303880131","age":"25"}`),
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
				mockService.EXPECT().DeleteByUserId(100).Return(errors.New("id not found")),
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
		handler.DeleteId(rw, r)
		if rw.Code != tcs.expCode {
			t.Errorf("%v, Expected %v got %v", tcs.desc, tcs.expCode, rw.Code)
		}
	}
}

func TestHandler_SearchingAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := services.NewMockUser(ctrl)
	handler := Handler{mockService}

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
				mockService.EXPECT().SearchAll().Return([]*model.User{&user}, nil),
			},
			expBody: []byte(`[{"id":1,"name":"sai","email":"Nayani@gmail.com","phone":"6303880131","age":"25"}]`),
		},
	}

	for _, tcs := range testCases {
		r := httptest.NewRequest("GET", fmt.Sprintf("/user"), nil)
		rw := httptest.NewRecorder()
		handler.GetAll(rw, r)
		if rw.Code != tcs.expCode {
			t.Errorf("%v, Expected %v got %v", tcs.desc, tcs.expCode, rw.Code)
		}
	}

}

func Test_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := services.NewMockUser(ctrl)
	handler := Handler{mockService}

	testCases := []struct {
		desc    string
		usr     model.User
		mock    []*gomock.Call
		expCode int
		expBody []byte
	}{
		{
			desc:    "testcase-1",
			usr:     user,
			expCode: http.StatusOK,
			mock: []*gomock.Call{
				mockService.EXPECT().InsertUserDetails(&user).Return(&user, nil),
			},
			expBody: []byte(`{"id":1,"name":"sai","email":"Nayani@gmail.com","phone":"6303880131","age":"25"}user created`),
		},
	}

	for _, tcs := range testCases {
		jsonUser, _ := json.Marshal(tcs.usr)
		r := httptest.NewRequest("POST", fmt.Sprintf("/user"), bytes.NewBuffer(jsonUser))
		rw := httptest.NewRecorder()
		handler.Create(rw, r)
		if rw.Code != tcs.expCode {
			t.Errorf("%v, Expected %v got %v", tcs.desc, tcs.expCode, rw.Code)
		}
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockService := services.NewMockUser(ctrl)
	handler := Handler{mockService}

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
				mockService.EXPECT().UpdateByUserId(&user).Return(&user, nil),
			},
			expBody: []byte(`{"id":1,"name":"sai","email":"Nayani@gmail.com","phone":"6303880131","age":"25"}`),
		},
	}

	for _, tcs := range testCases {
		jsonUser, _ := json.Marshal(model.User{
			Id:    1,
			Name:  "sai",
			Email: "Nayani@gmail.com",
			Phone: "6303880131",
			Age:   "25",
		})
		r := httptest.NewRequest("PUT", fmt.Sprintf("/user"), bytes.NewBuffer(jsonUser))
		rw := httptest.NewRecorder()
		handler.UpdateUser(rw, r)
		if rw.Code != tcs.expCode {
			t.Errorf("%v, Expected %v got %v", tcs.desc, tcs.expCode, rw.Code)
		}
	}
}
