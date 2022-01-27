package services

import "user-curd/models"

type Service interface {
	SearchByUserId(id int) (models.User, error)
	InsertUserDetails(models.User) (models.User, error)
	IsEmailValid(email string) bool
	DeleteByUserId(id int) error
	UpdateByUserId(models.User) error
	SearchAll() ([]models.User, error)
}
