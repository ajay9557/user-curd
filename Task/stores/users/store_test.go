package store

import (
	"database/sql"
//	"errors"
	"fmt"
	reflect "reflect"
	"testing"
	"zopsmart/Task/models"

	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/golang/mock/gomock"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		fmt.Println(fmt.Errorf("an error occured while opening user database"))
	}
	return db, mock
}

func TestGetAllUsers(t *testing.T) {
	db,mock := NewMock()
	s := DbStore{db:db}
	query := "select Id,Name,Email,Phone,Age from user;"
	tcs := []struct {
		expectedErr error
		expectedOut []models.User
		mockQuery *sqlmock.ExpectedQuery
	} {
		{
			expectedErr: nil,
			mockQuery: mock.ExpectQuery(query).
				WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).
					AddRow(1, "prasath", "prasath@gmail.com", "12345", 20).
					AddRow(2, "rishi kumar", "rishi.kumar@gmail.com", "67890", 25)),
			expectedOut: []models.User{
				{
					Id:    1,
					Name:  "prasath",
					Email: "prasath@gmail.com",
					Phone: "12345",
					Age:   20,
				},
				{
					Id:    2,
					Name:  "rishi kumar",
					Email: "rishi.kumar@gmail.com",
					Phone: "67890",
					Age:   25,
				},
			},
		},
	}
    for _,tc := range tcs {
		t.Run("testing",func(t *testing.T) {
		res, err := s.GetAllUsersStore()
		if !reflect.DeepEqual(res, tc.expectedOut) {
			t.Errorf("Expected: \t%v\nGot: \t%v\n",tc.expectedOut, res)
		}
		if err != tc.expectedErr {
			t.Errorf("Expected: \t%v\nGot: \t%v\n",tc.expectedErr, err)
		}
	})
			
	}
}

func TestGetUserById(t *testing.T) {
	db, mock := NewMock()
	s := DbStore{db: db}
	query := "select Id,Name,Email,Phone,Age  from user where id = ?"

	tcs := []struct {
		testCase    int
		id          int
		expectedErr error
		expectedOut models.User
	}{
		{
			testCase:    1,
			id:          1,
			expectedErr: nil,
			expectedOut: models.User{Id: 1, Name: "prasath", Email: "prasath@gmail.com", Phone: "12345", Age: 20},
		},
	}

	for _, tc := range tcs {
		mock.ExpectQuery(query).WithArgs(tc.id).WillReturnRows(sqlmock.NewRows([]string{"Id", "Name", "Email", "Phone", "Age"}).AddRow(1, "prasath", "prasath@gmail.com", "12345", 20))
		res, err := s.GetUserById(tc.id)
		if !reflect.DeepEqual(res, tc.expectedOut) {
			t.Errorf("Testcase[%v] Expected %v \n Got %v\n", tc.testCase, tc.expectedOut, res)
		}
		if err != nil && err != tc.expectedErr {
			t.Errorf("Testcase[%v] Expected %v \n Got %v\n", tc.testCase, tc.expectedErr, err)
		}
	}
}




func TestUpdate(t *testing.T) {

	db, mock,_ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	s := DbStore{db: db}
	tests := []struct {
		testCase    int
		input models.User
		expectedErr error

	}{
		{
			testCase:    1,
			input : models.User{
				Id : 1,
				Name : "jhon",
				Phone : "84732",
				Email: "jhon@gmail.com",
				Age: 21,
			},
			expectedErr: nil,
		
		},
	}
	

	
	for _, tc := range tests {
		t.Run("Success case testing ", func(t *testing.T) {
			query := "update user set"
			fields, values := find(tc.input)
			query += fields + " where id = ?"
			mock.ExpectExec(query).WithArgs(values[0], values[1], values[2], values[3], values[4]).WillReturnResult(sqlmock.NewResult(1, 1))
		
		//	fmt.Println(query)
			err := s.Update(tc.input)
			if err != tc.expectedErr {
				t.Errorf("Expected: \t%v\nGot: \t%v\n", tc.expectedErr, err)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	db, mock := NewMock()
	s := DbStore{db: db}
	query := "insert into users\\(Id, Name, Email, Phone, Age\\) VALUES\\(\\?,\\?,\\?,\\?,\\?\\)"

	tests := []struct {
		testCase    int
		Id          int
		Name        string
		Email       string
		Phone       string
		Age         int
		expectedErr error
		mock        []interface{}
	}{
		{
			testCase:    1,
			Id:          6,
			Name:        "Tony",
			Email:       "tony@gmail.com",
			Phone:       "30001",
			Age:         43,
			expectedErr: nil,
			mock:        []interface{}{mock.ExpectExec(query).WithArgs(6, "Tony", "tony@gmail.com", "30001", 43).WillReturnResult(sqlmock.NewResult(0, 1))},
		},
	}

	for _, tc := range tests {
		err := s.Create(tc.Id, tc.Name, tc.Email, tc.Phone, tc.Age)
		if err != nil && err != tc.expectedErr {
			t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.testCase, tc.expectedErr, err)
		}
	}
}

func TestDelete(t *testing.T) {

	db, mock := NewMock()
	s := DbStore{db: db}
	query := "delete from user where id = ?"

	tests := []struct {
		testCase    int
		Id          int
		expectedErr error
		mock        []interface{}
	}{
		{
			testCase:    1,
			Id:          6,
			expectedErr: nil,
			mock:        []interface{}{mock.ExpectExec(query).WithArgs(6).WillReturnResult(sqlmock.NewResult(1, 1))},
		},
	}

	for _, tc := range tests {
		err := s.Delete(tc.Id)
		if err != nil && err != tc.expectedErr {
			t.Errorf("Testcase[%v] Expected %v \n Got %v\n", tc.testCase, tc.expectedErr, err)
		}
	}

}