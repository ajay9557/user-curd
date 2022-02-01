package services

import "user-crud/models"

type User interface {
	Insert(usr *models.User) (*models.User, error)
	GetAll() ([]*models.User, error)
	GetById(id int) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id int) error
}
