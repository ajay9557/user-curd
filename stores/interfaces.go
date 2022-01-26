package stores

import (
	"user-curd/models"
)

type Store interface {
	InsertUser(models.User) error
	FetchAllUsers() ([]models.User, error)
	FetchUserById(id int) (models.User, error)
	UpdateUser(models.User) error
	DeleteUserById(id int) error
	GetEmail(email string) (bool, error)
}
