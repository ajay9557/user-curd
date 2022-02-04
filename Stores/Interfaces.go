package Stores

import "Icrud/TModels"

type IStore interface {
	UserById(id int) (TModels.User, error)
	GetUsers() ([]*TModels.User, error)
	InsertUser(u TModels.User) (int, error)
	// DeleteUserById(id int) (error)
	DeleteUserById(id int) error
	UpdateUserById(u TModels.User, id int) error
	GetEmail(email string) (bool, error)
}
