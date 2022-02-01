package user

import (
	"errors"
	"fmt"
	"go_lang/Assignment/user-curd/models"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_Create(t *testing.T) {
	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	dbHandler := SqlDb{database}
	if err != nil {
		fmt.Println("ERROR IN MOCKING DB")
	}

	query := "insert into users(name, email, phone, age) values(?, ?, ?, ?)"

	testCases := []struct {
		desc          string
		input         models.User
		mock          []interface{}
		expectedError error
	}{
		{
			desc:          "Test Case 1",
			input:         models.User{Id: 10, Name: "Akash Sharma", Email: "akash22@gmail.com", Phone: "7827876546", Age: 23},
			mock:          []interface{}{mock.ExpectExec(query).WithArgs("Akash Sharma", "akash22@gmail.com", "7827876546", 23).WillReturnResult(sqlmock.NewResult(1, 1))},
			expectedError: nil,
		},
		{
			desc:          "Test Case 2",
			input:         models.User{Id: 54, Name: "Akash Gupta", Email: "akash22@gmail.com", Phone: "7827876546", Age: 23},
			mock:          []interface{}{mock.ExpectExec(query).WithArgs("Akash Gupta", "akash22@gmail.com", "7827876546", 23).WillReturnError(errors.New("ERROR IN INSERTING DATA"))},
			expectedError: errors.New("ERROR IN INSERTING DATA"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			err = dbHandler.Create(testCase.input)
			if testCase.expectedError == nil && err != nil {
				t.Errorf("Error Occoured: %s", err)
			}
			if err != nil && !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf("Output: %v, Expected: %v", err, testCase.expectedError)
			}
		})
	}
}

func Test_ReadOne(t *testing.T) {
	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	dbHandler := SqlDb{database}
	if err != nil {
		fmt.Println("ERROR IN MOCKING DB")
	}

	query := "select * from users where id="

	testCases := []struct {
		desc          string
		input         int
		mock          []interface{}
		expectedError error
	}{
		{
			desc:  "Test Case 1",
			input: 3,
			mock: []interface{}{
				mock.ExpectQuery(query).WithArgs(3).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "age"}).AddRow(3, "Rakesh Richard", "rakesh43@rediffmail.com", "4324434343", 12)),
			},
			expectedError: nil,
		},
		{
			desc:  "Test Case 2",
			input: 1000,
			mock: []interface{}{
				mock.ExpectQuery(query).WithArgs(3).WillReturnError(errors.New("ERROR IN FETCHING ROW")),
			},
			expectedError: errors.New("ERROR IN FETCHING ROW"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			usr, err := dbHandler.ReadOne(testCase.input)
			if testCase.expectedError == nil && err != nil {
				fmt.Print("expected ", testCase.expectedError, "obtained", err)
			}
			if err != nil && !reflect.DeepEqual(testCase.expectedError, err) {
				fmt.Println("Output: ", err, "Expected: ", testCase.expectedError)
			}
			fmt.Println(usr)
		})
	}
}

func Test_ReadAll(t *testing.T) {
	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	dbHandler := SqlDb{database}
	if err != nil {
		fmt.Println("ERROR IN MOCKING DB")
	}

	query := "select * from users"

	testCases := []struct {
		desc          string
		mock          []interface{}
		expectedError error
	}{
		{
			desc: "Test Case 1",
			mock: []interface{}{
				mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "age"}).AddRow(1, "Devid Jhones", "devid54@gmail.com", "9876543563", 45)),
			},
			expectedError: nil,
		},
		{
			desc: "Test Case 2",
			mock: []interface{}{
				mock.ExpectQuery(query).WillReturnError(errors.New("ERROR IN FETCHING ROWS")),
			},
			expectedError: errors.New("ERROR IN FETCHING ROWS"),
		},
		{
			desc: "Test Case 3",
			mock: []interface{}{
				mock.ExpectQuery(query).WillReturnError(errors.New("ERROR IN FETCHING USER DATA")),
			},
			expectedError: errors.New("ERROR IN FETCHING USER DATA"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			usr, err := dbHandler.ReadAll()
			if testCase.expectedError == nil && err != nil {
				t.Errorf("%v %v", err, testCase.expectedError)
			}
			if len(usr) != 0 && err != nil && !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf("output: %v expected: %v", err, testCase.expectedError)
			}
			fmt.Println(usr)
		})
	}
}

