package users

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/tejas/user-crud/models"
)

func TestGetById(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "age"}).AddRow(
		1, "tejas", "tejas@gmail.com", "9876543210", 23,
	)

	testCases := []struct {
		desc      string
		id        int
		expected  *models.User
		mockQuery *sqlmock.ExpectedQuery
	}{
		{
			desc: "Case 1: Success case",
			id:   1,
			expected: &models.User{
				Id:    1,
				Name:  "tejas",
				Email: "tejas@gmail.com",
				Phone: "9876543210",
				Age:   23,
			},
			mockQuery: mock.ExpectQuery("select * from user where id = ?").WithArgs(1).WillReturnRows(rows),
		},
		{
			desc: "Case 2: Failure case",
			id:   2, expected: nil,
			mockQuery: mock.ExpectQuery("select * from user where id = ?").WithArgs(2).WillReturnError(errors.New("Invalid Id")),
		},
	}

	userStore := New(db)

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			u, _ := userStore.GetUserById(test.id)

			if !reflect.DeepEqual(u, test.expected) {
				t.Errorf("Expected: %v, Got: %v", test.expected, u)
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "age"}).AddRow(
		1, "tejas", "tejas@gmail.com", "9876543210", 23,
	).AddRow(
		2, "nithin", "nithin@gmail.com", "9988776655", 18,
	)

	testCases := []struct {
		desc      string
		expected  []models.User
		mockQuery *sqlmock.ExpectedQuery
	}{
		{
			desc: "Case 1: Success case",
			expected: []models.User{
				{
					Id:    1,
					Name:  "tejas",
					Email: "tejas@gmail.com",
					Phone: "9876543210",
					Age:   23,
				},
				{
					Id:    2,
					Name:  "nithin",
					Email: "nithin@gmail.com",
					Phone: "9988776655",
					Age:   18,
				},
			},
			mockQuery: mock.ExpectQuery("select * from user").WillReturnRows(rows),
		},
		{
			desc:      "Case 2: Failure Case",
			expected:  []models.User{},
			mockQuery: mock.ExpectQuery("select * from user").WillReturnError(errors.New("Cannot fetch users")),
		},
	}

	userStore := New(db)

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			userList, _ := userStore.GetUsers()

			if !reflect.DeepEqual(userList, test.expected) {
				t.Errorf("Expected: %v, Got: %v", test.expected, userList)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	defer db.Close()

	testCases := []struct {
		desc     string
		id       int
		expected int
		mockCall *sqlmock.ExpectedExec
	}{
		{
			desc:     "Case 1: Success case",
			id:       1,
			expected: 1,
			mockCall: mock.ExpectExec("update user set name = ?, email = ?, phone = ?, age = ? where id = ?").WithArgs("vivek", "vivek@gmail.com", "9988776655", 23, 1).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:     "Case 2: Failure case",
			id:       2,
			expected: 0,
			mockCall: mock.ExpectExec("update user set name = ?, email = ?, phone = ?, age = ? where id = ?").WithArgs("vivek", "vivek@gmail.com", "9988776655", 23, 2).WillReturnError(errors.New("Invalid Id")),
		},
	}
	userStore := New(db)

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			affectedRows, _ := userStore.UpdateUser(test.id, models.User{Name: "vivek", Email: "vivek@gmail.com", Phone: "9988776655", Age: 21})

			if affectedRows != test.expected {
				t.Errorf("Expected: %d, Got: %d", test.expected, affectedRows)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer db.Close()

	testCases := []struct {
		desc     string
		id       int
		expected int
		mockCall *sqlmock.ExpectedExec
	}{
		{
			desc:     "Case 1: Success case",
			id:       1,
			expected: 1,
			mockCall: mock.ExpectExec("delete from user where id = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:     "Case 2: Failure case",
			id:       2,
			expected: 0,
			mockCall: mock.ExpectExec("delete from user where id = ?").WithArgs(2).WillReturnError(errors.New("Invalid Id")),
		},
		{
			desc:     "Case 3",
			id:       2,
			expected: 0,
			mockCall: mock.ExpectExec("DELETE FROM user WHERE id = ?").WithArgs(2).WillReturnResult(sqlmock.NewResult(0, 0)),
		},
	}

	userStore := New(db)

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			affectedRows, _ := userStore.DeleteUser(test.id)

			if affectedRows != test.expected {
				t.Errorf("Expected: %d, Got: %d", test.expected, affectedRows)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer db.Close()

	testUser := models.User{
		Name:  "vivek",
		Email: "vivek@gmail.com",
		Phone: "9988776655",
		Age:   23,
	}

	testCases := []struct {
		desc     string
		expected int
		mockCall *sqlmock.ExpectedExec
	}{
		{
			desc:     "Case 1: Success Case",
			expected: 1,
			mockCall: mock.ExpectExec("insert into user(name, email, phone, age) values(?, ?, ?, ?)").WithArgs("vivek", "vivek@gmail.com", "9988776655", 23).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:     "Case 2: Failure Case",
			expected: 0,
			mockCall: mock.ExpectExec("insert into user(name, email, phone, age) values(?, ?, ?, ?)").WithArgs("vivek", "vivek@gmail.com", "9988776655", 23).WillReturnError(errors.New("Email Id already exist")),
		},
	}

	userStore := New(db)

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			user, _ := userStore.CreateUser(testUser)

			if user != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, user)
			}
		})
	}
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "age"}).AddRow(
		1, "tejas", "vivek@gmail.com", "9876543210", 21,
	)

	testCases := []struct {
		desc     string
		email    string
		expected bool
		mockCall *sqlmock.ExpectedQuery
	}{
		{
			desc:     "Case 1: Success Case",
			email:    "vivek@gmail.com",
			expected: false,
			mockCall: mock.ExpectQuery("select * from user where email = ?").WithArgs("vivek@gmail.com").WillReturnRows(rows),
		},
		{
			desc:     "Case 2: Failure case",
			email:    "tejas@gmail.com",
			expected: true,
			mockCall: mock.ExpectQuery("select * from user where email = ?").WithArgs("tejas@gmail.com").WillReturnError(errors.New("Email is already in use")),
		},
	}

	userStore := New(db)

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			isValid := userStore.GetUserByEmail(test.email)

			if isValid != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, isValid)
			}
		})
	}
}
