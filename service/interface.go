package service

import "user-curd/entities"

type UserService interface {
	GetUserByIdService(id int) (*entities.User, error)
	GetAllUsersService() ([]*entities.User, error)
	CreateUserService(user entities.User) error
	UpdateUserService(user entities.User) error
	DeleteUserService(id int) error
}
