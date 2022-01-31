package users

import (
	"database/sql"
	"errors"
	"strings"
	"user-curd/models"
	"user-curd/stores"
)

type UserStorer struct {
	db *sql.DB
}

func New(db *sql.DB) stores	.Store {
	return &UserStorer{db: db}
}

func (u UserStorer) InsertUser(user models.User) error {

	insertQ := "insert into user(Id, Name, Email, Phone, Age) VALUES(?,?,?,?,?)" //Insert query
	if user.Id == 0 {
		return errors.New("id can`t be 0")
	}
	_, err := u.db.Exec(insertQ, user.Id, user.Name, user.Email, user.Phone, user.Age)
	if err != nil {
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
		_ = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
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
		_ = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
	}
	return user, nil
}

func (u UserStorer) UpdateUser(user models.User) error {
	updateEntities := []interface{}{}
	updateQ := "update user set " //Update query
	if user.Age > 0 {
		updateQ += "Age=?,"
		updateEntities = append(updateEntities, user.Age)
	}
	if user.Name != "" {
		updateQ += "Name=?,"
		updateEntities = append(updateEntities, user.Name)
	}
	if user.Email != "" {
		updateQ += "Email=?,"
		updateEntities = append(updateEntities, user.Email)
	}
	if user.Phone != "" {
		updateQ += "Phone=?,"
		updateEntities = append(updateEntities, user.Phone)
	}
	updateQ = strings.TrimRight(updateQ, ",")
	if user.Id != 0 {
		updateQ += " where Id=?;"
		updateEntities = append(updateEntities, user.Id)
	}

	_, err := u.db.Exec(updateQ, updateEntities...)
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