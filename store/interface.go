package store

import (
	"github.com/tejas/user-crud/models"
)

type User interface {
	GetUserById(id int) (*models.User, error)
	GetUsers() ([]*models.User, error)
	CreateUsers(user models.User) error
	UpdateUser(user models.User) error
	DeleteUser(id int) error
}
