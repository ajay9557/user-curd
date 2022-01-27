package users

import (
	"errors"
	"reflect"
	"testing"
	"user-curd/models"
	"user-curd/stores"

	gomock "github.com/golang/mock/gomock"
)

func TestInserUserDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)
	testCases := []struct {
		desc     string
		User     models.User
		mock     []*gomock.Call
		expecErr error
	}{
		{
			desc: "Success case",
			User: models.User{
				Id:    1,
				Name:  "gopi",
				Email: "gopi@gmail.com",
				Phone: "1234567899",
				Age:   23,
			},
			mock: []*gomock.Call{
				mockStore.EXPECT().InsertUser(gomock.Any()).Return(nil),
				mockStore.EXPECT().GetEmail(gomock.Any()).Return(true, nil),
			},
			expecErr: nil,
		},
		{
			desc: "Failure case - 1",
			User: models.User{
				Id:    2,
				Name:  "gopi",
				Email: "gopi123@gmail",
				Phone: "1234567899",
				Age:   23,
			},
			expecErr: errors.New("email not valid"),
		},
		{
			desc: "Failure case - 2",
			User: models.User{
				Id:    2,
				Name:  "gopi",
				Email: "gopi123@gmail.com",
				Phone: "1234567899",
				Age:   23,
			},
			mock: []*gomock.Call{
				mockStore.EXPECT().GetEmail(gomock.Any()).Return(false, nil),
			},
			expecErr: errors.New("email already present"),
		},
		{
			desc: "Failure case - id",
			User: models.User{
				Id:    0,
				Name:  "gopi1",
				Email: "gopi1236@gmail.com",
				Phone: "12345267899",
				Age:   23,
			},
			expecErr: errors.New("id should not be zero"),
		},
		{
			desc: "Failure case - 3",
			User: models.User{
				Id:    2,
				Name:  "gopi",
				Email: "gopi@gmail.com",
				Phone: "123411567899",
				Age:   23,
			},
			mock: []*gomock.Call{
				mockStore.EXPECT().InsertUser(gomock.Any()).Return(errors.New("error generated")),
				mockStore.EXPECT().GetEmail(gomock.Any()).Return(true, nil),
			},
			expecErr: errors.New("error generated"),
		},
	}
	for _, tes := range testCases {
		t.Run(tes.desc, func(t *testing.T) {
			err := mock.InsertUserDetails(tes.User)
			if err != nil && !reflect.DeepEqual(err, tes.expecErr) {
				t.Error("Expected: ", tes.expecErr, "Obtained: ", err)
			}
		})
	}
}

func TestFetchAllUsersDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	testCases := []struct {
		desc      string
		mock      []*gomock.Call
		expectRes []models.User
		expecErr  error
	}{
		{
			desc: "Success case",
			mock: []*gomock.Call{mockStore.EXPECT().FetchAllUsers().Return([]models.User{
				{Id: 1,
					Name:  "gopi",
					Email: "gopi@gmail.com",
					Phone: "1234567899",
					Age:   23},
			}, nil)},
			expectRes: []models.User{
				{
					Id:    1,
					Name:  "gopi",
					Email: "gopi@gmail.com",
					Phone: "1234567899",
					Age:   23,
				},
			},
			expecErr: nil,
		},
		{
			desc: "Failure case",
			mock: []*gomock.Call{mockStore.EXPECT().FetchAllUsers().Return([]models.User{
				{Id: 2,
					Name:  "gopi",
					Email: "gopi123@gmail.com",
					Phone: "1234567899",
					Age:   23},
			}, errors.New("data fetching error"))},
			expectRes: []models.User{
				{
					Id:    2,
					Name:  "gopi",
					Email: "gopi123@gmail.com",
					Phone: "1234567899",
					Age:   23,
				},
			},
			expecErr: errors.New("data fetching error"),
		},
	}
	for _, tes := range testCases {
		t.Run(tes.desc, func(t *testing.T) {
			res, err := mock.FetchAllUserDetails()
			if err != nil && !reflect.DeepEqual(err, tes.expecErr) {
				t.Error("Expected: ", tes.expecErr, "Obtained: ", err)
			}
			if !reflect.DeepEqual(res, tes.expectRes) {
				t.Error("Expected: ", tes.expectRes, "Obtained: ", res)
			}
		})
	}
}

