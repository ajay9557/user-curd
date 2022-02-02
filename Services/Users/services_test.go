package Users

import (
	"Icrud/Stores"
	"Icrud/TModels"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := Stores.NewMockIStore(ctrl)
	testUserService := New(mockUserStore)

	tests := []struct {
		desc     string
		id       int
		expected TModels.User
		mockCall *gomock.Call
	}{
		{
			desc:     "Case1",
			id:       1,
			expected: TModels.User{Id: 1, Name: "Naruto", Email: "naruto@japan.com", Phone: "9999999999", Age: 18},
			mockCall: mockUserStore.EXPECT().UserById(1).Return(TModels.User{Id: 1, Name: "Naruto", Email: "naruto@japan.com", Phone: "9999999999", Age: 18}, nil),
		},
		{
			desc:     "Case2",
			id:       2,
			expected: TModels.User{},
			mockCall: mockUserStore.EXPECT().UserById(2).Return(TModels.User{}, errors.New("Cannot fetch user for given id")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			user, err := testUserService.UserById(test.id)

			if err != nil && !reflect.DeepEqual(test.expected, user) {
				t.Errorf("Expected: %v, Got: %v", test.expected, user)
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := Stores.NewMockIStore(ctrl)
	testUserService := New(mockUserStore)

	data1 := []TModels.User{
		{Id: 1, Name: "Naruto", Email: "naruto@gmail.com", Phone: "9999999999", Age: 18},
		{Id: 2, Name: "Itachi", Email: "itachi@gmail.com", Phone: "8320578360", Age: 24},
	}

	tests := []struct {
		desc     string
		expected []TModels.User
		mockCall *gomock.Call
	}{
		{
			desc:     "Case1",
			expected: data1,
			mockCall: mockUserStore.EXPECT().GetUsers().Return(data1, nil),
		},
		{
			desc:     "Case2",
			expected: []TModels.User{},
			mockCall: mockUserStore.EXPECT().GetUsers().Return([]TModels.User{}, errors.New("Cannot fetch users")),
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

func TestDeleteUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := Stores.NewMockIStore(ctrl)
	testUserService := New(mockUserStore)

	tests := []struct {
		desc     string
		id       int
		expected int
		mockCall *gomock.Call
	}{
		{
			desc:     "Case1",
			id:       1,
			expected: 1,
			mockCall: mockUserStore.EXPECT().DeleteUserById(1).Return(1, nil),
		},
		{
			desc:     "Case2",
			id:       2,
			expected: 0,
			mockCall: mockUserStore.EXPECT().DeleteUserById(2).Return(0, errors.New("Invalid id")),
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			rowsAffected, _ := testUserService.DeleteUserById(test.id)

			if rowsAffected != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, rowsAffected)
			}
		})
	}
}

func TestInsertUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := Stores.NewMockIStore(ctrl)
	testUserService := New(mockUserStore)

	testUser := TModels.User{Name: "Ridhdhish", Email: "ridhdhish@gmail.com", Phone: "8320578360", Age: 21}

	returnUser := TModels.User{
		Id:    1,
		Name:  "Ridhdhish",
		Email: "ridhdhish@gmail.com",
		Phone: "8320578360",
		Age:   21,
	}

	tests := []struct {
		desc string
		user TModels.User
		// expected int
		expectedUser  TModels.User
		expectedError error
		mockCall      []*gomock.Call
	}{
		{
			desc: "Case1",
			user: testUser,
			// expected: 1,
			expectedUser:  returnUser,
			expectedError: nil,
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().InsertUser(testUser).Return(1, nil),
				mockUserStore.EXPECT().GetEmail("ridhdhish@gmail.com").Return(false, nil),
				mockUserStore.EXPECT().UserById(1).Return(returnUser, nil),
			},
		},
		{
			desc: "Case2",
			user: testUser,
			// expected: 0,
			expectedUser:  TModels.User{},
			expectedError: errors.New("Email id is already in exist"),
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().GetEmail("ridhdhish@gmail.com").Return(true, nil),
			},
		},
		// {
		// 	desc:     "Case3",
		// 	user:     TModels.User{},
		// 	expected: 0,
		// 	mockCall: nil,
		// },
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			// lastInsertedId, _ := testUserService.InsertUser(test.user)
			user, err := testUserService.InsertUser(test.user)

			// if lastInsertedId != test.expected {
			// 	t.Errorf("Expected: %v, Got: %v", test.expected, lastInsertedId)
			// }
			if test.expectedError != nil && errors.Is(err, test.expectedError) && !reflect.DeepEqual(user, test.expectedUser) {
				t.Errorf("Expected: %v, Got: %v", test.expectedUser, user)
			}
		})
	}
}

func TestUpdateUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := Stores.NewMockIStore(ctrl)
	testUserService := New(mockUserStore)

	testUser := TModels.User{Name: "Ridhdhish", Email: "ridhdhish@gmail.com", Phone: "8320578360", Age: 21}
	expectedUser := TModels.User{
		Id:    1,
		Name:  "Ridhdhish",
		Email: "ridhdhish@gmail.com",
		Phone: "8320578360",
		Age:   21,
	}

	updatedUser := TModels.User{
		Id:    1,
		Name:  "Ridhdhish",
		Email: "ridhdhish@gmail.com",
		Phone: "8320578360",
		Age:   22,
	}

	tests := []struct {
		desc string
		id   int
		// expected int
		expectedUser  TModels.User
		expectedError error
		mockCall      []*gomock.Call
	}{
		{
			desc: "Case1",
			id:   1,
			// expected: 1,
			expectedUser: updatedUser,
			mockCall: []*gomock.Call{
				mockUserStore.EXPECT().UserById(1).Return(expectedUser, nil),
				mockUserStore.EXPECT().UpdateUserById(testUser, 1).Return(1, nil),
				mockUserStore.EXPECT().UserById(1).Return(updatedUser, nil),
			},
		},
		{
			desc: "Case2",
			id:   -1,
			// expected: 0,
			expectedUser:  TModels.User{},
			expectedError: errors.New("User id should be greater than 0"),
			mockCall:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			// lastInsertedId, _ := testUserService.UpdateUserById(testUser, test.id)

			// if lastInsertedId != test.expected {
			// 	t.Errorf("Expected: %v, Got: %v", test.expected, lastInsertedId)
			// }
			user, err := testUserService.UpdateUserById(testUser, test.id)
			if err != nil && err.Error() != test.expectedError.Error() && !reflect.DeepEqual(user, test.expectedUser) {
				t.Errorf("expectedUser: %v, Got: %v", test.expectedUser, user)
			}
		})
	}
}
