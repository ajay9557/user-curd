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
			err := mock.DeletebyIdService(tc.Id)
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
		Input    models.User
		mock      *gomock.Call
		expectErr error
	}{
		{
			Id:        1,
			Input:     models.User{
				Id : 1,
				Name : "andrew",
				Phone : "36785",
				Email : "andrew@gmail.com",
				Age : 21,

			},
			mock:      mockStore.EXPECT().Update(models.User{
				Id : 1,
				Name : "andrew",
				Phone : "36785",
				Email : "andrew@gmail.com",
				Age : 21,

			}).Return(nil),
			expectErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run("Sucess test case ", func(t *testing.T) {
			err := mock.UpdatebyIdService(tc.Input)
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
			res, err := mock.GetUserByIdService(tc.Id)
			if err != nil && (err != tc.expectErr) {
				t.Errorf("Expected %v: But got %v", tc.expectErr, err)
			}

			if !reflect.DeepEqual(res, tc.expectOut) {
				t.Errorf("Expected %v: But got %v", tc.expectOut, res)
			}
		})
	}

}

