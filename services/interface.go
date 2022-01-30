package services

import "go_lang/Assignment/user-curd/models"

type Services interface {
	AddUser(usr models.User) error
	GetUser(id int) (models.User, error)
	GetAllUser() ([]models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id int) error
}
