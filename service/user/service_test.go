package user

import (
	"fmt"
	"reflect"
	"testing"
	"zopsmart/user-curd/model"
	store "zopsmart/user-curd/store"

	gomock "github.com/golang/mock/gomock"
)

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
		users  []model.User
		user   model.User
		output model.User
		err    error
	}{
		{
			users: []model.User{
				{
					Name: "Naira", Email: "naira@gmail.com", Phone: "6303844857", Age: 20,
				},
			},
			user:   model.User{Name: "Naira", Email: "naa@gmail.com", Phone: "6303844857", Age: 20},
			output: model.User{Name: "Naira", Email: "naa@gmail.com", Phone: "6303844857", Age: 20},
			err:    nil,
		},
		{
			users: []model.User{
				{
					Name: "Naira", Email: "naa@gmail.com", Phone: "6303844857", Age: 20,
				},
			},
			user:   model.User{Name: "Naira", Email: "naa@gmail.com", Phone: "6303844857", Age: 20},
			output: model.User{Name: "Naira", Email: "naa@gmail.com", Phone: "6303844857", Age: 20},
			err:    fmt.Errorf("e,ail already exists"),
		},
		{
			users: []model.User{
				{
					Name: "Naira", Email: "naa@gmail.com", Phone: "6303844857", Age: 20,
				},
			},
			user:   model.User{Name: "Naira", Email: "naagmail.com", Phone: "6303844857", Age: 20},
			output: model.User{Name: "Naira", Email: "naagmail.com", Phone: "6303844857", Age: 20},
			err:    fmt.Errorf("enter valid email"),
		},
	}
	for _, tes := range tests {
		isValid, err := CheckMail(tes.user.Email)
		if isValid {
			call := mockSvc.EXPECT().GetAllUsers().Return(tes.users, tes.err).MinTimes(0)
			mockSvc.EXPECT().AddUser(tes.user).Return(tes.user.Id, tes.err).After(call).MinTimes(0)
			if tes.err == nil {

				op, err := mock.PostUser(tes.user)

				if !reflect.DeepEqual(op, tes.output) {
					t.Errorf("expected %v got %v", tes.output, op)
				}
				if !reflect.DeepEqual(err, tes.err) {
					t.Errorf("expected %s got %s", tes.err, err)
				}
			}
		} else {
			if !reflect.DeepEqual(err, tes.err) {
				t.Errorf("expected %s got %s", tes.err, err)
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
		id     int
		user   model.User
		output model.User
		users  []model.User
		err    error
	}{
		{
			id: 2,
			users: []model.User{
				{
					Id: 2, Name: "Naira", Email: "naa@gmail.com", Phone: "6303844857", Age: 20,
				},
			},
			user:   model.User{Id: 2, Name: "Naira", Email: "naira@gmail.com", Phone: "6303844857", Age: 20},
			output: model.User{Id: 2, Name: "Naira", Email: "naira@gmail.com", Phone: "6303844857", Age: 20},
			err:    nil,
		},
		{
			id: 0,
			users: []model.User{
				{
					Id: 2, Name: "Naira", Email: "naa@gmail.com", Phone: "6303844857", Age: 20,
				},
			},
			user:   model.User{},
			output: model.User{},
			err:    fmt.Errorf("id doesnt exist"),
		},
	}
	for _, tes := range tests {
		call := mockSvc.EXPECT().GetAllUsers().Return(tes.users, tes.err).MinTimes(0)
		mockSvc.EXPECT().GetUserById(tes.id).Return(tes.user, tes.err).After(call).MinTimes(0)
		isValid, _ := CheckMail(tes.user.Email)
		if isValid {
			thirdcall := mockSvc.EXPECT().GetAllUsers().Return(tes.users, tes.err).MinTimes(0)
			mockSvc.EXPECT().UpdateUser(tes.user).Return(tes.err).After(thirdcall).MinTimes(0)

			if tes.err == nil {

				op, err := mock.Update(tes.id, tes.user)

				if !reflect.DeepEqual(op, tes.output) {
					t.Errorf("expect %v got %v", tes.user, op)
				}

				if !reflect.DeepEqual(err, tes.err) {
					t.Errorf("expect %s got %s", tes.err, err)
				}
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
		id    int
		users []model.User
		err   error
	}{
		{
			id: 2,
			users: []model.User{
				{
					Id: 2, Name: "Naira", Email: "naa@gmail.com", Phone: "6303844857", Age: 20,
				},
			},
			err: nil,
		},
	}
	for _, tes := range tests {
		call := mockSvc.EXPECT().GetAllUsers().Return(tes.users, tes.err)
		mockSvc.EXPECT().DeleteUser(tes.id).Return(tes.err).After(call).MinTimes(0)
		if tes.err == nil {
			err := mock.DeleteByID(tes.id)
			if !reflect.DeepEqual(err, tes.err) {
				t.Errorf("expect %s got %s", tes.err, err)
			}
		}
	}

}
