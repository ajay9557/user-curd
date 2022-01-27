package datastore

import (
	"user-curd/entities"
)

type UserStore interface {
	GetUserById(id interface{}) (*entities.User, error)
	GetAllUsers() ([]*entities.User, error)
	CreateUser(user entities.User) error
	UpdateUser(user entities.User) error
	DeleteUser(id interface{}) error
}
