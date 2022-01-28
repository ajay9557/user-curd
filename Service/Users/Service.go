package Users

import (
	"errors"
	"fmt"
	"user-curd/model"
	store "user-curd/stores"

	"regexp"
)

type User_Details struct {
	u store.Store
}

func New(us store.Store) User_Details {
	return User_Details{us}
}

func (ud User_Details) SearchByUserId(id int) (model.User, error) {
	var newUser model.User
	res, err := ud.u.GetById(id)
	fmt.Println()
	if err != nil {
		return newUser, err
	}
	return res, nil
}

func (ud User_Details) InsertUserDetails(user model.User) (model.User, error) {

	var newUser model.User
	ud.u.InsertDetails(user)
	newUser.Id = user.Id
	newUser.Email = user.Email
	newUser.Phone = user.Phone
	newUser.Age = user.Age
	newUser.Name = user.Name

	return newUser, nil
}

func (ud User_Details) IsEmailValid(e string) bool {

	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(e) < 3 || len(e) > 254 {
		return false
	}
	if emailRegex.MatchString(e) {
		res, err := ud.u.GetAll()
		if err != nil {
			return false
		}
		for _, k := range res {
			if k.Email == e {
				return false
			}
		}
		return true
	}
	return false

}

func (ud User_Details) DeleteByUserId(i int) error {
	err := ud.u.DeleteById(i)
	if err != nil {
		return err
	}
	return nil
}

func (ud User_Details) UpdateByUserId(user model.User) error {
	err := ud.u.UpdateById(user)
	if err != nil {
		fmt.Println(err)
		return errors.New("t")
	}
	return nil
}

func (ud User_Details) SearchAll() ([]model.User, error) {
	var users []model.User
	res, err := ud.u.GetAll()
	if err != nil {
		return users, errors.New("t")
	}
	return res, err

}
