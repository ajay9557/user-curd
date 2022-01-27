package users

import (
	"database/sql"
	"errors"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
	"user-crud/models"
	"user-crud/stores"
)

var usr = models.User{
	Id:    1,
	Name:  "himanshu",
	Email: "himanshu8083@gmail.com",
	Phone: "8083860404",
	Age:   25,
}

func TestUserServ_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dbhandler := stores.NewMockUser(ctrl)
	serv := New(dbhandler)

	testcases := []struct {
		desc   string
		inp    models.User
		exp    models.User
		expErr error
		mock   []*gomock.Call
	}{
		{
			"testcase-1",
			usr,
			usr,
			nil,
			[]*gomock.Call{
				dbhandler.EXPECT().Insert(usr).Return(usr, nil),
			},
		},
	}
	for _, tcs := range testcases {
		out, err := serv.Insert(tcs.inp)
		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}
func TestUserServ_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dbhandler := stores.NewMockUser(ctrl)
	serv := New(dbhandler)

	testcases := []struct {
		desc    string
		inpId   int
		inpName string
		expErr  error
		mock    []*gomock.Call
	}{
		{
			"testcase-1",
			1,
			"rahul",
			nil,
			[]*gomock.Call{
				dbhandler.EXPECT().Update(1, "rahul").Return(nil),
			},
		},
	}
	for _, tcs := range testcases {
		err := serv.Update(tcs.inpId, tcs.inpName)
		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
	}
}
func TestUserServ_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dbhandler := stores.NewMockUser(ctrl)
	serv := New(dbhandler)

	testcases := []struct {
		desc   string
		inp    int
		exp    models.User
		expErr error
		mock   []*gomock.Call
	}{
		{
			"testcase-1",
			1,
			usr,
			nil,
			[]*gomock.Call{
				dbhandler.EXPECT().GetById(1).Return(usr, nil),
				dbhandler.EXPECT().Delete(1).Return(nil),
			},
		},
		{
			"testcase-2",
			1,
			models.User{},
			errors.New("id does not exists"),
			[]*gomock.Call{
				dbhandler.EXPECT().GetById(1).Return(models.User{}, errors.New("id does not exists")),
			},
		},
		{
			"testcase-3",
			1,
			models.User{},
			errors.New("unable to delete user"),
			[]*gomock.Call{
				dbhandler.EXPECT().GetById(1).Return(usr, nil),
				dbhandler.EXPECT().Delete(1).Return(errors.New("unable to delete user")),
			},
		},
	}
	for _, tcs := range testcases {
		out, err := serv.Delete(tcs.inp)
		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}
func TestUserServ_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dbhandler := stores.NewMockUser(ctrl)
	serv := New(dbhandler)

	testcases := []struct {
		desc   string
		inp    int
		exp    models.User
		expErr error
		mock   []*gomock.Call
	}{
		{
			"testcase-1",
			1,
			usr,
			nil,
			[]*gomock.Call{
				dbhandler.EXPECT().GetById(1).Return(usr, nil),
			},
		},
		{
			"testcase-2",
			1,
			models.User{},
			errors.New("id does not exists"),
			[]*gomock.Call{
				dbhandler.EXPECT().GetById(1).Return(models.User{}, errors.New("id does not exists")),
			},
		},
	}
	for _, tcs := range testcases {
		out, err := serv.GetById(tcs.inp)
		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}
func TestUserServ_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dbhandler := stores.NewMockUser(ctrl)
	serv := New(dbhandler)

	testcases := []struct {
		desc   string
		exp    []models.User
		expErr error
		mock   []*gomock.Call
	}{
		{
			"testcase-1",
			[]models.User{usr},
			nil,
			[]*gomock.Call{
				dbhandler.EXPECT().GetAll().Return([]models.User{usr}, nil),
			},
		},
		{
			"testcase-2",
			[]models.User{},
			sql.ErrNoRows,
			[]*gomock.Call{
				dbhandler.EXPECT().GetAll().Return([]models.User{}, sql.ErrNoRows),
			},
		},
	}
	for _, tcs := range testcases {
		out, err := serv.GetAll()
		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}
