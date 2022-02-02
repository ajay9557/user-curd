package Users

import (
	"errors"
	"fmt"
	"user-curd/model"
	store "user-curd/stores"
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

	checkEmail := true
	Idcheck := CheckId(user.Id)
	if Idcheck {
		ok := emailvalidation(user.Email)
		if ok {
			res, err := ud.u.GetAll()
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

func (ud User_Details) IsEmailValid(e string) bool {

	ok := emailvalidation(e)

	if ok {
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
	idcheck := CheckId(i)
	if idcheck {
		err := ud.u.DeleteById(i)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Id is not valid")
}

func (ud User_Details) UpdateByUserId(user model.User) error {
	checkEmail := true
	Idcheck := CheckId(user.Id)
	if Idcheck {
		ok := emailvalidation(user.Email)
		if ok {
			res, err := ud.u.GetAll()
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

func (ud User_Details) SearchAll() ([]model.User, error) {
	var users []model.User
	res, err := ud.u.GetAll()
	if err != nil {
		return users, errors.New("t")
	}
	return res, err

}
