package users

import (
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/tejas/user-crud/models"
)

func TestGetUserById(t *testing.T) {

	db, mock, _ := sqlmock.New()

	defer db.Close()

	query := "select id, name, email, phone, age form user where id = ?"

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
			desc: "Case 1: Success Case",
			id:   1,
			expected: &models.User{
				Id:    1,
				Name:  "tejas",
				Email: "tejas@gmail.com",
				Phone: "9876543210",
				Age:   23,
			},
			mockQuery: mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows),
		},
		{
			desc:      "Case 2: failure case",
			id:        4,
			expected:  nil,
			mockQuery: mock.ExpectQuery(query).WithArgs(4).WillReturnError(errors.New("error fetching while get users by id, invalid id")),
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
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Print(err)
	}

	defer db.Close()

	query := "select id, name, email, phone, age form user"

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "age"}).
		AddRow(1, "tejas", "tejas@gmail.com", "9876543210", 23).
		AddRow(2, "nithin", "nithin@gmail.com", "9988776655", 18)

	testCases := []struct {
		desc      string
		expected  []models.User
		mockQuery *sqlmock.ExpectedQuery
	}{
		{
			desc: "Case 1: Success Case",
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
			mockQuery: mock.ExpectQuery(query).WillReturnRows(rows),
		},
		{
			desc:      "Case 2: Failure Case",
			expected:  nil,
			mockQuery: mock.ExpectQuery(query).WillReturnError(errors.New("cannot fetch users")),
		},
	}

	userStore := New(db)

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			u, _ := userStore.GetUsers()

			if !reflect.DeepEqual(u, test.expected) {
				t.Errorf("Expected: %v, Got: %v", test.expected, u)
			}

		})
	}
}

func TestCreateUser(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("error connecting to mock database : %v", err)
	}
	defer db.Close()

	userStore := New(db)

	query := "insert into user values (?,?,?,?,?)"

	testCases := []struct {
		desc        string
		input       models.User
		mockQuery   *sqlmock.ExpectedExec
		expectedErr error
	}{
		{
			desc: "Case 1: Success Case",
			input: models.User{
				Id:    3,
				Name:  "vivek",
				Email: "vivek@gmail.com",
				Phone: "8979695949",
				Age:   23,
			},
			mockQuery: mock.ExpectExec(query).
				WithArgs(3, "vivek", "vivek@gmail.com", "8979695949", 23).
				WillReturnResult(sqlmock.NewResult(1, 1)),
			expectedErr: nil,
		},

		{
			desc: "Case 2: Failure Case",
			input: models.User{
				Id:    3,
				Name:  "vivek",
				Email: "vivek@gmail.com",
				Phone: "8979695949",
				Age:   23,
			},
			mockQuery: mock.ExpectExec(query).
				WithArgs(3, "vivek", "vivek@gmail.com", "8979695949", 23).
				WillReturnError(errors.New("invalid data")),
			expectedErr: errors.New("invalid data"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+tc.desc, func(t *testing.T) {
			err := userStore.CreateUsers(tc.input)
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("Expected: %v, Got: %v, ", tc.expectedErr, err)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("error connecting to mock database : %v", err)
	}
	defer db.Close()

	store := New(db)

	testCases := []struct {
		desc        string
		input       models.User
		expectedErr error
	}{

		{
			desc: "Case 1: Success case",
			input: models.User{
				Id:    1,
				Name:  "madhu",
				Email: "madhu@gmail.com",
				Phone: "8899776655",
				Age:   22,
			},
			expectedErr: nil,
		},

		{
			desc: "Case 2: Failure Case",
			input: models.User{
				Name: "Hello",
			},
			expectedErr: errors.New("error"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+tc.desc, func(t *testing.T) {

			query := "update user set"

			if tc.expectedErr == nil {
				fields, values := preCheck(tc.input)
				query += fields + " where id = ?"

				mock.ExpectExec(query).WithArgs(values[0], values[1], values[2], values[3], values[4]).WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectExec(query).WithArgs().WillReturnError(errors.New("error"))
			}
			err := store.UpdateUser(tc.input)
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("Expected: %v, Got: %v, ", tc.expectedErr, err)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("error connecting to mock database : %v", err)
	}
	defer db.Close()

	store := New(db)

	query := "delete from user where id = ?"

	testCases := []struct {
		desc        string
		id          int
		mockQuery   *sqlmock.ExpectedExec
		expectedErr error
	}{
		{
			desc: "Case 1: Success Case",
			id:   2,
			mockQuery: mock.ExpectExec(query).WithArgs(2).
				WillReturnResult(sqlmock.NewResult(1, 1)),
			expectedErr: nil,
		},
		{
			desc: "Case 2: Failure Case",
			mockQuery: mock.ExpectExec(query).
				WillReturnError(errors.New("cannot delete user data for the given id")),
			expectedErr: errors.New("cannot delete user data for the given id"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+tc.desc, func(t *testing.T) {
			err := store.DeleteUser(tc.id)
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("Expected: %v, Got: %v, ", tc.expectedErr, err)
			}
		})
	}
}
