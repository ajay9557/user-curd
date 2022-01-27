package service

import (
	"errors"
	"strconv"
	"zopsmart/Task/models"
	"zopsmart/Task/services"
	"zopsmart/Task/stores"
)

type UserServiceHandler struct {
	stores store.Store
}

func New(store store.Store) services.Services {
	return UserServiceHandler { stores:store}
}



func (u UserServiceHandler) GetUserById(Id int) (models.User, error) {

	res,err := u.stores.GetUserById(Id)
	if err != nil {
		return res, errors.New("Error occured while getting the data")
	}

	return res, nil
}

func (u UserServiceHandler) DeletebyId(Id int) error {
    err := u.stores.Delete(Id)
	if err != nil {
		return errors.New("Error occured while deleting data")
	}

	return nil
}

func (u UserServiceHandler) UpdatebyId(Id int, Phone string) error {
	err := u.stores.Update(Id, Phone)

	if err != nil {
		return errors.New("Error")
	}

	return err
}

func (u UserServiceHandler) ValidateEmail(mail string) (bool, error) {
	ok, err := u.stores.GetMail(mail)
	if err != nil {
		return ok, errors.New("Error")
	}
	return ok, nil
}

func (u UserServiceHandler) ValidateId(Id int) (bool, error) {
	if Id < 1 {
		return false, errors.New("Id cannot be less than zero")
	}
	id := strconv.Itoa(Id)

	if id == "" {
		return false, errors.New("Id cannot be empty")
	}
	return true, nil
}

