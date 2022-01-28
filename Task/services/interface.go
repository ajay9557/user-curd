package service

import (
	"zopsmart/Task/models"
)

type Services interface {
	GetUserById(Id int) (models.User, error)
	DeletebyId(Id int) error
	ValidateId(Id int) (bool, error)
	ValidateEmail(email string) (bool, error)
	UpdatebyId(Id int, Phone string) error
	GetAllUsersService() ([]models.User,error)
}
