package users

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/usercurd/models"
)

 var usr = models.User {
	 Id: 1,
	 Name: "Amrish",
	 Email: "am12@gmail.com",
	 Phone: "11",
	 Age: 23,
 }
 func NewMock() (*sql.DB, sqlmock.Sqlmock){
	 db,mock,err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	 if err != nil {
		 log.Fatalf("an error '%s' was not expected when opening a stub connection",err)
	 }
	 return db,mock
 }
  func TestUserStore_Insert(t *testing.T) {
	  db, mock := NewMock()
	  query := "insert into User values(?,?,?,?,?)"

	  testcases := []struct {
		  inp models.User
		  expOut models.User
		  expErr error
		  mock []interface{}
	  }{
		  {
			  usr,
			  usr,
			  nil,
			  []interface{}{
				  mock.ExpectPrepare(query).ExpectExec().
				  WithArgs(usr.Id,usr.Name,usr.Email,usr.Phone,usr.Age). 
				  WillReturnResult(sqlmock.NewResult(1,1)),
			  },
		  },
		  {
			  usr,
			  models.User{},
			  errors.New("duplicate id"),
			  []interface{}{
				  mock.ExpectPrepare(query).ExpectExec().
				  WithArgs(usr.Id,usr.Name,usr.Email,usr.Phone,usr.Age).
				  WillReturnError(errors.New("duplicate id")),
			  },
		  },
		  {
			  usr,
			  models.User{},
			  errors.New("error in executing statement"),
			  []interface{}{
				  mock.ExpectPrepare(query).
				    WillReturnError(errors.New("error in executing statement")),
			  },
		  },
	  }
	  dbhandler := New(db)
	  for _,tcs := range testcases {
		  out, err := dbhandler.Insert(tcs.inp)
		  if !reflect.DeepEqual(tcs.expErr, err){
			  t.Errorf("expected %v, got %v", tcs.expOut, out)
		  }
		  if !reflect.DeepEqual(tcs.expOut,out) {
			  t.Errorf("expected error %v, got %v",tcs.expErr,err)
		  }
	  }
  }
  func TestUserStore_Update(t *testing.T) {
	  	db, mock := NewMock()
		query := "insert into User values(?,?,?,?,?)"
	
		testcases := []struct {
			inp    models.User
			expOut models.User
			expErr error
			mock   []interface{}
		}{
			{
				usr,
				usr,
				nil,
				[]interface{}{
					mock.ExpectPrepare(query).ExpectExec().
						WithArgs(usr.Id, usr.Name, usr.Email, usr.Phone, usr.Age).
						WillReturnResult(sqlmock.NewResult(1, 1)),
				},
			},
			{
				usr,
				models.User{},
				errors.New("duplicate id"),
				[]interface{}{
					mock.ExpectPrepare(query).ExpectExec().
						WithArgs(usr.Id, usr.Name, usr.Email, usr.Phone, usr.Age).
						WillReturnError(errors.New("duplicate id")),
				},
			},
			{
				usr,
				models.User{},
				errors.New("error in executing statement"),
				[]interface{}{
					mock.ExpectPrepare(query).
						WillReturnError(errors.New("error in executing statement")),
				},
			},
		}
		dbhandler := New(db)
		for _, tcs := range testcases {
			out, err := dbhandler.Insert(tcs.inp)
			if !reflect.DeepEqual(out, tcs.expOut) {
				t.Errorf("expected %v, got %v", tcs.expOut, out)
			}
			if !reflect.DeepEqual(tcs.expErr, err) {
				t.Errorf("expected error %v, got %v", tcs.expErr, err)
			}
		}
	}

	func TestUserStore_Delete(t *testing.T) {
		db, mock := NewMock()
		query := "delete from User where id = (?)"
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
			err := dbhandler.Delete(tcs.inp)
	
			if !reflect.DeepEqual(err, tcs.expErr) {
				t.Errorf("expected error %v, got %v", tcs.expErr, err)
			} 
		} 
	} 
	func TestUserStore_GetById(t *testing.T) {
		db, mock := NewMock()
		query := "select * from User where id = ?"
		dbhandler := New(db)
	
		testcases := []struct {
			desc   string
			inp    int
			exp    models.User
			expErr error
			mock   []interface{}
		}{
			{ 
				"testcase-1",
				1,
				usr,
				nil,
				[]interface{}{
					mock.ExpectQuery(query).WithArgs(1).
						WillReturnRows(
							sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).
								AddRow(
									1,
									"Amrish",
									"am12@gmail.com",
									"11",
									23,
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

	func TestUserStore_GetAll(t *testing.T) {
		db, mock := NewMock()
		query := "select * from User"
		dbhandler := New(db)
	
		testcases := []struct {
			desc   string
			inp    int
			exp    []models.User
			expErr error
			mock   []interface{}
		}{
			{
				"testcase-1",
				1,
				[]models.User{usr},
				nil,
				[]interface{}{
					mock.ExpectQuery(query).
						WillReturnRows(
							sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).
								AddRow(
									1,
									"Amrish",
									"am12@gmail.com",
									"11",
									23,
								),
						),
				},
			},
			{
				"testcase-2",
				1,
				[]models.User{},
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
	
