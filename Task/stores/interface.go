package store


import (
	"zopsmart/Task/models"
)

type Store interface {
	GetAllUsersStore() ([]models.User,error)
	GetUserById(id int) (models.User, error)
	Create(Id int, Name string, Email string, Phone string, Age int) error
	Update(user models.User) error
	Delete(id int) error
}
