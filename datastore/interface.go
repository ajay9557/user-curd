package datastore

import (
	"user-curd/entities"
)

type (
	UserStoresHandler interface {
		GetUserById(interface{}) (*entities.User, error)
		GetAllUsers() ([]*entities.User, error)
		CreateUser(entities.User) error
		UpdateUser(entities.User) error
		DeleteUser(interface{}) error
	}
)
