package Users

import (
	"Icrud/Services"
	"Icrud/Stores"
	"Icrud/TModels"
	"errors"

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

func (su *ServUser) GetUsers() ([]TModels.User, error) {
	var u []TModels.User
	res, err := su.isu.GetUsers()
	if err != nil {
		fmt.Println(err)
		return []TModels.User{}, err
	}
	u = res
	return u, nil

}

func (su *ServUser) InsertUser(users TModels.User) (int, error) {
	var iid int
	isEmailExist, _ := su.isu.GetEmail(users.Email)
	if isEmailExist {
		return iid, errors.New("email id is already in use")
	}
	res, err := su.isu.InsertUser(users)
	if err != nil {
		return iid, err
	}
	iid = res
	return iid, nil
}

func (su *ServUser) DeleteUserById(id int) (int, error) {
	var iid int
	res, err := su.isu.DeleteUserById(id)
	if err != nil {
		return iid, err
	}
	iid = res
	return iid, nil
}

func (su *ServUser) UpdateUserById(u TModels.User, id int) (int, error) {
	// var u TModels.User
	var iid int
	res, err := su.isu.UpdateUserById(u, id)
	if err != nil {
		fmt.Println(err)
		return iid, err
	}
	iid = res
	return iid, nil
}

// func (su *ServUser) EmailValidation(email string) (bool, error) {
// 	ok, err := su.isu.GetEmail(email)
// 	if err != nil {
// 		return ok, errors.New("error Generated")
// 	}
// 	return ok, nil
// }
