package users

import (
	"errors"
	"layer/user/models"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetById(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "age"}).AddRow(
		1, "Naruto", "naruto@gmail.com", "9999999999", 21,
	)

	tests := []struct {
		desc      string
		id        int
		expected  *models.User
		mockQuery *sqlmock.ExpectedQuery
	}{
		{desc: "Case1", id: 1, expected: &models.User{Id: 1, Name: "Naruto", Email: "naruto@gmail.com", Phone: "9999999999", Age: 21}, mockQuery: mock.ExpectQuery("SELECT * FROM user WHERE id = ?").WithArgs(1).WillReturnRows(rows)},
		{desc: "Case2", id: 2, expected: nil, mockQuery: mock.ExpectQuery("SELECT * FROM user WHERE id = ?").WithArgs(2).WillReturnError(errors.New("Invalid Id"))},
	}

	userStore := New(db)

	for _, test := range tests {
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
		1, "Naruto", "naruto@gmail.com", "9999999999", 18,
	).AddRow(
		2, "Itachi", "itachi@gmail.com", "8320578360", 24,
	)

	tests := []struct {
		desc      string
		expected  []models.User
		mockQuery *sqlmock.ExpectedQuery
	}{
		{
			desc: "Case1",
			expected: []models.User{
				{Id: 1, Name: "Naruto", Email: "naruto@gmail.com", Phone: "9999999999", Age: 18},
				{Id: 2, Name: "Itachi", Email: "itachi@gmail.com", Phone: "8320578360", Age: 24},
			},
			mockQuery: mock.ExpectQuery("SELECT * FROM user").WillReturnRows(rows),
		},
		{
			desc:      "Case2",
			expected:  []models.User{},
			mockQuery: mock.ExpectQuery("SELECT * FROM user").WillReturnError(errors.New("Cannot fetch users")),
		},
	}

	userStore := New(db)

	for _, test := range tests {
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

	tests := []struct {
		desc          string
		id            int
		user          models.User
		expectedError error
		mockCall      []*sqlmock.ExpectedExec
	}{
		{
			desc: "Case1",
			id:   1,
			user: models.User{
				Name:  "Ridhdhish",
				Email: "rid@gmail.com",
				Phone: "8320578360",
				Age:   21,
			},
			expectedError: nil,
			mockCall: []*sqlmock.ExpectedExec{
				mock.ExpectExec("UPDATE user SET name = ?, email = ?, phone = ?, age = ? WHERE id = ?").WithArgs("Ridhdhish", "rid@gmail.com", "8320578360", 21, 1).WillReturnResult(sqlmock.NewResult(1, 1)),
			},
		},
		{
			desc: "Case2",
			id:   1,
			user: models.User{
				Name: "Ridhdhish",
			},
			expectedError: errors.New("Internal server error"),
			mockCall: []*sqlmock.ExpectedExec{
				mock.ExpectExec("UPDATE user SET name = ?, WHERE id = ?").WithArgs("Ridhdhish", 1).WillReturnError(errors.New("Internal server error")),
			},
		},
	}
	userStore := New(db)

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			// mock.ExpectBegin()
			err := userStore.UpdateUser(test.id, test.user)

			if test.expectedError != nil && errors.Is(err, test.expectedError) {
				t.Errorf("Expected: %d, Got: %d", test.expectedError, err)
			}
			// mock.ExpectCommit()
		})
	}
}

func TestDeleteUser(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer db.Close()

	tests := []struct {
		desc          string
		id            int
		expectedError error
		mockCall      *sqlmock.ExpectedExec
	}{
		{
			desc:          "Case1",
			id:            1,
			expectedError: nil,
			mockCall:      mock.ExpectExec("DELETE FROM user WHERE id = ?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:          "Case2",
			id:            2,
			expectedError: errors.New("Could not delete user for given id"),
			mockCall:      mock.ExpectExec("DELETE FROM user WHERE id = ?").WithArgs(2).WillReturnError(errors.New("Could not delete user for given id")),
		},
	}

	userStore := New(db)

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := userStore.DeleteUser(test.id)

			if test.expectedError != nil && errors.Is(err, test.expectedError) {
				t.Errorf("Expected: %v, Got: %v", test.expectedError.Error(), err.Error())
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer db.Close()

	testUser := models.User{Name: "Ridhdhish", Email: "ridhdhish@gmail.com", Phone: "8320578360", Age: 21}

	tests := []struct {
		desc     string
		expected int
		mockCall *sqlmock.ExpectedExec
	}{
		{
			desc:     "Case1",
			expected: 1,
			mockCall: mock.ExpectExec("INSERT INTO user(name, email, phone, age) VALUES(?, ?, ?, ?)").WithArgs("Ridhdhish", "ridhdhish@gmail.com", "8320578360", 21).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:     "Case2",
			expected: 0,
			mockCall: mock.ExpectExec("INSERT INTO user(name, email, phone, age) VALUES(?, ?, ?, ?)").WithArgs("Ridhdhish", "ridhdhish@gmail.com", "8320578360", 21).WillReturnError(errors.New("Email Id already exist")),
		},
	}

	userStore := New(db)

	for _, test := range tests {
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

	rows := sqlmock.NewRows([]string{"id"}).AddRow(
		1,
	)

	tests := []struct {
		desc     string
		email    string
		expected bool
		mockCall *sqlmock.ExpectedQuery
	}{
		{
			desc:     "Case1",
			email:    "ridhdhish@gmail.com",
			expected: false,
			mockCall: mock.ExpectQuery("SELECT id FROM user WHERE email = ?").WithArgs("ridhdhish@gmail.com").WillReturnRows(rows),
		},
		{
			desc:     "Case2",
			email:    "naruto@gmail.com",
			expected: true,
			mockCall: mock.ExpectQuery("SELECT id FROM user WHERE email = ?").WithArgs("naruto@gmail.com").WillReturnError(errors.New("Email is already in use")),
		},
	}

	userStore := New(db)

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			isValid := userStore.GetUserByEmail(test.email)

			if isValid != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, isValid)
			}
		})
	}
}
