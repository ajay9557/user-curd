package stores

import (
	"go_lang/Assignment/user-curd/models"
)

type Crud interface {
	Create(name string, email string, phone string, age int) error
	ReadOne(id int) (models.User, error)
	ReadAll() ([]models.User, error)
	Update(id int, name string, email string, phone string, age int) error
	Delete(id int) error
}
