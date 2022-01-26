package users

import (
	"user-curd/models"
	"user-curd/services"
	"user-curd/stores"
	"errors"
)

type UserServiceHandler struct {
	stores stores.Store //Users interface
}

func New(store stores.Store) services.Services {
	return UserServiceHandler{stores: store}
}

func (u UserServiceHandler) EmailValidation(email string) (bool, error) {
	ok, err := u.stores.GetEmail(email)
	if err != nil {
		return ok, errors.New("error generated")
	}
	return ok, nil //true if email not present in db
}

func (u UserServiceHandler) InsertUserDetails(user models.User) error {

	err := u.stores.InsertUser(user)
	if err != nil {
		return errors.New("error generated")
	}
	return nil
}

func (u UserServiceHandler) FetchAllUserDetails() ([]models.User, error) {
	var allUsers []models.User
	allUsers, err := u.stores.FetchAllUsers()
	if err != nil {
		return allUsers, errors.New("data fetching error")
	}
	return allUsers, nil
}

func (u UserServiceHandler) FetchUserDetailsById(id int) (models.User, error) {
	user, err := u.stores.FetchUserById(id)
	if err != nil {
		return user, errors.New("data fetching error")
	}
	return user, nil
}

func (u UserServiceHandler) UpdateUserDetails(user models.User) error {
	err := u.stores.UpdateUser(user)
	if err != nil {
		return errors.New("error generated")
	}
	return nil
}

func (u UserServiceHandler) DeleteUserDetailsById(id int) error {
	err := u.stores.DeleteUserById(id)
	if err != nil {
		return errors.New("error generated")
	}
	return nil
}
