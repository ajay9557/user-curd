package users

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tejas/user-crud/models"
	"github.com/tejas/user-crud/stores"
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
			desc: "Case 1: Success case",
			id:   1,
			expected: models.User{
				Id:    1,
				Name:  "tejas",
				Email: "tejas@gmail.com",
				Phone: "9876543230",
				Age:   23,
			},
			mockCall: mockUserStore.EXPECT().GetUserById(1).Return(&models.User{Id: 1, Name: "tejas", Email: "tejas@gmail.com", Phone: "9876543230", Age: 23}, nil),
		},
		{
			desc:     "Case 2: Failure case",
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

	data1 := []models.User{
		{
			Id:    1,
			Name:  "tejas",
			Email: "tejas@gmail.com",
			Phone: "9876543230",
			Age:   23,
		},
		{
			Id:    2,
			Name:  "nithin",
			Email: "nithin@gmail.com",
			Phone: "9988776655",
			Age:   18,
		},
	}

	tests := []struct {
		desc     string
		expected []models.User
		mockCall *gomock.Call
	}{
		{
			desc:     "Case 1: Success case",
			expected: data1,
			mockCall: mockUserStore.EXPECT().GetUsers().Return(data1, nil),
		},
		{
			desc:     "Case 2: Failure case",
			expected: []models.User{},
			mockCall: mockUserStore.EXPECT().GetUsers().Return([]models.User{}, errors.New("Cannot fetch users")),
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

	testUser := models.User{Name: "vivek", Email: "vivek@gmail.com", Phone: "9988776655", Age: 23}

	tests := []struct {
		desc     string
		id       int
		expected int
		mockCall *gomock.Call
	}{
		{
			desc:     "Case 1: Success case",
			id:       1,
			expected: 1,
			mockCall: mockUserStore.EXPECT().UpdateUser(1, testUser).Return(1, nil),
		},
		{
			desc:     "Case 2: Failure case",
			id:       2,
			expected: 0,
			mockCall: mockUserStore.EXPECT().UpdateUser(2, testUser).Return(0, errors.New("Invalid id")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			lastInsertedId, _ := testUserService.UpdateUser(test.id, testUser)

			if lastInsertedId != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, lastInsertedId)
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
		desc     string
		id       int
		expected int
		mockCall *gomock.Call
	}{
		{
			desc:     "Case 1: Success case",
			id:       1,
			expected: 1,
			mockCall: mockUserStore.EXPECT().DeleteUser(1).Return(1, nil),
		},
		{
			desc:     "Case 2: Failure case",
			id:       2,
			expected: 0,
			mockCall: mockUserStore.EXPECT().DeleteUser(2).Return(0, errors.New("Invalid id")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			rowsAffected, _ := testUserService.DeleteUser(test.id)

			if rowsAffected != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, rowsAffected)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := stores.NewMockUser(ctrl)
	testUserService := New(mockUserStore)

	testUser := models.User{Name: "vivek", Email: "vivek@gmail.com", Phone: "9988776655", Age: 23}

	tests := []struct {
		desc     string
		user     models.User
		expected int
		mockCall []*gomock.Call
	}{
		{
			desc:     "Case 1",
			user:     testUser,
			expected: 1,
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().CreateUser(testUser).Return(1, nil),
				mockUserStore.EXPECT().GetUserByEmail("vivek@gmail.com").Return(true),
			},
		},
		{
			desc:     "Case 2",
			user:     testUser,
			expected: 0,
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().GetUserByEmail("vivek@gmail.com").Return(false),
			},
		},
		{
			desc:     "Case 3",
			user:     models.User{},
			expected: 0,
			mockCall: nil,
		},
		{
			desc:     "Case 4",
			user:     testUser,
			expected: 0,
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().CreateUser(testUser).Return(0, errors.New("Could not able to create new user")),
				mockUserStore.EXPECT().GetUserByEmail("vivek@gmail.com").Return(true),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			lastInsertedId, _ := testUserService.CreateUser(test.user)

			if lastInsertedId != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, lastInsertedId)
			}
		})
	}
}
