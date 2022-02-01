package service

import (
	"errors"
	"zopsmart/Task/models"
	service "zopsmart/Task/services"
	store "zopsmart/Task/stores"
)

type UserServiceHandler struct {
	stores store.Store
}

func New(store store.Store) service.Services {
	return UserServiceHandler{stores: store}
}

func (u UserServiceHandler) GetAllUsersService() ([]models.User,error){
	return u.stores.GetAllUsersStore()
}

func (u UserServiceHandler) GetUserByIdService(Id int) (models.User, error) {
	ok := validateId(Id)
	if !ok{
		return models.User{},errors.New("invalid Id")
	}
	res, err := u.stores.GetUserById(Id)
	if err != nil {
		return res, errors.New("error occured while getting the data")
	}

	return res, nil
}

func (u UserServiceHandler) DeletebyIdService(Id int) error {
	ok := validateId(Id)
	if !ok {
		return errors.New("invalid Id")
	}
	err := u.stores.Delete(Id)
	if err != nil {
		return errors.New("error occured while deleting data")
	}

	return nil
}

func (u UserServiceHandler) UpdatebyIdService(us models.User) error {
	okid := validateId(us.Id)
	if !okid {
		return errors.New("invalid Id")
	}
	okmail := validateEmail(us.Email)
	if !okmail {
		return errors.New("invalid Email")
	}
	okphone := validatePhone(us.Phone)
	if !okphone {
		return errors.New("invalid Phone")
	}
	err := u.stores.Update(us)

	if err != nil {
		return errors.New("Error")
	}

	return nil
}
func (u UserServiceHandler) CreateUserService(us models.User) error{
	okid := validateId(us.Id)
	if !okid {
		return errors.New("invalid Id")
	}
	okmail := validateEmail(us.Email)
	if !okmail {
		return errors.New("invalid Email")
	}
	okphone := validatePhone(us.Phone)
	if !okphone {
		return errors.New("invalid Phone")
	}
	err := u.stores.Create(us.Id,us.Name,us.Phone,us.Email,us.Age)
	if err!=nil {
		return errors.New("Error")
	}
	return nil
}
