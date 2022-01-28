package Users

import (
	"Icrud/TModels"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUserById(t *testing.T) {
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
		expected  TModels.User
		mockQuery *sqlmock.ExpectedQuery
	}{
		{desc: "Case1", id: 1, expected: TModels.User{Id: 1, Name: "Naruto", Email: "naruto@gmail.com", Phone: "9999999999", Age: 21}, mockQuery: mock.ExpectQuery("select id,name,email,phone,age from HUser where id = ?").WithArgs(1).WillReturnRows(rows)},
		{desc: "Case2", id: 2, expected: TModels.User{}, mockQuery: mock.ExpectQuery("SELECT * from user where id = ?").WithArgs(2).WillReturnError(errors.New("Invalid Id"))},
	}

	userStore := New(db)

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			u, _ := userStore.UserById(test.id)

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
		expected  []TModels.User
		mockQuery *sqlmock.ExpectedQuery
	}{
		{
			desc: "Case1",
			expected: []TModels.User{
				{Id: 1, Name: "Naruto", Email: "naruto@gmail.com", Phone: "9999999999", Age: 18},
				{Id: 2, Name: "Itachi", Email: "itachi@gmail.com", Phone: "8320578360", Age: 24},
			},
			mockQuery: mock.ExpectQuery("select * from HUser").WillReturnRows(rows),
		},
		{
			desc:      "Case2",
			expected:  []TModels.User{},
			mockQuery: mock.ExpectQuery("select * from HUser").WillReturnError(errors.New("Cannot fetch users")),
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

func TestDeleteUserById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	tests := []struct {
		desc     string
		id       int
		expected int
		mockCall *sqlmock.ExpectedExec
	}{
		{
			desc:     "Case1",
			id:       1,
			expected: 1,
			mockCall: mock.ExpectExec("delete from HUser where id=?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:     "Case2",
			id:       2,
			expected: 0,
			mockCall: mock.ExpectExec("delete from HUser where id=?").WithArgs(2).WillReturnError(errors.New("Invalid Id")),
		},
	}

	userStore := New(db)

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			affectedRows, _ := userStore.DeleteUserById(test.id)

			if affectedRows != test.expected {
				t.Errorf("Expected: %d, Got: %d", test.expected, affectedRows)
			}
		})
	}
}

func TestInsertUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	testUser := TModels.User{Name: "Ridhdhish", Email: "ridhdhish@gmail.com", Phone: "8320578360", Age: 21}

	tests := []struct {
		desc     string
		expected int
		mockCall *sqlmock.ExpectedExec
	}{
		{desc: "Case1", expected: 1, mockCall: mock.ExpectExec("Insert into HUser (name,email,phone,age) values (?,?,?,?)").WithArgs("Ridhdhish", "ridhdhish@gmail.com", "8320578360", 21).WillReturnResult(sqlmock.NewResult(1, 1))},
		{desc: "Case2", expected: -1, mockCall: mock.ExpectExec("Insert into HUser (name,email,phone,age) values (?,?,?,?)").WithArgs("Ridhdhish", "ridhdhish1@gmail.com", "8320578360", 21).WillReturnError(errors.New("Connection Lost"))},
	}

	userStore := New(db)

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			user, _ := userStore.InsertUser(testUser)

			if user != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, user)
			}
		})
	}
}

func TestUpdateUserById(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	tests := []struct {
		desc     string
		id       int
		expected int
		mockCall *sqlmock.ExpectedExec
	}{
		{
			desc:     "Case1",
			id:       1,
			expected: 1,
			mockCall: mock.ExpectExec("Update HUser set  name = ?,email = ?,phone = ?,age = ? where id =?").WithArgs("Ridhdhish", "ridhdhish@gmail.com", "8320578360", 21, 1).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			desc:     "Case2",
			id:       2,
			expected: -1,
			mockCall: mock.ExpectExec("Update HUser set  name = ?,email = ?,phone = ?,age = ? where id =?").WithArgs("Ridhdhish", "ridhdhish@gmail.com", "8320578360", 21, 2).WillReturnError(errors.New("Invalid Id")),
		},
	}
	userStore := New(db)

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			affectedRows, _ := userStore.UpdateUserById(TModels.User{Name: "Ridhdhish", Email: "ridhdhish@gmail.com", Phone: "8320578360", Age: 21}, test.id)

			if affectedRows != test.expected {
				t.Errorf("Expected: %d, Got: %d", test.expected, affectedRows)
			}
		})
	}
}

func TestGetEmail(t *testing.T) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "age"}).AddRow(
		1, "Naruto", "ridhdhish@gmail.com", "9999999999", 21,
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
			mockCall: mock.ExpectQuery("Select email from HUser where email=?").WithArgs("ridhdhish@gmail.com").WillReturnRows(rows),
		},
		{
			desc:     "Case2",
			email:    "naruto@gmail.com",
			expected: true,
			mockCall: mock.ExpectQuery("Select email from HUser where email=?").WithArgs("naruto@gmail.com").WillReturnError(errors.New("Email is already in use")),
		},
	}

	userStore := New(db)

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			isEmailExist, _ := userStore.GetEmail(test.email)
			if test.expected != isEmailExist {
				t.Errorf("Expected: %t, Got: %t", test.expected, isEmailExist)
			}
		})
	}

}
