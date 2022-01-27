package Services

import "Icrud/TModels"

type ISUser interface {
	UserById(id int) (TModels.User, error)
	GetUsers() ([]TModels.User, error)
	InsertUser(u TModels.User) (int, error)
	DeleteUserById(id int) (int, error)
	UpdateUserById(u TModels.User, id int) (int, error)
	// EmailValidation(email string) (bool, error)
}
