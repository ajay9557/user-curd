package users

import (
	"errors"
	"user-crud/models"
	"user-crud/services"
	"user-crud/stores"
)

type userServ struct {
	u stores.User
}

func New(st stores.User) services.User {
	return &userServ{st}
}
func (u userServ) Insert(usr models.User) (models.User, error) {
	usr, err := u.u.Insert(usr)
	return usr, err
}

func (u userServ) GetAll() ([]models.User, error) {
	var users []models.User
	users, err := u.u.GetAll()
	return users, err
}

func (u userServ) GetById(id int) (models.User, error) {
	usr, err := u.u.GetById(id)
	return usr, err
}

func (u userServ) Update(id int, name string) error {
	err := u.u.Update(id, name)
	return err
}

func (u userServ) Delete(id int) (models.User, error) {
	// check id exists
	usr, err := u.u.GetById(id)

	if err != nil {
		return models.User{}, errors.New("id does not exists")
	}
	err = u.u.Delete(id)
	if err != nil {
		return models.User{}, errors.New("unable to delete user")
	}
	return usr, err
}
