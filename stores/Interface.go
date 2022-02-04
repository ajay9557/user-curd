package stores

import "user-curd/model"

type User interface {
	GetById(Id int) (*model.User, error)
	GetAll() ([]*model.User, error)
	UpdateById(user *model.User) (*model.User, error)
	DeleteById(Id int) (err error)
	InsertDetails(user *model.User) (*model.User, error)
}
