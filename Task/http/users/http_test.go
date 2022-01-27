package http

import (
	"zopsmart/Task/models"
	"zopsmart/Task/services"

	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"

)

func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := services.NewMockServices(ctrl)
	mock := New(mockService)
	testCases := []struct {
		id        int
		mock      []*gomock.Call
		expecErr  error
		expecBody []byte
	}{
		{
			id:       1,
			expecErr: nil,
			mock: []*gomock.Call{
				mockService.EXPECT().GetUserById(1).Return(models.User{
					Id:    1,
					Name:  "prasath",
					Email: "prasath@gmail.com",
					Phone: "12345",
					Age:   20,
				}, nil),
			},
			expecBody: []byte(`{"Id":1,"Name":"prasath","Email":"prasath@gmail.com","Phone":"12345","Age":20}`),
		},
	}

	for _,tc := range testCases {
		req := httptest.NewRequest("GET","/user="+tc.id,nil)
		wr := httptest.NewRecorder()
		mock.GetUserById(req,wr)
		if wr.Body.String()!= string(tc.expecBody) {
			t.Errorf("Expected %v but got %v",string(tc.expecBody),wr.Body.String())
		}
	}

}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := services.NewMockServices(ctrl)
	mock := New(mockService)

	testCases := []struct {
		id        int
		mock      []*gomock.Call
		expectedErr  error
		expectedOut []byte
	}{
		{
			id:       1,
			expectedErr: nil,
			mock: []*gomock.Call{
				mockService.EXPECT().DeleteUser(1).Return(nil),
			},
			expectedOut: []byte(`{user deleted}`),
		},
	}
	for _,tc := range testCases {
		t.Run("",func (t *testing.T)  {
			req:= httptest.NewRequest("DELETE", "/delete?id="+tc.id, nil)
			writer := httptest.NewRecorder()
			mock.DeleteUser(writer, req)
			fmt.Println(writer.Body.String())
			if writer.Body.String() != string(tc.expectedOut) {
				t.Errorf("Expected %v Obtained %v", string(tc.expectedOut), writer.Body.String())
			}
			
		})
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := services.NewMockServices(ctrl)
	mock := New(mockService)

	testCases := []struct {
		id        int
		mock      []*gomock.Call
		expectedErr  error
		expectedOut []byte
	}{
		{
			id:       1,
			expectedErr: nil,
			mock: []*gomock.Call{
				mockService.EXPECT().DeleteUser(1).Return(nil),
			},
			expectedOut: []byte(`{user updated}`),
		},
	}
	for _,tc := range testCases {
		t.Run("",func (t *testing.T)  {
			req:= httptest.NewRequest("UPDATE", "/update?id="+tc.id, nil)
			writer := httptest.NewRecorder()
			mock.DeleteUser(writer, req)
			fmt.Println(writer.Body.String())
			if writer.Body.String() != string(tc.expectedOut) {
				t.Errorf("Expected %v Obtained %v", string(tc.expectedOut), writer.Body.String())
			}
			
		})
	}

}


func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	 mockService := services.NewMockServices(ctrl)
	 mock := New(mockService)

	testCases := []struct {
		id int
		mock []*gomock.Call
		expectedErr error
		expectedOut []byte

	}{
		{
			id:       1,
			expectedErr: nil,
			mock: []*gomock.Call{
				mockService.EXPECT().DeleteUser(1).Return(nil),
			},
			expectedOut: []byte(`{user created}`),
		},
	}

	for _,tc := range testCases {
		t.Run("",func(t *testing.T) {
			req:= httptest.NewRequest("INSERT", "/insert="+tc.id, nil)
			writer := httptest.NewRecorder()
			mock.DeleteUser(writer, req)
			fmt.Println(writer.Body.String())
			if writer.Body.String() != string(tc.expectedOut) {
				t.Errorf("Expected %v Obtained %v", string(tc.expectedOut), writer.Body.String())
			}

		})
	}
}
