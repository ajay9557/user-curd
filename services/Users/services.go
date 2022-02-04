package Users

import (
	"errors"
	"user-curd/model"
	"user-curd/services/validations"
	store "user-curd/stores"
)

type UserDetails struct {
	u store.User
}

func New(us store.User) UserDetails {
	return UserDetails{us}
}

func (ud UserDetails) SearchByUserId(id int) (*model.User, error) {
	usr, err := ud.u.GetById(id)
	return usr, err
}

func (ud UserDetails) InsertUserDetails(user *model.User) (*model.User, error) {
	if !validations.ValidateEmail(user.Email) {
		return &model.User{}, errors.New("invalid emailID")
	}
	if !validations.ValidatePhone(user.Phone) {
		return &model.User{}, errors.New("invalid Phone number")
	}
	user, err := ud.u.InsertDetails(user)
	return user, err
}

func (ud UserDetails) DeleteByUserId(Id int) error {
	_, err := ud.u.GetById(Id)

	if err != nil {
		return err
	}
	err = ud.u.DeleteById(Id)
	if err != nil {
		return errors.New("unable to delete user")
	}
	return nil
}

func (ud UserDetails) UpdateByUserId(usr *model.User) (*model.User, error) {
	if !validations.ValidateEmail(usr.Email) {
		return &model.User{}, errors.New("invalid emailID")
	}
	if !validations.ValidatePhone(usr.Phone) {
		return &model.User{}, errors.New("invalid Phone number")
	}
	usr, err := ud.u.UpdateById(usr)
	return usr, err

}

func (ud UserDetails) SearchAll() ([]*model.User, error) {
	var users []*model.User
	users, err := ud.u.GetAll()
	return users, err
}
