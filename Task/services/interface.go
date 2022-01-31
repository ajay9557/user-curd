package service

import (
	"zopsmart/Task/models"
)

type Services interface {
	GetUserByIdService(Id int) (models.User, error)
	DeletebyIdService(Id int) error
	UpdatebyIdService(models.User) error
	GetAllUsersService() ([]models.User,error)
	CreateUserService(us models.User) error
}
