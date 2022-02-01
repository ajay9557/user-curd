package stores

import (
	"go_lang/Assignment/user-curd/models"
)

type Crud interface {
	Create(user models.User) error
	ReadOne(id int) (models.User, error)
	ReadAll() ([]models.User, error)
	Update(user models.User) error
	Delete(id int) error
}
