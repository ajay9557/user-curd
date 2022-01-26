package services

import "user-curd/models"

type Services interface {
	EmailValidation(email string) (bool, error)
	InsertUserDetails(models.User) error
	FetchAllUserDetails() ([]models.User, error)
	FetchUserDetailsById(id int) (models.User, error)
	UpdateUserDetails(user models.User) error
	DeleteUserDetailsById(id int) error
}
