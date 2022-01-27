package users

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
	"user-curd/datastore"
	"user-curd/entities"

	"github.com/golang/mock/gomock"
)

func TestGetUserByIdService(t *testing.T) {

	// create a mock for datastore layer
	ctrl := gomock.NewController(t)
	mockUserStore := datastore.NewMockUserStoresHandler(ctrl)
	usrServ := New(mockUserStore)
	defer ctrl.Finish()

	testCases := []struct {
		caseId      int
		input       interface{}
		mockCall    *gomock.Call
		expectedErr error
		expectedOut entities.User
	}{
		// Success case
		{
			caseId:      1,
			input:       1,
			expectedErr: nil,
			expectedOut: entities.User{
				Id:    1,
				Name:  "john",
				Email: "john91@example.com",
				Phone: "8095765432",
				Age:   21,
			},
			mockCall: mockUserStore.
				EXPECT().
				GetUserById(1).
				Return(&entities.User{
					Id:    1,
					Name:  "john",
					Email: "john91@example.com",
					Phone: "8095765432",
					Age:   21,
				}, nil),
		},
		// Error validating id
		{
			caseId:      2,
			expectedErr: errors.New("error invalid id type"),
		},
		// Error from datastore
		{
			caseId:      3,
			input:       10,
			expectedErr: errors.New("error fetching data from database"),
			mockCall: mockUserStore.
				EXPECT().
				GetUserById(10).
				Return(&entities.User{}, errors.New("error fetching data from database")),
			expectedOut: entities.User{},
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {

			resp, err := usrServ.GetUserByIdService(tc.input)
			if !reflect.DeepEqual(resp, &tc.expectedOut) {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedOut, resp)
			}
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedErr, err)
			}
		})
	}

}

func TestGetAllUsersService(t *testing.T) {

	// create a mock for datastore layer
	ctrl := gomock.NewController(t)
	mockUserStore := datastore.NewMockUserStoresHandler(ctrl)
	usrService := New(mockUserStore)
	defer ctrl.Finish()

	testCases := []struct {
		caseId      int
		expectedErr error
		expectedOut []*entities.User
	}{
		// Success case
		{
			caseId:      1,
			expectedErr: nil,
			expectedOut: []*entities.User{
				{
					Id:    1,
					Name:  "john",
					Email: "john19@example.com",
					Phone: "9823218772",
					Age:   21,
				},
				{
					Id:    2,
					Name:  "jess",
					Email: "jessH99@example.com",
					Phone: "9821328772",
					Age:   19,
				},
			},
		},
		// Error case
		{
			caseId:      2,
			expectedErr: errors.New("error fetching data from databases"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {
			mockUserStore.EXPECT().GetAllUsers().Return(tc.expectedOut, tc.expectedErr)
			resp, err := usrService.GetAllUsersService()
			if tc.expectedOut != nil && !reflect.DeepEqual(resp, tc.expectedOut) {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedOut, resp)
			}
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedErr, err)
			}
		})
	}
}

