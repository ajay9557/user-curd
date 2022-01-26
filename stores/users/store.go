package users

import (
	"user-curd/models"
	"user-curd/stores"
	"database/sql"
	"errors"
	"fmt"
)

type UserStorer struct {
	db *sql.DB
}

func New(db *sql.DB) stores.Store {
	return &UserStorer{db: db}
}

func (u UserStorer) InsertUser(user models.User) error {

	insertQ := "INSERT INTO user(Id, Name, Email, Phone, Age) VALUES(?,?,?,?,?)" //Insert query
	if user.Id == 0 {
		return errors.New("id can`t be 0")
	}
	_, err := u.db.Exec(insertQ, user.Id, user.Name, user.Email, user.Phone, user.Age)
	if err != nil {
		fmt.Println(err)
		return errors.New("database error")
	}
	return nil

}

func (u UserStorer) FetchAllUsers() ([]models.User, error) {
	var allUsers []models.User
	var user models.User
	ReadQ := "Select Id,Name,Email,Phone,Age from user"
	rows, err := u.db.Query(ReadQ)
	if err != nil {
		return nil, errors.New("database error")
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
		if err != nil {
			return nil, errors.New("database row scan error")
		}
		allUsers = append(allUsers, user)
	}
	return allUsers, nil
}

func (u UserStorer) FetchUserById(id int) (models.User, error) {
	ReadQ := "Select Id,Name,Email,Phone,Age from user where Id=?"
	var user models.User
	rows, err := u.db.Query(ReadQ, id)
	if err != nil {
		return user, errors.New("database error")
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
		if err != nil {
			return user, errors.New("database error")
		}
	}
	return user, nil
}

func (u UserStorer) UpdateUser(user models.User) error {
	updateQ := "Update user set Name=?,Age=?,Email=?,Phone=? where Id=?;"
	_, err := u.db.Exec(updateQ, user.Name, user.Age, user.Email, user.Phone, user.Id)
	if err != nil {
		return errors.New("database error")
	}
	return nil
}

func (u UserStorer) DeleteUserById(id int) error {
	deleteQ := "DELETE From user where Id=?;"
	_, err := u.db.Exec(deleteQ, id)
	if err != nil {
		return errors.New("database error")
	}
	return nil
}

func (u UserStorer) GetEmail(email string) (bool, error) {
	Emails := []string{}
	searchQ := "select Email from user where email=?;"
	rows, err := u.db.Query(searchQ, email)
	if err != nil {
		return false, errors.New("database error")
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&email)
		if err != nil {
			return false, errors.New("database error")
		}
		Emails = append(Emails, email)
	}
	return len(Emails) == 0, nil
}
