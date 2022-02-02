package users

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tejas/user-crud/models"
	"github.com/tejas/user-crud/store"
)

func TestFindUserById(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockUserStore := store.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	testCases := []struct {
		desc     string
		id       int
		expected models.User
		mockCall *gomock.Call
	}{
		{
			desc: "Case 1: Success Case",
			id:   1,
			expected: models.User{
				Id:    1,
				Name:  "tejas",
				Email: "tejas@gmail.com",
				Phone: "9876543210",
				Age:   23,
			},
			mockCall: mockUserStore.EXPECT().GetUserById(1).Return(&models.User{
				Id:    1,
				Name:  "tejas",
				Email: "tejas@gmail.com",
				Phone: "9876543210",
				Age:   23,
			}, nil),
		},
		{
			desc:     "Case 2: Failure Case",
			id:       2,
			expected: models.User{},
			mockCall: mockUserStore.EXPECT().GetUserById(2).Return(&models.User{}, errors.New("Cannot fetch user for the given id")),
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			user, err := testUserService.GetUserById(test.id)

			if err != nil && !reflect.DeepEqual(test.expected, user) {
				t.Errorf("Expected : %v, Got : %v", test.expected, *user)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockUserStore := store.NewMockUser(ctrl)
	usrService := New(mockUserStore)
	defer ctrl.Finish()

	testCases := []struct {
		desc        string
		expectedErr error
		expectedOut []*models.User
	}{

		{
			desc:        "Case 1: Success case",
			expectedErr: nil,
			expectedOut: []*models.User{
				{
					Id:    1,
					Name:  "virat",
					Email: "virat@gmail.com",
					Phone: "9080706050",
					Age:   21,
				},
				{
					Id:    2,
					Name:  "rahul",
					Email: "rahul@gmail.com",
					Phone: "8799767570",
					Age:   19,
				},
			},
		},

		{
			desc:        "Case 2: Failure Case",
			expectedErr: errors.New("cannot fetch users"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+tc.desc, func(t *testing.T) {
			mockUserStore.EXPECT().GetUsers().Return(tc.expectedOut, tc.expectedErr)
			resp, err := usrService.GetUsers()
			if tc.expectedOut != nil && !reflect.DeepEqual(resp, tc.expectedOut) {
				t.Errorf("TestCase%v Expected: %v, Got: %v, ", tc.desc, tc.expectedOut, resp)
			}
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("Expected: %v, Got: %v, ",tc.expectedErr, err)
			}
		})
	}
}

func TestCreateUserService(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockUserStore := store.NewMockUser(ctrl)
	usrService := New(mockUserStore)
	defer ctrl.Finish()

	testCases := []struct {
		desc        string
		input       models.User
		mockCall    *gomock.Call
		expectedErr error
	}{

		{
			desc: "Case 1: Success case",
			input: models.User{
				Id:    3,
				Name:  "ross",
				Email: "ross@gmail.com",
				Phone: "6575859555",
				Age:   40,
			},
			mockCall: mockUserStore.
				EXPECT().
				CreateUsers(models.User{
					Id:    3,
					Name:  "ross",
					Email: "ross@gmail.com",
					Phone: "6575859555",
					Age:   40,
				}).
				Return(nil),
			expectedErr: nil,
		},
		{
			desc: "Case 2:invalid id",
			input: models.User{
				Id:    -1,
				Name:  "joey",
				Email: "",
				Phone: "",
				Age:   21,
			},
			mockCall:    nil,
			expectedErr: errors.New("invalid id"),
		},

		{
			desc: "Case 3: invalid email",
			input: models.User{
				Id:    4,
				Name:  "rachel",
				Email: "helloemail",
				Phone: "9876545670",
				Age:   21,
			},
			mockCall:    nil,
			expectedErr: errors.New("invalid email"),
		},
		{
			desc: "Case 4: invalid email",
			input: models.User{
				Id:    4,
				Name:  "rock",
				Email: "rockgmail.com",
				Phone: "7878787878",
				Age:   21,
			},
			mockCall:    nil,
			expectedErr: errors.New("invalid email"),
		},
		{
			desc: "Case 5: invalid email",
			input: models.User{
				Id:    4,
				Name:  "arun",
				Email: "arunA22N$@gmail.com",
				Phone: "6767676767",
				Age:   21,
			},
			mockCall:    nil,
			expectedErr: errors.New("invalid email"),
		},

		{
			desc: "Case 6: invalid phone",
			input: models.User{
				Id:    4,
				Name:  "arun",
				Email: "arun@gmail.com",
				Phone: "",
				Age:   21,
			},
			expectedErr: errors.New("invalid phone number"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+tc.desc, func(t *testing.T) {
			err := usrService.CreateUser(tc.input)
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("Expected: %v, Got: %v, ",tc.expectedErr, err)
			}
		})
	}
}

func TestUpdateUserService(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockUserStore := store.NewMockUser(ctrl)
	userService := New(mockUserStore)
	defer ctrl.Finish()

	testCase := []struct {
		desc        string
		input       models.User
		mockCall    *gomock.Call
		expectedErr error
	}{

		{
			desc: "Case 1: Success case",
			input: models.User{
				Id:    3,
				Name:  "paul",
				Email: "paul22@gmail.com",
				Phone: "9998887777",
				Age:   24,
			},
			mockCall: mockUserStore.
				EXPECT().
				UpdateUser(models.User{
					Id:    3,
					Name:  "paul",
					Email: "paul22@gmail.com",
					Phone: "9998887777",
					Age:   24,
				}).
				Return(nil),
			expectedErr: nil,
		},

		{
			desc: "Case 2: invalid id",
			input: models.User{
				Id:    -1,
				Name:  "paul",
				Email: "",
				Phone: "",
				Age:   24,
			},
			mockCall:    nil,
			expectedErr: errors.New("invalid id"),
		},

		{
			desc: "Case 3: invalid email",
			input: models.User{
				Id:    4,
				Name:  "paul",
				Email: "paul@sss",
				Phone: "9998887777",
				Age:   24,
			},
			mockCall:    nil,
			expectedErr: errors.New("invalid email"),
		},

		{
			desc: "Case 4: invalid phone",
			input: models.User{
				Id:    4,
				Name:  "paul",
				Email: "paul22@gmail.com",
				Phone: "8888t8888s",
				Age:   24,
			},
			mockCall:    nil,
			expectedErr: errors.New("invalid phone number"),
		},
	}

	for _, tc := range testCase {
		t.Run("testing "+tc.desc, func(t *testing.T) {
			err := userService.UpdateUser(tc.input)
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("Expected: %v, Got: %v, ",tc.expectedErr, err)
			}
		})
	}
}

func TestDeleteUserService(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockUserStore := store.NewMockUser(ctrl)
	userService := New(mockUserStore)
	defer ctrl.Finish()

	testCase := []struct {
		desc        string
		input       int
		mockCall    *gomock.Call
		expectedErr error
	}{

		{
			desc:  "Case 1: Success case",
			input: 3,
			mockCall: mockUserStore.
				EXPECT().
				DeleteUser(3).Return(nil),
			expectedErr: nil,
		},
		{
			desc:        "Case 2: invalid id",
			input:       -2,
			mockCall:    nil,
			expectedErr: errors.New("invalid id"),
		},
	}

	for _, tc := range testCase {
		t.Run("testing "+tc.desc, func(t *testing.T) {
			err := userService.DeleteUser(tc.input)
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("Expected: %v, Got: %v, ", tc.expectedErr, err)
			}
		})
	}
}
