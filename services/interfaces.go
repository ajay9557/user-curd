package services

import "user-curd/models"

type Service interface {
	GetByUserId(id int) (models.User, error)
	InsertUserDetails(models.User) (models.User, error)
	DeleteByUserId(id int) error
	UpdateByUserId(models.User) error
	GetAll() ([]models.User, error)
}
