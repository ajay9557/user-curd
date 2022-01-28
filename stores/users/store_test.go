package users

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"user-curd/models"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	testCases := []struct {
		desc     string
		User     models.User
		Mock     []interface{}
		expecErr error
	}{
		{
			desc: "Success case",
			User: models.User{
				Id:    1,
				Name:  "gopi",
				Email: "gopi@gmail.com",
				Phone: "1234567899",
				Age:   23,
			},
			Mock: []interface{}{
				mock.ExpectExec("insert into user").
					WithArgs(1, "gopi", "gopi@gmail.com", "1234567899", 23).
					WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			expecErr: nil,
		},
		{
			desc: "Failure - ID case",
			User: models.User{
				Id:    0,
				Name:  "gopi",
				Email: "gopi@gmail.com",
				Phone: "1234567899",
				Age:   23,
			},
			Mock: []interface{}{
				mock.ExpectExec("insert into user").
					WithArgs(0, "gopi", "gopi@gmail.com", "1234567899", 23).
					WillReturnError(errors.New("id can`t be 0")),
			},
			expecErr: errors.New("id can`t be 0"),
		},
		{
			desc: "Failure - DB case",
			User: models.User{
				Id:    1,
				Name:  "chand",
				Email: "gopi@gmail.com",
				Phone: "1234567899",
				Age:   23,
			},
			Mock: []interface{}{
				mock.ExpectExec("insert into user").
					WithArgs(1, "chand", "gopi@gmail.com", "1234567899", 23).
					WillReturnError(errors.New("database error")),
			},
			expecErr: errors.New("database error"),
		},
	}
	handler := New(db)
	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			err := handler.InsertUser(ts.User)
			if err != nil && !reflect.DeepEqual(ts.expecErr, err) {
				fmt.Print("expected ", ts.expecErr, "obtained", err)
			}
		})
	}
}

func TestFetchAllUsers(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testFetch := []struct {
		desc     string
		Mock     []interface{}
		expecRes []models.User
		expecErr error
	}{
		{
			desc: "Success case",
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Email,Phone,Age from user").
					WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).
						AddRow(1, "gopi", "gopi@gmail.com", "1234567899", 23)),
			},
			expecRes: []models.User{
				{
					Id:    1,
					Name:  "gopi",
					Email: "gopi@gmail.com",
					Phone: "1234567899",
					Age:   23,
				},
			},
			expecErr: nil,
		},
		{
			desc: "Failure case ",
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Email,Phone,Age from user").
					WillReturnError(errors.New("database error")),
			},
			expecRes: nil,
			expecErr: errors.New("database error"),
		},
	}
	handler := New(db)
	for _, ts := range testFetch {
		t.Run(ts.desc, func(t *testing.T) {
			res, err := handler.FetchAllUsers()
			if err != nil && !reflect.DeepEqual(ts.expecErr, err) {
				fmt.Print("expected ", ts.expecErr, "obtained", err)
			}
			if res != nil && !reflect.DeepEqual(ts.expecRes, res) {
				fmt.Print("expected ", ts.expecRes, "obtained", res)
			}
		})
	}
}

func TestFetchUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testFetch := []struct {
		desc     string
		Id       int
		Mock     []interface{}
		expecRes models.User
		expecErr error
	}{
		{
			desc: "Success case",
			Id:   1,
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Email,Phone,Age from user").WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).
						AddRow(1, "gopi", "gopi@gmail.com", "1234567899", 23)),
			},
			expecRes: models.User{
				Id:    1,
				Name:  "gopi",
				Email: "gopi@gmail.com",
				Phone: "1234567899",
				Age:   23,
			},
			expecErr: nil,
		},
		{
			desc: "Failure case ",
			Id:   2,
			Mock: []interface{}{
				mock.ExpectQuery("Select Id,Name,Email,Phone,Age from user").
					WillReturnError(errors.New("database error")),
			},
			expecRes: models.User{
				Id:    0,
				Name:  "",
				Email: "",
				Phone: "",
				Age:   0,
			},
			expecErr: errors.New("database error"),
		},
	}
	handler := New(db)
	for _, ts := range testFetch {
		t.Run(ts.desc, func(t *testing.T) {
			res, err := handler.FetchUserById(ts.Id)
			if err != nil && !reflect.DeepEqual(ts.expecErr, err) {
				fmt.Print("expected ", ts.expecErr, "obtained", err)
			}
			if !reflect.DeepEqual(ts.expecRes, res) {
				fmt.Print("expected ", ts.expecRes, "obtained", res)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	testCases := []struct {
		desc     string
		User     models.User
		Mock     []interface{}
		expecErr error
	}{
		{
			desc: "Success case",
			User: models.User{
				Id:    1,
				Name:  "gopi",
				Email: "gopi123@gmail.com",
				Phone: "1234567699",
				Age:   23,
			},
			Mock: []interface{}{
				mock.ExpectExec("update user set").
					WithArgs(23, "gopi", "gopi123@gmail.com", "1234567699", 1).
					WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			expecErr: nil,
		},
		{
			desc: "Failure case",
			User: models.User{
				Id:    2,
				Name:  "gopi",
				Email: "gopi123@gmail.com",
				Phone: "1234567699",
				Age:   23,
			},
			Mock: []interface{}{
				mock.ExpectExec("Update user set").
					WithArgs(23, "gopi", "gopi123@gmail.com", "1234567699", 2).
					WillReturnError(errors.New("database error")),
			},
			expecErr: errors.New("database error"),
		},
	}
	handler := New(db)
	for _, ts := range testCases {
		t.Run(ts.desc, func(t *testing.T) {
			err := handler.UpdateUser(ts.User)
			if err != nil && !reflect.DeepEqual(ts.expecErr, err) {
				fmt.Print("expected ", ts.expecErr, "obtained", err)
			}
		})
	}
}

func TestDeleteUserById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testDelete := []struct {
		desc     string
		Id       int
		Mock     []interface{}
		expecErr error
	}{
		{
			desc: "Success case",
			Id:   1,
			Mock: []interface{}{
				mock.ExpectExec("DELETE From user").
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			expecErr: nil,
		},
		{
			desc: "Failure case",
			Id:   2,
			Mock: []interface{}{
				mock.ExpectExec("DELETE From user").
					WithArgs(1).
					WillReturnError(errors.New("database error")),
			},
			expecErr: errors.New("database error"),
		},
	}
	handler := New(db)
	for _, ts := range testDelete {
		t.Run(ts.desc, func(t *testing.T) {
			err := handler.DeleteUserById(ts.Id)
			if err != nil && !reflect.DeepEqual(ts.expecErr, err) {
				fmt.Print("expected ", ts.expecErr, "obtained", err)
			}
		})
	}
}

func TestGetEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testEmail := []struct {
		desc     string
		Email    string
		Mock     []interface{}
		expecRes bool
		expecErr error
	}{
		{
			desc:  "Success case",
			Email: "gopi1234@gmail.com",
			Mock: []interface{}{
				mock.ExpectQuery("select Email from user").
					WithArgs("gopi1234@gmail.com").WillReturnRows(sqlmock.NewRows([]string{"Email"}).AddRow("")),
			},
			expecRes: true,
			expecErr: nil,
		},
		{
			desc:  "Success case - 2",
			Email: "gopi123@gmail.com",
			Mock: []interface{}{
				mock.ExpectQuery("select Email from user").
					WithArgs("gopi123@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{"Email"}).AddRow("gopi123@gmail.com")),
			},
			expecRes: false,
			expecErr: nil,
		},
		{
			desc:  "Failure case",
			Email: "gopi12@gmail.com",
			Mock: []interface{}{
				mock.ExpectQuery("select Email from user").
					WithArgs("gopi12@gmail.com").
					WillReturnError(errors.New("database error")),
			},
			expecRes: false,
			expecErr: errors.New("database error"),
		},
	}
	handler := New(db)
	for _, ts := range testEmail {
		t.Run(ts.desc, func(t *testing.T) {
			res, err := handler.GetEmail(ts.Email)
			if !reflect.DeepEqual(ts.expecRes, res) {
				fmt.Print("expected ", ts.expecRes, "obtained", res)
			}
			if err != nil && !reflect.DeepEqual(ts.expecErr, err) {
				fmt.Print("expected ", ts.expecErr, "obtained", err)
			}
		})
	}

}
