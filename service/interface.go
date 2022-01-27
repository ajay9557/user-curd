package service

import "user-curd/entities"

type UserServiceHandler interface {
	GetUserByIdService(id interface{}) (*entities.User, error)
	GetAllUsersService() ([]*entities.User, error)
	CreateUserService(user entities.User) error
	UpdateUserService(user entities.User) error
	DeleteUserService(id interface{}) error
}
