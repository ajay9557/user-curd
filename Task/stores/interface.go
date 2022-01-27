package store

import "zopsmart/Task/models"

type Store interface {
	GetUserById(id int) (models.User, error)
	Update(Id int, Phone string) error
	Delete(id int) error
	Create(Id int, Name string, Email string, Phone string, Age int) error
	GetMail(mail string) (bool,error)
}


