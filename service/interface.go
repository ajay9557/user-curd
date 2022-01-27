package service

import "user-curd/entities"

type (
	UserServiceHandler interface {
		GetUserByIdService(interface{}) (*entities.User, error)
		GetAllUsersService() ([]*entities.User, error)
		CreateUserService(user entities.User) error
		UpdateUserService(user entities.User) error
		DeleteUserService(interface{}) error
	}
)
