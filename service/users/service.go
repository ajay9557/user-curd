package users

import (
	"errors"
	"user-curd/datastore"
	"user-curd/entities"
)

type UserService struct {
	usrStoreHandler datastore.UserStore
}

func New(store datastore.UserStore) *UserService {
	return &UserService{
		usrStoreHandler: store,
	}
}

func (us *UserService) GetUserByIdService(id interface{}) (*entities.User, error) {

	// check if id is valid
	if validateId(id) {
		resp, err := us.usrStoreHandler.GetUserById(id)
		if err != nil {
			return &entities.User{}, err
		}
		return resp, nil
	}
	return &entities.User{}, errors.New("error invalid id type")
}

func (us *UserService) GetAllUsersService() ([]*entities.User, error) {

	return us.usrStoreHandler.GetAllUsers()
}

func (us *UserService) CreateUserService(user entities.User) error {

	// validate id, email and phone
	if !validateId(user.Id) {
		return errors.New("error invalid id")
	}
	if !validateEmail(user.Email) {
		return errors.New("error invalid email")
	}
	if !validatePhone(user.Phone) {
		return errors.New("error invalid phone")
	}
	return us.usrStoreHandler.CreateUser(user)
}

func (us *UserService) UpdateUserService(user entities.User) error {

	// Validate id, email and phone only if they are to be updated
	if !validateId(user.Id) {
		return errors.New("error invalid id")
	}
	if user.Email != "" && !validateEmail(user.Email) {
		return errors.New("error invalid email")
	}
	if user.Phone != "" && !validatePhone(user.Phone) {
		return errors.New("error invalid phone")
	}
	return us.usrStoreHandler.UpdateUser(user)
}

func (us *UserService) DeleteUserService(id interface{}) error {

	// validation of id
	if !validateId(id) {
		return errors.New("error invalid id")
	}
	return us.usrStoreHandler.DeleteUser(id)
}
