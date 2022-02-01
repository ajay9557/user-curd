package store

import "zopsmart/user-curd/model"

type Store interface {
	GetUserById(id int) (model.User, error)
	AddUser(usr model.User) (int, error)
	DeleteUser(id int) error
	GetAllUsers() ([]model.User, error)
	UpdateUser(model.User) error
}
