package services

import "layer/user/models"

type User interface {
	GetUserById(id int) (models.User, error)
	GetUsers() ([]*models.User, error)
	UpdateUser(id int, user models.User) (*models.User, error)
	DeleteUser(id int) error
	CreateUser(models.User) (*models.User, error)
}
