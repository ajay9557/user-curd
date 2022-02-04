package Users

import (
	"Icrud/Services"
	"Icrud/Stores"
	"Icrud/TModels"
	"errors"
	"reflect"

	"fmt"
)

type ServUser struct {
	isu Stores.IStore
}

func New(su Stores.IStore) Services.ISUser {
	return &ServUser{su}
}

func (su *ServUser) UserById(id int) (TModels.User, error) {
	var u TModels.User
	res, err := su.isu.UserById(id)
	if err != nil {
		fmt.Println(err)
		return u, err
	}
	u.Name, u.Email, u.Phone, u.Age = res.Name, res.Email, res.Phone, res.Age
	u.Id = id

	return u, nil
}

func (su *ServUser) GetUsers() ([]*TModels.User, error) {
	var u []*TModels.User
	res, err := su.isu.GetUsers()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	u = res
	return u, nil

}

func (su *ServUser) InsertUser(users TModels.User) (TModels.User, error) {
	// var iid int
	var uuser TModels.User
	if reflect.DeepEqual(users, TModels.User{}) {
		return uuser, errors.New("Need an non-empty user data")
	}
	validEmail := validateEmail(users.Email)
	if !validEmail {
		return uuser, errors.New("Invalid EMail")
	}

	validPhone := validatePhone(users.Phone)
	if !validPhone {
		return uuser, errors.New("Invalid Phone")
	}
	isEmailExist, _ := su.isu.GetEmail(users.Email)
	if isEmailExist {
		return uuser, errors.New("email id is already in use")
	}
	newId, err := su.isu.InsertUser(users)
	if err != nil {
		return uuser, err
	}
	// iid = res
	// return iid, nil
	updatedUser, _ := su.isu.UserById(newId)
	return updatedUser, nil
}

func (su *ServUser) DeleteUserById(id int) error {
	// var iid int
	// if id < 1 {
	// 	return errors.New("Id must be >0")
	// }
	err := su.isu.DeleteUserById(id)
	if err != nil {
		return err
	}
	// iid = res
	return nil
}

func (su *ServUser) UpdateUserById(u TModels.User, id int) (TModels.User, error) {
	var uu TModels.User
	// var iid int

	if id < 1 {
		return uu, errors.New("User id should be greater than 0")
	}
	_, err := su.isu.UserById(id)
	if err != nil {
		return uu, err
	}
	err = su.isu.UpdateUserById(u, id)
	if err != nil {
		// fmt.Println(err)
		return uu, err
	}
	// iid = res
	// return iid, nil
	updatedUser, _ := su.isu.UserById(id)

	return updatedUser, nil

}

// func (su *ServUser) EmailValidation(email string) (bool, error) {
// 	ok, err := su.isu.GetEmail(email)
// 	if err != nil {
// 		return ok, errors.New("error Generated")
// 	}
// 	return ok, nil
// }
