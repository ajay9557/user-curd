package services

import "user-curd/model"

type User interface {
	SearchByUserId(id int) (*model.User, error)
	InsertUserDetails(user *model.User) (*model.User, error)
	DeleteByUserId(Id int) error
	UpdateByUserId(usr *model.User) (*model.User, error)
	SearchAll() ([]*model.User, error)
}
