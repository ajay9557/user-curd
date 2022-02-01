package service

import "zopsmart/user-curd/model"

type User interface {
	GetByID(id int) (model.User, error)
	PostUser(model.User) (model.User, error)
	DeleteByID(id int) error
	GetUsers() ([]model.User, error)
	Update(int, model.User) (model.User, error)
}
