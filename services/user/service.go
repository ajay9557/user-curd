package user

import (
	"errors"
	"go_lang/Assignment/user-curd/models"
	"go_lang/Assignment/user-curd/services"
	"go_lang/Assignment/user-curd/stores"
)

type UserService struct {
	storeInterface stores.Crud
}

func New(si stores.Crud) services.Services {
	return &UserService{si}
}

func (userservice *UserService) AddUser(usr models.User) error {

	if !IsNumberValid(usr.Phone) {
		return errors.New("INVALID PHONE NUMBER")
	}

	if !IsUniqueEmail(usr.Email) {
		return errors.New("INVALID EMAIL")
	}
	result := userservice.storeInterface.Create(usr)
	if result != nil {
		return errors.New("FAILED TO ADD USER")
	}
	return nil
}

func (userservice *UserService) GetUser(id int) (models.User, error) {

	user, _ := userservice.storeInterface.ReadOne(id)

	return user, nil
}

func (userservice *UserService) GetAllUser() ([]models.User, error) {
	return userservice.storeInterface.ReadAll()
}

func (userservice *UserService) UpdateUser(usr models.User) error {

	if !IsNumberValid(usr.Phone) {
		return errors.New("INVALID PHONE NUMBER")
	}

	if !IsUniqueEmail(usr.Email) {
		return errors.New("INVALID EMAIL")
	}

	err := userservice.storeInterface.Update(usr)
	if err != nil {
		return err
	}
	return nil
}

func (userservice *UserService) DeleteUser(id int) error {

	err := userservice.storeInterface.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
