package users

import (
	"user-curd/models"
	"user-curd/stores"
	"errors"
	"reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc string
		Id   int
		mock []*gomock.Call
		err  error
	}{
		{
			desc: "Success",
			Id:   1,
			mock: []*gomock.Call{mockStore.EXPECT().DeleteById(gomock.Any()).Return(nil)},
			err:  nil,
		},
		{
			desc: "Failure",
			Id:   7,
			mock: []*gomock.Call{mockStore.EXPECT().DeleteById(gomock.Any()).Return(errors.New("t"))},
			err:  errors.New("t"),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			err := mock.DeleteByUserId(tc.Id)
			if err != nil && !reflect.DeepEqual(err, tc.err) {
				t.Error("Expected: ", tc.err, "Obtained: ", err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc  string
		input models.User
		mock  []*gomock.Call
		err   error
	}{
		{
			desc: "Success",
			input: models.User{
				Id:    2,
				Name:  "Zopsmart",
				Email: "sudheerpuppala@gmail.com",
				Phone: "+919908577405",
				Age:   12,
			},
			mock: []*gomock.Call{mockStore.EXPECT().UpdateById(gomock.Any()).Return(nil)},
			err:  nil,
		},
		{
			desc: "Failure",
			input: models.User{
				Id:    3,
				Name:  "test",
				Email: "sudhee@gmail.com",
				Phone: "+919908577405",
				Age:   12,
			},
			mock: []*gomock.Call{mockStore.EXPECT().UpdateById(gomock.Any()).Return(errors.New("t"))},
			err:  errors.New("t"),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			err := mock.UpdateByUserId(tc.input)
			if err != nil && !reflect.DeepEqual(err, tc.err) {
				t.Error("Expected: ", tc.err, "Obtained: ", err)
			}
		})
	}
}

func TestSearchById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc     string
		Id       int
		expected models.User
		mock     []*gomock.Call
		err      error
	}{
		{
			desc: "Success",
			Id:   1,
			expected: models.User{
				Id:    1,
				Name:  "Puppala",
				Email: "puppala@gmail.com",
				Phone: "919908577405",
				Age:   22,
			},
			mock: []*gomock.Call{mockStore.EXPECT().SearchById(gomock.Any()).Return(models.User{1, "Puppala", "puppala@gmail.com", "919908577405", 22}, nil)},
			err:  nil,
		},
		{
			desc: "Failure",
			Id:   1,
			expected: models.User{
				Id:    1,
				Name:  "Puppala",
				Email: "puppala@gmail.com",
				Phone: "919908577405",
				Age:   22,
			},
			mock: []*gomock.Call{mockStore.EXPECT().SearchById(gomock.Any()).Return(models.User{1, "Puppala", "puppala@gmail.com", "919908577405", 22}, nil)},
			err:  errors.New("t"),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.SearchByUserId(tc.Id)
			if err != nil && !reflect.DeepEqual(tc.err, err) {
				t.Error("Expected: ", tc.err, "Obtained: ", err)
			}
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expected, res)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc     string
		expected []models.User
		mock     []*gomock.Call
		err      error
	}{
		{
			desc: "Success",
			expected: []models.User{
				{Id: 2,
					Name:  "Zopsmart",
					Email: "sudheerpuppala@gmail.com",
					Phone: "+919908577405",
					Age:   12,
				},
			},
			mock: []*gomock.Call{mockStore.EXPECT().SearchDetails().Return([]models.User{
				{Id: 2,
					Name:  "Zopsmart",
					Email: "sudheerpuppala@gmail.com",
					Phone: "+919908577405",
					Age:   12,
				},
			}, nil)},
			err: nil,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.SearchAll()
			if err != nil && !reflect.DeepEqual(err, tc.err) {
				t.Error("Expected: ", tc.err, "Obtained: ", err)
			}
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expected, res)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc     string
		input    models.User
		expected models.User
		mock     []*gomock.Call
		err      error
	}{
		{
			desc: "Success",
			input: models.User{
				Id:    1,
				Name:  "Puppala",
				Email: "puppala@gmail.com",
				Phone: "919908577405",
				Age:   22,
			},
			expected: models.User{
				Id:    1,
				Name:  "Puppala",
				Email: "puppala@gmail.com",
				Phone: "919908577405",
				Age:   22,
			},
			mock: []*gomock.Call{mockStore.EXPECT().InsertDetails(gomock.Any()).Return(nil)},
			err:  nil,
		},
		{
			desc: "Failure",
			input: models.User{
				Id:    1,
				Name:  "Puppala",
				Email: "puppala@gmail.com",
				Phone: "919908577405",
				Age:   22,
			},
			expected: models.User{
				Id:    1,
				Name:  "Puppala",
				Email: "puppala@gmail.com",
				Phone: "919908577405",
				Age:   22,
			},
			mock: []*gomock.Call{mockStore.EXPECT().InsertDetails(gomock.Any()).Return(errors.New("t"))},
			err:  errors.New("t"),
		},
	}
	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := mock.InsertUserDetails(tc.input)
			if err != nil && !reflect.DeepEqual(err, tc.err) {
				t.Error("Expected: ", tc.err, "Obtained: ", err)
			}
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expected, res)
			}
		})
	}
}

func TestEmailValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := stores.NewMockStore(ctrl)
	mock := New(mockStore)

	tcs := []struct {
		desc     string
		Email    string
		input    []models.User
		mock     []*gomock.Call
		expected bool
		err      error
	}{
		{
			desc:  "Success",
			Email: "minnu01@gmail.com",
			input: []models.User{
				{Id: 2,
					Name:  "Zopsmart",
					Email: "sudheerpuppala@gmail.com",
					Phone: "+919908577405",
					Age:   12,
				},
			},
			mock: []*gomock.Call{mockStore.EXPECT().SearchDetails().Return([]models.User{
				{Id: 2,
					Name:  "Zopsmart",
					Email: "sudheerpuppala@gmail.com",
					Phone: "+919908577405",
					Age:   12,
				},
			}, nil)},
			expected: true,
			err:      nil,
		},
		{
			desc:  "Success-3",
			Email: "sudheerpuppala@gmail.com",
			input: []models.User{
				{Id: 2,
					Name:  "Zopsmart",
					Email: "sudheerpuppala@gmail.com",
					Phone: "+919908577405",
					Age:   12,
				},
				{
					Id:    3,
					Name:  "Zopsmart",
					Email: "sudheerpuppala@gmail.com",
					Phone: "+919908577405",
					Age:   12,
				},
			},
			mock: []*gomock.Call{mockStore.EXPECT().SearchDetails().Return([]models.User{
				{Id: 2,
					Name:  "Zopsmart",
					Email: "sudheerpuppala@gmail.com",
					Phone: "+919908577405",
					Age:   12,
				},
				{Id: 3,
					Name:  "Zopsmart",
					Email: "sudheerpuppala@gmail.com",
					Phone: "+919908577405",
					Age:   12,
				},
			}, nil)},
			expected: false,
			err:      nil,
		},
		{
			desc:  "Failure",
			Email: "minnu01@gmail.com",
			input: []models.User{
				{Id: 2,
					Name:  "Zopsmart",
					Email: "sudheerpuppala@gmail.com",
					Phone: "+919908577405",
					Age:   12,
				},
			},
			mock: []*gomock.Call{mockStore.EXPECT().SearchDetails().Return([]models.User{
				{Id: 2,
					Name:  "Zopsmart",
					Email: "sudheerpuppala@gmail.com",
					Phone: "+919908577405",
					Age:   12,
				},
			}, errors.New("t"))},
			expected: false,
			err:      errors.New("t"),
		},
		{
			desc:     "Failure=2",
			Email:    "mi",
			expected: false,
			err:      errors.New("t"),
		},
		{
			desc:     "Failure-3",
			Email:    "minnugmail.com",
			expected: false,
			err:      errors.New("t"),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			res := mock.IsEmailValid(tc.Email)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Error("Expected: ", tc.expected, "Obtained: ", res)
			}
		})
	}
}
