package user

import (
	"fmt"
	"reflect"
	"testing"
	"zopsmart/user-curd/model"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_GetUserById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}
	defer db.Close()
	hndlr := New(db)
	rows := sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).
		AddRow(1, "Sarah Vaughan", "Sarah@gmail.com", "9866895296", 21)
	tests := []struct {
		desc string
		id   int
		usr  model.User
		err  error
		mock []interface{}
	}{
		{
			desc: "case-1",
			id:   1,
			usr:  model.User{Id: 1, Name: "Sarah Vaughan", Email: "Sarah@gmail.com", Phone: "9866895296", Age: 21},
			err:  nil,
			mock: []interface{}{
				mock.ExpectQuery("SELECT * FROM User WHERE id = ?").WithArgs(1).WillReturnRows(rows),
			},
		},
	}
	for _, tc := range tests {
		usr, err := hndlr.GetUserById(tc.id)
		fmt.Printf("%T,%T", err, tc.err)
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("expected error: %s, got error: %s", tc.err, err)
		}
		if !reflect.DeepEqual(tc.usr, usr) {
			t.Errorf("expected %v, got :  %v", tc.usr, usr)
		}

	}

}

func Test_AddUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}
	defer db.Close()
	hndlr := New(db)

	tests := []struct {
		desc string
		usr  model.User
		id   int
		err  error
		mock []interface{}
	}{
		{
			desc: "case-1",
			usr:  model.User{Name: "Sarah Vaughan", Email: "Sarah@gmail.com", Phone: "9866895296", Age: 21},
			id:   1,
			err:  nil,
			mock: []interface{}{
				mock.ExpectExec("INSERT INTO User(Name,Email,Phone,Age) VALUES (?,?,?,?)").WithArgs("Sarah Vaughan", "Sarah@gmail.com", "9866895296", 21).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
	}
	for _, tc := range tests {
		u, err := hndlr.AddUser(tc.usr)

		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("expected error: %s, got error: %s", tc.err, err)
		}
		if tc.id != u {
			t.Errorf("expected %v, got :  %v", tc.id, u)
		}

	}

}

func Test_UpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}
	defer db.Close()
	hndlr := New(db)
	tests := []struct {
		id   int
		usr  model.User
		err  error
		mock []interface{}
	}{
		{

			id:  1,
			usr: model.User{Id: 1, Name: "Naira", Email: "naira@gmail.com", Phone: "9866895296", Age: 20},
			err: nil,
			mock: []interface{}{
				mock.ExpectExec("Update User set name=?, email=?,phone=?,age=? Where id=?").WithArgs("Naira", "naira@gmail.com", "9866895296", 20, 1).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
	}
	for _, tc := range tests {
		err := hndlr.UpdateUser(tc.usr)

		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("expected error: %s, got error: %s", tc.err, err)
		}

	}

}

func Test_Delete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}
	defer db.Close()
	hndlr := New(db)
	tests := []struct {
		desc string
		usr  []model.User
		id   int
		err  error
		mock []interface{}
	}{
		{
			desc: "case-1",
			usr:  []model.User{{1, "Mahi", "Mahi@gmail.com", "9866895296", 20}},
			id:   1,
			err:  nil,
			mock: []interface{}{
				mock.ExpectExec("Delete FROM User WHERE id = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
		{
			desc: "case-3",
			usr:  []model.User{{1, "Mahi", "Mahi@gmail.com", "9866895296", 20}},
			id:   5 / 2,
			err:  fmt.Errorf("ExecQuery 'Delete FROM User WHERE id = ?', arguments do not match: argument 0 expected [int64 - 0] does not match actual [int64 - 2]"),
			mock: []interface{}{
				mock.ExpectExec("Delete FROM User WHERE id = ?").WithArgs(1 / 2).WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("ExecQuery 'Delete FROM User WHERE id = ?', arguments do not match: argument 0 expected [int64 - 0] does not match actual [int64 - 2]"))),
			},
		},
	}
	for _, tc := range tests {
		err := hndlr.DeleteUser(tc.id)

		if !reflect.DeepEqual(err, tc.err) {
			t.Fatalf("expected error: %s, got %s", tc.err, err)
		}

	}

}

func Test_GetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while opening a stub database connection", err)
	}
	defer db.Close()
	hndlr := New(db)
	rows := sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).
		AddRow(1, "Sarah Vaughan", "Sarah@gmail.com", "9866895296", 21).
		AddRow(2, "Naira", "Naira@gmail.com", "6303844857", 20)
	tests := []struct {
		desc string
		usr  []model.User
		err  error
		mock []interface{}
	}{
		{
			desc: "case-1",
			usr:  []model.User{{Id: 1, Name: "Sarah Vaughan", Email: "Sarah@gmail.com", Phone: "9866895296", Age: 21}, {Id: 2, Name: "Naira", Email: "Naira@gmail.com", Phone: "6303844857", Age: 20}},
			err:  nil,
			mock: []interface{}{
				mock.ExpectQuery("Select * From User").WillReturnRows(rows),
			},
		},
	}
	for _, tc := range tests {
		rows, err := hndlr.GetAllUsers()

		if !reflect.DeepEqual(err, tc.err) {
			t.Fatalf("expected error: %s, got %s", tc.err, err)
		}
		if !reflect.DeepEqual(rows, tc.usr) {
			t.Fatalf("expected error: %v, got %v", tc.usr, rows)
		}

	}
}
