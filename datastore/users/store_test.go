package users

import (
	"errors"
	"log"
	"reflect"
	"strconv"
	"testing"
	"user-curd/entities"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetUserById(t *testing.T) {

	// Create a mock db
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("error connecting to mock database : %v", err)
	}
	defer db.Close()

	store := New(db)

	// define the query
	query := "select id, name, email, phone, age from user where id = ?"

	testCases := []struct {
		caseId      int
		id          int
		mockQuery   *sqlmock.ExpectedQuery
		expectedErr error
		expectedOut *entities.User
	}{
		// Success Case
		{
			caseId: 1,
			id:     1,
			mockQuery: mock.ExpectQuery(query).WithArgs(1).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "age"}).
					AddRow(1, "john", "john21@example.com", "9728810299", 21)),
			expectedErr: nil,
			expectedOut: &entities.User{
				Id:    1,
				Name:  "john",
				Email: "john21@example.com",
				Phone: "9728810299",
				Age:   21,
			},
		},
		// Error case
		{
			caseId: 2,
			id:     3,
			mockQuery: mock.ExpectQuery(query).WithArgs(3).
				WillReturnError(errors.New("error fetching from database, id not found")),
			expectedErr: errors.New("error fetching from database, id not found"),
			expectedOut: &entities.User{},
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {

			resp, err := store.GetUserById(tc.id)
			if !reflect.DeepEqual(resp, tc.expectedOut) {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedOut, resp)
			}
			if err != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedErr, err)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {

	// Create a mock db
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("error connecting to mock database : %v", err)
	}
	defer db.Close()

	// get store struct
	store := New(db)

	// define query
	query := "select id,name,email,phone,age from user"

	testCases := []struct {
		caseId      int
		mockQuery   *sqlmock.ExpectedQuery
		expectedErr error
		expectedOut []*entities.User
	}{
		// Success Case
		{
			caseId:      1,
			expectedErr: nil,
			mockQuery: mock.ExpectQuery(query).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone", "age"}).
					AddRow(1, "john", "john19@example.com", "9823218772", 21).
					AddRow(2, "jess", "jessH99@example.com", "9821328772", 19)),
			expectedOut: []*entities.User{
				{
					Id:    1,
					Name:  "john",
					Email: "john19@example.com",
					Phone: "9823218772",
					Age:   21,
				},
				{
					Id:    2,
					Name:  "jess",
					Email: "jessH99@example.com",
					Phone: "9821328772",
					Age:   19,
				},
			},
		},

		// Error case
		{
			caseId: 2,
			mockQuery: mock.ExpectQuery(query).
				WillReturnError(errors.New("error fetching data from database")),
			expectedErr: errors.New("error fetching data from database"),
			expectedOut: []*entities.User{},
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {

			resp, _ := store.GetAllUsers()
			if !reflect.DeepEqual(resp, tc.expectedOut) {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedOut, resp)
			}
			if err != nil && tc.expectedErr != nil && err.Error() != tc.expectedErr.Error() {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedErr, err)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {

	// Create a mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("error connecting to mock database : %v", err)
	}
	defer db.Close()

	store := New(db)

	// define the query
	query := "insert into user values (?,?,?,?,?)"

	testCases := []struct {
		caseId      int
		input       entities.User
		mockQuery   *sqlmock.ExpectedExec
		expectedErr error
	}{
		// Success case
		{
			caseId: 1,
			input: entities.User{
				Id:    3,
				Name:  "James",
				Email: "james89@example.com",
				Phone: "7897895621",
				Age:   34,
			},
			mockQuery: mock.ExpectExec(query).
				WithArgs(3, "James", "james89@example.com", "7897895621", 34).
				WillReturnResult(sqlmock.NewResult(1, 1)),
			expectedErr: nil,
		},
		// Error case (duplicate entry)
		{
			caseId: 2,
			input: entities.User{
				Id:    3,
				Name:  "James",
				Email: "james89@example.com",
				Phone: "7897895621",
				Age:   34,
			},
			mockQuery: mock.ExpectExec(query).
				WithArgs(3, "James", "james89@example.com", "7897895621", 34).
				WillReturnError(errors.New("illegal entry in the database denied")),
			expectedErr: errors.New("illegal entry in the database denied"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {
			err := store.CreateUser(tc.input)
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedErr, err)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {

	// Create a mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("error connecting to mock database : %v", err)
	}
	defer db.Close()

	// get store struct
	store := New(db)

	testCases := []struct {
		caseId      int
		input       entities.User
		expectedErr error
	}{
		// Success case with all fields updated
		{
			caseId: 2,
			input: entities.User{
				Id:    1,
				Name:  "jane",
				Email: "jane123@example.com",
				Phone: "9808912345",
				Age:   22,
			},
			expectedErr: nil,
		},
		// Error case
		{
			caseId: 3,
			input: entities.User{
				Name: "Hello",
			},
			expectedErr: errors.New("error, no id provided, cannot update"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {
			// define base query
			query := "update user set"

			if tc.expectedErr == nil {
				fields, values := formQuery(tc.input)
				query += fields + " where id = ?"

				mock.ExpectExec(query).WithArgs(values[0], values[1], values[2], values[3], values[4]).WillReturnResult(sqlmock.NewResult(1, 1))
			} else {
				mock.ExpectExec(query).WithArgs().WillReturnError(errors.New("error, no id provided, cannot update"))
			}
			err := store.UpdateUser(tc.input)
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedErr, err)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	// Create a mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Printf("error connecting to mock database : %v", err)
	}
	defer db.Close()

	store := New(db)

	// define the query
	query := "delete from user where id = ?"

	testCases := []struct {
		caseId      int
		id          int
		mockQuery   *sqlmock.ExpectedExec
		expectedErr error
	}{
		// Success Case
		{
			caseId: 1,
			id:     2,
			mockQuery: mock.ExpectExec(query).WithArgs(2).
				WillReturnResult(sqlmock.NewResult(1, 1)),
			expectedErr: nil,
		},
		// Error case
		{
			caseId: 2,
			mockQuery: mock.ExpectExec(query).
				WillReturnError(errors.New("error, not able to delete data")),
			expectedErr: errors.New("error, not able to delete data"),
		},
	}

	for _, tc := range testCases {
		t.Run("testing "+strconv.Itoa(tc.caseId), func(t *testing.T) {
			err := store.DeleteUser(tc.id)
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", tc.caseId, tc.expectedErr, err)
			}
		})
	}
}
