package stores

import "user-curd/models"

type Store interface {
	InsertDetails(models.User) error
	UpdateById(models.User) error
	DeleteById(id int) error
	SearchById(id int) (models.User, error)
	SearchDetails() ([]models.User, error)
}
