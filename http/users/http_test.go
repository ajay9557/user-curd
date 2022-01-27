package users

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-curd/models"
	"user-curd/services"

	"github.com/golang/mock/gomock"
)

func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := services.NewMockServices(ctrl)
	mock := New(mockService)

	testCases := []struct {
		desc      string
		id        string
		mock      []*gomock.Call
		expecCode int
		expecErr  error
		expecBody []byte
	}{
		{
			desc:     "Success case ",
			id:       "2",
			expecErr: nil,
			mock: []*gomock.Call{
				mockService.EXPECT().FetchUserDetailsById(2).Return(models.User{
					Id:    2,
					Name:  "gopi",
					Email: "gopi@gmail.com",
					Phone: "1234567899",
					Age:   23,
				}, nil),
			},
			expecCode: http.StatusOK,
			expecBody: []byte(`{"Id":2,"Name":"gopi","Email":"gopi@gmail.com","Phone":"1234567899","Age":23}`),
		},
		{
			desc:      "Failure case-1",
			id:        "1a",
			expecCode: http.StatusBadRequest,
			expecErr:  errors.New("invalid parameter id"),
			expecBody: []byte("invalid parameter id"),
		},
		{
			desc:     "Failure case - 2",
			id:       "2",
			expecErr: errors.New("internal error"),
			mock: []*gomock.Call{
				mockService.EXPECT().FetchUserDetailsById(2).Return(models.User{}, errors.New("internal error")),
			},
			expecCode: http.StatusInternalServerError,
			expecBody: []byte("internal error"),
		},
	}

	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/user?id="+v.id, nil)
			rw := httptest.NewRecorder()
			mock.GetUserById(rw, r)
			if rw.Code != v.expecCode {
				t.Errorf("Expected %v Obtained %v", v.expecCode, rw.Code)
			}
			if rw.Body.String() != string(v.expecBody) {
				t.Errorf("Expected %v Obtained %v", string(v.expecBody), rw.Body.String())
			}
		})
	}
}

func TestPostUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := services.NewMockServices(ctrl)
	mock := New(mockService)

	testCases := []struct {
		desc      string
		user      []byte
		mock      []*gomock.Call
		expecCode int
		expecErr  error
		expecRes  []byte
	}{
		{
			desc: "Success case",
			user: []byte(`{
				"Id":    1,
				"Name":  "gopi",
				"Email": "gopi@gmail.com",
				"Phone": "1234567899",
				"Age":   23
			}`),
			mock: []*gomock.Call{
				mockService.EXPECT().InsertUserDetails(models.User{
					Id:    1,
					Name:  "gopi",
					Email: "gopi@gmail.com",
					Phone: "1234567899",
					Age:   23,
				}).Return(nil).MaxTimes(5),
			},
			expecErr:  nil,
			expecCode: http.StatusOK,
			expecRes:  []byte("User created"),
		},
		{
			desc: "Failure case -1",
			user: []byte(`{
				"Id":    1,
				"Name":  "gopi1",
				"Email": "gopi@gmail.com",
				"Phone": "1234567899",
				"Age":   23,
			}`),
			expecCode: http.StatusBadRequest,
			expecErr:  errors.New("invalid body"),
			expecRes:  []byte("invalid body"),
		},
	}
	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {

			r := httptest.NewRequest("POST", "/insert", bytes.NewReader(v.user))
			rw := httptest.NewRecorder()
			mock.PostUser(rw, r)
			if rw.Code != v.expecCode {
				t.Errorf("Expected %v Obtained %v", v.expecCode, rw.Code)
			}
			if rw.Body.String() != string(v.expecRes) {
				t.Errorf("Expected %v Obtained %v", string(v.expecRes), rw.Body.String())
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := services.NewMockServices(ctrl)
	mock := New(mockService)
	testCases := []struct {
		desc      string
		mock      []*gomock.Call
		expecErr  error
		expecBody []byte
	}{
		{
			desc:     "Success case ",
			expecErr: nil,
			mock: []*gomock.Call{
				mockService.EXPECT().FetchAllUserDetails().Return([]models.User{{
					Id:    2,
					Name:  "gopi",
					Email: "gopi@gmail.com",
					Phone: "1234567899",
					Age:   23,
				},
				}, nil),
			},
			expecBody: []byte(`[{"Id":2,"Name":"gopi","Email":"gopi@gmail.com","Phone":"1234567899","Age":23}]`),
		},
		{
			desc:     "Failure case-1 ",
			expecErr: errors.New("error generated"),
			mock: []*gomock.Call{
				mockService.EXPECT().FetchAllUserDetails().Return([]models.User{{
					Id:    2,
					Name:  "gopi",
					Email: "gopi@gmail.com",
					Phone: "1234567899",
					Age:   23,
				},
				}, errors.New("error generated")),
			},
			expecBody: []byte("error generated"),
		},
	}
	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/users", nil)
			rw := httptest.NewRecorder()
			mock.GetUsers(rw, r)
			if rw.Body.String() != string(v.expecBody) {
				t.Errorf("Expected %v Obtained %v", string(v.expecBody), rw.Body.String())
			}
		})
	}

}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := services.NewMockServices(ctrl)
	mock := New(mockService)
	testCases := []struct {
		desc     string
		id       string
		expecErr error
		mock     []*gomock.Call
		expecRes []byte
	}{
		{
			desc:     "Success case",
			id:       "1",
			expecErr: nil,
			mock: []*gomock.Call{
				mockService.EXPECT().DeleteUserDetailsById(1).Return(nil),
			},
			expecRes: []byte("User deleted successfully"),
		},
		{
			desc:     "Failure case - 1",
			id:       "0",
			expecErr: errors.New("Id shouldn't be zero"),
			expecRes: []byte("Id shouldn't be zero"),
		},
		{
			desc:     "Failure case - 2",
			id:       "1",
			expecErr: errors.New("error generated"),
			mock: []*gomock.Call{
				mockService.EXPECT().DeleteUserDetailsById(1).Return(errors.New("error generated")),
			},
			expecRes: []byte("error generated"),
		},
	}
	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			r := httptest.NewRequest("DELETE", "/delete?id="+v.id, nil)
			rw := httptest.NewRecorder()
			mock.DeleteUser(rw, r)
			fmt.Println(rw.Body.String())
			if rw.Body.String() != string(v.expecRes) {
				t.Errorf("Expected %v Obtained %v", string(v.expecRes), rw.Body.String())
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := services.NewMockServices(ctrl)
	mock := New(mockService)

	testCases := []struct {
		desc     string
		user     []byte
		mock     []*gomock.Call
		expecErr error
		expecRes []byte
	}{
		{
			desc: "Success case",
			user: []byte(`{
				"Id":    1,
				"Name":  "gopi",
				"Email": "gopi@gmail.com",
				"Phone": "1234567899",
				"Age":   23
			}`),
			mock: []*gomock.Call{
				mockService.EXPECT().UpdateUserDetails(models.User{
					Id:    1,
					Name:  "gopi",
					Email: "gopi@gmail.com",
					Phone: "1234567899",
					Age:   23,
				}).Return(nil).MaxTimes(5),
			},
			expecErr: nil,
			expecRes: []byte("User updated"),
		},
		{
			desc: "Failure case -1",
			user: []byte(`{
				"Id":    1,
				"Name":  "gopi1",
				"Email": "gopi@gmail.com",
				"Phone": "1234567899",
				"Age":   23,
			}`),
			expecErr: errors.New("invalid body"),
			expecRes: []byte("invalid body"),
		},
	}
	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			r := httptest.NewRequest("PUT", "/update", bytes.NewReader(v.user))
			rw := httptest.NewRecorder()
			mock.UpdateUser(rw, r)
			fmt.Println(rw.Body.String())
			if rw.Body.String() != string(v.expecRes) {
				t.Errorf("Expected %v Obtained %v", string(v.expecRes), rw.Body.String())
			}
		})
	}
}
