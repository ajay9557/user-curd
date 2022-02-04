package stores

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
	"testing"
	"user-curd/model"

	"github.com/DATA-DOG/go-sqlmock"
)

var use = model.User{
	Id:    1,
	Name:  "sai",
	Email: "Nayani@gmail.com",
	Phone: "6303880131",
	Age:   "25",
}

func NMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}
func Test_Delete(t *testing.T) {
	db, mock := NMock()
	query := "delete from User where Id=?"
	dbhandler := New(db)
	testcases := []struct {
		desc   string
		inp    int
		expErr error
		mock   []interface{}
	}{
		{
			"testcase-1",
			1,
			nil,
			[]interface{}{
				mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1)),
			},
		},
	}
	for _, tcs := range testcases {
		err := dbhandler.DeleteById(tcs.inp)

		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("expected error %v, got %v", tcs.expErr, err)
		}
	}
}

func Test_Update(t *testing.T) {
	db, mock := NMock()
	dbhandler := New(db)
	testcases := []struct {
		desc   string
		inp    model.User
		expErr error
		mock   []interface{}
	}{

		{
			"testcase-2",
			use,
			errors.New("error in updating user name"),
			[]interface{}{
				mock.ExpectExec("update User set Name = ?, Email = ?, Phone = ?, Age = ? where Id = ?").WithArgs("sai", "Nayani@gmail.com", "6303880131", "25", 1).
					WillReturnError(errors.New("error in updating user name")),
			},
		},
	}
	for _, tcs := range testcases {
		_, err := dbhandler.UpdateById(&tcs.inp)
		if !reflect.DeepEqual(err, tcs.expErr) {
			t.Errorf("expect error %v, got %v", tcs.expErr, err)
		}
	}
}
func Test_Id(t *testing.T) {
	db, mock := NMock()
	query := "SELECT Id,Name,Email,Phone,Age FROM User where Id=?"
	dbhandler := New(db)

	testcases := []struct {
		desc   string
		inp    int
		exp    *model.User
		expErr error
		mock   []interface{}
	}{
		{
			"testcase-1",
			1,
			&use,
			nil,
			[]interface{}{
				mock.ExpectQuery(query).WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).
							AddRow(
								1,
								"sai",
								"Nayani@gmail.com",
								"6303880131",
								25,
							),
					),
			},
		},
	}
	for _, tcs := range testcases {
		out, err := dbhandler.GetById(tcs.inp)
		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("expected %v, got %v", tcs.exp, out)
		}
		if !reflect.DeepEqual(tcs.expErr, err) {
			t.Errorf("expected error %v , got %v", tcs.expErr, err)
		}
	}
}
func Test_SearchAll(t *testing.T) {
	db, mock := NMock()
	query := "SELECT Id,Name,Email,Phone,Age FROM User"
	dbhandler := New(db)

	testcases := []struct {
		desc   string
		inp    int
		exp    []*model.User
		expErr error
		mock   []interface{}
	}{
		{
			"testcase-1",
			1,
			[]*model.User{&use},
			nil,
			[]interface{}{
				mock.ExpectQuery(query).
					WillReturnRows(
						sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).
							AddRow(
								1,
								"sai",
								"Nayani@gmail.com",
								"6303880131",
								25,
							),
					),
			},
		},
		{
			"testcase-2",
			1,
			[]*model.User{},
			errors.New("error in preparing statement"),
			[]interface{}{
				mock.ExpectQuery(query).WillReturnError(errors.New("error in preparing statement")),
			},
		},
	}
	for _, tcs := range testcases {
		out, err := dbhandler.GetAll()
		if !reflect.DeepEqual(out, tcs.exp) {
			t.Errorf("expected %v, got %v", tcs.exp, out)
		}
		if !reflect.DeepEqual(tcs.expErr, err) {
			t.Errorf("expected error %v , got %v", tcs.expErr, err)
		}
	}
}
func Test_Insert(t *testing.T) {
	db, mock := NMock()
	query := "insert into User(Name,Email,Phone,Age) values(?,?,?,?)"

	testcases := []struct {
		desc   string
		inp    model.User
		expOut *model.User
		expErr error
		mock   []interface{}
	}{
		{
			"testcase-1",
			use,
			&use,
			nil,
			[]interface{}{
				mock.ExpectExec(query).
					WithArgs(use.Name, use.Email, use.Phone, use.Age).
					WillReturnResult(sqlmock.NewResult(1, 0)),
			},
		},
		{
			"testcase-2",
			use,
			&model.User{},
			errors.New("error in executing insert"),
			[]interface{}{
				mock.ExpectPrepare(query).ExpectExec().
					WithArgs(use.Id, use.Name, use.Email, use.Phone, use.Age).
					WillReturnError(errors.New("error in executing insert")),
			},
		},
		{
			"testcase-3",
			use,
			&model.User{},
			errors.New("error in executing insert"),
			[]interface{}{
				mock.ExpectPrepare(query).
					WillReturnError(errors.New("error in executing insert")),
			},
		},
	}
	dbhandler := New(db)
	for _, tcs := range testcases {
		out, err := dbhandler.InsertDetails(&tcs.inp)
		if !reflect.DeepEqual(out, tcs.expOut) {
			t.Errorf("expected %v, got %v", tcs.expOut, out)
		}
		if !reflect.DeepEqual(tcs.expErr, err) {
			t.Errorf("expected error %v, got %v", tcs.expErr, err)
		}
	}
}
