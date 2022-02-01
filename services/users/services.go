package users

import (
	"errors"
	"user-crud/models"
	"user-crud/services"
	"user-crud/services/validations"
	"user-crud/stores"
)

type userServ struct {
	u stores.User
}

func New(st stores.User) services.User {
	return &userServ{st}
}
func (u userServ) Insert(usr *models.User) (*models.User, error) {
	if !validations.ValidateEmail(usr.Email) {
		return &models.User{}, errors.New("invalid emailID")
	}
	if !validations.ValidatePhone(usr.Phone) {
		return &models.User{}, errors.New("invalid Phone number")
	}
	user, err := u.u.Insert(usr)
	return user, err
}

func (u userServ) GetAll() ([]*models.User, error) {
	var users []*models.User
	users, err := u.u.GetAll()
	return users, err
}

func (u userServ) GetById(id int) (*models.User, error) {
	usr, err := u.u.GetById(id)
	return usr, err
}

func (u userServ) Update(usr *models.User) (*models.User, error) {
	if !validations.ValidateEmail(usr.Email) {
		return &models.User{}, errors.New("invalid emailID")
	}
	if !validations.ValidatePhone(usr.Phone) {
		return &models.User{}, errors.New("invalid Phone number")
	}
	usr, err := u.u.Update(usr)
	return usr, err
}

func (u userServ) Delete(id int) error {
	// check id exists
	_, err := u.u.GetById(id)

	if err != nil {
		return err
	}
	err = u.u.Delete(id)
	if err != nil {
		return errors.New("unable to delete user")
	}
	return nil
}
