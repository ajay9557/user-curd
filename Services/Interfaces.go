package Services

import "Icrud/TModels"

type ISUser interface {
	UserById(id int) (TModels.User, error)
	GetUsers() ([]*TModels.User, error)
	InsertUser(u TModels.User) (TModels.User, error)
	DeleteUserById(id int) error
	UpdateUserById(u TModels.User, id int) (TModels.User, error)
	// EmailValidation(email string) (bool, error)
}
