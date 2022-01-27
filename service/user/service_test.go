package user

import (
	"fmt"
	"reflect"
	"testing"
	"zopsmart/user-curd/model"
	store "zopsmart/user-curd/store"

	gomock "github.com/golang/mock/gomock"
)

func Test_CheckMail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := store.NewMockStore(ctrl)
	mock := New(mockSvc)
	tests := []struct {
		desc   string
		usr    []model.User
		input  string
		output bool
		err    error
	}{
		{desc: "Case-1", usr: []model.User{{1, "Mahi", "Mahi@gmail.com", "9866895296", 20}}, input: "Naira@gmail.com", output: true},
		{desc: "Case-2", usr: []model.User{{1, "Mahi", "Mahi@gmail.com", "9866895296", 20}}, input: "Mahi@gmail.com", output: false},
		{desc: "Case-2", usr: []model.User{{1, "Mahi", "Mahi@gmail.com", "9866895296", 20}}, input: "Mahigmail.com", output: false},
	}
	for _, tes := range tests {
		mockSvc.EXPECT().GetAllUsers().Return(tes.usr, tes.err).MinTimes(0)
		if tes.err == nil {
			fmt.Println(tes.input)
			op, _ := mock.CheckMail(tes.input)

			if op != tes.output {
				t.Errorf("expect %v got %v", tes.output, op)
			}

		}
	}
}

func Test_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := store.NewMockStore(ctrl)
	mock := New(mockSvc)
	tests := []struct {
		id     int
		output model.User
		user   []model.User
		err    error
	}{
		{id: 1, output: model.User{1, "Mahi", "Mahi@gmail.com", "9866895296", 20}, user: []model.User{{1, "Mahi", "Mahi@gmail.com", "9866895296", 20}}, err: nil},
	}
	for _, tes := range tests {
		mockSvc.EXPECT().GetUserById(tes.id).Return(tes.output, tes.err)
		if tes.err == nil {

			op, err := mock.GetByID(tes.id)

			if op != tes.output {
				t.Errorf("expect %v got %v", tes.output, op)
			}
			if !reflect.DeepEqual(err, tes.err) {
				t.Errorf("expect %s got %s", tes.err, err)
			}
		}
	}
}

func Test_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := store.NewMockStore(ctrl)
	mock := New(mockSvc)
	tests := []struct {
		users []model.User
		err   error
	}{
		{users: []model.User{{1, "Mahi", "Mahi@gmail.com", "9866895296", 20}}, err: nil},
	}
	for _, tes := range tests {
		mockSvc.EXPECT().GetAllUsers().Return(tes.users, tes.err)
		if tes.err == nil {

			op, err := mock.GetUsers()

			if !reflect.DeepEqual(op, tes.users) {
				t.Errorf("expect %v got %v", tes.users, op)
			}
			if !reflect.DeepEqual(err, tes.err) {
				t.Errorf("expect %s got %s", tes.err, err)
			}
		}
	}

}

func Test_PostUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := store.NewMockStore(ctrl)
	mock := New(mockSvc)
	tests := []struct {
		name  string
		email string
		phone string
		age   int
		usr   model.User
		err   error
	}{
		{
			name:  "Naira",
			email: "naira@gmail.com",
			phone: "6303844857",
			age:   20,
			usr:   model.User{Name: "Naira", Email: "naira@gmail.com", Phone: "6303844857", Age: 20},
			err:   nil,
		},
	}
	for _, tes := range tests {
		mockSvc.EXPECT().AddUser(tes.usr).Return(tes.usr.Id, tes.err)
		if tes.err == nil {

			op, err := mock.PostUser(tes.name, tes.email, tes.phone, tes.age)

			if !reflect.DeepEqual(op, tes.usr) {
				t.Errorf("expect %v got %v", tes.usr, op)
			}
			if !reflect.DeepEqual(err, tes.err) {
				t.Errorf("expect %s got %s", tes.err, err)
			}
		}
	}

}

func Test_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := store.NewMockStore(ctrl)
	mock := New(mockSvc)
	tests := []struct {
		u   model.User
		usr model.User
		us  []model.User
		op  model.User
		err error
	}{
		{
			u:   model.User{Id: 1, Name: "Nair", Email: "naira@gmail.com", Phone: "6303844857", Age: 20},
			usr: model.User{Id: 1, Name: "Naira", Email: "naira@gmail.com", Phone: "6303844857", Age: 20},
			op:  model.User{Id: 1, Name: "Naira", Email: "naira@gmail.com", Phone: "6303844857", Age: 20},
			us:  []model.User{{Id: 1, Name: "Naira", Email: "naira@gmail.com", Phone: "6303844857", Age: 20}, {Id: 2, Name: "Naira", Email: "Naira@gmail.com", Phone: "6303844857", Age: 20}},
			err: nil,
		},
	}
	for _, tes := range tests {
		call := mockSvc.EXPECT().UpdateUser(tes.usr).Return(tes.err)
		mockSvc.EXPECT().GetAllUsers().Return(tes.us, tes.err).After(call)

		if tes.err == nil {

			op, err := mock.Update(tes.u, tes.usr)

			if !reflect.DeepEqual(op, tes.op) {
				t.Errorf("expect %v got %v", tes.usr, op)
			}
			if !reflect.DeepEqual(err, tes.err) {
				t.Errorf("expect %s got %s", tes.err, err)
			}
		}
	}

}

func Test_DeleteByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSvc := store.NewMockStore(ctrl)
	mock := New(mockSvc)
	tests := []struct {
		id  int
		err error
	}{
		{
			id:  1,
			err: nil,
		},
	}
	for _, tes := range tests {
		mockSvc.EXPECT().DeleteUser(tes.id).Return(tes.err)
		if tes.err == nil {

			err := mock.DeleteByID(tes.id)
			if !reflect.DeepEqual(err, tes.err) {
				t.Errorf("expect %s got %s", tes.err, err)
			}
		}
	}

}
