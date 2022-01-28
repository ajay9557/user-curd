package service

import (
	reflect "reflect"
	"testing"
	"zopsmart/Task/models"
	store "zopsmart/Task/stores"

	gomock "github.com/golang/mock/gomock"
)

func TestGetAllUsersService(t*testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	testCases := []struct {
		caseId      int
		expectedErr error
		expectedOut []models.User
	}{

		{
			caseId:      1,
			expectedErr: nil,
			expectedOut: []models.User{
				{
					Id:    1,
					Name:  "prasath",
					Email: "prasath@example.com",
					Phone: "12345",
					Age:   20,
				},
				{
					Id:    2,
					Name:  "rishi kumar",
					Email: "rishi.kumar@example.com",
					Phone: "67890",
					Age:   25,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run("running success test case ", func(t *testing.T) {
			mockStore.EXPECT().GetAllUsersStore().Return(tc.expectedOut, tc.expectedErr)
			resp, err := mock.GetAllUsersService()
			if tc.expectedOut != nil && !reflect.DeepEqual(resp, tc.expectedOut) {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedOut, resp)
			}
			if err !=tc.expectedErr {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedErr, err)
			}
		})
	}
}
func TestValidateEmail(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()
	mockStore := store.NewMockStore(ctrl)
	mock := New(mockStore)

	testCases := []struct {
		Email     string
		mockCall  *gomock.Call
		expectRes bool
		expectErr error
	}{
		{
			Email:     "prasath@gmail.com",
			mockCall:  mockStore.EXPECT().GetMail("prasath@gmail.com").Return(true, nil),
			expectRes: true,
			expectErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run("Sucess test case ", func(t *testing.T) {
			res, err := mock.ValidateEmail(tc.Email)
			if err != nil && err != tc.expectErr {
				t.Errorf("Expected: %v but got %v", tc.expectErr, err)
			}
			if !reflect.DeepEqual(res, tc.expectRes) {
				t.Errorf("Expected: %v but got %v", tc.expectRes, res)
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
		Id        int
		mock      *gomock.Call
		expectErr error
	}{
		{
			Id:        1,
			mock:      mockStore.EXPECT().Delete(1).Return(nil),
			expectErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run("Sucess test case ", func(t *testing.T) {
			err := mock.DeletebyId(tc.Id)
			if err != nil && err != tc.expectErr {
				t.Errorf("Expected: %v but got %v", tc.expectErr, err)
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
		Id        int
		Phone     string
		mock      *gomock.Call
		expectErr error
	}{
		{
			Id:        1,
			Phone:     "12345",
			mock:      mockStore.EXPECT().Update(1, "12345").Return(nil),
			expectErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run("Sucess test case ", func(t *testing.T) {
			err := mock.UpdatebyId(tc.Id, tc.Phone)
			if err != nil && (err != tc.expectErr) {
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
		Id        int
		mock      *gomock.Call
		expectOut models.User
		expectErr error
	}{
		{
			Id: 1,
			mock: mockStore.EXPECT().GetUserById(1).Return(models.User{
				Id:    1,
				Name:  "prasath",
				Email: "prasath@gmail.com",
				Phone: "12345",
				Age:   20,
			}, nil),
			expectErr: nil,
			expectOut: models.User{
				Id:    1,
				Name:  "prasath",
				Email: "prasath@gmail.com",
				Phone: "12345",
				Age:   20},
		},
	}

	for _, tc := range testCases {
		t.Run("Sucess test case ", func(t *testing.T) {
			res, err := mock.GetUserById(tc.Id)
			if err != nil && (err != tc.expectErr) {
				t.Errorf("Expected %v: But got %v", tc.expectErr, err)
			}

			if !reflect.DeepEqual(res, tc.expectOut) {
				t.Errorf("Expected %v: But got %v", tc.expectOut, res)
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
		Id        int
		expectOut bool
		expectErr error
	}{
		{
			Id:        1,
			expectOut: true,
			expectErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run("Sucess test case ", func(t *testing.T) {
			res, err := mock.ValidateId(tc.Id)
			if err != nil && (err != tc.expectErr) {
				t.Error("Expected: ", tc.expectErr, "but Got: ", err)
			}

			if !reflect.DeepEqual(res, tc.expectOut) {
				t.Error("Expected :", tc.expectOut, "But Got:", res)
			}
		})
	}

}
