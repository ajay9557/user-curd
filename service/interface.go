package service

import "zopsmart/user-curd/model"

type User interface {
	CheckMail(email string) (bool, error)
	GetByID(id int) (model.User, error)
	PostUser(name string, email string, phone string, age int) (model.User, error)
	DeleteByID(id int) error
	GetUsers() ([]model.User, error)
	Update(model.User, model.User) (model.User, error)
}
