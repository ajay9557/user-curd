package users

import (
	"database/sql"
	"errors"
	"user-crud/models"
	"user-crud/stores"
)

type userStore struct {
	db *sql.DB
}

func New(db *sql.DB) stores.User {
	return &userStore{db}
}

func (u *userStore) Insert(usr *models.User) (*models.User, error) {
	query := "insert into user1(name,email,phone,age) values(?,?,?,?)"
	res, err := u.db.Exec(query, usr.Name, usr.Email, usr.Phone, usr.Age)
	if err != nil {
		return &models.User{}, errors.New("error in executing insert")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return &models.User{}, errors.New("could not get last inserted id")
	}
	usr.Id = int(id)
	return usr, nil
}

func (u *userStore) GetAll() ([]*models.User, error) {
	var users []*models.User
	query := "select * from user1"
	rows, err := u.db.Query(query)
	defer func() {
		if err == nil {
			rows.Close()
		}
	}()

	if err != nil {
		return []*models.User{}, err
	}
	for rows.Next() {
		usr := models.User{}
		rows.Scan(&usr.Id, &usr.Name, &usr.Email, &usr.Phone, &usr.Age)
		users = append(users, &usr)
	}
	println(users)
	return users, nil
}

func (u *userStore) GetById(id int) (*models.User, error) {
	var usr models.User
	query := "select * from user1 where id = ?"
	row := u.db.QueryRow(query, id)
	err := row.Scan(&usr.Id, &usr.Name, &usr.Email, &usr.Phone, &usr.Age)
	if err == sql.ErrNoRows {
		return &models.User{}, errors.New("no id found")
	}
	//err := rows.Scan(usr.Id, usr.Name, usr.Email, usr.Phone, usr.Age)
	return &usr, err
}

func (u *userStore) Update(user *models.User) (*models.User, error) {
	feilds, values := stores.BuildQuery(*user)
	query := "update user1 set " + feilds + "where id = ?"
	_, err := u.db.Exec(query, values...)
	if err != nil {
		return &models.User{}, errors.New("error in updating user name")
	}
	// fetch the updated details
	usr, err := u.GetById(user.Id)
	return usr, nil
}

func (u *userStore) Delete(id int) error {
	query := "delete from user1 where id = (?)"
	_, err := u.db.Exec(query, id)
	return err
}
