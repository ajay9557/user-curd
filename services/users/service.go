package users

import (
	"errors"
	"user-curd/models"
	"user-curd/services"
	"user-curd/stores"
)

type UserServiceHandler struct {
	stores stores.Store //Users interface
}

func New(store stores.Store) services.Services {
	return UserServiceHandler{stores: store}
}

func (u UserServiceHandler) InsertUserDetails(user models.User) error {
	checkId := idCheck(user.Id)
	checkEmail := true
	if !checkId {
		return errors.New("id should not be zero")
	}
	ok := emailValidation(user.Email)
	if !ok {
		return errors.New("email not valid")
	}
	userDetails, _ := u.stores.FetchAllUsers()
	for _, v := range userDetails {
		if v.Email == user.Email {
			checkEmail = false
			break
		}
	}
	if !checkEmail {
		return errors.New("email already present")
	}
	err := u.stores.InsertUser(user)
	if err != nil {
		return errors.New("error generated")
	}
	return nil
}

func (u UserServiceHandler) FetchAllUserDetails() ([]models.User, error) {
	var allUsers []models.User
	allUsers, err := u.stores.FetchAllUsers()
	if err != nil {
		return allUsers, errors.New("data fetching error")
	}
	return allUsers, nil
}

func (u UserServiceHandler) FetchUserDetailsById(id int) (models.User, error) {
	checkId := idCheck(id)
	if !checkId {
		return models.User{}, errors.New("id should not be zero")
	}
	user, err := u.stores.FetchUserById(id)
	if err != nil {
		return user, errors.New("data fetching error")
	}
	return user, nil
}

func (u UserServiceHandler) UpdateUserDetails(user models.User) error {
	checkId := idCheck(user.Id)
	checkEmail := true
	if !checkId {
		return errors.New("id should not be zero")
	}
	ok := emailValidation(user.Email)
	if !ok {
		return errors.New("email not valid")
	}
	userDetails, _ := u.stores.FetchAllUsers()
	for _, v := range userDetails {
		if v.Email == user.Email {
			checkEmail = false
			break
		}
	}
	if !checkEmail {
		return errors.New("email already present")
	}
	err := u.stores.UpdateUser(user)
	if err != nil {
		return errors.New("error generated")
	}
	return nil
}

func (u UserServiceHandler) DeleteUserDetailsById(id int) error {
	checkId := idCheck(id)
	if !checkId {
		return errors.New("id should not be zero")
	}
	err := u.stores.DeleteUserById(id)
	if err != nil {
		return errors.New("error generated")
	}
	return nil
}