func TestFetchAllUserDetailsById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	testCases := []struct {
		desc      string
		Id        int
		mock      []*gomock.Call
		expectRes models.User
		expecErr  error
	}{
		{
			desc: "Success case",
			Id:   1,
			mock: []*gomock.Call{mockStore.EXPECT().FetchUserById(gomock.Any()).Return(models.User{
				Id:    1,
				Name:  "gopi",
				Email: "gopi@gmail.com",
				Phone: "1234567899",
				Age:   23,
			}, nil)},
			expectRes: models.User{
				Id:    1,
				Name:  "gopi",
				Email: "gopi@gmail.com",
				Phone: "1234567899",
				Age:   23,
			},
			expecErr: nil,
		},
		{
			desc:      "Failure case - id",
			Id:        0,
			expectRes: models.User{},
			expecErr:  errors.New("id should not be zero"),
		},
		{
			desc: "Failure case",
			Id:   2,
			mock: []*gomock.Call{mockStore.EXPECT().FetchUserById(gomock.Any()).Return(models.User{
				Id:    2,
				Name:  "gopi",
				Email: "gopi@gmail.com",
				Phone: "1234567899",
				Age:   23,
			}, errors.New("data fetching error"))},
			expectRes: models.User{
				Id:    2,
				Name:  "gopi",
				Email: "gopi@gmail.com",
				Phone: "1234567899",
				Age:   23,
			},
			expecErr: errors.New("data fetching error"),
		},
	}
	for _, tes := range testCases {
		t.Run(tes.desc, func(t *testing.T) {
			res, err := mock.FetchUserDetailsById(tes.Id)
			if err != nil && !reflect.DeepEqual(err, tes.expecErr) {
				t.Error("Expected: ", tes.expecErr, "Obtained: ", err)
			}
			if !reflect.DeepEqual(res, tes.expectRes) {
				t.Error("Expected: ", tes.expectRes, "Obtained: ", res)
			}
		})
	}
}

func TestUpdateUserDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)
	testCases := []struct {
		desc     string
		User     models.User
		mock     []*gomock.Call
		expecErr error
	}{
		{
			desc: "Success case",
			User: models.User{
				Id:    1,
				Name:  "gopi",
				Email: "gopi@gmail.com",
				Phone: "1234567899",
				Age:   23,
			},
			mock: []*gomock.Call{
				mockStore.EXPECT().UpdateUser(gomock.Any()).Return(nil),
				mockStore.EXPECT().GetEmail(gomock.Any()).Return(true, nil),
			},
			expecErr: nil,
		},
		{
			desc: "Failure case - id",
			User: models.User{
				Id:    0,
				Name:  "gopi",
				Email: "gopi123@gmail.com",
				Phone: "1234567899",
				Age:   23,
			},
			expecErr: errors.New("id should not be zero"),
		},
		{
			desc: "Failure case - 1",
			User: models.User{
				Id:    2,
				Name:  "gopi",
				Email: "gopi123@gmail",
				Phone: "1234567899",
				Age:   23,
			},
			expecErr: errors.New("email not valid"),
		},
		{
			desc: "Failure case - 2",
			User: models.User{
				Id:    2,
				Name:  "gopi",
				Email: "gopi123@gmail.com",
				Phone: "1234567899",
				Age:   23,
			},
			mock: []*gomock.Call{
				mockStore.EXPECT().GetEmail(gomock.Any()).Return(false, nil),
			},
			expecErr: errors.New("email already present"),
		},
		{
			desc: "Failure case - 3",
			User: models.User{
				Id:    2,
				Name:  "gopi",
				Email: "gopi@gmail.com",
				Phone: "123411567899",
				Age:   23,
			},
			mock: []*gomock.Call{
				mockStore.EXPECT().UpdateUser(gomock.Any()).Return(errors.New("error generated")),
				mockStore.EXPECT().GetEmail(gomock.Any()).Return(true, nil),
			},
			expecErr: errors.New("error generated"),
		},
	}
	for _, tes := range testCases {
		t.Run(tes.desc, func(t *testing.T) {
			err := mock.UpdateUserDetails(tes.User)
			if err != nil && !reflect.DeepEqual(err, tes.expecErr) {
				t.Error("Expected: ", tes.expecErr, "Obtained: ", err)
			}
		})
	}
}

func TestDeleteUserDetailsById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	testCases := []struct {
		desc     string
		Id       int
		mock     []*gomock.Call
		expecErr error
	}{
		{
			desc:     "Success case",
			Id:       1,
			mock:     []*gomock.Call{mockStore.EXPECT().DeleteUserById(gomock.Any()).Return(nil)},
			expecErr: nil,
		},
		{
			desc:     "Failure case",
			Id:       2,
			mock:     []*gomock.Call{mockStore.EXPECT().DeleteUserById(gomock.Any()).Return(errors.New("error generated"))},
			expecErr: errors.New("error generated"),
		},
		{
			desc:     "Failure case - id",
			Id:       0,
			expecErr: errors.New("id should not be zero"),
		},
	}
	for _, tes := range testCases {
		t.Run(tes.desc, func(t *testing.T) {
			err := mock.DeleteUserDetailsById(tes.Id)
			if err != nil && !reflect.DeepEqual(err, tes.expecErr) {
				t.Error("Expected: ", tes.expecErr, "Obtained: ", err)
			}
		})
	}
}
