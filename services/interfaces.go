package services

import "user-crud/models"

type User interface {
	Insert(usr models.User) (models.User, error)
	GetAll() ([]models.User, error)
	GetById(id int) (models.User, error)
	Update(id int, name string) error
	Delete(id int) (models.User, error)
}
