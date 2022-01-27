package users

import (
	"user-curd/models"
	"errors"

	//"fmt"

	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_Insert(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("database error :%s", err)
	}

	tcsInsert := []struct {
		desc string
		user models.User
		Mock []interface{}
		err  error
	}{
		{
			desc: "Success",
			user: models.User{
				Id:    1,
				Name:  "test",
				Email: "sudheer01@gmail.com",
				Phone: "+919908577405",
				Age:   22,
			},
			Mock: []interface{}{
				mock.ExpectExec(`INSERT INTO User VALUES`).WithArgs(1, "test", "sudheer01@gmail.com", "+919908577405", 22).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			err: nil,
		},
		{
			desc: "Failure",
			user: models.User{
				Id:    2,
				Name:  "Zopsmart",
				Email: "sudheerpuppala@gmail.com",
				Phone: "+919908577405",
				Age:   12,
			},
			Mock: []interface{}{
				mock.ExpectExec("INSERT INTO User VALUES").WithArgs(2, "Zopsmart", "sudheerpuppala@gmail.com", "+919908577405", 12).WillReturnError(errors.New("t")),
			},
			err: errors.New("t"),
		},
	}
	handler := New(db)
	for _, tc := range tcsInsert {
		t.Run(tc.desc, func(t *testing.T) {
			err := handler.InsertDetails(tc.user)
			if err != nil && !reflect.DeepEqual(tc.err, err) {
				t.Errorf("Expected : %s,Obtained : %s ", tc.err, err)
			}
		})
	}
}
func Test_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("database error :%s", err)
	}
	tcsUpdate := []struct {
		desc  string
		input models.User
		Mock  []interface{}
		err   error
	}{
		{
			desc: "Success",
			input: models.User{
				Id:    1,
				Name:  "test",
				Email: "sudheer01@gmail.com",
				Phone: "+919908577405",
				Age:   22,
			},
			Mock: []interface{}{
				mock.ExpectExec(`update User set`).WithArgs("test", "sudheer01@gmail.com", "+919908577405", 22, 1).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			err: nil,
		},
		{
			desc: "Failure",
			input: models.User{
				Id:    2,
				Name:  "Zopsmart",
				Email: "sudheerpuppala@gmail.com",
				Phone: "+919908577405",
				Age:   12,
			},
			Mock: []interface{}{
				mock.ExpectExec(`update User set`).WithArgs("Zopsmart", "sudheerpuppala@gmail.com", "+919908577405", 12, 2).WillReturnError(errors.New("t")),
			},
			err: errors.New("t"),
		},
	}
	handler := New(db)
	for _, tc := range tcsUpdate {
		t.Run(tc.desc, func(t *testing.T) {
			err := handler.UpdateById(tc.input)
			if err != nil && !reflect.DeepEqual(tc.err, err) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.err, err)
			}
		})
	}
}

func Test_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("database error :%s", err)
	}

	tcsDelete := []struct {
		desc  string
		Id    int
		Name  string
		Email string
		Phone string
		Age   int
		Mock  []interface{}
		err   error
	}{
		{
			desc:  "Success",
			Id:    1,
			Name:  "test",
			Email: "sudheer01@gmail.com",
			Phone: "+919908577405",
			Age:   22,
			Mock: []interface{}{
				mock.ExpectExec(`delete from User`).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			err: nil,
		},
		{
			desc:  "Failure",
			Id:    2,
			Name:  "Zopsmart",
			Email: "sudheerpuppala@gmail.com",
			Phone: "+919908577405",
			Age:   12,
			Mock: []interface{}{
				mock.ExpectExec(`delete from User`).WithArgs(2).WillReturnError(errors.New("t")),
			},
			err: errors.New("t"),
		},
	}
	handler := New(db)
	for _, tc := range tcsDelete {
		t.Run(tc.desc, func(t *testing.T) {
			err := handler.DeleteById(tc.Id)
			if tc.Id < 1 {
				return
			}
			if err != nil && !reflect.DeepEqual(tc.err, err) {
				t.Errorf("Expected : %s,Obtained : %s ", tc.err, err)
			}
		})
	}
}

func Test_Search(t *testing.T) {
	query := "SELECT Id,Name,Email,Phone,Age FROM Userwhere Id=?"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("database error :%s", err)
	}
	tcsSearch := []struct {
		desc     string
		Id       int
		expected models.User
		Mock     []interface{}
		err      error
	}{
		{
			desc: "Success",
			Id:   2,
			expected: models.User{
				Id:    2,
				Name:  "Zopsmart",
				Email: "sudheerpuppala@gmail.com",
				Phone: "+919908577405",
				Age:   12,
			},
			Mock: []interface{}{mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).AddRow(2, "Zopsmart", "sudheerpuppala@gmail.com", "+919908577405", 12))},
			err:  nil,
		},
		{
			desc: "Failure",
			Id:   2,
			expected: models.User{
				Id:    2,
				Name:  "Zopsmart",
				Email: "sudheerpuppala@gmail.com",
				Phone: "+919908577405",
				Age:   12,
			},
			Mock: []interface{}{mock.ExpectQuery(query).WillReturnError(errors.New("t"))},
			err:  errors.New("t"),
		},
	}
	handler := New(db)
	for _, tc := range tcsSearch {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := handler.SearchById(tc.Id)
			if err != nil {
				return
			}
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expected, res)
			}
		})
	}
}

func Test_SearchAll(t *testing.T) {
	query := "SELECT Id,Name,Email,Phone,Age FROM User"
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("database error :%s", err)
	}
	tcsSearch := []struct {
		desc     string
		Id       int
		expected []models.User
		Mock     []interface{}
		err      error
	}{
		{
			desc: "Success",
			Id:   2,
			expected: []models.User{
				{Id: 2,
					Name:  "Zopsmart",
					Email: "sudheerpuppala@gmail.com",
					Phone: "+919908577405",
					Age:   12,
				},
			},
			Mock: []interface{}{mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).AddRow(2, "Zopsmart", "sudheerpuppala@gmail.com", "+919908577405", 12))},
			err:  nil,
		},
		{
			desc: "Success-2",
			Id:   2,
			expected: []models.User{
				{Id: 2,
					Name:  "Zopsmart",
					Email: "sudheerpuppala@gmail.com",
					Phone: "+919908577405",
					Age:   12,
				},
			},
			Mock: []interface{}{mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).AddRow(2, "Zopsmart", "sudheerpuppala@gmail.com", "+919908577405", 12))},
			err:  nil,
		},
		{
			desc: "Failure",
			Id:   2,
			expected: []models.User{
				{Id: 2,
					Name:  "Zopsmart",
					Email: "sudheerpuppala@gmail.com",
					Phone: "+919908577405",
					Age:   12,
				},
			},
			Mock: []interface{}{mock.ExpectQuery(query).WillReturnError(errors.New("t"))},
			err:  errors.New("t"),
		},
	}
	handler := New(db)
	for _, tc := range tcsSearch {
		t.Run(tc.desc, func(t *testing.T) {
			res, err := handler.SearchDetails()
			if err != nil {
				return
			}
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("Expected : %v,Obtained : %v ", tc.expected, res)
			}
		})
	}
}
