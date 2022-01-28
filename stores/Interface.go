package store

import "user-curd/model"

type Store interface {
	GetById(Id int) (model.User, error)
	GetAll() ([]model.User, error)
	UpdateById(user model.User) (err error)
	DeleteById(Id int) (err error)
	InsertDetails(user model.User) error
}
