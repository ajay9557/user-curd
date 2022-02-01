package users

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"layer/user/models"
	"layer/user/stores"
)

func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	tests := []struct {
		desc     string
		id       int
		expected models.User
		mockCall *gomock.Call
	}{
		{
			desc:     "Case1",
			id:       1,
			expected: models.User{Id: 1, Name: "Naruto", Email: "naruto@japan.com", Phone: "9999999999", Age: 18},
			mockCall: mockUserStore.EXPECT().GetUserById(1).Return(&models.User{Id: 1, Name: "Naruto", Email: "naruto@japan.com", Phone: "9999999999", Age: 18}, nil),
		},
		{
			desc:     "Case2",
			id:       2,
			expected: models.User{},
			mockCall: mockUserStore.EXPECT().GetUserById(2).Return(&models.User{}, errors.New("Cannot fetch user for given id")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			user, err := testUserService.GetUserById(test.id)

			if err != nil && !reflect.DeepEqual(test.expected, user) {
				t.Errorf("Expected: %v, Got: %v", test.expected, user)
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	data1 := []*models.User{
		{Id: 1, Name: "Naruto", Email: "naruto@gmail.com", Phone: "9999999999", Age: 18},
		{Id: 2, Name: "Itachi", Email: "itachi@gmail.com", Phone: "8320578360", Age: 24},
	}

	tests := []struct {
		desc     string
		expected []*models.User
		mockCall *gomock.Call
	}{
		{
			desc:     "Case1",
			expected: data1,
			mockCall: mockUserStore.EXPECT().GetUsers().Return(data1, nil),
		},
		{
			desc:     "Case2",
			expected: nil,
			mockCall: mockUserStore.EXPECT().GetUsers().Return(nil, errors.New("Cannot fetch users")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			users, err := testUserService.GetUsers()

			if err != nil && !reflect.DeepEqual(test.expected, users) {
				t.Errorf("Expected: %v, Got: %v", test.expected, users)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	testUser := models.User{Name: "Ridhdhish", Email: "ridhdhish@gmail.com", Phone: "8320578360", Age: 21}
	expectedUser := models.User{
		Id:    1,
		Name:  "Ridhdhish",
		Email: "ridhdhish@gmail.com",
		Phone: "8320578360",
		Age:   21,
	}
	updatedUser := models.User{
		Id:    1,
		Name:  "Ridhdhish",
		Email: "ridhdhish@gmail.com",
		Phone: "8320578360",
		Age:   22,
	}

	tests := []struct {
		desc          string
		id            int
		expectedUser  models.User
		expectedError error
		mockCall      []*gomock.Call
	}{
		{
			desc:          "Case1",
			id:            1,
			expectedUser:  updatedUser,
			expectedError: nil,
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().GetUserById(1).Return(&expectedUser, nil),
				mockUserStore.EXPECT().UpdateUser(1, testUser).Return(nil),
				mockUserStore.EXPECT().GetUserById(1).Return(&updatedUser, nil),
			},
		},
		{
			desc:          "Case2",
			id:            -1,
			expectedUser:  models.User{},
			expectedError: errors.New("User id should be greater than 0"),
			mockCall:      nil,
		},
		{
			desc:          "Case3",
			id:            100,
			expectedUser:  models.User{},
			expectedError: errors.New("User not exists for given user id"),
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().GetUserById(100).Return(nil, errors.New("User not exists for given user id")),
			},
		},
		{
			desc:          "Case4",
			id:            1,
			expectedUser:  updatedUser,
			expectedError: errors.New("Connection lost"),
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().GetUserById(1).Return(&expectedUser, nil),
				mockUserStore.EXPECT().UpdateUser(1, testUser).Return(errors.New("Connection lost")),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			user, err := testUserService.UpdateUser(test.id, testUser)

			if err != nil && err.Error() != test.expectedError.Error() && !reflect.DeepEqual(user, test.expectedUser) {
				t.Errorf("expectedUser: %v, Got: %v", test.expectedUser, user)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	tests := []struct {
		desc          string
		id            int
		expectedError error
		mockCall      []*gomock.Call
	}{
		{
			desc:          "Case1",
			id:            1,
			expectedError: nil,
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().GetUserById(1).Return(&models.User{}, nil),
				mockUserStore.EXPECT().DeleteUser(1).Return(nil),
			},
		},
		{
			desc:          "Case2",
			id:            2,
			expectedError: errors.New("User does not exist for given id"),
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().GetUserById(2).Return(nil, errors.New("User does not exist for given id")),
			},
		},
		{
			desc:          "Case3",
			id:            1,
			expectedError: errors.New("Connection lost"),
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().GetUserById(1).Return(&models.User{}, nil),
				mockUserStore.EXPECT().DeleteUser(1).Return(errors.New("Connection lost")),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := testUserService.DeleteUser(test.id)

			if err != nil && errors.Is(err, test.expectedError) {
				t.Errorf("Expected: %v, Got: %v", test.expectedError, err)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	testUser := models.User{Name: "Ridhdhish", Email: "ridhdhish@gmail.com", Phone: "8320578360", Age: 21}
	returnUser := models.User{
		Id:    1,
		Name:  "Ridhdhish",
		Email: "ridhdhish@gmail.com",
		Phone: "8320578360",
		Age:   21,
	}

	tests := []struct {
		desc          string
		user          models.User
		expectedUser  models.User
		expectedError error
		mockCall      []*gomock.Call
	}{
		{
			desc:          "Case1",
			user:          testUser,
			expectedUser:  returnUser,
			expectedError: nil,
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().GetUserByEmail("ridhdhish@gmail.com").Return(true),
				mockUserStore.EXPECT().CreateUser(testUser).Return(1, nil),
				mockUserStore.EXPECT().GetUserById(1).Return(&returnUser, nil),
			},
		},
		{
			desc:          "Case2",
			user:          testUser,
			expectedUser:  models.User{},
			expectedError: errors.New("Email id is already in exist"),
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().GetUserByEmail("ridhdhish@gmail.com").Return(false),
			},
		},
		{
			desc:          "Case3",
			user:          models.User{},
			expectedUser:  models.User{},
			expectedError: errors.New("Need user data to create new user"),
			mockCall:      nil,
		},
		{
			desc:          "Case4",
			user:          testUser,
			expectedUser:  models.User{},
			expectedError: errors.New("Could not able to create new user"),
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().GetUserByEmail("ridhdhish@gmail.com").Return(true),
				mockUserStore.EXPECT().CreateUser(testUser).Return(0, errors.New("Could not able to create new user")),
			},
		},
		{
			desc:          "Case5",
			user:          models.User{Name: "Ridhdhish", Email: "ridhdhish", Phone: "8320578360", Age: 21},
			expectedUser:  models.User{},
			expectedError: errors.New("Invalid email address"),
			mockCall:      nil,
		},
		{
			desc:          "Case5",
			user:          models.User{Name: "Ridhdhish", Email: "ridhdhish@gmail.com", Phone: "796889", Age: 21},
			expectedUser:  models.User{},
			expectedError: errors.New("Invalid phone number"),
			mockCall:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			user, err := testUserService.CreateUser(test.user)

			if test.expectedError != nil && errors.Is(err, test.expectedError) && !reflect.DeepEqual(user, test.expectedUser) {
				t.Errorf("Expected: %v, Got: %v", test.expectedUser, user)
			}
		})
	}
}
