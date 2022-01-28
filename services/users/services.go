package users

import (
	"errors"
	"reflect"
	"regexp"

	"github.com/tejas/user-crud/models"
	"github.com/tejas/user-crud/services"
	"github.com/tejas/user-crud/stores"
)

type User struct {
	u stores.User
}

func New(u stores.User) services.User {
	return &User{u}
}

func (st *User) GetUserById(id int) (models.User, error) {
	user, err := st.u.GetUserById(id)

	if err != nil {
		return models.User{}, errors.New("Cannot fetch user for given id")
	}

	return *user, nil
}

func (st *User) GetUsers() ([]models.User, error) {
	users, err := st.u.GetUsers()

	if err != nil {
		return []models.User{}, errors.New("Cannot fetch users")
	}

	return users, nil

}

func (st *User) UpdateUser(id int, user models.User) (int, error) {
	updateid, err := st.u.UpdateUser(id, user)

	if err != nil {
		return 0, errors.New("Could not able to update user for given id")
	}

	return updateid, nil
}

func (st *User) DeleteUser(id int) (int, error) {
	rowcount, err := st.u.DeleteUser(id)

	if err != nil {
		return 0, errors.New("Could not able to delete user for given id")
	}

	return rowcount, nil
}

func (st *User) CreateUser(user models.User) (int, error) {

	if reflect.DeepEqual(user, models.User{}) {
		return 0, errors.New("Need user data to create new user")
	}

	validEmail := validateEmail(user.Email)

	if !validEmail {
		return 0, errors.New("Invalid email address")
	}

	verifyEmail := st.u.GetUserByEmail(user.Email)

	if !verifyEmail {
		return 0, errors.New("Email id is already in use")
	}

	lastInsertedId, err := st.u.CreateUser(user)

	if err != nil {
		return 0, errors.New("Could not able to create new user")
	}

	return lastInsertedId, nil
}

func validateEmail(email string) bool {
	validEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return validEmail.MatchString(email)
}
