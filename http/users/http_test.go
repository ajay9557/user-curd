package users

import (
	"user-curd/models"
	"user-curd/services"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var user = []models.User{
	{
		Id:    1,
		Name:  "Sudheer",
		Email: "sudheer01@gmail.com",
		Phone: "9908577405",
		Age:   22,
	},
	{
		Id:    2,
		Name:  "Puppala",
		Email: "puppala@gmail.com",
		Phone: "9908577405",
		Age:   24,
	},
}

func TestHandler_Search(t *testing.T) {
	tcs := []struct {
		Id       string
		expected models.User
		stCode   int
		err      error
	}{
		{"1", user[0], http.StatusOK, nil},
		{"2", user[1], http.StatusOK, nil},
	}

	ctrl := gomock.NewController(t)

	serv := services.NewMockService(ctrl)
	//mock := New(serv)
	h := Handler{serv}
	for _, tc := range tcs {
		link := "/users/%s"
		r := httptest.NewRequest("GET", fmt.Sprintf(link, tc.Id), nil)
		w := httptest.NewRecorder()
		//router := mux.NewRouter()
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
		expected models.User
		stCode   int
		err      error
	}{
		{"1", user[0], http.StatusOK, nil},
		{"2", user[1], http.StatusOK, nil},
	}

	ctrl := gomock.NewController(t)

	serv := services.NewMockService(ctrl)
	//mock := New(serv)
	h := Handler{serv}
	for _, tc := range tcs {
		link := "/users/%s"
		r := httptest.NewRequest("DELETE", fmt.Sprintf(link, tc.Id), nil)
		w := httptest.NewRecorder()
		//router := mux.NewRouter()
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
		expected []models.User
		stCode   int
		err      error
	}{
		{"1", []models.User{
			{Id: 2,
				Name:  "Zopsmart",
				Email: "sudheerpuppala@gmail.com",
				Phone: "+919908577405",
				Age:   12,
			}}, http.StatusOK, nil},
	}

	ctrl := gomock.NewController(t)

	serv := services.NewMockService(ctrl)
	//mock := New(serv)
	h := Handler{serv}
	for _, tc := range tcs {
		link := "/users"
		r := httptest.NewRequest("GET", fmt.Sprintf(link), nil)
		w := httptest.NewRecorder()
		//router := mux.NewRouter()
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
	serv := services.NewMockService(ctrl)
	h := Handler{serv}
	tests := []struct {
		desc     string
		usr      models.User
		output   models.User
		stcode   int
		validity bool
		err      error
	}{
		{
			desc:     "Success",
			usr:      models.User{Id: 1, Name: "Sudheer", Email: "Sudheer@gmail.com", Phone: "9908577405", Age: 22},
			output:   models.User{Id: 1, Name: "Sudheer", Email: "Sudheer@gmail.com", Phone: "9908577405", Age: 22},
			stcode:   http.StatusOK,
			validity: true,
			err:      nil,
		},
		{
			desc:     "Failure",
			usr:      models.User{Id: 3, Name: "Sudheer", Email: "Sudheergmail.com", Phone: "9908577405", Age: 22},
			output:   models.User{},
			stcode:   http.StatusBadRequest,
			validity: false,
			err:      nil,
		},
	}
	for _, tc := range tests {

		l, _ := json.Marshal(tc.usr)
		m := bytes.NewBuffer(l)
		link := "/users"
		serv.EXPECT().IsEmailValid(tc.usr.Email).Return(tc.validity)
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
	serv := services.NewMockService(ctrl)
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
				"Name":  "sudheer",
				"Email": "sudheer@gmail.com",
				"Phone": "9908577405",
				"Age":   23
			}`),
			mock: []*gomock.Call{
				serv.EXPECT().IsEmailValid("sudheer@gmail.com").Return(true).MaxTimes(5),
				serv.EXPECT().UpdateByUserId(models.User{
					Id:    1,
					Name:  "sudheer",
					Email: "sudheer@gmail.com",
					Phone: "9908577405",
					Age:   23,
				}).Return(nil).MaxTimes(5),
			},
			expecErr: nil,
			expecRes: []byte("user updated"),
		},
		{
			desc: "Failure case -1",
			user: []byte(`{
				"Id":    1,
				"Name":  "sudheer",
				"Email": "sudheer@gmail.com",
				"Phone": "9908577405",
				"Age":   23,
			}`),
			expecErr: errors.New("invalid body"),
			expecRes: []byte("invalid body"),
		},
		{
			desc: "Failure case -2",
			user: []byte(`{
				"Id":    0,
				"Name":  "sudheer",
				"Email": "sudheer@gmail.com",
				"Phone": "9908577405",
				"Age":   23
			}`),
			expecErr: errors.New("Id shouldn't be zero"),
			expecRes: []byte("Id shouldn't be zero"),
		},
		{
			desc: "Failure case -3",
			user: []byte(`{
				"Id":    1,
				"Name":  "sudheer",
				"Email": "sudheerpup@gmail.com",
				"Phone": "9908577405",
				"Age":   23
			}`),
			mock: []*gomock.Call{
				serv.EXPECT().IsEmailValid("sudheerpup@gmail.com").Return(false),
			},
			expecErr: errors.New("Email already there,create new email"),
			expecRes: []byte("Email already there,create new email"),
		},
	}
	for _, v := range testCases {
		t.Run(v.desc, func(t *testing.T) {
			r := httptest.NewRequest("PUT", "/update", bytes.NewReader(v.user))
			rw := httptest.NewRecorder()
			h.UpdateUserDetails(rw, r)
			fmt.Println(rw.Body.String())
			if rw.Body.String() != string(v.expecRes) {
				t.Errorf("Expected %v Obtained %v", string(v.expecRes), rw.Body.String())
			}
		})
	}
}
