package services

import "github.com/tejas/user-crud/models"

type User interface {
	GetUserById(id int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	UpdateUser(user models.User) error
	CreateUser(user models.User) error
	DeleteUser(id int) error
}
