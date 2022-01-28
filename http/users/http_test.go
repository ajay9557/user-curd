package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/tejas/user-crud/models"
	"github.com/tejas/user-crud/services"
)

func TestUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := services.NewMockUser(ctrl)
	mock := New(mockService)

	testCases := []struct {
		desc      string
		id        string
		mock      []*gomock.Call
		expecErr  error
		expecBody []byte
	}{
		{
			desc:     "Case 1: Success case",
			id:       "1",
			expecErr: nil,
			mock: []*gomock.Call{
				mockService.EXPECT().GetUserById(1).Return(models.User{
					Id:    1,
					Name:  "tejas",
					Email: "tejas@gmail.com",
					Phone: "9876543210",
					Age:   23,
				}, nil),
			},
			expecBody: []byte(`{"Id":1,"Name":"tejas","Email":"tejas@gmail.com","Phone":"9876543210","Age":23}`),
		},
		{
			desc:      "Case 2: Failure case - 1",
			id:        "1a",
			expecErr:  nil,
			expecBody: []byte("Invalid User Id"),
		},
		{
			desc:     "Case 3: Failure case - 2",
			id:       "100000000",
			expecErr: nil,
			mock: []*gomock.Call{
				mockService.EXPECT().GetUserById(2).Return(models.User{
					Id:    0,
					Name:  "",
					Email: "",
					Phone: "",
					Age:   0,
				}, errors.New("User id not found")),
			},
			expecBody: []byte("User id not found"),
		},
	}

	for _, test := range testCases {
		req := httptest.NewRequest("GET", "/api/users/"+test.id, nil)
		res := httptest.NewRecorder()

		req = mux.SetURLVars(req, map[string]string{
			"id": test.id,
		})

		mock.FindUserById(res, req)

	}
}

func TestGetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := services.NewMockUser(ctrl)
	mock := New(mockService)
	testCases := []struct {
		desc      string
		mock      []*gomock.Call
		expecErr  error
		expecBody []byte
	}{
		{
			desc:     "Case 1: Success case ",
			expecErr: nil,
			mock: []*gomock.Call{
				mockService.EXPECT().GetUsers().Return([]models.User{{
					Id:    1,
					Name:  "tejas",
					Email: "tejas@gmail.com",
					Phone: "9876543210",
					Age:   23,
				},
				}, nil),
			},
			expecBody: []byte(`[{"Id":1,"Name":"tejas","Email":"tejas@gmail.com","Phone":"9876543210","Age":23}]`),
		},
		{
			desc:     "Failure case-1 ",
			expecErr: errors.New("Could not fetch users"),
			mock: []*gomock.Call{
				mockService.EXPECT().GetUsers().Return([]models.User{{
					Id:    2,
					Name:  "tejas",
					Email: "tejas@gmail.com",
					Phone: "9876543210",
					Age:   23,
				},
				}, errors.New("Could not fetch users")),
			},
			expecBody: []byte("Could not fetch users"),
		},
	}

	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/api/users/", nil)
			rw := httptest.NewRecorder()
			mock.FindAllUsers(rw, r)
			if rw.Body.String() != string(v.expecBody) {
				t.Errorf("Expected %v Obtained %v", string(v.expecBody), rw.Body.String())
			}
		})
	}

}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := services.NewMockUser(ctrl)
	mock := New(mockService)

	testCases := []struct {
		desc     string
		user     models.User
		mock     []*gomock.Call
		expecErr error
		expecRes []byte
	}{
		{
			desc: "Case 1: Success case",
			user: models.User{
				Id:    1,
				Name:  "tejas",
				Email: "tejas@gmail.com",
				Phone: "9876543210",
				Age:   23,
			},
			mock: []*gomock.Call{
				mockService.EXPECT().UpdateUser(0, models.User{
					Id:    1,
					Name:  "tejas",
					Email: "tejas@gmail.com",
					Phone: "9876543210",
					Age:   23,
				}).Return(nil),
			},
			expecErr: nil,
			// expecRes: []byte("User updated"),
		},
		{
			desc: "Case 1: Failure case -1",
			user: models.User{
				Id:    -2,
				Name:  "tejas",
				Email: "tejas@gmail.com",
				Phone: "9876543210",
				Age:   23,
			},
			expecErr: errors.New("Invalid Id"),
			expecRes: []byte("Invalid Id"),
		},
	}
	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			b, _ := json.Marshal(v.user)
			r := httptest.NewRequest("PUT", "/api/users/"+v.id, bytes.NewBuffer(b))
			rw := httptest.NewRecorder()
			mock.UpdateUserById(rw, r)
			fmt.Println(rw.Body.String())
			if rw.Body.String() != string(v.expecRes) {
				t.Errorf("Expected %v Obtained %v", string(v.expecRes), rw.Body.String())
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := services.NewMockUser(ctrl)
	mock := New(mockService)
	testCases := []struct {
		desc     string
		id       string
		expecErr error
		mock     []*gomock.Call
		expecRes []byte
	}{
		{
			desc:     "Case 1: Success case",
			id:       "1",
			expecErr: nil,
			mock: []*gomock.Call{
				mockService.EXPECT().DeleteUser(1).Return(nil),
			},
			expecRes: []byte("user deleted successfully"),
		},
		{
			desc:     "Case 2: Failure case - 1",
			id:       "-2",
			expecErr: errors.New("Invalid Id"),
			expecRes: []byte("Invalid Id"),
		},
		{
			desc:     "Case 3: Failure case - 2",
			id:       "0",
			expecErr: errors.New("error while deleting user"),
			mock: []*gomock.Call{
				mockService.EXPECT().DeleteUser(0).Return(errors.New("error while deleting user")),
			},
			expecRes: []byte("error while deleting user"),
		},
	}
	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			r := httptest.NewRequest("DELETE", "/api/users/"+v.id, nil)
			rw := httptest.NewRecorder()
			mock.DeleteUserById(rw, r)
			fmt.Println(rw.Body.String())
			if rw.Body.String() != string(v.expecRes) {
				t.Errorf("Expected %v Obtained %v", string(v.expecRes), rw.Body.String())
			}
		})
	}
}