func TestCreateUserService(t *testing.T) {

	// create a mock for datastore layer
	ctrl := gomock.NewController(t)
	mockUserStore := datastore.NewMockUserStoresHandler(ctrl)
	usrService := New(mockUserStore)
	defer ctrl.Finish()

	testCases := []struct {
		caseId      int
		input       entities.User
		mockCall    *gomock.Call
		expectedErr error
	}{
		// Success case
		{
			caseId: 1,
			input: entities.User{
				Id:    3,
				Name:  "jenny",
				Email: "jenny27@example.com",
				Phone: "8765234589",
				Age:   40,
			},
			mockCall: mockUserStore.
				EXPECT().
				CreateUser(entities.User{
					Id:    3,
					Name:  "jenny",
					Email: "jenny27@example.com",
					Phone: "8765234589",
					Age:   40,
				}).
				Return(nil),
			expectedErr: nil,
		},
		// Case invalid id
		{
			caseId: 2,
			input: entities.User{
				Id:    -1,
				Name:  "josh",
				Email: "",
				Phone: "",
				Age:   21,
			},
			mockCall:    nil,
			expectedErr: errors.New("error invalid id"),
		},
		// Case invalid emails
		{
			caseId: 3,
			input: entities.User{
				Id:    4,
				Name:  "josh",
				Email: "bad-email",
				Phone: "1289312798",
				Age:   21,
			},
			mockCall:    nil,
			expectedErr: errors.New("error invalid email"),
		},
		{
			caseId: 4,
			input: entities.User{
				Id:    4,
				Name:  "josh",
				Email: "josh-example.com",
				Phone: "1289312798",
				Age:   21,
			},
			mockCall:    nil,
			expectedErr: errors.New("error invalid email"),
		},
		{
			caseId: 5,
			input: entities.User{
				Id:    4,
				Name:  "josh",
				Email: "joshH56D$@example.com",
				Phone: "1289312798",
				Age:   21,
			},
			mockCall:    nil,
			expectedErr: errors.New("error invalid email"),
		},
		// Case invalid phone
		{
			caseId: 6,
			input: entities.User{
				Id:    4,
				Name:  "josh",
				Email: "josh65@gmail.com",
				Phone: "",
				Age:   21,
			},
			expectedErr: errors.New("error invalid phone"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {
			err := usrService.CreateUserService(tc.input)
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedErr, err)
			}
		})
	}
}

func TestUpdateUserService(t *testing.T) {

	// create a mock for datastore layer
	ctrl := gomock.NewController(t)
	mockUserStore := datastore.NewMockUserStoresHandler(ctrl)
	userService := New(mockUserStore)
	defer ctrl.Finish()

	testCase := []struct {
		caseId      int
		input       entities.User
		mockCall    *gomock.Call
		expectedErr error
	}{
		// Success case
		{
			caseId: 1,
			input: entities.User{
				Id:    3,
				Name:  "josh",
				Email: "josh22@example.com",
				Phone: "7892678212",
				Age:   42,
			},
			mockCall: mockUserStore.
				EXPECT().
				UpdateUser(entities.User{
					Id:    3,
					Name:  "josh",
					Email: "josh22@example.com",
					Phone: "7892678212",
					Age:   42,
				}).
				Return(nil),
			expectedErr: nil,
		},
		// Error cases
		// Case invalid id
		{
			caseId: 2,
			input: entities.User{
				Id:    -1,
				Name:  "josh",
				Email: "",
				Phone: "",
				Age:   21,
			},
			mockCall:    nil,
			expectedErr: errors.New("error invalid id"),
		},
		// Case invalid email
		{
			caseId: 3,
			input: entities.User{
				Id:    4,
				Name:  "josh",
				Email: "sdai@sss",
				Phone: "1289312798",
				Age:   21,
			},
			mockCall:    nil,
			expectedErr: errors.New("error invalid email"),
		},
		// Case invalid phone
		{
			caseId: 4,
			input: entities.User{
				Id:    4,
				Name:  "josh",
				Email: "josh29@example.com",
				Phone: "789798s999",
				Age:   21,
			},
			mockCall:    nil,
			expectedErr: errors.New("error invalid phone"),
		},
	}

	for _, tc := range testCase {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {
			err := userService.UpdateUserService(tc.input)
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedErr, err)
			}
		})
	}
}

func TestDeleteUserService(t *testing.T) {

	// create a mock for datastore layer
	ctrl := gomock.NewController(t)
	mockUserStore := datastore.NewMockUserStoresHandler(ctrl)
	userService := New(mockUserStore)
	defer ctrl.Finish()

	testCase := []struct {
		caseId      int
		input       interface{}
		mockCall    *gomock.Call
		expectedErr error
	}{
		// Success case
		{
			caseId: 1,
			input:  3,
			mockCall: mockUserStore.
				EXPECT().
				DeleteUser(3).Return(nil),
			expectedErr: nil,
		},
		// Error case invalid id
		{
			caseId:      2,
			input:       'x',
			mockCall:    nil,
			expectedErr: errors.New("error invalid id"),
		},
	}

	for _, tc := range testCase {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {
			err := userService.DeleteUserService(tc.input)
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedErr, err)
			}
		})
	}
}
