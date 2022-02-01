package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"zopsmart/user-curd/model"
	service "zopsmart/user-curd/service"

	gomock "github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func Test_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serv := service.NewMockUser(ctrl)
	h := Handler{serv}
	tests := []struct {
		usr      model.User
		output   model.User
		stcode   int
		validity bool
		err      error
		er       error
	}{
		{
			usr:      model.User{Name: "Naira", Email: "naira@gmail.com", Phone: "9866895296", Age: 20},
			output:   model.User{Name: "Naira", Email: "naira@gmail.com", Phone: "9866895296", Age: 20},
			stcode:   http.StatusCreated,
			validity: true,
			err:      nil,
			er:       nil,
		},
		{
			usr:      model.User{Name: "Naira", Email: "nairagmail.com", Phone: "9866895296", Age: 20},
			output:   model.User{},
			stcode:   http.StatusBadRequest,
			validity: false,
			err:      fmt.Errorf("Enter Valid email address"),
			er:       nil,
		},
	}
	for _, tes := range tests {

		l, _ := json.Marshal(tes.usr)
		m := bytes.NewBuffer(l)
		link := "/user"
		call := serv.EXPECT().CheckMail(tes.usr.Email).Return(tes.validity, tes.err)

		r := httptest.NewRequest("POST", link, m)
		w := httptest.NewRecorder()

		serv.EXPECT().PostUser(tes.usr.Name, tes.usr.Email, tes.usr.Phone, tes.usr.Age).After(call).Return(tes.output, tes.er).MinTimes(0)
		h.AddUser(w, r)
		if w.Code != tes.stcode {
			t.Fatalf("AddUser()=%v,want %v", w.Code, tes.stcode)
		}

	}

}

func Test_UserWithId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serv := service.NewMockUser(ctrl)
	h := Handler{serv}
	tests := []struct {
		id     int
		output model.User
		stcode int
		err    error
	}{
		{
			id:     1,
			output: model.User{Id: 1, Name: "Naira", Email: "naira@gmail.com", Phone: "9866895296", Age: 20},
			stcode: http.StatusOK,
			err:    nil,
		},
		{
			id:     0,
			output: model.User{},
			stcode: http.StatusBadRequest,
			err:    fmt.Errorf("User not found"),
		},
	}
	for _, tes := range tests {
		link := "/user/%v"
		serv.EXPECT().GetByID(tes.id).Return(tes.output, tes.err)
		r := httptest.NewRequest("GET", fmt.Sprintf(link, tes.id), nil)
		w := httptest.NewRecorder()
		i := strconv.Itoa(tes.id)
		r = mux.SetURLVars(r, map[string]string{
			"id": i,
		})

		h.UserWithId(w, r)
		if w.Code != tes.stcode {
			t.Fatalf("UserWithId() = %v , want %v", w.Code, tes.stcode)
		}

	}

}

func Test_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serv := service.NewMockUser(ctrl)
	h := Handler{serv}
	tests := []struct {
		id     int
		output model.User
		stcode int
		err    error
	}{
		{
			id:     1,
			output: model.User{Id: 1, Name: "Naira", Email: "naira@gmail.com", Phone: "9866895296", Age: 20},
			stcode: http.StatusAccepted,
			err:    nil,
		},
		{
			id:     0,
			stcode: http.StatusBadRequest,
			err:    fmt.Errorf("User not found"),
		},
	}
	for _, tes := range tests {

		link := "/user/delete/%v"
		r := httptest.NewRequest("DELETE", fmt.Sprintf(link, tes.id), nil)
		w := httptest.NewRecorder()
		i := strconv.Itoa(tes.id)
		r = mux.SetURLVars(r, map[string]string{
			"id": i,
		})
		serv.EXPECT().DeleteByID(tes.id).Return(tes.err)
		h.DeleteUser(w, r)
		if w.Code != tes.stcode {
			t.Fatalf("DeleteUser() = %v , want %v", w.Code, tes.stcode)
		}
	}

}

func TestHandler_GetAllUsers(t *testing.T) {
	tcs := []struct {
		expected []model.User
		stCode   int
		err      error
	}{
		{[]model.User{
			{Id: 2,
				Name:  "Zopsmart",
				Email: "sudheerpuppala@gmail.com",
				Phone: "+919908577405",
				Age:   12,
			}}, http.StatusOK, nil},
		{[]model.User{{}}, http.StatusOK, nil},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serv := service.NewMockUser(ctrl)
	h := Handler{serv}
	for _, tc := range tcs {

		link := "/user"
		r := httptest.NewRequest("GET", fmt.Sprintf(link), nil)
		w := httptest.NewRecorder()
		serv.EXPECT().GetUsers().Return(tc.expected, tc.err)

		h.GetAllUsers(w, r)
		if w.Code != tc.stCode {
			t.Fatalf("SearchAll() = %v , want %v", w.Code, tc.stCode)
		}
	}
}

func Test_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serv := service.NewMockUser(ctrl)
	h := Handler{serv}
	tests := []struct {
		id     int
		usrlst []model.User
		usr    model.User
		updusr model.User
		stcode int
		err    error
		er     error
		e      error
	}{
		{
			id:     1,
			usrlst: []model.User{{1, "Mahi", "Mahi@gmail.com", "9866895296", 20}},
			usr:    model.User{Id: 1, Name: "Mahi", Email: "Mahi@gmail.com", Phone: "9866895296", Age: 20},
			updusr: model.User{Id: 1, Name: "Mahir", Email: "Mahi@gmail.com", Phone: "9866895296", Age: 20},
			stcode: http.StatusOK,
			err:    nil,
			er:     nil,
			e:      nil,
		},
		{
			id:     0,
			usrlst: []model.User{{1, "Mahi", "Mahi@gmail.com", "9866895296", 20}},
			usr:    model.User{},
			updusr: model.User{},
			stcode: http.StatusBadRequest,
			err:    fmt.Errorf("User not found"),
			er:     nil,
			e:      nil,
		},
	}
	for _, tes := range tests {
		l, _ := json.Marshal(tes.updusr)
		m := bytes.NewBuffer(l)
		link := "/user/update/%v"
		r := httptest.NewRequest("PATCH", fmt.Sprintf(link, tes.id), m)
		w := httptest.NewRecorder()
		i := strconv.Itoa(tes.id)
		r = mux.SetURLVars(r, map[string]string{
			"id": i,
		})

		call := serv.EXPECT().GetByID(tes.id).Return(tes.usr, tes.err)
		call2 := serv.EXPECT().Update(tes.usr, tes.updusr).Return(tes.updusr, tes.er).After(call).MinTimes(0)
		serv.EXPECT().GetUsers().Return(tes.usrlst, tes.e).After(call2).MinTimes(0)

		h.UpdateUser(w, r)
		if w.Code != tes.stcode {
			t.Fatalf("UpdateUser() = %v , want %v", w.Code, tes.stcode)
		}
	}

}
