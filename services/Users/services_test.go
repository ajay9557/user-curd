package Users

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"user-curd/model"
	"user-curd/stores"

	gomock "github.com/golang/mock/gomock"
)

var user = model.User{
	Id:    1,
	Name:  "sai",
	Email: "Nayani@gmail.com",
	Phone: "6303880131",
	Age:   "25",
}

func TestUserServ_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dbhandler := stores.NewMockUser(ctrl)
	serv := New(dbhandler)

	testcases := []struct {
		desc   string
		inp    int
		expErr error
		mock   []*gomock.Call
	}{
		{
			"testcase-1",
			1,
			nil,
			[]*gomock.Call{
				dbhandler.EXPECT().GetById(1).Return(&user, nil),
				dbhandler.EXPECT().DeleteById(1).Return(nil),
			},
		},
		{
			"testcase-2",
			1,
			errors.New("id does not exists"),
			[]*gomock.Call{
				dbhandler.EXPECT().GetById(1).Return(&model.User{}, errors.New("id does not exists")),
			},
		},
		/* {
			"testcase-3",
			1,
			errors.New("unable to delete user"),
			[]*gomock.Call{
				dbhandler.EXPECT().GetById(1).Return(&user, nil),
				dbhandler.EXPECT().DeleteById(1).Return(errors.New("unable to delete user")),
			},
		}, */
	}
	for _, tcs := range testcases {
		err := serv.DeleteByUserId(tcs.inp)
		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dbhandler := stores.NewMockUser(ctrl)
	serv := New(dbhandler)

	testcases := []struct {
		desc   string
		inp    *model.User
		expErr error
		mock   []*gomock.Call
	}{
		{
			"testcase-1",
			&user,
			nil,
			[]*gomock.Call{
				dbhandler.EXPECT().UpdateById(&user).Return(&user, nil),
			},
		},
	}
	for _, tcs := range testcases {
		_, err := serv.UpdateByUserId(tcs.inp)
		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
	}
}

func TestSearchById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dbhandler := stores.NewMockUser(ctrl)
	serv := New(dbhandler)

	testcases := []struct {
		desc   string
		inp    int
		exp    *model.User
		expErr error
		mock   []*gomock.Call
	}{
		{
			"testcase-1",
			1,
			&user,
			nil,
			[]*gomock.Call{
				dbhandler.EXPECT().GetById(1).Return(&user, nil),
			},
		},
		{
			"testcase-2",
			1,
			&model.User{},
			errors.New("id does not exists"),
			[]*gomock.Call{
				dbhandler.EXPECT().GetById(1).Return(&model.User{}, errors.New("id does not exists")),
			},
		},
	}
	for _, tcs := range testcases {
		out, err := serv.SearchByUserId(tcs.inp)
		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}

func TestSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dbhandler := stores.NewMockUser(ctrl)
	serv := New(dbhandler)

	testcases := []struct {
		desc   string
		exp    []*model.User
		expErr error
		mock   []*gomock.Call
	}{
		{
			"testcase-1",
			[]*model.User{&user},
			nil,
			[]*gomock.Call{
				dbhandler.EXPECT().GetAll().Return([]*model.User{&user}, nil),
			},
		},
		{
			"testcase-2",
			[]*model.User{},
			sql.ErrNoRows,
			[]*gomock.Call{
				dbhandler.EXPECT().GetAll().Return([]*model.User{}, sql.ErrNoRows),
			},
		},
	}
	for _, tcs := range testcases {
		out, err := serv.SearchAll()
		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}

func TestInsert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dbhandler := stores.NewMockUser(ctrl)
	serv := New(dbhandler)

	testcases := []struct {
		desc   string
		inp    model.User
		exp    *model.User
		expErr error
		mock   []*gomock.Call
	}{
		{
			"testcase-1",
			user,
			&user,
			nil,
			[]*gomock.Call{
				dbhandler.EXPECT().InsertDetails(&user).Return(&user, nil),
			},
		},
	}
	for _, tcs := range testcases {
		out, err := serv.InsertUserDetails(&tcs.inp)
		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.expErr, err)
		}
		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("%v, expected %v, got %v", tcs.desc, tcs.exp, out)
		}
	}
}
