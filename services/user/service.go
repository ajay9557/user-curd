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

func New(si stores.Crud) Services {
	return &UserService{si}
}

func (userservice *UserService) AddUser(usr models.User) error {

	if !services.IsNumberValid(usr.Phone) {
		return errors.New("INVALID PHONE NUMBER")
	}

	if !services.IsUniqueEmail(usr.Email) {
		return errors.New("INVALID EMAIL")
	}
	result := userservice.storeInterface.Create(usr.Name, usr.Email, usr.Phone, usr.Age)
	if result != nil {
		return errors.New("FAILED TO ADD USER")
	}
	return nil
}

func (userservice *UserService) GetUser(id int) (models.User, error) {
	user, err := userservice.storeInterface.ReadOne(id)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (userservice *UserService) GetAllUser() ([]models.User, error) {

	userList, err := userservice.storeInterface.ReadAll()
	if err != nil {
		return userList, err
	}
	return userList, nil
}

func (userservice *UserService) UpdateUser(usr models.User) error {

	err := userservice.storeInterface.Update(usr.Id, usr.Name, usr.Email, usr.Phone, usr.Age)
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
