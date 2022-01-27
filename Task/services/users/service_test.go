package service

import (

	"reflect"
	"testing"
	"zopsmart/Task/stores"
	"zopsmart/Task/models"

	gomock "github.com/golang/mock/gomock"
)

func TestValidateEmail(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	testCases := []struct {
		Email     string
		mock      []*gomock.Call
		expectRes bool
		expectErr  error
	}{
		{
			Email: "prasath@gmail.com",
			mock: []*gomock.Call(mockStore.EXPECT().GetMail(gomock.Any()).Return(true,nil)),
			expectRes: true,
			expectErr: nil,
		},
	}

	for _,tc := range testCases {
		t.Run("Sucess test case ",func (t *testing.T) {
			res, err := mock.ValidateEmail(tc.Email)
			if err != nil && !reflect.DeepEqual(err, tc.expectErr) {
				t.Error("Expected: ", tc.expectErr, "Obtained: ", err)
			}
			if !reflect.DeepEqual(res, tc.expectRes) {
				t.Error("Expected: ", tc.expectRes, "Obtained: ", res)
			}
		})
	}

}

func TestDeletebyId(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	testCases := []struct {
		Id    int
		mock      []*gomock.Call
		expectErr  error
	}{
		{
			Id: 1,
			mock: []*gomock.Call(mockStore.EXPECT().DeletebyId(gomock.Any()).Return(true,nil)),
			expectErr: nil,
		},
	}

	for _,tc := range testCases {
		t.Run("Sucess test case ",func (t *testing.T) {
			err := mock.DeletebyId(tc.Id)
			if err != nil && !reflect.DeepEqual(err, tc.expectErr) {
				t.Error("Expected: ", tc.expectErr, "but Got: ", err)
			}
		})
	}


}


func TestUpdatebyId(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	testCases := []struct {
		Id    int
		mock      []*gomock.Call
		expectErr  error
	}{
		{
			Id: 1,
			mock: []*gomock.Call(mockStore.EXPECT().UpdatebyId(gomock.Any()).Return(true,nil)),
			expectErr: nil,
		},
	}

	for _,tc := range testCases {
		t.Run("Sucess test case ",func (t *testing.T) {
			err := mock.DeletebyId(tc.Id)
			if err != nil && !reflect.DeepEqual(err, tc.expectErr) {
				t.Error("Expected: ", tc.expectErr, "but Got: ", err)
			}
		})
	}

}

func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	testCases := []struct {
		Id    int
		mock      []*gomock.Call
		expectOut models.User
		expectErr  error
	}{
		{
			Id: 1,
			mock: []*gomock.Call(mockStore.EXPECT().GetUserById(gomock.Any()).Return(models.User{
				Id: 1,
				Name: "Prasath",
				Email:"prasath@gmail.com",
				Phone: "12345",
				Age: 20,
			})),
			expectErr: nil,
		},
	}

	for _,tc := range testCases {
		t.Run("Sucess test case ",func (t *testing.T) {
			res, err := mock.GetUserById(tc.Id)
			if err != nil && !reflect.DeepEqual(err, tc.expectErr) {
				t.Error("Expected: ", tc.expectErr, "but Got: ", err)
			}

			if !reflect.DeepEqual(res,tc.expectOut) {
				t.Errorf("Expected :",tc.expectOut,"But Got:",res)
			}
		})
	}

}

func TestValidateId(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	testCases := []struct {
		Id    int
		mock      []*gomock.Call
		expectOut bool
		expectErr  error
	}{
		{
			Id: 1,
			mock: []*gomock.Call(mockStore.EXPECT().ValidateId(gomock.Any()).Return(true,nil)),
			expectOut: true,
			expectErr: nil,
		},
	}

	for _,tc := range testCases {
		t.Run("Sucess test case ",func (t *testing.T) {
			res, err := mock.ValidateId(tc.Id)
			if err != nil && !reflect.DeepEqual(err, tc.expectErr) {
				t.Error("Expected: ", tc.expectErr, "but Got: ", err)
			}

			if !reflect.DeepEqual(res,tc.expectOut) {
				t.Errorf("Expected :",tc.expectOut,"But Got:",res)
			}
		})
	}


}

