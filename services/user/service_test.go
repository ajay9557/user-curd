package user

import (
	"errors"
	"go_lang/Assignment/user-curd/models"
	"go_lang/Assignment/user-curd/stores"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_AddUser(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockUserStore := stores.NewMockCrud(controller)
	userMockUser := New(mockUserStore)

	testCase := []struct {
		desc     string
		input    models.User
		name     string
		email    string
		phone    string
		age      int
		expected error
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case 1",
			input:    models.User{Id: 10, Name: "Ram Swami", Email: "ramswami43@gmail.com", Phone: "7827878984", Age: 54},
			expected: nil,
			mock:     []*gomock.Call{mockUserStore.EXPECT().Create("Ram Swami", "ramswami43@gmail.com", "7827878984", 54).Return(nil)},
		},
		{
			desc:     "Test Case 2",
			input:    models.User{Id: 11, Name: "Ram Swami", Email: "ramswi43@gmail.com", Phone: "78278", Age: 54},
			expected: errors.New("INVALID PHONE NUMBER"),
			mock:     nil,
		},
		{
			desc:     "Test Case 3",
			input:    models.User{Id: 18, Name: "Anna Swapi", Email: "anna.com", Phone: "7878789845", Age: 34},
			expected: errors.New("INVALID EMAIL"),
			mock:     nil,
		},
		{
			desc:     "Test Case 4",
			input:    models.User{Id: 10, Name: "Ram Swami", Email: "ramswami43@gmail.com", Phone: "7827878984", Age: 54},
			expected: errors.New("FAILED TO ADD USER"),
			mock:     []*gomock.Call{mockUserStore.EXPECT().Create("Ram Swami", "ramswami43@gmail.com", "7827878984", 54).Return(errors.New("FAILED TO ADD USER"))},
		},
	}

	for _, tcs := range testCase {
		t.Run(tcs.desc, func(t *testing.T) {
			result := userMockUser.AddUser(tcs.input)
			if !reflect.DeepEqual(tcs.expected, result) {
				t.Errorf("output: %v, expected: %v", result, tcs.expected)
			}
		})
	}
}

func Test_GetUser(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockUserStore := stores.NewMockCrud(controller)
	userMockUser := New(mockUserStore)

	testCase := []struct {
		desc     string
		input    int
		expected interface{}
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case 1",
			input:    10,
			expected: models.User{Id: 10, Name: "Ram Swami", Email: "ramswami43@gmail.com", Phone: "7827878984", Age: 54},
			mock:     []*gomock.Call{mockUserStore.EXPECT().ReadOne(10).Return(models.User{Id: 10, Name: "Ram Swami", Email: "ramswami43@gmail.com", Phone: "7827878984", Age: 54}, nil)},
		},
		{
			desc:     "Test Case 2",
			input:    1000,
			expected: models.User{},
			mock:     []*gomock.Call{mockUserStore.EXPECT().ReadOne(1000).Return(models.User{}, errors.New("ERROR IN FETCHING ROW"))},
		},
	}

	for _, tcs := range testCase {
		t.Run(tcs.desc, func(t *testing.T) {
			result, err := userMockUser.GetUser(tcs.input)
			if err != nil && !reflect.DeepEqual(tcs.expected, result) {
				t.Errorf("output: %v, expected: %v", result, tcs.expected)
			}
		})
	}
}

func Test_GetAllUser(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockUserStore := stores.NewMockCrud(controller)
	userMockUser := New(mockUserStore)

	testCase := []struct {
		desc     string
		expected interface{}
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case 1",
			expected: []models.User{{Id: 10, Name: "Ram Swami", Email: "ramswami43@gmail.com", Phone: "7827878984", Age: 54}},
			mock:     []*gomock.Call{mockUserStore.EXPECT().ReadAll().Return([]models.User{{Id: 10, Name: "Ram Swami", Email: "ramswami43@gmail.com", Phone: "7827878984", Age: 54}}, nil)},
		},
		{
			desc:     "Test Case 2",
			expected: []models.User{{}},
			mock:     []*gomock.Call{mockUserStore.EXPECT().ReadAll().Return([]models.User{{}}, nil)},
		},
	}

	for _, tcs := range testCase {
		t.Run(tcs.desc, func(t *testing.T) {
			result, err := userMockUser.GetAllUser()
			if err != nil && !reflect.DeepEqual(tcs.expected, result) {
				t.Errorf("output: %v, expected: %v", result, tcs.expected)
			}
		})
	}
}

func Test_UpdateUser(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockUserStore := stores.NewMockCrud(controller)
	userMockUser := New(mockUserStore)

	testCase := []struct {
		desc     string
		input    models.User
		expected interface{}
		mock     []*gomock.Call
	}{
		{
			desc:     "Test Case 1",
			input:    models.User{Id: 10, Name: "Ram Swami", Email: "ramswami43@gmail.com", Phone: "7827878984", Age: 54},
			expected: nil,
			mock:     []*gomock.Call{mockUserStore.EXPECT().Update(10, "Ram Swami", "ramswami43@gmail.com", "7827878984", 54).Return(nil)},
		},
		{
			desc:     "Test Case 2",
			input:    models.User{Id: 10, Name: "Ram Swami", Email: "ramswami43@gmail.com", Phone: "78278", Age: 54},
			expected: errors.New("INVALID PHONE NUMBER"),
			mock:     nil,
		},
		{
			desc:     "Test Case 3",
			input:    models.User{Id: 10, Name: "Ram Swami", Email: "ramswami4.com", Phone: "7827878984", Age: 54},
			expected: errors.New("INVALID EMAIL"),
			mock:     nil,
		},
		{
			desc:     "Test Case 4",
			input:    models.User{Id: 10000, Name: "Ram Swami", Email: "ramswami43@gmail.com", Phone: "7827878984", Age: 54},
			expected: errors.New("FAILED TO UPDATE USER DATA"),
			mock:     []*gomock.Call{mockUserStore.EXPECT().Update(10000, "Ram Swami", "ramswami43@gmail.com", "7827878984", 54).Return(errors.New("FAILED TO UPDATE USER DATA"))},
		},
	}

	for _, tcs := range testCase {
		t.Run(tcs.desc, func(t *testing.T) {
			err := userMockUser.UpdateUser(tcs.input)
			if err != nil && !reflect.DeepEqual(tcs.expected, err) {
				t.Errorf("output: %v, expected: %v", err, tcs.expected)
			}
		})
	}
}

func Test_DeleteUser(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()
	mockUserStore := stores.NewMockCrud(controller)
	userMockUser := New(mockUserStore)

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
			mock:     []*gomock.Call{mockUserStore.EXPECT().Delete(10).Return(nil)},
		},
		{
			desc:     "Test Case 2",
			input:    1000,
			expected: errors.New("FAILED TO DELETE USER"),
			mock:     []*gomock.Call{mockUserStore.EXPECT().Delete(1000).Return(errors.New("FAILED TO DELETE USER"))},
		},
	}

	for _, tcs := range testCase {
		t.Run(tcs.desc, func(t *testing.T) {
			err := userMockUser.DeleteUser(tcs.input)
			if err != nil && !reflect.DeepEqual(tcs.expected, err) {
				t.Errorf("output: %v, expected: %v", err, tcs.expected)
			}
		})
	}
}
