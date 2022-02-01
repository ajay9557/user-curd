package users

import (
	"database/sql"
	"errors"

	"layer/user/models"
	"layer/user/stores"
)

type dbStore struct {
	db *sql.DB
}

func New(db *sql.DB) stores.User {
	return &dbStore{db: db}
}

/*
GET /api/users/{id}
Fetch user for given id
*/
func (u *dbStore) GetUserById(id int) (*models.User, error) {
	db := u.db

	row := db.QueryRow("SELECT * FROM user WHERE id = ?", id)

	var user models.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Age)
	if err != nil {
		return nil, errors.New("Invalid Id")
	}

	return &user, nil
}

/*
GET /api/users
Fetch all users
*/
func (u *dbStore) GetUsers() ([]models.User, error) {
	db := u.db

	rows, err := db.Query("SELECT * FROM user")

	if err != nil {
		return []models.User{}, errors.New("Cannot fetch users")
	}

	var users []models.User

	for rows.Next() {
		var u models.User

		_ = rows.Scan(&u.Id, &u.Name, &u.Email, &u.Phone, &u.Age)

		users = append(users, u)
	}

	return users, nil
}

/*
PUT /api/users/{id}
Update user for given id
*/
func (u *dbStore) UpdateUser(id int, user models.User) (int, error) {
	db := u.db

	queryString := "UPDATE user"

	fields, args := formUpdateQuery(id, user)

	if fields != "" {
		queryString += " SET" + fields + " WHERE id = ?"
		_, err := db.Exec(queryString, args...)

		if err != nil {
			return -1, errors.New("Internal server error")
		}

		return id, nil
	}

	return -1, errors.New("Nothing to update. Provide user data to update the user")
}

/*
DELETE /api/users/{id}
Delete user for given id
*/
func (u *dbStore) DeleteUser(id int) error {
	db := u.db

	_, err := db.Exec("DELETE FROM user WHERE id = ?", id)

	if err != nil {
		return errors.New("Could not delete user for given id")
	}

	return nil
}

// Only used for email validation (email exists or not)
func (u *dbStore) GetUserByEmail(email string) bool {
	db := u.db

	row := db.QueryRow("SELECT id FROM user WHERE email = ?", email)

	var user int
	err := row.Scan(&user)

	return err != nil
}

/*
POST /api/users
Creating new user
*/
func (u *dbStore) CreateUser(user models.User) (int, error) {
	db := u.db

	result, err := db.Exec("INSERT INTO user(name, email, phone, age) VALUES(?, ?, ?, ?)", user.Name, user.Email, user.Phone, user.Age)

	if err != nil {
		return 0, errors.New("Could not create new user")
	}

	lastInsertedId, _ := result.LastInsertId()

	return int(lastInsertedId), nil
}
