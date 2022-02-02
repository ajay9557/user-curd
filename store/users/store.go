package users

import (
	"database/sql"
	"errors"
	"log"

	"github.com/tejas/user-crud/models"
	"github.com/tejas/user-crud/store"
)

type dbStore struct {
	db *sql.DB
}

func New(db *sql.DB) store.User {
	return &dbStore{db: db}
}

func (u *dbStore) GetUserById(id int) (*models.User, error) {

	row := u.db.QueryRow("select id, name, email, phone, age from user where id = ?", id)

	var user models.User

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)

	if err != nil {
		return &models.User{}, errors.New("cannot fetch user details, invalid id")
	}

	return &user, nil
}

func (u *dbStore) GetUsers() ([]*models.User, error) {

	rows, err := u.db.Query("select id, name, email, phone, age from user")

	if err != nil {
		return []*models.User{}, errors.New("cannot fetch users")
	}

	var users []*models.User

	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Phone, &u.Age)

		if err != nil {
			return nil, errors.New("error")
		}

		users = append(users, &u)
	}

	return users, nil
}

func (u *dbStore) UpdateUser(user models.User) error {

	fields, values := preCheck(user)

	// _, err := u.db.Exec("update user set name = ?, email = ?, phone = ?, age = ? where id = ?", user.Name, user.Email, user.Phone, user.Age, user.Id)

	query := "update user set" + fields + " where id = ?"

	_, err := u.db.Exec(query, values...)

	if err != nil {
		return errors.New("error")
	}

	return nil
}

func (u *dbStore) CreateUsers(user models.User) error {

	_, err := u.db.Exec("insert into user values (?,?,?,?,?)", user.Id, user.Name, user.Email, user.Phone, user.Age)

	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (u *dbStore) DeleteUser(id int) error {

	_, err := u.db.Exec("delete from user where id = ?", id)

	if err != nil {
		return errors.New("cannot delete user data for the given id")
	}

	return nil
}
