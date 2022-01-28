package services

import "github.com/tejas/user-crud/models"

type User interface {
	GetUserById(id int) (models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(id int, user models.User) (int, error)
	DeleteUser(id int) (int, error)
	CreateUser(user models.User) (int, error)
}
