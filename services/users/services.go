package users

import (
	"errors"

	"github.com/tejas/user-crud/models"
	"github.com/tejas/user-crud/services"
	"github.com/tejas/user-crud/store"
)

type User struct {
	u store.User
}

func New(u store.User) services.User {
	return &User{u: u}
}

func (st *User) GetUserById(id int) (*models.User, error) {

	if ValidId(id) {
		row, err := st.u.GetUserById(id)

		if err != nil {
			return nil, errors.New("cannot fetch user by id")
		}

		return row, nil
	}

	return nil, errors.New("cannot fetch user by id")

}

func (st *User) GetUsers() ([]*models.User, error) {
	users, err := st.u.GetUsers()
	if err != nil {
		return nil, errors.New("cannot fetch users")
	}

	return users, nil
}

func (st *User) UpdateUser(user models.User) error {

	if !ValidId(user.Id) {
		return errors.New("invalid id")
	}

	if !ValidEmail(user.Email) {
		return errors.New("invalid email")
	}

	if !ValidPhone(user.Phone) {
		return errors.New("invalid phone number")
	}

	u1, _ := st.u.GetUserById(user.Id)

	if u1 == nil {
		return errors.New("user not exits")
	}

	users := st.u.UpdateUser(user)

	return users
}

func (st *User) CreateUser(user models.User) error {
	if !ValidId(user.Id) {
		return errors.New("invalid id")
	}

	if !ValidEmail(user.Email) {
		return errors.New("invalid email")
	}

	if !ValidPhone(user.Phone) {
		return errors.New("invalid phone number")
	}

	createuser := st.u.CreateUsers(user)

	return createuser

}

func (st *User) DeleteUser(id int) error {

	if !ValidId(id) {
		return errors.New("invalid id")
	}

	var user models.User


	u1, _ := st.u.GetUserById(user.Id)

	if u1 == nil {
		return errors.New("user not exits")
	}

	deleteuser := st.u.DeleteUser(id)

	return deleteuser
}