func Test_Update(t *testing.T) {
	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	dbHandler := SqlDb{database}
	if err != nil {
		fmt.Println("ERROR IN MOCKING DB")
	}

	query := "update users set name=?, email=?, phone=?, age=? where id=?"

	testCases := []struct {
		desc          string
		input         models.User
		mock          []interface{}
		expectedError error
	}{
		{
			desc:  "Test Case 1",
			input: models.User{Id: 10, Name: "Akash Sharma", Email: "akash22@gmail.com", Phone: "7827876546", Age: 23},
			mock: []interface{}{
				mock.ExpectExec(query).WithArgs("Akash Sharma", "akash22@gmail.com", "7827876546", 23, 10).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			expectedError: nil,
		},
		{
			desc:  "Test Case 2",
			input: models.User{Id: 10, Name: "Akash Sharma", Email: "akash22@gmail.com", Phone: "7827876546", Age: 23},
			mock: []interface{}{
				mock.ExpectExec(query).WithArgs("Akash Sharma", "akash22@gmail.com", "7827876546", 23, 10).WillReturnError(errors.New("FAILED TO UPDATE USER DATA")),
			},
			expectedError: errors.New("FAILED TO UPDATE USER DATA"),
		},
		{
			desc:  "Test Case 3",
			input: models.User{Id: -1, Name: "Akash Sharma", Email: "akash22@gmail.com", Phone: "7827876546", Age: 23},
			mock: []interface{}{
				mock.ExpectExec(query).WithArgs("Akash Sharma", "akash22@gmail.com", "7827876546", 23, -1).WillReturnError(errors.New("INVALID ID")),
			},
			expectedError: errors.New("INVALID ID"),
		},
		{
			desc:  "Test Case 4",
			input: models.User{Id: 10, Name: "", Email: "akash22@gmail.com", Phone: "", Age: 23},
			mock: []interface{}{
				mock.ExpectExec("update users set email=?, age=? where id=?").WithArgs("akash22@gmail.com", 23, 10).WillReturnError(errors.New("FAILED TO UPDATE USER DATA")),
			},
			expectedError: errors.New("FAILED TO UPDATE USER DATA"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			err := dbHandler.Update(testCase.input)
			if testCase.expectedError == nil && err != nil {
				t.Errorf("%v %v", err, testCase.expectedError)
			}
			if err != nil && !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf("output: %v expected: %v", err, testCase.expectedError)
			}
		})
	}
}

func Test_Delete(t *testing.T) {
	database, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	dbHandler := New(database)
	if err != nil {
		fmt.Println("ERROR IN MOCKING DB")
	}

	query := "delete from users where id=?"

	testCases := []struct {
		desc          string
		input         int
		mock          []interface{}
		expectedError error
	}{
		{
			desc:  "Test Case 1",
			input: 1,
			mock: []interface{}{
				mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
			expectedError: nil,
		},
		{
			desc:  "Test Case 2",
			input: -1,
			mock: []interface{}{
				mock.ExpectExec(query).WithArgs(-1).WillReturnError(errors.New("FAILED TO DELETE USER")),
			},
			expectedError: errors.New("FAILED TO DELETE USER"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			err := dbHandler.Delete(testCase.input)
			if testCase.expectedError == nil && err != nil {
				t.Errorf("%v %v", err, testCase.expectedError)
			}
			if err != nil && !reflect.DeepEqual(testCase.expectedError, err) {
				t.Errorf("output: %v expected: %v", err, testCase.expectedError)
			}
		})
	}
}
