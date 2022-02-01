package users

import (
	"errors"
	"fmt"
	"user-curd/models"
	"user-curd/stores"
	//"fmt"
)

type UserDetails struct {
	u stores.Store
}

func New(us stores.Store) UserDetails {
	return UserDetails{us}
}

func (ud UserDetails) GetByUserId(id int) (models.User, error) {
	var newUser models.User
	Idcheck := CheckId(id)
	if Idcheck {
		res, err := ud.u.GetById(id)
		if err != nil {
			return newUser, err
		}
		return res, nil
	}
	return newUser, errors.New("invalid id")
}

func (ud UserDetails) InsertUserDetails(user models.User) (models.User, error) {
	var newUser models.User
	newUser.Id = user.Id
	newUser.Email = user.Email
	newUser.Phone = user.Phone
	newUser.Age = user.Age
	newUser.Name = user.Name
	checkEmail := true
	Idcheck := CheckId(user.Id)
	if Idcheck {
		ok := isEmailValid(user.Email)
		if ok {
			res, err := ud.u.GetDetails()
			if err != nil {
				checkEmail = false
			}
			for _, k := range res {
				if k.Email == user.Email {
					checkEmail = false
				}
			}
			if checkEmail {
				err := ud.u.InsertDetails(user)
				if err != nil {
					return newUser, errors.New("t")
				}
			}
		} else {
			return newUser, errors.New("invalid email")
		}
	} else {
		return newUser, errors.New("invalid id")
	}
	return newUser, nil
}

func (ud UserDetails) DeleteByUserId(i int) error {
	Idcheck := CheckId(i)
	if Idcheck {
		err := ud.u.DeleteById(i)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("invalid id")

}

func (ud UserDetails) UpdateByUserId(user models.User) error {
	checkEmail := true
	Idcheck := CheckId(user.Id)
	if Idcheck {
		ok := isEmailValid(user.Email)
		if ok {
			res, err := ud.u.GetDetails()
			if err != nil {
				checkEmail = false
			}
			for _, k := range res {
				if k.Email == user.Email {
					checkEmail = false
				}
			}
			if checkEmail {
				err := ud.u.UpdateById(user)
				if err != nil {
					fmt.Println(err)
					return errors.New("t")
				}
			}
		} else {
			return errors.New("invalid email")
		}
	} else {
		return errors.New("invalid id")
	}
	return nil
}

func (ud UserDetails) GetAll() ([]models.User, error) {
	var users []models.User
	res, err := ud.u.GetDetails()
	if err != nil {
		return users, errors.New("t")
	}
	return res, err

}
