package users

import (
	"errors"
	"fmt"
	"layer/user/models"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Println(err)
	}
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
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Println(err)
	}
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
