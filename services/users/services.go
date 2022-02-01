package users

import (
	"errors"
	"layer/user/models"
	"layer/user/services"
	"layer/user/stores"
	"reflect"
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
		return models.User{}, err
	}

	return *user, nil
}

func (st *User) GetUsers() ([]*models.User, error) {
	users, err := st.u.GetUsers()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (st *User) UpdateUser(id int, user models.User) (*models.User, error) {

	// Check for negative user id
	if id < 0 {
		return nil, errors.New("User id should be greater than 0")
	}

	// Check for user exists or not for given user id
	_, err := st.u.GetUserById(id)
	if err != nil {
		return nil, err
	}

	err = st.u.UpdateUser(id, user)

	if err != nil {
		return nil, err
	}

	// Fetching user after update
	updatedUser, _ := st.u.GetUserById(id)

	return updatedUser, nil
}

func (st *User) DeleteUser(id int) error {

	// Checking user is exist or not before deleing
	_, err := st.u.GetUserById(id)
	if err != nil {
		return err
	}

	err = st.u.DeleteUser(id)

	if err != nil {
		return err
	}

	return nil
}

func (st *User) CreateUser(user models.User) (*models.User, error) {

	if reflect.DeepEqual(user, models.User{}) {
		return nil, errors.New("Need user data to create new user")
	}

	validEmail := validateEmail(user.Email)

	if !validEmail {
		return nil, errors.New("Invalid email address")
	}

	validPhone := validatePhone(user.Phone)

	if !validPhone {
		return nil, errors.New("Invalid phone number")
	}

	isValid := st.u.GetUserByEmail(user.Email)

	if !isValid {
		return nil, errors.New("Email id is already exist")
	}

	id, err := st.u.CreateUser(user)

	if err != nil {
		return nil, err
	}

	updatedUser, _ := st.GetUserById(id)

	return &updatedUser, nil
}
