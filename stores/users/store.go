package users

import (
	"database/sql"
	"errors"
	"github.com/usercurd/models"
	"github.com/usercurd/stores"
)

type userStore struct {
	db *sql.DB
}

func New(db *sql.DB) stores.User {
	return &userStore{db}
}

func (u *userStore) Insert(usr models.User) (models.User, error) {
	query := "insert into User values(?,?,?,?,?)"
	stmt, err := u.db.Prepare(query)
	defer func() {
		if err == nil {
			stmt.Close()
		}
	}()
	if err != nil {
		return models.User{}, err
	}
	_, err = stmt.Exec(usr.Id, usr.Name, usr.Email, usr.Phone, usr.Age)
	if err != nil {
		return models.User{}, err
	}
	return usr, nil
}

func (u *userStore) GetAll() ([]models.User, error) {
	var users []models.User
	query := "select * from User"
	rows, err := u.db.Query(query)
	defer func() {
		if err == nil {
			rows.Close()
		}
	}()

	if err != nil {
		return []models.User{}, err
	}
	for rows.Next() {
		usr := models.User{}
		rows.Scan(&usr.Id, &usr.Name, &usr.Email, &usr.Phone, &usr.Age)
		users = append(users, usr)
	}
	println(users)
	return users, nil
}

func (u *userStore) GetById(id int) (models.User, error) {
	var usr models.User
	query := "select * from User where id = ?"
	rows := u.db.QueryRow(query, id)
	err := rows.Scan(&usr.Id, &usr.Name, &usr.Email, &usr.Phone, &usr.Age)
	return usr, err
}

func (u *userStore) Update(id int, name string) error {
	query := "update User set name = ? where id = ?"
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, id)
	if err != nil {
		return errors.New("error in updating user name")
	}
	return nil
}

func (u *userStore) Delete(id int) error {
	query := "delete from User where id = (?)"
	_, err := u.db.Exec(query, id)
	return err
}
