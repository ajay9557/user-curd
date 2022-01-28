package Service

import "user-curd/model"

type Service interface {
	SearchByUserId(id int) (model.User, error)
	InsertUserDetails(model.User) (model.User, error)
	IsEmailValid(email string) bool
	DeleteByUserId(id int) error
	UpdateByUserId(model.User) error
	SearchAll() ([]model.User, error)
}
