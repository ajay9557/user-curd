package Users

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	service "user-curd/Service"
	"user-curd/model"

	gomock "github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var user = []model.User{
	{
		Id:    1,
		Name:  "Nayani",
		Email: "Sai@gmail.com",
		Phone: "9908577405",
		Age:   "22",
	},
	{
		Id:    2,
		Name:  "Sai",
		Email: "Nayani@gmail.com",
		Phone: "9908577405",
		Age:   "24",
	},
}

func TestHandler_Search(t *testing.T) {
	tcs := []struct {
		Id       string
		expected model.User
		stCode   int
		err      error
	}{
		{"1", user[0], http.StatusOK, nil},
		{"2", user[1], http.StatusOK, nil},
	}

	ctrl := gomock.NewController(t)

	serv := service.NewMockService(ctrl)

	h := Handler{serv}
	for _, tc := range tcs {
		link := "/users/%s"
		r := httptest.NewRequest("GET", fmt.Sprintf(link, tc.Id), nil)
		w := httptest.NewRecorder()

		r = mux.SetURLVars(r, map[string]string{
			"id": tc.Id,
		})
		id, err := strconv.Atoi(tc.Id)
		if err == nil {
			serv.EXPECT().SearchByUserId(id).Return(tc.expected, tc.err)
		}
		h.Search(w, r)
		if w.Code != tc.stCode {
			t.Fatalf("SearchId() = %v , want %v", w.Code, tc.stCode)
		}
	}
}

func TestHandler_Delete(t *testing.T) {
	tcs := []struct {
		Id       string
		expected model.User
		stCode   int
		err      error
	}{
		{"1", user[0], http.StatusOK, nil},
		{"2", user[1], http.StatusOK, nil},
	}

	ctrl := gomock.NewController(t)

	serv := service.NewMockService(ctrl)

	h := Handler{serv}
	for _, tc := range tcs {
		link := "/users/%s"
		r := httptest.NewRequest("DELETE", fmt.Sprintf(link, tc.Id), nil)
		w := httptest.NewRecorder()

		r = mux.SetURLVars(r, map[string]string{
			"id": tc.Id,
		})
		id, err := strconv.Atoi(tc.Id)
		if err == nil {
			serv.EXPECT().DeleteByUserId(id).Return(tc.err)
		}
		h.DeleteId(w, r)
		if w.Code != tc.stCode {
			t.Fatalf("DeleteId() = %v , want %v", w.Code, tc.stCode)
		}
	}
}

func TestHandler_SearchingAll(t *testing.T) {
	tcs := []struct {
		Id       string
		expected []model.User
		stCode   int
		err      error
	}{
		{"1", []model.User{
			{Id: 2,
				Name:  "Nayani",
				Email: "sai@gmail.com",
				Phone: "9908577405",
				Age:   "12",
			}}, http.StatusOK, nil},
	}

	ctrl := gomock.NewController(t)

	serv := service.NewMockService(ctrl)

	h := Handler{serv}
	for _, tc := range tcs {
		link := "/users"
		r := httptest.NewRequest("GET", fmt.Sprintf(link), nil)
		w := httptest.NewRecorder()

		r = mux.SetURLVars(r, map[string]string{
			"id": tc.Id,
		})
		_, err := strconv.Atoi(tc.Id)
		if err == nil {
			serv.EXPECT().SearchAll().Return(tc.expected, tc.err)
		}
		h.GetAll(w, r)
		if w.Code != tc.stCode {
			t.Fatalf("SearchAll() = %v , want %v", w.Code, tc.stCode)
		}
	}
}
func Test_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serv := service.NewMockService(ctrl)
	h := Handler{serv}
	tests := []struct {
		desc     string
		usr      model.User
		output   model.User
		stcode   int
		validity bool
		err      error
	}{
		{
			desc:     "Success",
			usr:      model.User{Id: 1, Name: "Sai", Email: "Nayani@gmail.com", Phone: "9908577405", Age: "22"},
			output:   model.User{Id: 1, Name: "Sai", Email: "Nayani@gmail.com", Phone: "9908577405", Age: "22"},
			stcode:   http.StatusOK,
			validity: true,
			err:      nil,
		},
	}
	for _, tc := range tests {

		l, _ := json.Marshal(tc.usr)
		m := bytes.NewBuffer(l)
		link := "/users"
		r := httptest.NewRequest("POST", link, m)
		w := httptest.NewRecorder()

		serv.EXPECT().InsertUserDetails(tc.usr).Return(tc.output, tc.err).MinTimes(0)
		h.Create(w, r)
		if w.Code != tc.stcode {
			t.Fatalf("InsertUser()=%v,want %v", w.Code, tc.stcode)
		}
	}

}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serv := service.NewMockService(ctrl)
	h := Handler{serv}
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
				"Name":  "Sai",
				"Email": "Nayani@gmail.com",
				"Phone": "9908577405",
				"Age":   "23"
			}`),
			mock: []*gomock.Call{
				serv.EXPECT().IsEmailValid("Nayani@gmail.com").Return(true).MaxTimes(5),
				serv.EXPECT().UpdateByUserId(model.User{
					Id:    1,
					Name:  "Sai",
					Email: "Nayani@gmail.com",
					Phone: "9908577405",
					Age:   "23",
				}).Return(nil).MaxTimes(5),
			},
			expecErr: nil,
			expecRes: []byte("User updated"),
		},
		{
			desc: "Failure case -1",
			user: []byte(`{
				"Id":    1,
				"Name":  "Sai",
				"Email": "Nayani@gmail.com",
				"Phone": "9908577405",
				"Age":   "23",
			}`),
			expecErr: errors.New("invalid body"),
			expecRes: []byte("invalid body"),
		},
	}
	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			r := httptest.NewRequest("PUT", "/update", bytes.NewReader(v.user))
			rw := httptest.NewRecorder()
			h.UpdateUser(rw, r)
			fmt.Println(rw.Body.String())
			if rw.Body.String() != string(v.expecRes) {
				t.Errorf("Expected %v Obtained %v", string(v.expecRes), rw.Body.String())
			}
		})
	}
}
