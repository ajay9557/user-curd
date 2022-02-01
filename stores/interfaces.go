package stores

import "user-curd/models"

type Store interface {
	InsertDetails(models.User) error
	UpdateById(models.User) error
	DeleteById(id int) error
	GetById(id int) (models.User, error)
	GetDetails() ([]models.User, error)
}
