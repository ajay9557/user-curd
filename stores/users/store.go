package users

import (
	"database/sql"
	"errors"

	"github.com/tejas/user-crud/models"
	"github.com/tejas/user-crud/stores"
)

type dbStore struct {
	db *sql.DB
}

func New(db *sql.DB) stores.User {
	return &dbStore{db: db}
}

func (u *dbStore) GetUserById(id int) (*models.User, error) {
	db := u.db

	row := db.QueryRow("select * from user where id = ?", id)

	var user models.User

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)

	if err != nil {
		return nil, errors.New("Invalid Id")
	}

	return &user, nil
}

func (u *dbStore) GetUsers() ([]models.User, error) {
	db := u.db

	row, err := db.Query("select * from user")

	if err != nil {
		return []models.User{}, errors.New("Cannot fetch users")
	}

	var users []models.User

	for row.Next() {
		var u models.User

		_ = row.Scan(&u.Id, &u.Name, &u.Email, &u.Phone, &u.Age)

		users = append(users, u)
	}

	return users, nil
}

func (u *dbStore) UpdateUser(id int, user models.User) (int, error) {
	db := u.db

	_, err := db.Exec("update user set name = ?, email = ?, phone = ?, age = ?, where id = ?", user.Name, user.Email, user.Phone, user.Age, id)

	if err != nil {
		return 0, errors.New("Could not update user for given id")
	}

	return id, nil
}

func (u *dbStore) DeleteUser(id int) (int, error) {
	db := u.db

	res, err := db.Exec("delete from user where id = ?", id)

	if err != nil {
		return 0, errors.New("Could not delete user for given id")
	}

	rowcount, _ := res.RowsAffected()

	if rowcount == 0 {
		return 0, errors.New("Could not delete user for given id")
	}

	return int(rowcount), nil
}

func (u *dbStore) GetUserByEmail(email string) bool {
	db := u.db

	row, _ := db.Query("select * from user where email = ?", email)

	var user models.User

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)

	return err != nil

}

func (u *dbStore) CreateUser(user models.User) (int, error) {
	db := u.db

	res, err := db.Exec("insert into user(name, email, phone, age) values(?,?,?,?)", user.Name, user.Email, user.Phone, user.Age)

	if err != nil {
		return 0, errors.New("Could not create new user")
	}

	lastInsertId, _ := res.LastInsertId()

	return int(lastInsertId), nil
}
