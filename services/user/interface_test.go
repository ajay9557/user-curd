package user

import (
	"errors"
	"go_lang/Assignment/user-curd/models"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_AddUser(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	userMockService := NewMockServices(controller)

	testCase := []struct {
		desc     string
		input    models.User
		expected error
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case 1",
			input:    models.User{11, "Ram Swami", "ramswami43@gmail.com", "7827878984", 54},
			expected: nil,
			mock:     []*gomock.Call{userMockService.EXPECT().AddUser(models.User{11, "Ram Swami", "ramswami43@gmail.com", "7827878984", 54}).Return(nil)},
		},
		{
			desc:     "Test Case 2",
			input:    models.User{11, "Ram Swami", "ramswami43@gmail.com", "7878984", 54},
			expected: errors.New("INVALID PHONE NUMBER"),
			mock:     []*gomock.Call{userMockService.EXPECT().AddUser(models.User{11, "Ram Swami", "ramswami43@gmail.com", "7878984", 54}).Return(errors.New("INVALID PHONE NUMBER"))},
		},
		{
			desc:     "Test Case 3",
			input:    models.User{18, "Anna Swapi", "anna.com", "7878789845", 34},
			expected: errors.New("INVALID EMAIL"),
			mock:     []*gomock.Call{userMockService.EXPECT().AddUser(models.User{18, "Anna Swapi", "anna.com", "7878789845", 34}).Return(errors.New("INVALID EMAIL"))},
		},
	}

	for _, tcs := range testCase {
		t.Run(tcs.desc, func(t *testing.T) {
			result := userMockService.AddUser(tcs.input)
			if !reflect.DeepEqual(tcs.expected, result) {
				t.Errorf("output: %v, expected: %v", result, tcs.expected)
			}
		})
	}
}

func Test_GetUser(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	userMockService := NewMockServices(controller)

	testCase := []struct {
		desc     string
		input    int
		expected interface{}
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case 1",
			input:    10,
			expected: models.User{10, "Ram Swami", "ramswami43@gmail.com", "7827878984", 54},
			mock:     []*gomock.Call{userMockService.EXPECT().GetUser(10).Return(models.User{10, "Ram Swami", "ramswami43@gmail.com", "7827878984", 54}, nil)},
		},
		{
			desc:     "Test Case 2",
			input:    1000,
			expected: models.User{},
			mock:     []*gomock.Call{userMockService.EXPECT().GetUser(1000).Return(models.User{}, nil)},
		},
	}

	for _, tcs := range testCase {
		t.Run(tcs.desc, func(t *testing.T) {
			result, err := userMockService.GetUser(tcs.input)
			if err != nil && !reflect.DeepEqual(tcs.expected, result) {
				t.Errorf("output: %v, expected: %v", result, tcs.expected)
			}
		})
	}
}

func Test_GetAllUser(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	userMockService := NewMockServices(controller)

	testCase := []struct {
		desc     string
		expected interface{}
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case 1",
			expected: []models.User{{10, "Ram Swami", "ramswami43@gmail.com", "7827878984", 54}},
			mock:     []*gomock.Call{userMockService.EXPECT().GetAllUser().Return([]models.User{{10, "Ram Swami", "ramswami43@gmail.com", "7827878984", 54}}, nil)},
		},
		{
			desc:     "Test Case 2",
			expected: []models.User{{}},
			mock:     []*gomock.Call{userMockService.EXPECT().GetAllUser().Return([]models.User{{}}, nil)},
		},
	}

	for _, tcs := range testCase {
		t.Run(tcs.desc, func(t *testing.T) {
			result, err := userMockService.GetAllUser()
			if err != nil && !reflect.DeepEqual(tcs.expected, result) {
				t.Errorf("output: %v, expected: %v", result, tcs.expected)
			}
		})
	}
}

func Test_UpdateUser(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	userMockService := NewMockServices(controller)

	testCase := []struct {
		desc     string
		input    models.User
		expected interface{}
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case 1",
			input:    models.User{10, "Ram Swami", "ramswami43@gmail.com", "7827878984", 54},
			expected: nil,
			mock:     []*gomock.Call{userMockService.EXPECT().UpdateUser(models.User{10, "Ram Swami", "ramswami43@gmail.com", "7827878984", 54}).Return(nil)},
		},
		{
			desc:     "Test Case 2",
			input:    models.User{1000000000, "Ram Swami", "ramswami43@gmail.com", "7827878984", 54},
			expected: errors.New("FAILED TO UPDATE USER DATA"),
			mock:     []*gomock.Call{userMockService.EXPECT().UpdateUser(models.User{1000000000, "Ram Swami", "ramswami43@gmail.com", "7827878984", 54}).Return(errors.New("FAILED TO UPDATE USER DATA"))},
		},
	}

	for _, tcs := range testCase {
		t.Run(tcs.desc, func(t *testing.T) {
			err := userMockService.UpdateUser(tcs.input)
			if err != nil && !reflect.DeepEqual(tcs.expected, err) {
				t.Errorf("output: %v, expected: %v", err, tcs.expected)
			}
		})
	}
}

func Test_DeleteUser(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	userMockService := NewMockServices(controller)

	testCase := []struct {
		desc     string
		input    int
		expected interface{}
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case 1",
			input:    10,
			expected: nil,
			mock:     []*gomock.Call{userMockService.EXPECT().DeleteUser(10).Return(nil)},
		},
		{
			desc:     "Test Case 2",
			input:    10000000000000,
			expected: errors.New("FAILED TO DELETE USER"),
			mock:     []*gomock.Call{userMockService.EXPECT().DeleteUser(10000000000000).Return(errors.New("FAILED TO DELETE USER"))},
		},
	}

	for _, tcs := range testCase {
		t.Run(tcs.desc, func(t *testing.T) {
			err := userMockService.DeleteUser(tcs.input)
			if err != nil && !reflect.DeepEqual(tcs.expected, err) {
				t.Errorf("output: %v, expected: %v", err, tcs.expected)
			}
		})
	}
}
