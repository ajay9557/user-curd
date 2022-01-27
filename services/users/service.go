package users

import (
	"user-curd/models"
	"user-curd/stores"
	"errors"
	"fmt"

	//"fmt"
	"regexp"
)

type UserDetails struct {
	u stores.Store
}

func New(us stores.Store) UserDetails {
	return UserDetails{us}
}

func (ud UserDetails) SearchByUserId(id int) (models.User, error) {
	var newUser models.User
	res, err := ud.u.SearchById(id)
	fmt.Println()
	if err != nil {
		return newUser, err
	}
	return res, nil
}

func (ud UserDetails) InsertUserDetails(user models.User) (models.User, error) {
	//fmt.Println(user)
	var newUser models.User
	ud.u.InsertDetails(user)
	//fmt.Print(err)
	newUser.Id = user.Id
	newUser.Email = user.Email
	newUser.Phone = user.Phone
	newUser.Age = user.Age
	newUser.Name = user.Name

	return newUser, nil
}

func (ud UserDetails) IsEmailValid(e string) bool {

	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(e) < 3 || len(e) > 254 {
		return false
	}
	if emailRegex.MatchString(e) {
		res, err := ud.u.SearchDetails()
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

func (ud UserDetails) DeleteByUserId(i int) error {
	err := ud.u.DeleteById(i)
	if err != nil {
		return err
	}
	return nil
}

func (ud UserDetails) UpdateByUserId(user models.User) error {
	err := ud.u.UpdateById(user)
	if err != nil {
		fmt.Println(err)
		return errors.New("t")
	}
	return nil
}

func (ud UserDetails) SearchAll() ([]models.User, error) {
	var users []models.User
	res, err := ud.u.SearchDetails()
	if err != nil {
		return users, errors.New("t")
	}
	return res, err

}
