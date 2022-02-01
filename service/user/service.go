package user

import (
	"fmt"
	"zopsmart/user-curd/model"
	"zopsmart/user-curd/store"
)

type UserStore struct {
	store store.Store
}

func New(userStore store.Store) UserStore {
	return UserStore{store: userStore}
}

func (user UserStore) GetByID(id int) (model.User, error) {
	var us model.User
	res, err := user.store.GetUserById(id)
	if err != nil {
		return us, fmt.Errorf("%v", err)
	}
	return res, nil
}

func (user UserStore) GetUsers() ([]model.User, error) {
	var us []model.User
	res, err := user.store.GetAllUsers()
	if err != nil {
		return us, fmt.Errorf("%v", err)
	}
	return res, nil

}
func (user UserStore) PostUser(usr model.User) (model.User, error) {
	isValid, _ := CheckMail(usr.Email)
	if !isValid {
		return model.User{}, fmt.Errorf("enter valid email")
	}
	users, err := user.store.GetAllUsers()
	for _, u := range users {
		if u.Email == usr.Email {
			return model.User{}, fmt.Errorf("Email already exists")
		}
	}
	usr.Id, _ = user.store.AddUser(usr)
	if err != nil {
		return model.User{}, err
	}

	return usr, err

}

func (user UserStore) DeleteByID(id int) error {
	if id < 1 {
		return fmt.Errorf("Id doesn't exist")
	}

	users, _ := user.store.GetAllUsers()
	for _, usr := range users {
		if usr.Id == id {
			err := user.store.DeleteUser(id)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
		}
	}
	return nil
}

func (user UserStore) Update(id int, ur model.User) (model.User, error) {
	if id < 1 {
		return model.User{}, fmt.Errorf("Id doesn't exist")
	}
	fmt.Println(ur)
	users, _ := user.store.GetAllUsers()
	for _, u := range users {
		if u.Id == id {
			userdata, err := user.store.GetUserById(id)
			if err != nil {
				return model.User{}, err
			}
			if ur.Name != "" {
				userdata.Name = ur.Name
			}
			isValid, _ := CheckMail(ur.Email)
			if ur.Email != "" {
				if isValid {
					users, _ := user.store.GetAllUsers()
					for _, u := range users {
						if u.Email == ur.Email {
							return model.User{}, fmt.Errorf("Email already exists")
						}
					}
					userdata.Email = ur.Email
				}
			}

			if ur.Phone != "" {
				userdata.Phone = ur.Phone
			}
			if ur.Age != 0 {
				userdata.Age = ur.Age
			}

			err = user.store.UpdateUser(userdata)
			return userdata, err

		}
	}
	return model.User{}, fmt.Errorf("Not a valid ID")
}
