package user

import (
	"fmt"
	"net/mail"
	"zopsmart/user-curd/model"
	"zopsmart/user-curd/store"
)

type UserDetails struct {
	u store.Store
}

func New(userStore store.Store) UserDetails {
	return UserDetails{u: userStore}
}

func (user UserDetails) CheckMail(email string) (bool, error) {
	_, err := mail.ParseAddress(email)
	if err != nil{
		return false, err
	}
	users, err := user.u.GetAllUsers()
	for _, usr := range users {
		if usr.Email == email {
			return false, err
		}
	}
	return true, nil

}

func (user UserDetails) GetByID(id int) (model.User, error) {
	var us model.User
	res, err := user.u.GetUserById(id)
	if err != nil {
		return us, fmt.Errorf("%v", err)
	}
	us.Id = res.Id
	us.Name = res.Name
	us.Email = res.Email
	us.Phone = res.Phone
	us.Age = res.Age
	return us, nil
}

func (user UserDetails) GetUsers() ([]model.User, error) {
	var us []model.User
	res, err := user.u.GetAllUsers()
	if err != nil {
		return us, fmt.Errorf("%v", err)
	}
	us = res
	return us, nil

}
func (user UserDetails) PostUser(name string, email string, phone string, age int) (model.User, error) {
	usr := model.User{0, name, email, phone, age}
	var newusr model.User
	var err error
	usr.Id, _ = user.u.AddUser(usr)
	if err != nil {
		return newusr, err
	}
	newusr.Id = usr.Id
	newusr.Name = usr.Name
	newusr.Email = usr.Email
	newusr.Phone = usr.Phone
	newusr.Age = usr.Age
	return newusr, err

}

func (user UserDetails) DeleteByID(id int) error {
	err := user.u.DeleteUser(id)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func (user UserDetails) Update(userdata model.User, ur model.User) (model.User, error) {
	err := user.u.UpdateUser(ur)
	ur.Id = userdata.Id
	if ur.Name != "" {
		userdata.Name = ur.Name
	}
	isValid, _ := user.CheckMail(ur.Email)
	if isValid {
		userdata.Email = ur.Email
	}
	if ur.Phone != "" {
		userdata.Phone = ur.Phone
	}
	if ur.Age != 0 {
		userdata.Age = ur.Age
	}

	return userdata, err
}
